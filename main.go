package main

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func readExcel(f string) {
	xlsx, err := excelize.OpenFile(f)
	if err != nil {
		log.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell, err := xlsx.GetCellValue("Sheet1", "B2")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			log.Print(colCell, "\t")
		}
		log.Println()
	}
}

func main() {
	log.Println("The program has started...")
}
