package main

import "fmt"

func main() {
	const (
		applePrice   = 5.99
		pearPrice    = 7
		amountMoney  = 23
		totalApples  = 9
		totalPears   = 8
		numberFruits = 2
	)

	fmt.Printf("How much will be cost %d apples and %d pears?\n", totalApples, totalPears)
	totalPrice := applePrice*totalApples + pearPrice*totalPears
	fmt.Printf("Total price is: %.2f\n", totalPrice)

	fmt.Println("How many pears or apples can we buy?")
	numbersApple := amountMoney / applePrice
	numbersPears := amountMoney / pearPrice
	fmt.Printf("We can buy %d pears or %d apples\n", numbersPears, int(numbersApple))

	fmt.Printf("Can we buy %[1]d apples and %[1]d pears?\n", numberFruits)
	totalPrice = numberFruits * (applePrice + pearPrice)
	result := amountMoney > totalPrice
	fmt.Printf("We can buy it's %t\n", result)
}
