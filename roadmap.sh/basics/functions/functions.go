package main

import (
	"fmt"
)

// simple function that returns an integer
func buy(amount int) int {
	if amount >= 300 {
		if amount != 300 {
			change := amount - 300
			fmt.Printf("Oz is only 300, %d is your change", change)
			return change
		} else {
			change := 0
			fmt.Printf("Oz is only 300, %d is your change", change)
			return change
		}
	} else if amount >= 150 {
		if amount != 150 {
			change := amount - 150
			fmt.Printf("Half Oz is only 150, %d is your change", change)
			return change
		} else {
			change := 0
			fmt.Printf("Half Oz is only 150, so your change is %d", change)
			return change
		}
	} else {
		fmt.Println("Can't afford anything")
		return amount
	}
}
func main() {
	money := 310
	buy(money)
}
