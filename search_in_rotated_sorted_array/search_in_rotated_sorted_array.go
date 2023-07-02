package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var target int

	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &target)

	ar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &ar[i])
	}

	res := -1

	l, r := 0, n-1
	for l < r {
		mid := (l + r) / 2
		if ar[mid] == target {
			res = mid
			break
		} else if ar[mid] > ar[r] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	fmt.Printf("%d is the index of the value", res)
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
