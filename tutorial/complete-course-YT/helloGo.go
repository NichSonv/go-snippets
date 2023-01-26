package main

// Tut video link https://www.youtube.com/watch?v=YzLrWHZa-Kc
// Complemental learning...
// https://www.geeksforgeeks.org/how-to-take-input-from-the-user-in-golang/ -- .Scanln

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
	"time"
)

var (
	pl = fmt.Println	
	pf = fmt.Printf
	// p = fmt.Print
)

func main() {
	pl("Hello Go!")
	pl("What's your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}
	// About variables... [var name type] is the basic, and if it starts with capital letter, can be exported(!)
	// Accepted formats for declaration...
	// var vName string = "Derek"
	// var v1, v2 = 1.2, 3.4
	// var v3 = "hello"
	pl("About variables =======")
	v4 := 5.2
	pl(v4)
	v4 = 3
	pl(v4)
	// Data types -> int, float64, bool, string, rune
	pl("About data types =======")
	pl(reflect.TypeOf(25), reflect.TypeOf(2.3), reflect.TypeOf(true), reflect.TypeOf("Hello"), reflect.TypeOf('😬'))

	// Casting
	pl("Now about casting =======")
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2)
	cV3 := "50000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV3, reflect.TypeOf(cV3), cV4, err, reflect.TypeOf(cV4))
	cV5 := strconv.Itoa(cV4)
	pl(cV5, reflect.TypeOf(cV5))
	cV6 := "3.14"
	if cV7, err := strconv.ParseFloat(cV6, 64); err == nil {
		pl(cV7)
	} else {
		log.Fatal(err)
	}
	cV8 := fmt.Sprintf("%f", 3.14)
	pl(cV8)

	pl("About conditionals ======")
	pl("What's your age?")
	var age int
	fmt.Scanln(&age)
	if (age >= 1) && (age <= 15) {
		pl("Yougster")
	} else if (age > 15) && (age <= 30) {
		pl("Young adult")
	} else if (age > 30) && (age <= 55) {
		pl("Adult")
	} else if age > 55 {
		pl("Golden years")
	} else {
		pl("Is it possible?")
	}
	pl("File:", name, "age", age)

	// String methods

	sv1 := "A word"
	replacer := strings.NewReplacer("a", "Another")
	sv2 := replacer.Replace(sv1)
	pl(sv2)
	// len(string)
	// strings.Contains(sv2, "Another")
	// strings.Index(sv2, "o")
	pl("Replace :", strings.Replace(sv2, "o", "0", -1))
	sv3 := "\nSome words\n"
	sv3 = strings.TrimSpace(sv3)
	pl(sv3)
	// strings.Split
	// strings.ToLower
	// strings.ToUpper
	// strings.HasPrefix("tacocat", "taco")
	// strings.HasSuffix("tacocat", "cat")

	// Runes
	rStr := "abcdefg"
	pl("Rune Count :", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		pf("%d : %#U : %c\n", i, runeVal, runeVal)
	}

	// Time (very neat)
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())
}