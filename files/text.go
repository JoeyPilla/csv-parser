package files

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strings"
)

func HandleTXT(name string, f *multipart.FileHeader) []string {
	processedFileName := fmt.Sprintf("processed/%s.csv", name)
	pFile, err := os.Create(processedFileName)
	defer pFile.Close()
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	file, _ := f.Open()
	defer file.Close()
	writer := csv.NewWriter(pFile)
	defer writer.Flush()
	scanner := bufio.NewScanner(file)
	ledgerNumber, ledgerName := "", ""
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "|")
		ledgerNumber, ledgerName = processRow(row, ledgerNumber, ledgerName, writer)
	}
	return []string{processedFileName}
}
