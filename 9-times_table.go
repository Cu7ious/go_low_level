package main

import (
	"fmt"
)


// - times_table - prints the 9 times table, starting with 0
// - Return: void
func main() {
	const NEWLINE = "\n"
	const SPACE   = " "
	const COMMA   = ","
	var tens, ones, result int

	for tens = 0; tens <= 9; tens++ {
		for ones = 0; ones <= 9; ones++ {
			result = tens * ones

			if result < 10 {
				if ones != 0 {
					fmt.Printf(SPACE)
				}

				fmt.Printf("%d", result)
			} else {
				fmt.Printf("%d", int(result / 10))
				fmt.Printf("%d", int(result % 10))
			}

			if ones != 9 {
				fmt.Printf(COMMA)
				fmt.Printf(SPACE)
			}
		}
		fmt.Printf(NEWLINE)
	}
}
