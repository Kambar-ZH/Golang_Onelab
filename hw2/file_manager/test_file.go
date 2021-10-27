package filemanager

import (
	"sort"
	"math"
	"fmt"
)

func Do() {
	arr := []int {1, 2, 3}
	sort.Ints(arr)
	for i := range arr {
		arr[i] = int(math.Pow(float64(arr[i]), 2))
	}
	fmt.Println(arr)
}