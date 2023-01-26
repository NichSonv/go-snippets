package main

import (
	"fmt"
	"strings"
	// "log"
	"reflect"
	// "strconv"
)

var (
	pl = fmt.Println
	pf = fmt.Printf
	p = fmt.Print
	sl = fmt.Scanln
	// sf = fmt.Scanf
	// s = fmt.Scan
)

func main() {
	// Collatz conjecture => if a number is odd, 3x+1, if it's even, x/2
	var input, check int

	pl("State your initial positive integer")
	sl(&input)
	
	if (input <= 0) || (reflect.TypeOf(input) != reflect.TypeOf(check)) { // terminal gives out another error when +float is used, doesn't even seem to proccess this bit of code here
		pl("ERR - this is not a valid input! Try again.\nRemember: a positive integer!")
		main()
	}

	results := []int{}
	pf("your input is %d\nBelow is the whole sequence of Collatz's algorythm's results, until pattern repeat...\n", input)
	for i := input; i != 1; { // WIP
		// check if i is even...
		iFloat := float64(i)
		temp := iFloat/2
		temp2 := fmt.Sprintf("%f", temp)
		// actual algorythm
		if tempAr := strings.Split(temp2, "."); tempAr[1] == "000000" { // needs a more solid solution...
			i = int(temp)
			results = append(results, i)

		} else {
			i *= 3; i++
			results = append(results, i)
		}
	}
	pl(results)
	pf("size of list : %d\n", len(results))
	pl("restarting... If you wish to end the program, press CTRL+C.")
	main()
}