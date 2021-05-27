package main

import (
	"fmt"
	"sort"
)

func myfunc(a []int) int {
	sort.Ints(a)
	fmt.Println("Ints:   ", a)
	min := a[12-1]
	return min
}

func main() {
	values := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17, 
		89, 100, 77,55
	}
	var res int
	res = myfunc(values)
	fmt.Println("12th lowest number in a list of integer is:   ", res)

}
