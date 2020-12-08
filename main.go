package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gallzoltan/ppcsv/excel"
)

func saveCsv(datas []excel.Excelrow, fn string) {
	f, err := os.Create(fn)
	if err != nil {
		log.Fatal(err)
		f.Close()
		return
	}
	for i, row := range datas {
		if i == 0 {
			row.City = "Helység"
			row.Pcode = "Irsz"
		}
		fmt.Fprintf(f, "%s;%s;%s;%s;%s;%s;%s;%s;%s\r\n",
			row.Name,
			row.Pcode,
			row.City,
			row.Address,
			row.Device,
			row.Sn,
			row.Buyingdate,
			row.Startdate,
			row.PartnerID)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("file %s written successfully", fn)
}

func main() {
	xlsxFile := flag.String("source", "./568588.xlsx", "a forrás xlsx fájl helye")
	//xlsxFile := "./bin/testfiles/568588.xlsx"
	flag.Parse()
	log.Println("The program has started...")

	d := new(excel.Excelrow)
	d.ReadXLSX(*xlsxFile)
	saveCsv(d.GetReadyMap(), "readylist.csv")

	log.Println("The program has been finished.")
}
