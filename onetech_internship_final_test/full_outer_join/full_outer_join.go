package full_outer_join

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strings"
)

type File struct {
	rows map[string]bool
}

func GetFileWithRows(fPath string) (*File, error) {
	f, err := os.OpenFile(fPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	rows := make(map[string]bool, 0)
	for sc.Scan() {
		rows[sc.Text()] = true
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return nil, err
	}
	return &File{rows: rows}, nil
}

func SortedText(rows []string) string {
	sort.Strings(rows)
	var sb strings.Builder
	for idx, str := range rows {
		sb.WriteString(str)
		if idx == len(rows) - 1 {
			continue
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func UniqueRowsToSlice(f1 *File, f2 *File) []string {
	mp := make([]string, 0)
	for key := range f1.rows {
		if _, found := f2.rows[key]; !found {
			mp = append(mp, key)
		}
	}
	for key := range f2.rows {
		if _, found := f1.rows[key]; !found {
			mp = append(mp, key)
		}
	}
	return mp
}

func FullOuterJoin(f1Path, f2Path, resultPath string) {
	f1, err := GetFileWithRows(f1Path)
	if err != nil {
		panic(err)
	}
	f2, err := GetFileWithRows(f2Path)
	if err != nil {
		panic(err)
	}
	uniqueRows := UniqueRowsToSlice(f1, f2)

	sortedText := SortedText(uniqueRows)
	err = os.WriteFile(resultPath, []byte(sortedText), 0644)
	if err != nil {
		panic(err)
	}
}
