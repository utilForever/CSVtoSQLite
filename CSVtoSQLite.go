package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var programVersion = "0.1"

var tableName string
var csvFileName string
var showHelpMenu bool

func init() {
	flag.StringVar(&tableName, "t", "", ": '-t tableName' where tableName is the name of the SQLite table to hold your CSV file data [MANDATORY]")
	flag.StringVar(&csvFileName, "f", "", ": '-f fileName.csv' where fileName.csv is the name and path to a CSV file that contains your data for conversion [MANDATORY]")
	flag.BoolVar(&showHelpMenu, "h", false, ": '-h' to provide more detailed help on using this program")
}

func printAbout() {
	fmt.Println("===================================")
	fmt.Println("   CSV to SQLite program Ver", programVersion)
	fmt.Println("===================================")
}

func getSQLFileName() (fileName string) {
	var onlyFileName = filepath.Base(csvFileName)
	var extension = filepath.Ext(csvFileName)
	onlyFileName = onlyFileName[0 : len(onlyFileName)-len(extension)]

	sqlFileName := onlyFileName + ".sql"
	return sqlFileName
}

func main() {
	flag.Parse()

	if showHelpMenu {
		printAbout()
		flag.Usage()
		os.Exit(-3)
	}

	if csvFileName == "" || tableName == "" {
		fmt.Println("Error: Please provide both a '-t tableName' and '-f fileName.csv'")
		fmt.Println("Run 'CSVtoSQLite -h' for more information")
		os.Exit(-2)
	}

	csvFile, err := os.Open(csvFileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-3)
	}

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	sqlFileName := getSQLFileName()

	sqlFile, err := os.Create(sqlFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer sqlFile.Close()

	numFields := 0

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}

		numFields = len(record)

		for i := 0; i < numFields; i++ {
			// Something...
		}
	}
}
