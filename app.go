package main

import "fmt"

type Show struct {
	top bool
}

func main() {
	newMap := make(map[string]Show)
	newMap["top"] = Show{top: true}

	var topp bool

	newShow, ok := newMap["toppp"]

	fmt.Println(newShow)     // {show}
	fmt.Println(newShow.top) // false
	fmt.Println(ok)          // false
	fmt.Println(topp)        // false

	newShow, ok = newMap["top"]

	fmt.Println(newShow)     // {true}
	fmt.Println(newShow.top) // true
	fmt.Println(ok)          // true
	fmt.Println(topp)        // false
}
