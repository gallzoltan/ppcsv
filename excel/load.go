package excel

import (
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Excelrow struct {
	Name       string
	Address    string
	Email      string
	Tel        string
	Device     string
	Sn         string
	Srvname    string
	Buyingdate string
	Startdate  string
}

func init() {}

func (d Excelrow) ReadXLSX(f string) []Excelrow {
	var mapping []Excelrow
	xlsx, err := excelize.OpenFile(f)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := xlsx.Rows("Worksheet")

	for rows.Next() {
		row := rows.Columns()
		mapping = append(mapping, Excelrow{
			Name:       row[0],
			Address:    row[1],
			Email:      row[2],
			Tel:        row[3],
			Device:     row[4],
			Sn:         row[5],
			Srvname:    row[6],
			Buyingdate: row[7],
			Startdate:  row[8],
		})
	}

	return mapping
}

// func GetXLSXRows(f string) []ExcelRow {
// 	return readXLSX(f)
// }
