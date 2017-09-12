package main

import (
	"encoding/json"
	"fmt"
)

type fcode struct {
	Timestamp int64
	vvv       string
	eee       string
	fff       string
	sss       string
}

func main() {

	fcodelist := [...]fcode{
		{vvv: "vvv", eee: "eee", fff: "fff", sss: "sss", Timestamp: 1350008880111},
		{vvv: "vvv", eee: "eee", fff: "ffff", sss: "ssss", Timestamp: 1350008880111},
	}

	fj, _ := json.Marshal(fcodelist)
	fmt.Println(string(fj))

	fcodelistb := []fcode{}
	fcodelistb = append(fcodelistb, fcode{vvv: "vvv", eee: "eee", fff: "f", sss: "s", Timestamp: 1350008880111})

	fcodelistb = append(fcodelistb, fcode{vvv: "vvv", eee: "eeee", fff: "ff", sss: "ss", Timestamp: 1350008880111})

	fcodelistb = append(fcodelistb, fcode{vvv: "vvv", eee: "eeeee", fff: "fff", sss: "sss", Timestamp: 1350008880111})

	jsonFoo, _ := json.Marshal(fcodelist)
	fmt.Println(string(jsonFoo))
}
