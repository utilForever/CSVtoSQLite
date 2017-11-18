package main

import (
	"bufio"
	"bytes"
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

	sqlFileBuffer := bufio.NewWriter(sqlFile)

	lineNum := 0
	numRecord := 0

	var strBuffer bytes.Buffer;

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}

		numRecord = len(record)

		if lineNum == 0 {
			strBuffer.WriteString("PRAGMA foreign_keys=OFF;\nBEGIN TRANSACTION;\nCREATE TABLE " + tableName + " (")
		} else if lineNum > 0 {
			strBuffer.WriteString("INSERT INTO " + tableName + " VALUES (")
		}

		for i := 0; i < numRecord; i++ {
			// Something...
		}

		strBuffer.WriteString(");\n")

		isWritten, err := sqlFileBuffer.WriteString(strBuffer.String())	
		if (err != nil) || (isWritten != len(strBuffer.Bytes())) {
			fmt.Printf("Warning: Error writing to SQL file line %d: %s", lineNum, err)
			return
		}

		strBuffer.Reset()

		lineNum++
	}

	strBuffer.WriteString("COMMIT;\n")

	isWritten, err := sqlFileBuffer.WriteString(strBuffer.String())
	if (err != nil) || (isWritten != len(strBuffer.Bytes())) {
		fmt.Printf("Warning: Error writing to SQL file line %d: %s", lineNum, err)
		return
	}

	sqlFileBuffer.Flush()
	strBuffer.Reset()
}
