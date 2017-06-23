package basic

import (
	"fmt"
)

func Variable() {
	foo := "value0"
	fmt.Println(foo)
	fmt.Println(&foo)
}
