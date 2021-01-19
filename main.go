package main

import (
	"csv-parser/files"
	"encoding/csv"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/octet-stream")

	err := r.ParseMultipartForm(200000) // grab the multipart form
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	formdata := r.MultipartForm // ok, no problem so far, read the Form data

	//get the *fileheaders
	uploads := formdata.File["multiplefiles"] // grab the filenames

	fileNames := []string{}

	for _, file := range uploads { // loop through the files one by one
		fileNames = append(fileNames, processFile(file)...)
	}

	files.ZipFiles("done.zip", fileNames)
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote("done.zip"))
	http.ServeFile(w, r, "done.zip")
}

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/receive_multiple").Methods("POST").HandlerFunc(uploadHandler)
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir("frontend/build/static")))
	r.PathPrefix("/static/").Handler(staticHandler)
	buildHandler := http.FileServer(http.Dir("frontend/build"))
	r.PathPrefix("/").Handler(buildHandler)

	createDirectory("processed")
	createDirectory("processed/csv")
	createDirectory("processed/xlsx")

	http.ListenAndServe(":8000", r)
}

func createDirectory(name string) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(name, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}

func processFile(f *multipart.FileHeader) []string {
	splitName := strings.Split(f.Filename, ".")
	fileExt := splitName[len(splitName)-1]
	fileName := strings.Join(splitName[:len(splitName)-1], ".")

	fmt.Println(fileExt, fileName)
	switch fileExt {
	case "csv":
		return files.HandleCSV(fileName, f)
	case "txt":
		return files.HandleTXT(fileName, f)
	case "xlsx":
		return files.HandleXLSX(fileName, f)
	default:
		return []string{}
	}
}

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
