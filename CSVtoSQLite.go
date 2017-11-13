package main

import (
	"flag"
	"fmt"
	"os"
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
	fmt.Println("===================================");
	fmt.Println("   CSV to SQLite program Ver", programVersion);
	fmt.Println("===================================");
}

func main() {
	flag.Parse();

	if showHelpMenu {
		printAbout()
		flag.Usage()
		os.Exit(-3)
	}

	if csvFileName == "" || tableName == "" {
		fmt.Println("Error: Please provice both a '-t tableName' and '-f fileName.csv'")
		fmt.Println("Run 'CSVtoSQLite -h' for more information")
		os.Exit(-2)
	}
}