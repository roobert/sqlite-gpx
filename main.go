package main

import (
	"fmt"
	. "github.com/roobert/sqlite-gpx/db"
	. "github.com/roobert/sqlite-gpx/error"
	"github.com/twpayne/go-gpx"
	"os"
	"time"
)

// read gps tracker data from sqlite db and output gpx file
func main() {
	CreateDB("data.db")

	query := "select timestamp, latitude, longitude from data"
	rows, err := DB.Query(query)
	CheckErr(err)

	var wpts []*gpx.WptType
	var timestamp time.Time
	var latitude float64
	var longitude float64

	for rows.Next() {
		err = rows.Scan(&timestamp, &latitude, &longitude)
		CheckErr(err)

		wpt := &gpx.WptType{
			Lat:  latitude,
			Lon:  longitude,
			Time: timestamp,
		}

		wpts = append(wpts, wpt)
	}

	rows.Close()

	g := &gpx.GPX{
		Version: "1.0",
		Creator: "Whatever",
		Wpt:     wpts,
	}

	if err := g.WriteIndent(os.Stdout, "", "  "); err != nil {
		fmt.Printf("err == %v", err)
	}
}
