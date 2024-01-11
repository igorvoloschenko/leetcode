package main

import "fmt"

func convert(s string, numRows int) string {
	var i_cur, i_next, i_diag int

	input := []rune(s)
	output := make([]rune, len(input))
	offset := 2
	if numRows == 1 {
		offset = 1
	}

	i_output := 0
	for n := 0; n < numRows; n++ {
		i_next = n
		i_diag = n
		for i_next < len(input) && i_output < len(output) {
			i_cur = i_next
			output[i_output] = input[i_cur]
			i_next += (numRows * 2) - offset
			i_diag = i_next - (n * 2)
			if i_diag != i_cur && i_diag != i_next && i_diag < len(output) {
				i_output++
				output[i_output] = input[i_diag]
			}
			i_output++
		}
	}
	return string(output)
}

func main() {
	input := "PAYPALISHIRING"
	numRows := 4

	fmt.Printf(
		"convert(%s, %d) -> %s\n",
		input,
		numRows,
		convert(input, numRows),
	)
}
