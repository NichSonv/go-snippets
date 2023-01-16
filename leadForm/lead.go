package main

import (
	"fmt"
	"os"
)

var (
	p = fmt.Print
	pl = fmt.Println
	pf = fmt.Printf
	sc = fmt.Scan
	scl = fmt.Scanln
	scf = fmt.Scanf
	consign string
	pers Person
)
type Person struct {
	name string
	age int
	email string
}

func printPerson(pers Person) {
	pl("Name:", pers.name)
	pl("Age:", pers.age)
	pl("Email:", pers.email)
}

func main() {
	pl("Hello! Do you wish to fill this form? Y/N")
	sc(&consign); p("\n")
	if !(consign == "Y" || consign == "y" || consign == "N" || consign == "n") {
		pl("Error! Invalid answer. Try again")
		main()
	}
	switch consign {
	case "Y","y":
		pl("Alright! What's your name?")
		scf("%s", &pers.name); p("\n")
		// if pers.name != {
			
		pl("Your age?")
		scf("%d", &pers.age); p("\n")
		pl("And your email?")
		scf("%s", &pers.email); p("\n")
	case "N","n":
		pl("Alright, have a nice day!")
		os.Exit(0)
	}

	confirm()
}

func confirm() {
	pl("Ok, all set! Please review your info:")
	printPerson(pers); p("\n")

	pl("Confirm? Y/N")
	sc(&consign); p("\n")
	if !(consign == "Y" || consign == "y" || consign == "N" || consign == "n") {
		pl("Error! Invalid answer. Try again")
		confirm()
	} else {
		pl("Alright! Thank you very much!")
		os.Exit(0)
	}
}