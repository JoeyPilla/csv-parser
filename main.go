package main

import (
	"csv-parser/files"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strconv"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	year := r.URL.Query()["year"]
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
		fileNames = append(fileNames, files.HandleFile(file, year)...)
	}

	files.ZipFiles("processed.zip", fileNames)
	w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote("processed.zip"))
	http.ServeFile(w, r, "processed.zip")
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
