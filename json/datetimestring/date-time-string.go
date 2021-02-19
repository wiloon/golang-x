package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type DateTime time.Time

func (t DateTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type Good struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t Good) MarshalJSON() ([]byte, error) {
	type TmpJSON Good
	return json.Marshal(&struct {
		TmpJSON
		CreatedAt DateTime `json:"created_at"`
		UpdatedAt DateTime `json:"updated_at"`
	}{
		TmpJSON:   (TmpJSON)(t),
		CreatedAt: DateTime(t.CreatedAt),
		UpdatedAt: DateTime(t.UpdatedAt),
	})
}

func main() {
	good := Good{123, "chenqiognhe", time.Now(), time.Now()}
	bytes, _ := json.Marshal(good)
	fmt.Printf("%s", bytes)
}
