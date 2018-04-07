package autotesting

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

//defaultLateDays is the number of late days
//a given student gets by default. Use this
//when we don't find a student's NetID in
//the map.
const defaultLateDays = 4

//LateDaysMap tracks late days remaining
//for each student NetID. Note that this
//is NOT safe for concurrent use because
//Go maps are not safe for concurrent use.
type LateDaysMap map[string]int

//LoadLateDaysFromCSV loads and returns a LateDaysMap
//from CSV data read from `reader`. The CSV data MUST NOT
//contain a header row, and each record should contain just
//two fields: a string NetID; and a number of late days remaining
//for that NetID.
func LoadLateDaysFromCSV(reader io.Reader) (LateDaysMap, error) {
	//BUG: there's a bug in this code that will
	//cause this function to fail in some cases. Write
	//tests to discover the problem and fix it!
	//If your tests all pass without changing this code,
	//you aren't testing all the cases!

	//create a new CSV reader over `reader`
	csvReader := csv.NewReader(reader)
	//create a new LateDaysMap
	lateDaysMap := LateDaysMap{}
	//loop until no more records
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		numDays, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, fmt.Errorf("error converting %s to an integer: %v", record[1], err)
		}
		lateDaysMap[record[0]] = numDays
	}
	return lateDaysMap, nil
}

//Consume consumes a late day for the student identified
//by `netID`, and returns the number of late days remaining.
//Once the number of late days for a student falls to zero,
//it will remain at zero.
func (ldm LateDaysMap) Consume(netID string) int {
	//BUG: there's a bug in here that you
	//need to discover and fix! If all your tests
	//pass without modifying this code, you're not
	//testing all the possible cases!

	//if there is no entry for `netID` in the map,
	//create one with the default number of late days
	if ldm[netID] == 0 {
		ldm[netID] = defaultLateDays
	}
	//decrement the late days for the netID
	ldm[netID]--
	//return the number of late days left
	return ldm[netID]
}
