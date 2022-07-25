package main

import (
	"fmt"
	"sync"
)

func main() {
	// Розкоментуй мене)
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	// Ваша реалізація
	var wg sync.WaitGroup
	for key, val := range n {
		wg.Add(1)
		go func(number int, sl []int) {
			defer wg.Done()
			fmt.Printf("slise %d ", number)
			sum(sl)
			fmt.Println()
		}(key, val)
	}
	wg.Wait()

	//task2
	ch := make(chan int)
	go totalSum(n, ch)
	fmt.Println("sum", <-ch)
}

func sum(sl []int) {
	var res int
	for _, val := range sl {
		res += val
	}
	fmt.Printf("%d", res)
}

//task2
func totalSum(sl [][]int, ch chan int) {
	var temp int
	for _, row := range sl {
		for _, val := range row {
			temp += val
		}
	}
	ch <- temp
}
