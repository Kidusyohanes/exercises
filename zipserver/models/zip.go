package models

import (
	"io"
	"bufio"
	"strings"
	//"fmt"
)

//Zip represents a zip code record.
//The `json:"..."` field tags allow us to change
//the name of the field when it is encoded into JSON
//see https://drstearns.github.io/tutorials/gojson/
type Zip struct {
	Code  string `json:"code,omitempty"`
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

//ZipSlice is a slice of *Zip,
//that is, pointers to Zip struct instances
type ZipSlice []*Zip

//ZipIndex is a map from strings to ZipSlices
type ZipIndex map[string]ZipSlice


func extractZipCode(line string) *Zip {
	buf := make([]rune, len(line))

	shouldCopy := true
	for i, v := range line {
		if v == '"' {
			shouldCopy = !shouldCopy
		} else {
			if shouldCopy {
				buf[i] = v
			}
		}
	}

	zSlice := strings.Split(string(buf), ",")

	return &Zip{zSlice[0], zSlice[3], zSlice[6]}
}


//LoadZips loads the zip code records from CSV stream
//returning a ZipSlice or an error.
func LoadZips(r io.Reader, expectedNumber int) (ZipSlice, error) {
	//TODO: process the `r` as a stream of CSV records
	//creating a Zip struct for each record, and
	//appending that to a ZipSlice that you return

	zList := make(ZipSlice, expectedNumber)

	scanner := bufio.NewScanner(r)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text() // Println will add back the final '\n'

		z := extractZipCode(line)

		if i >= len(zList) {
			zList = append(zList, z)
		} else {
			zList[i] = z
		}
	}

	return zList, nil
}


func BuildIndex(zList ZipSlice) ZipIndex {

	zInd := ZipIndex{}

	for _, zip := range zList {
			_, found := zInd[zip.City]

			if !found {
				zInd[zip.City] = ZipSlice{}
			}

			zInd[zip.City] = append(zInd[zip.City], zip)
	}

	return zInd
}
