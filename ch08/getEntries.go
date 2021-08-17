package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"
)

var PORT = ":8765"
var DATAFILE = "/tmp/data.csv"

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

type PhoneBook []Entry

var data = PhoneBook{}

func readCSVFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		temp := Entry{
			Name:       line[0],
			Surname:    line[1],
			Tel:        line[2],
			LastAccess: line[3],
		}
		// Storing to global variable
		data = append(data, temp)
	}

	return nil
}

func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	for _, row := range data {
		temp := []string{row.Name, row.Surname, row.Tel, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}

func getFileHandler(w http.ResponseWriter, r *http.Request) {
	var tempFileName string

	// Create temporary file name
	f, err := os.CreateTemp("", "data*.txt")
	tempFileName = f.Name()

	// Remove the file
	defer os.Remove(tempFileName)

	// Save data to it
	err = saveCSVFile(tempFileName)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Cannot create: "+tempFileName)
		return
	}

	fmt.Println("Serving ", tempFileName)
	// Serve it to the client
	http.ServeFile(w, r, tempFileName)

	// 30 seconds to get the file
	time.Sleep(30 * time.Second)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	Body := "Thanks for visiting!\n"
	fmt.Fprintf(w, "%s", Body)
}

func main() {
	err := readCSVFile(DATAFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	mux.HandleFunc("/getContents/", getFileHandler)

	fmt.Println("Starting server on:", PORT)
	err = http.ListenAndServe(PORT, mux)
	fmt.Println(err)
}
