package reverse

import "fmt"

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; {
		r[i], r[j] = r[j], r[i]
		i++
		j--
	}
	return string(r)
}

type Test struct {
	input    string
	expected string
}

func ReverseTest() {
	tests := []Test{
		{
			input:    "Hello",
			expected: "olleH",
		},
		{
			input:    "Алихан",
			expected: "нахилА",
		},
		{
			input:    "Палинdrom",
			expected: "mordнилаП",
		},
		{
			input:    "шалаш",
			expected: "шалаш",
		},
	}
	fmt.Printf("\n---testing reverse function\n")
	for id, test := range tests {
		actual := Reverse(test.input)
		if actual != test.expected {
			fmt.Printf("%d) incorrect result\nexpected: %v\nfound: %v\n", id+1, test.expected, actual)
			return
		}
		fmt.Printf("%d) test passed\nexpected: %v\nfound: %v\n", id+1, test.expected, actual)
	}
	fmt.Printf("---passed all testcases\n")
}