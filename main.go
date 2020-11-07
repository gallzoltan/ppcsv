package main

import (
	"log"

	"github.com/gallzoltan/ppcsv/excel"
)

func main() {
	log.Println("The program has started...")

	d := new(excel.Excelrow)
	d.ReadXLSX("./bin/testfiles/568588.xlsx")
	for _, row := range d.GetReadyMap() {
		log.Printf("%s;%s;%s", row.Address, row.Name, row.Buyingdate)
	}

	log.Println("The program has been finished.")
}
