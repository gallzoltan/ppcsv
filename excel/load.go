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
		readymap = append(readymap, Excelrow{
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
}

func (d Excelrow) GetReadyMap() []Excelrow {
	return readymap
}

func (d Excelrow) GetFailedMap() []Excelrow {
	return failedmap
}
