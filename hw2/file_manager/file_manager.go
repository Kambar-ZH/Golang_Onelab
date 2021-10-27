package filemanager

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func CleanFile(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	if err = f.Truncate(0); err != nil {
		return err
	}
	if _, err = f.Seek(0, 0); err != nil {
		return err
	}
	return nil
}

func SortFileImports(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	start := -1
	for id, ln := range text {
		if ln == "import (" {
			id++
			start = id
			var imports []string
			for text[id] != ")" && id < len(text) {
				imports = append(imports, text[id])
				id++
			}

			sort.Strings(imports)

			for iter := start; iter < id; iter++ {
				text[iter] = imports[iter-start]
			}
			break
		}
	}
	var b strings.Builder
	for _, ln := range text {
		b.WriteString(ln)
		b.WriteString("\n")
	}
	if err = CleanFile(fileName); err != nil {
		return err
	}
	file, _ = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	file.WriteString(b.String())
	file.Close()
	return nil
}

func TestSortFileImports(fileName string) {
	// test_file.go is unsorted before call of the function
	fmt.Printf("\n---testing fibonacci function\n")
	err := SortFileImports(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\n---file imports has been sorted succesfully\n")
}