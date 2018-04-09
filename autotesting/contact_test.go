package autotesting

import (
	"testing"
	"time"
)

func TestContactSetTracking(t *testing.T) {
	//TODO: write automated tests that create
	//an instance of a Contact struct, call
	//the .SetTrackingData() method to update
	//the tracking fields (.InsertedAt and .UpdatedAt),
	//and verify that the fields were set to the
	//correct values. You can't know *exactly*
	//what those fields will be set to, but if you
	//capture the current time just before you call
	//.SetTrackingData(), and again just after that
	//call returns, you can expect any modified field
	//values to be between those two times.
	cases := []struct {
		name    string
		contact Contact
	}{
		{"Both Fields Zero", Contact{FirstName: "test", LastName: "test"}},
	}

	for _, c := range cases {
		before := time.Now()
		c.contact.SetTrackingData()
		end := time.Now()
		if c.contact.InsertedAt.Before(before) || c.contact.InsertedAt.After(end) {
			t.Errorf("case %s: incorrect InsertedAt: expected between %v and %v, but got %v", c.name, before, end, c.contact.InsertedAt)
		}
	}
}
