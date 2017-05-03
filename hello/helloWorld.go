package hello

import "fmt"

func main() {
	hw("go")
}

/* block comment */
func hw(name string) {
	fmt.Println("Hello," + name)
}
