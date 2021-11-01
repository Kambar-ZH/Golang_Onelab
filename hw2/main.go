package main

import (
	"example/atoi"
	"example/fibonacci"
	"example/itoa"
	"example/reverse"
	"example/rune"
	filemanager "example/file_manager"
)

func main() {
	// see terminal log
	atoi.AtoiTest()
	itoa.ItoaTest()
	reverse.ReverseTest()
	filemanager.TestSortFileImports("./file_manager/test_file.go")
	fibonacci.TestFibonacci()
	rune.TestRuneByIndex("")
}
