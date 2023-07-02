package main

import (
	"fmt"
	"math"

	"golang.org/x/exp/constraints"
)

func main() {
	var n int

	fmt.Scanf("%d", &n)

	ar := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &ar[i])
	}

	var left, right int

	maxi := math.MinInt
	for right < len(ar) {
		if ar[left] < ar[right] {
			profit := ar[right] - ar[left]
			maxi = max(maxi, profit)
		} else {
			left = right
		}
		right++
	}

	fmt.Printf("%d is the max profit.", maxi)
}

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}
