package atoi

import (
	"example/reverse"
	"fmt"
)

var (
	errEmpty = fmt.Errorf("cannot convert empty string to int")
	errParse = fmt.Errorf("string cannot be parsed to int")
)

func Atoi(s string) (int, error) {
	if len(s) == 0 {
		return 0, errEmpty
	}
	ans := 0
	pos := true
	upTo := len(s)
	if s[0] == '-' {
		pos = false
		upTo--
	}
	s = reverse.Reverse(s)
	for iterator, power := 0, 1; iterator < upTo; iterator, power = iterator + 1, power * 10 {
		if s[iterator] < '0' || s[iterator] > '9' {
			return 0, errParse
		}
		ans += int(s[iterator] - '0') * power
	}
	if !pos {
		ans = -ans
	}
	return ans, nil
}

type Test struct {
	input string
	expected Result
} 

type Result struct {
	value int
	err error
}

func (r Result) String() string {
	return fmt.Sprintf("value: {%d} err: {%v}", r.value, r.err)
}

func AtoiTest() {
	tests := []Test {
		{
			input: "228",
			expected: Result {
				value: 228,
				err: nil,
			},
		},
		{
			input: "-228",
			expected: Result {
				value: -228,
				err: nil,
			},
		},
		{
			input: "",
			expected: Result {
				value: 0,
				err: errEmpty,
			},
		},
		{
			input: "228a",
			expected: Result {
				value: 0,
				err: errParse,
			},
		},
	}
	fmt.Printf("\n---testing atoi function\n")
	for id, test := range tests {
		actual, err := Atoi(test.input)
		if actual != test.expected.value || err != test.expected.err {
			fmt.Printf("%d) incorrect result\nexpected: %v\nfound: %v\n", id + 1, test.expected, Result{actual, err})
			return
		}
		fmt.Printf("%d) test passed\nexpected: %v\nfound: %v\n", id + 1, test.expected, Result{actual, err})
	}
	fmt.Printf("---passed all testcases\n")
}