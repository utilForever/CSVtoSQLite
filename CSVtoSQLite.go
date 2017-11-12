package main

import (
	"flag"
	"fmt"
	"os"
)

var programVersion = "0.1"

var showHelpMenu bool;

func init() {
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
}