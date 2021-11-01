package utils

import (
	"sort"
	"strings"
)

type KeyValue struct {
	key   string
	value int
}

func topWords(s string, n int) []string {
	symbols := []string{":", ",", "!", ".", "?"}
	for _, symbol := range symbols {
		s = strings.Replace(s, symbol, "", -1)
	}
	list := strings.Split(s, " ")
	mp := make(map[string]int)
	for _, word := range list {
		mp[word]++
	}
	var keyValues []KeyValue
	for key, value := range mp {
		keyValues = append(keyValues, KeyValue{
			key:   key,
			value: value,
		})
	}
	sort.Slice(keyValues, func(i, j int) bool {
		return keyValues[i].value > keyValues[j].value
	})
	var top []string
	for i := 0; i < n; i++ {
		if i >= len(keyValues) {
			break
		}
		top = append(top, keyValues[i].key)
	}
	sort.Strings(top)
	return top
}
