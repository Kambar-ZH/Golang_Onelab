package fibonacci

import "fmt"

func Fibonacci() func() int {
	a, b := 1, 1
	return func() int {
		defer func() {
			a, b = b, a+b
		}()
		return a
	}
}

func TestFibonacci() {
	fmt.Printf("\n---testing fibonacci function\n")
	F := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d = %d\n", i, F())
	}
}
