package itoa

import (
	"example/reverse"
	"fmt"
)

func Itoa(n int) string {
	var b []byte
	pos := true
	if n == 0 {
		return "0"
	}
	if n < 0 {
		pos = false
		n = -n
	}
	for n != 0 {
		b = append(b, byte(n%10+'0'))
		n /= 10
	}
	if !pos {
		b = append(b, '-')
	}
	res := string(b)
	return reverse.Reverse(res)
}

type Test struct {
	input    int
	expected string
}

func ItoaTest() {
	tests := []Test{
		{
			input:    228,
			expected: "228",
		},
		{
			input:    -1223,
			expected: "-1223",
		},
		{
			input:    0,
			expected: "0",
		},
		{
			input:    26,
			expected: "26",
		},
	}
	fmt.Printf("\n---testing itoa function\n")
	for id, test := range tests {
		actual := Itoa(test.input)
		if actual != test.expected {
			fmt.Printf("%d) incorrect result\nexpected: %v\nfound: %v\n", id+1, test.expected, actual)
			return
		}
		fmt.Printf("%d) test passed\nexpected: %v\nfound: %v\n", id+1, test.expected, actual)
	}
	fmt.Printf("---passed all testcases\n")
}
