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
	Device     string
	Sn         string
	Buyingdate string
	Startdate  string
	PartnerID  string
}

var rawmap []Excelrow

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
			Device:     row[2],
			Sn:         row[3],
			Buyingdate: row[4],
			Startdate:  row[5],
			PartnerID:  row[6],
		})
	}
}

func (d Excelrow) GetReadyMap() []Excelrow {
	return rawmap
}

func splitAddress(line string) [3]string {
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
