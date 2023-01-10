package main

import (
	"fmt"
)

var (
	pl = fmt.Println
	p = fmt.Print
	pf = fmt.Printf
)

func main() {
	// Var without initial value will fall back to default type value
	// Also, multiple vars can be declared inside a block for simpler codes
	var (
		a string
		b int
		c bool
	)

	pl(a)
	pl(b)
	pl(c)

	/*
	Variable naming rules in Go...
		A variable name must start with a letter or an underscore character (_)
		A variable name cannot start with a digit
		A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
		Variable names are case-sensitive (age, Age and AGE are three different variables)
		There is no limit on the length of the variable name
		A variable name cannot contain spaces
		The variable name cannot be any Go keywords
	*/

	// Constant variable (read-only)
	// Constant are usually named in UPPERCASE, for easier identification - good practice
	const PI = 3.14
	pl(PI)

	// Outputs ===
	// Print
	i, j := "Hello", "Go"
	
	p(i)
	p(j)

	p(i, "\n")
	p(j, "\n")

	p(i, "\n", j)

	p(i, " ", j)

	x, y := 10, 20

	p(x,y)

	// Println is similar to Print, but a whitespace is added between arguments, and a new line at the end

	// Printf
}
