package files

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func HandleCSV(name string, f *multipart.FileHeader, year string) []string {
	processedFileName := fmt.Sprintf("processed/csv/%s.csv", name)
	xlsxFileName := fmt.Sprintf("processed/xlsx/%s.xlsx", name)
	csvfile, _ := f.Open()
	// Open the file
	file, err := os.Create(processedFileName)
	defer file.Close()
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	writer := csv.NewWriter(file)
	// Parse the file
	r := csv.NewReader(csvfile)

	processFile(r, writer, year)
	writer.Flush()
	CSVtoXLSX(processedFileName, xlsxFileName)
	return []string{processedFileName, xlsxFileName}
}

func processFile(r *csv.Reader, writer *csv.Writer, year string) {
	// Iterate through the records
	ledgerNumber, ledgerName := "", ""
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		ledgerNumber, ledgerName = processRow(record, ledgerNumber, ledgerName, year, writer)
	}
}
