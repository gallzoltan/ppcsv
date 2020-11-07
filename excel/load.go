package excel

import (
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Excelrow struct {
	Name       string
	Pcode      string
	City       string
	Address    string
	Email      string
	Tel        string
	Device     string
	Sn         string
	Srvname    string
	Buyingdate string
	Startdate  string
}

var rawmap []Excelrow
var readymap []Excelrow
var failedmap []Excelrow

func init() {}

func (d Excelrow) ReadXLSX(f string) {
	xlsx, err := excelize.OpenFile(f)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := xlsx.Rows("Worksheet")

	for rows.Next() {
		row := rows.Columns()
		rawmap = append(rawmap, Excelrow{
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
	validateDatas()
}

func (d Excelrow) GetReadyMap() []Excelrow {
	return readymap
}

func (d Excelrow) GetFailedMap() []Excelrow {
	return failedmap
}

func splitAddress(adr string) []string {
	var result []string
	return result
}

func validateDatas() {
	readymap = rawmap
	// for _, row := range rawmap {
	// 	log.Println(row.Address)
	// }
}
