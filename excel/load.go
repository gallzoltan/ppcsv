package excel

import (
	"log"
	"regexp"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Excelrow struct {
	Name       string
	Pcode      string
	City       string
	Address    string
	DefAdr     string
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
		adr := splitAddress(row[1])
		rawmap = append(rawmap, Excelrow{
			Name:       row[0],
			Pcode:      adr[0],
			City:       adr[1],
			Address:    adr[2],
			DefAdr:     row[1],
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

func splitAddress(line string) [3]string {
	//\b(\d{4})(?!(\/|\)))\b
	//\b[1-9]\d{3}\b(?!(\/|\)))
	var valid = regexp.MustCompile(`\b^[1-9]\d{3}\b`)
	result := [3]string{"", "", ""}

	if valid.MatchString(line) {
		result[0] = valid.FindString(line)
		line = valid.ReplaceAllString(line, "")
	}

	if strings.Contains(line, ",") {
		adr := strings.Split(line, ",")
		result[1] = adr[0]
		result[2] = adr[1]
	} else {
		result[2] = line
	}
	return result
}

func validateDatas() {
	//readymap = rawmap
	for _, row := range rawmap {
		if row.City == "" || row.Pcode == "" || len(row.Sn) != 28 {
			failedmap = append(failedmap, row)
		} else {
			readymap = append(readymap, row)
		}
	}
}
