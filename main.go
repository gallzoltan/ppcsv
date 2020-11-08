package main

import (
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
	for _, row := range datas {
		fmt.Fprintf(f, "%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s;%s\r\n",
			row.Name,
			row.Pcode,
			row.City,
			row.Address,
			row.DefAdr,
			row.Email,
			row.Tel,
			row.Device,
			row.Sn,
			row.Srvname,
			row.Buyingdate,
			row.Startdate)
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
	log.Println("The program has started...")

	d := new(excel.Excelrow)
	d.ReadXLSX("./bin/testfiles/568588.xlsx")
	saveCsv(d.GetReadyMap(), "readylist.csv")
	saveCsv(d.GetFailedMap(), "failedlist.csv")

	log.Println("The program has been finished.")
}
