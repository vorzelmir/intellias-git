package main

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int

	keys := make(map[int]bool)
	for _, val := range arr {
		if _, ok := keys[val]; !ok {
			keys[val] = false
			result = append(result, val)
		}
	}
}
