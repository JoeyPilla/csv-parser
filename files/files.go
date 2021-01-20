package files

import (
	"mime/multipart"
	"strings"
)

func HandleFile(f *multipart.FileHeader, year []string) []string {
	splitName := strings.Split(f.Filename, ".")
	fileExt := splitName[len(splitName)-1]
	fileName := strings.Join(splitName[:len(splitName)-1], ".")

	switch fileExt {
	case "csv":
		return HandleCSV(fileName, f, year[0])
	case "txt":
		return HandleTXT(fileName, f, year[0])
	case "xlsx":
		return HandleXLSX(fileName, f, year[0])
	default:
		return []string{}
	}
}
