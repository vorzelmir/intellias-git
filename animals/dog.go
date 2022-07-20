package main

import (
	"fmt"
)

const dogFeed float32 = 10

type Dog struct {
	title      string
	weight     float32
	number 		 int
	amountFeed float32
}

func (c *Dog) newDog(weight float32, amount int) *Dog {
	c = &Dog{}
	c.title = "Spike"
	c.weight = weight
	c.number = amount
	c.amountFeed = dogFeed * float32(c.number) * c.weight
	return c
}


func (d *Dog) getTotalFeed(weight float32, amount int) float32 {
	  d = d.newDog(weight, amount)
		fmt.Printf("Dog name: %s, dog weight: %.2f, dogs number: %d, feed: %.2f\n",
		d.title, d.weight, d.number, d.amountFeed)
	return d.amountFeed
}
