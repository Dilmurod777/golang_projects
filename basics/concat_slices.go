package main

import "fmt"

func Concat(slices [][]int) []int {
	var result []int

	for _, s := range slices {
		result = append(result, s...)
	}

	return result
}
//
//func main() {
//	type pair struct {
//		s [][]int
//		r []int
//	}
//
//	test := []pair{
//		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
//		{[][]int{{1, 2}, {3, 4}, {5, 6}}, []int{1, 2, 3, 4, 5, 6}},
//		{[][]int{{1, 2}, {}, {5, 6}}, []int{1, 2, 5, 6}},
//	}
//
//	for _, t := range test {
//		s := t.s
//		r := t.r
//		r2 := Concat(s)
//		fmt.Printf("Test %x = %x | %x\n", s, r, r2)
//	}
//}