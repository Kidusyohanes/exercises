package models

import (
	"encoding/csv"
	"fmt"
	"io"
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

//LoadZips loads the zip code records from CSV stream
//returning a ZipSlice or an error.
func LoadZips(r io.Reader, expectedNumber int) (ZipSlice, error) {
	//TODO: process the `r` as a stream of CSV records
	//creating a Zip struct for each record, and
	//appending that to a ZipSlice that you return
	reader := csv.NewReader(r)
	fields, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading header row: %v", err)
	}
	if len(fields) < 7 {
		return nil, fmt.Errorf("CSV has %d fields but we require %d", len(fields), 7)
	}

	zips := make(ZipSlice, 0, expectedNumber)
	for {
		fields, err := reader.Read()
		if err == io.EOF {
			return zips, nil
		}
		if err != nil {
			return nil, fmt.Errorf("Error parsing CSV: %v", err)
		}
		z := &Zip{
			Code:  fields[0],
			City:  fields[3],
			State: fields[6],
		}
		zips = append(zips, z)
	}
}
