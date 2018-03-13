package main

import "fmt"
import (
	"github.com/jung-kurt/gofpdf"
	"time"
)

func main() {
	a := "HELLO,assh1le"
	stringb := "oooooooo"

	fmt.Println(a)
	fmt.Print(a)
	fmt.Println()

	fmt.Printf("HELLO,%+v hhhh%+vhhhh\n", "asshole", "00")

	createPdf(a, stringb)
}

func createPdf(text string, textb string) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	content := text + ", " + time.Now().Format("2006-01-02 15:04:05") + ", " + textb

	pdf.Cell(40, 10, content)
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("create pdf with content:", content)
}
