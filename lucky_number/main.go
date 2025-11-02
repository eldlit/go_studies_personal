package main

import (
	"github.com/eldlit/go_studies_personal/lucky_number/randon"

	"github.com/fatih/color"
)

func main() {

	greenColor := color.New(color.FgGreen)

	n := random.Number()
	greenColor.Println("Hey there, ")
	greenColor.Printf("Your lucky number is %d!\n", n)
}
