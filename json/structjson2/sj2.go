package main

import (
	"encoding/json"
	"os"
	"time"
)

type MyUser struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	LastSeen time.Time `json:"lastSeen"`
}

func (u *MyUser) MarshalJSON() ([]byte, error) {
	type Alias MyUser
	return json.Marshal(&struct {
		LastSeen int64 `json:"lastSeen"`
		*Alias
	}{
		LastSeen: u.LastSeen.Unix(),
		Alias:    (*Alias)(u),
	})
}

type AnotherUser struct {
	MyUser
	Foo string `json:"foo"`
}

func (u *AnotherUser) MarshalJSON() ([]byte, error) {
	type Alias AnotherUser
	return json.Marshal(&struct {
		LastSeen int64 `json:"lastSeen"`
		*Alias
	}{
		LastSeen: u.LastSeen.Unix(),
		Alias:    (*Alias)(u),
	})
}
func main() {
	_ = json.NewEncoder(os.Stdout).Encode(
		&MyUser{1, "Ken", time.Now()},
	)

	u := &AnotherUser{Foo: "bar"}
	u.Name = "name0"
	u.ID = 3
	u.LastSeen = time.Now()
	_ = json.NewEncoder(os.Stdout).Encode(
		u,
	)
}
