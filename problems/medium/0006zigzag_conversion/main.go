package main

import "fmt"

// Zigzag Conversion
// Complexity - Time: O(n) - Space: O(n)
// Pattern: Simulation
// Go through each character and put it in the right row bucket.
// Direction flips at top and bottom rows, like a pendulum.

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]string, numRows)
	currentRow := 0
	goingDown := true

	for _, char := range s {
		rows[currentRow] += string(char)

		if currentRow == numRows-1 {
			goingDown = false
		} else if currentRow == 0 {
			goingDown = true
		}

		if goingDown {
			currentRow++
		} else {
			currentRow--
		}
	}

	result := ""
	for _, row := range rows {
		result += row
	}
	return result
}

// -----------------------------------------------------------------

func main() {
	s := "PAYPALISHIRING"
	result := convert(s, 3)
	fmt.Println(result) // PAHNAPLSIIGYIR
}
