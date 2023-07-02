package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)

	ar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &ar[i])
	}

	l, r := 0, n-1
	for l < r {
		mid := (l + r) / 2
		if ar[mid] > ar[r] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	fmt.Printf("%d is the minimum in the array", ar[l])
}

func min(values []int) int {
	min := math.MaxInt
	n := len(values)
	for i := 0; i < n; i++ {
		if values[i] < min {
			min = values[i]
		}
	}

	return min
}
