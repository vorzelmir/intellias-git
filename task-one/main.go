package main

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int

	//values bool but does not matter, needs keys only
	var keys = make(map[int]bool)

	//keys are unique they can't repeat
	for _, val := range arr {
		keys[val] = false
	}

	//add to the result slise unique keys
	for key, _ := range keys {
		result = append(result, key)
	}
}
