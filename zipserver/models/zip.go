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
	/*
		Converts a line of string into a Zip struct.
	*/
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
	zList := make(ZipSlice, expectedNumber)

	/*
		TODO: Read the CSV file line by line, call extractZipCode function above, to
		create a master slice of all the zip codes. Reference below stackoverflow question
		to figure out how to read a file line by line.

		https://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
	*/

	return zList, nil
}


func BuildIndex(zList ZipSlice) ZipIndex {

	zInd := ZipIndex{}

	/*
		TODO: Finish this function that returns a ZipIndex, which is a map[string]ZipSlice,
		where city names are keys, and slice of Zip's are values.
	*/

	return zInd
}
