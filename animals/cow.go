package main

import "fmt"

const cowFeed float32 = 25

type Cow struct {
	title      string
	weight     float32
	number int
	amountFeed float32
}

func (c *Cow) newCow(weight float32, amount int) *Cow {
	c = &Cow{}
	c.title = "Milk"
	c.weight = weight
	c.number = amount
	c.amountFeed = cowFeed * float32(c.number) * c.weight
	return c
}


func (c *Cow) getTotalFeed(weight float32, number int) float32 {
	c = c.newCow(weight, number)
	fmt.Printf("Cow name: %s, cow weight: %.2f, cows number: %d, cow feed: %.2f\n",
	c.title, c.weight, c.number, c.amountFeed)
	return c.amountFeed
}
