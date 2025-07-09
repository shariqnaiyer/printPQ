package main

import (
	"fmt"
)

type printer struct {
	ID string `json:"id"`
}

func main() {
	fmt.Println("Hello, welcome to printPQ")
	a := printer{ID: "Maverick12"}
	printPrinterID(a)

}

func printPrinterID(printer printer) {
	fmt.Println(printer.ID)
}
