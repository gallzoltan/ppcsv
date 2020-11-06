package excel

import (
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type ExcelRow struct {
	name       string
	address    string
	email      string
	tel        string
	device     string
	sn         string
	srvname    string
	buyingdate string
	startdate  string
}

func readXLSX(f string) []ExcelRow {
	var mapping []ExcelRow
	xlsx, err := excelize.OpenFile(f)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := xlsx.Rows("Worksheet")

	for rows.Next() {
		row := rows.Columns()
		mapping = append(mapping, ExcelRow{
			name:       row[0],
			address:    row[1],
			email:      row[2],
			tel:        row[3],
			device:     row[4],
			sn:         row[5],
			srvname:    row[6],
			buyingdate: row[7],
			startdate:  row[8],
		})
	}
	// m := map[string]interface{}{
	// 	"Data": mapping,
	// }
	// fmt.Println(m["Data"])
	return mapping
}

func GetXLSXRows(f string) []ExcelRow {
	return readXLSX(f)
}
