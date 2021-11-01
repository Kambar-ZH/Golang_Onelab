package protobuf

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

func GetRows(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}

type Field struct {
	name  string
	_type string
	cost  int
}

func (f Field) String() string {
	return fmt.Sprintf("\n| name: %s | type: %s | cost: %d |", f.name, f._type, f.cost)
}

func TotalCost(fields []Field) int {
	bucket := 0
	counter := 0
	for _, field := range fields {
		if bucket+field.cost >= 8 {
			bucket = field.cost
			counter++
			} else {
			bucket += field.cost
		}
	}
	return counter * 8
}

func GetFields(rows []string) []Field {
	mp := map[string]int{
		"bool":   1,
		"int8":   1,
		"int16":  2,
		"int32":  4,
		"int64":  8,
		"uint8":  1,
		"uint16": 2,
		"uint32": 4,
		"uint64": 8,
		"byte":   1,
		"rune":   4,
		"int":    8,
		"uint":   8,
		"string": 4,
	}

	var fields []Field
	for id, row := range rows {
		if strings.Contains(row, "struct") {
			id++
			for rows[id] != "}" && id < len(rows) {
				field := Field{}
				data := strings.Fields(rows[id])
				fmt.Println(len(data))

				for key, value := range mp {
					if strings.Contains(data[1], key) {
						field.name = data[0]
						field._type = data[1]
						field.cost = value
						break
					}
				}
				fields = append(fields, field)
				id++
			}
			break
		}
	}
	return fields
}

func Perm(a []Field, f func([]Field)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []Field, f func([]Field), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		b := make([]Field, len(a))
		copy(b, a)
		perm(b, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func SortStructs(fileName string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	rows := GetRows(file)

	_fields := GetFields(rows)

	var fieldsPerms [][]Field
	Perm(_fields, func(a []Field) {
		fieldsPerms = append(fieldsPerms, a)
	})

	sort.Slice(fieldsPerms, func(i, j int) bool {
		return TotalCost(fieldsPerms[i]) < TotalCost(fieldsPerms[j])
	})
	top := 3
	for i, fields := range fieldsPerms {
		if i == top {
			break
		}
		fmt.Printf("ID: [%d]\n Fields: [%v] \nCost: [%d]\n", i+1, fields, TotalCost(fields))
	}

	return nil
}
