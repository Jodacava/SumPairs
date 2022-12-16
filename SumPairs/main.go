package main

import "fmt"

func main() {
	list := []int{1, 9, 5, 0, 20, -4, 12, 16, 7}
	SumPairs(list, 12)
}

func SumPairs(list []int, sum int) {
	for i, np := range list {
		for j := i + 1; j < len(list); j++ {
			if np+list[j] == sum {
				fmt.Printf("%d, %d\n", np, list[j])
			}
		}
	}
}
