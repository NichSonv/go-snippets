package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	p = fmt.Print
	pl = fmt.Println
	pf = fmt.Printf
)

// func printPerson(pers Person) {
	// pl("Name:", pers.name)
// }

func main() {
	pl("Hello! Do you wish to fill this form? Y/N")
	reader := bufio.NewReader(os.Stdin)
	consign, err := reader.ReadString('\n')
	var check bool
	if err != nil {
		log.Fatal(err)
	}
	switch consign {
	case "Y", "y":
		check = true
	case "N", "n":
		check = false
	default:
		pl("Not a valid answer")
		
	}
	if !(check) {
		pl("Alright then, have a great day!")
	}


}