package main

import (
	"hw3/protobuf"
	"log"
)

func main() {
	err := protobuf.SortStructs("./protobuf/test_file.go")
	if err != nil {
		log.Fatal(err)
	}
}