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
			if i_diag != i_cur && i_diag != i_next && i_diag < len(input) {
				i_output++
				output[i_output] = input[i_diag]
			}
			i_output++
		}
	}
	return string(output)
}

func main() {
	type TestData struct {
		s       string
		numRows int
	}

	tests := []TestData{
		{"PAYPALISHIRING", 3},
		{"PAYPALISHIRING", 4},
		{"A", 1},
		{"AB", 1},
	}

	for _, test := range tests {
		fmt.Printf(
			"convert('%s', %d) -> '%s'\n",
			test.s,
			test.numRows,
			convert(test.s, test.numRows),
		)

	}
}
