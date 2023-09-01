package main

import (
	"fmt"
)

func main() {
	x := 2

	// example of a simple if statment
	if x > 0 {
		fmt.Println("true")
	}
	fmt.Println(x)

	//example of chain conditions.  not = !, and = &&, or = ||
	val := false && true
	fmt.Println(val)

	//example of if else statment
	age := 20

	if age > 21 {
		fmt.Println("can drink")
	} else if age > 18 {
		fmt.Println("can enter, but not drink")
	} else {
		fmt.Println("too young, cant enter")
	}

	//converting previous if else to a switch statment
	switch age {
	case 20:
		fmt.Println("can enter, but not drink")
	case 21:
		fmt.Println("can drink")
	case 19:
		fmt.Println("too young, cant enter")
	}
}
