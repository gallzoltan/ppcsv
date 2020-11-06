package main

import (
	"log"

	"github.com/gallzoltan/ppcsv/excel"
)

func main() {
	log.Println("The program has started...")
	excel.GetXLSXRows("./bin/testfiles/568588.xlsx")
	log.Println("The program has been finished.")
}
