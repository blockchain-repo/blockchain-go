package common

import (
	"fmt"
	"testing"
)

func Test_Sort(t *testing.T) {
	count := make(map[string]int, 0)
	count["a"] = 1
	count["c"] = 3
	count["b"] = 2
	sorted_count := SortMapByValue(count)
	fmt.Println(sorted_count[len(sorted_count)-1])
}
