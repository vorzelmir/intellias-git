package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"
	var result string

	//convert string to slise of integer
	var strSlise []string = strings.Split(input, " ")
	var intSlise []int
	for _, val := range strSlise {
		if res, err := strconv.Atoi(val); err == nil {
			intSlise = append(intSlise, res)
		}
	}

	//find out the max and min value in the slise
	max := intSlise[0]
	min := intSlise[len(strSlise)-1]
	for _, val := range intSlise {
		switch {
		case val > max:
			max = val
		case val < min:
			min = val
		}
	}

	//create slise of max and min
	var strArr []string = []string{
		strconv.Itoa(max),
		strconv.Itoa(min),
	}

	//convert slise to string and store in result
	result = strings.Join(strArr, " ")
	fmt.Printf("Result: %s\n", result)
}
