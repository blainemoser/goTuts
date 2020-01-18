package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Timestamp ... Create a custom time-stamp
type Timestamp time.Time

// UnmarshalJSON This unmarshalling function will run whenever json.Unmarshal is run on the type Timestamp
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	fmt.Println("*t1:")
	fmt.Println(&t)
	// Note that the second param is sliced - the outer "" marks are removed
	value, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	// Handle possible error
	if err != nil {
		return err
	}

	*t = Timestamp(value)
	fmt.Println("*t2:")
	fmt.Println(&t) // same point in memory as before
	return nil
}

func main() {

	// Represents a timestamp from the Twitter API
	var input = `{
		"created_at": "Thu May 31 00:00:01 +0000 2012"
	}`

	// We will attempt to unmarshal the JSON param 'created_at' into a time value
	var val map[string]Timestamp

	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}

	fmt.Println(val)
	for index, value := range val {
		fmt.Println(index)
		fmt.Println(value)
	}

}
