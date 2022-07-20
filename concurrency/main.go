package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sum(sl [][]int) {
	for key, row := range sl {
		go func(arr []int, k int) {
			var sum int
			wg.Add(1)
			for _, item := range arr {
				sum += item
			}
			fmt.Println("sum ", k+1, ":", sum)
			defer wg.Done()
		}(row, key)
	}
}

func main() {
	n := [][]int{
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 6, 7, 8},
		{3, 5, 7, 8},
	}

	sum(n)
	wg.Wait()
}

//output the next
// sum  1 : 15
// sum  2 : 35
// sum  3 : 23
