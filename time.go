package hdnsclient

import (
	"fmt"
	"strings"
	"time"
)

//CustomTime struct to unmarshal time data from response
type CustomTime struct {
	time.Time
}

const ctLayout = "2006-01-02 15:04:05 -0700 MST"

//UnmarshalJSON for conversion from json
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

//MarshalJSON for conversion to json
func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == nilTime {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

var nilTime = (time.Time{}).UnixNano()

func (ct *CustomTime) isSet() bool {
	return ct.UnixNano() != nilTime
}
