package main

import (
	"log"

	"github.com/gallzoltan/ppcsv/excel"
)

func main() {
	log.Println("The program has started...")
	d := new(excel.Excelrow)
	datas := d.ReadXLSX("./bin/testfiles/568588.xlsx")
	//log.Println(datas[0].Name)
	for _, row := range datas {
		log.Printf("%s;%s;%s", row.Address, row.Name, row.Buyingdate)
	}
	// m := map[string]interface{}{
	// 	"Data": excel.ReadXLSX("./bin/testfiles/568588.xlsx"),
	// }
	// log.Println(m["Data"])
	log.Println("The program has been finished.")
}
