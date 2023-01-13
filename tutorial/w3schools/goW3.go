package main

import (
	"fmt"
)

var (
	pl = fmt.Println
	p = fmt.Print
	pf = fmt.Printf
)

func test(param1 int, param2 string, param3 bool) {
	pl("the results are...", param1, param2, param3)
}

func sum(x int, y int) int {
	return x + y
}

func naked(a string) (name string) {
	name = a
	return
}

func loop(x int) string {
	if x < 10 && x > 0 {
		pl(x)
		x++
		loop(x)
	} else {
		pl("== error, invalid int ==")
	}
	return "done"
}

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

	// Printf - %v - value; %#v - value in Go-syntax format ; %T - Type ; %% - % ; and several others for use with int, string... (see tut link for ref)
	pftxt1 := "Hello"
	pftxt2 := 15

	pf("pftxt1 has value: %#v and type: %T\n", pftxt1, pftxt1)
	pf("pftxt2 has value: %v and type: %T\n", pftxt2, pftxt2)

	// Data types =======
	// int and uint - signed (both + and -) and unsigned (only +)
	// int (auto 32 or 64), int8, int16, int32 and int64 (bitsize)
	// 8bits - 2^8 max size ; 16bits - 2^16 ; 32bits - 2^32 ; 64bits - 2^64

	// Arrays
	var array_ex = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	var array_ex2 = [...]float32{1.4, 2.3}
	array_ex3 := [5]int{1:2, 2:4}
	pl(array_ex)
	pl(array_ex2)
	pl(array_ex3)
	pl(len(array_ex2))

	// Slices
	mySlice := []int{}
	fromArray := array_ex[0:8]

	pl(len(mySlice), cap(mySlice))
	pf("fromArray = %v\n", fromArray)
	pf("fromArray = %d\n", len(fromArray))
	pf("fromArray = %d\n", cap(fromArray))

	slice2 := make([]int, 5, 9)
	pl(slice2)
	pf("lenght = %d\n", len(slice2))
	pf("cap = %d\n", cap(slice2))

	slice2 = append(slice2, 20, 22, 26, 55, 43)
	pf("slice2 = %#v\n", slice2)
	pf("lenght = %d\n", len(slice2))
	pf("cap = %d\n", cap(slice2))

	slice3 := append(fromArray, slice2...)
	pl("fromArray =", fromArray, "\n", "slice2 =", slice2)
	pl(slice3)
	pf("lenght = %d\n", len(slice3))
	pf("cap = %d\n", cap(slice3))

	cutSlice := slice2[:len(slice2)-4]
	shortSlice := make([]int, len(cutSlice))
	copy(shortSlice, cutSlice)
	pf("slice2 = %v\n%d\n%d\n", slice2, len(slice2), cap(slice2))
	pf("cutSlice = %v\n%d\n%d\n", cutSlice, len(cutSlice), cap(cutSlice))
	pf("shortSlice = %v\n%d\n%d\n", shortSlice, len(shortSlice), cap(shortSlice))

	slice4 := make([]int, len(slice2)-4)
	copy(slice4, slice2)
	pf("slice4 = %v\n%d\n%d\n", slice4, len(slice4), cap(slice4))

	for x := 0; x < 7 ; x++ {
		if x < 2 {
			continue
		}
		if x == 6 {
			break
		}
		pl(x)
	}

	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
	fmt.Printf("%v\t%v\n", idx, val)
	}

	test(3, "very", true)

	pl(sum(5, 9))

	pl(naked("LOL"))

	pl(loop(1))

	// Struct
	type Person struct {
		name string
		age int
		job string
		salary float64
	}

	var pers1 Person
	pers1.name = "Jeff"
	pers1.age = 32
	pers1.job = "Mechanic"
	pers1.salary = 3245.5

	pl(pers1)

}
