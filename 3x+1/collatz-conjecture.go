package main

import (
	"fmt"
	"reflect"
)

var (
	pl = fmt.Println
	pf = fmt.Printf
	// p = fmt.Print
	sl = fmt.Scanln
	// sf = fmt.Scanf
	// s = fmt.Scan
)

func main() {
	// Collatz conjecture => if a number is odd, 3x+1, if it's even, x/2
	var input, check int

	pl("State your initial positive integer")
	sl(&input)
	
	if (input <= 0) || (reflect.TypeOf(input) != reflect.TypeOf(check)) {
		pl("ERR - this is not a valid input! Try again.\nRemember: a positive integer!")
		main()
	}
	pf("your input: %d\n", input)
}