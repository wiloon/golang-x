package basic

import "log"

func sliceX() {
	foo := []string{}

	//foo[0]="foov"
	foo = append(foo, "v0")
	log.Println(foo)
}
