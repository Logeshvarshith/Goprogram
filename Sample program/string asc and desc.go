package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"ccc", "aaa", "zzz", "ddd", "mmm"}
	fmt.Println("Given names :", names)
	sort.Strings(names)
	fmt.Println("Names after sorting ascending :", names)
	sort.Sort(sort.Reverse(sort.StringSlice(names)))
	fmt.Println("Names after sorting descending :", names)
}
