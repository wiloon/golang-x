package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Datetime time.Time

func (t Datetime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type Good struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	CreatedAt Datetime `json:"created_at"`
	UpdatedAt Datetime `json:"updated_at"`
}

func main() {
	good := Good{123, "chenqiognhe", Datetime(time.Now()), Datetime(time.Now())}
	bytes, _ := json.Marshal(good)
	fmt.Printf("%s", bytes)
}
