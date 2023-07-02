package main

import "fmt"

func main() {
	var n, target int
	if _, err := fmt.Scanf("%d", &n); err != nil {
		fmt.Println(err)
	}

	if _, err := fmt.Scanf("%d", &target); err != nil {
		fmt.Println(err)
	}
	ar := make([]int, n)

	compliments := make(map[int]int)
	for i := 0; i < n; i++ {
		if _, err := fmt.Scanf("%d", &ar[i]); err != nil {
			fmt.Println(err)
		}
	}

	for i := 0; i < n; i++ {
		if elem, ok := compliments[target-ar[i]]; !ok {
			compliments[ar[i]] = i
		} else {
			fmt.Printf("%d at position %d and %d at position %d can sum up to %d\n", ar[i], i, target-ar[i], elem, target)
		}
	}
}
