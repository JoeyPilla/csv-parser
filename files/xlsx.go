package files

import (
	"encoding/csv"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/tealeg/xlsx/v3"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
)

func HandleXLSX(name string, f *multipart.FileHeader, year string) []string {
	// open an existing file
	file, _ := f.Open()
	bytes, _ := ioutil.ReadAll(file)
	xlFile, err := xlsx.OpenBinary(bytes)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	sheetLen := len(xlFile.Sheets)
	names := []string{}
	if sheetLen == 1 {
		for _, sh := range xlFile.Sheets {
			names = append(names, handleSheet(sh, name, year)...)
		}
	} else {
		for _, sh := range xlFile.Sheets {
			names = append(names, handleSheet(sh, fmt.Sprintf("%s_sheet1", name), year)...)
		}
	}
	return names
}

func handleSheet(sheet *xlsx.Sheet, name, year string) []string {
	_, err := os.Stat("temp")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("temp", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}

	pFile, err := os.Create(fmt.Sprintf("temp/%s.csv", name))
	defer pFile.Close()
	writer := csv.NewWriter(pFile)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	var vals []string
	err = sheet.ForEachRow(func(row *xlsx.Row) error {
		if row != nil {
			vals = vals[:0]
			err := row.ForEachCell(func(cell *xlsx.Cell) error {
				str, err := cell.FormattedValue()
				if err != nil {
					return err
				}
				vals = append(vals, str)
				return nil
			})
			if err != nil {
				return err
			}
		}
		writer.Write(vals)
		return nil
	})
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	writer.Flush()
	processedFileName := fmt.Sprintf("processed/csv/%s.csv", name)
	file, err := os.Create(processedFileName)
	defer file.Close()
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	csvfile, err := os.Open(fmt.Sprintf("temp/%s.csv", name))
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	w := csv.NewWriter(file)
	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))
	processFile(r, w, year)
	w.Flush()
	xlsxFileName := fmt.Sprintf("processed/xlsx/%s.xlsx", name)

	CSVtoXLSX(processedFileName, xlsxFileName)

	return []string{processedFileName, xlsxFileName}
}

func CSVtoXLSX(csvPath, XLSXPath string) {
	csvFile, err := os.Open(csvPath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	f := excelize.NewFile()
	// Create a new worksheet.
	f.NewSheet("Sheet1")
	row := 1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		f.InsertRow("Sheet1", row)
		err = f.SetSheetRow("Sheet1", fmt.Sprintf("A%d", row), &record)
		row += 1
	}
	if err := f.SaveAs(XLSXPath); err != nil {
		fmt.Println(err)
	}
}
