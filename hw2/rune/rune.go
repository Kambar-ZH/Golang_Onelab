package rune

import "fmt"

var (
	errIndex = fmt.Errorf("index out of range")
	errNilIndex = fmt.Errorf("index is nil")
	errNilString = fmt.Errorf("string is nil")
)

func RuneByIndex(s *string, i *int) (rune, error) {
	runes := []rune(*s)
	if i == nil {
		return 0, errNilIndex
	}
	if s == nil {
		return 0, errNilString
	}
	if *i < 0 || *i >= len(runes) {
		return 0, errIndex
	}
	return runes[*i], nil
}

func TestRuneByIndex(str string) {
	fmt.Printf("\n---testing runeByIndex function\n")
	runes := []rune(str)
	for i := -1; i <= len(runes); i++ {
		r, err := RuneByIndex(&str, &i)
		fmt.Printf("%d %c %v\n", i, r, err)
	}
}