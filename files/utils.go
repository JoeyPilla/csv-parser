package files

import (
	"encoding/csv"
	"log"
	"regexp"
	"strings"
)

func processRow(row []string, ledgerNumber, ledgerName string, writer *csv.Writer) (string, string) {
	re := regexp.MustCompile(`\d+`)
	for _, col := range row {
		if strings.Contains(col, "Ledger Account:") {
			split := strings.Split(col, "  ")
			return re.FindString(col), strings.TrimSpace(split[len(split)-1])
		} else if strings.Contains(col, "2020-") {
			newRow := append([]string{ledgerNumber, ledgerName}, row...)
			err := writer.Write(newRow)
			if err != nil {
				log.Fatalln("Error writing to file", err)
			}
		}
	}
	return ledgerNumber, ledgerName
}
