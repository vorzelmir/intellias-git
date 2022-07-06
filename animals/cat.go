package main

import "fmt"

const catFood float32 = 7

type Cat struct {
	title      string
	weight     float32
	number int
	amountFeed float32
}

func (c *Cat) newCat(weight float32, amount int) *Cat {
	c = &Cat{}
	c.title = "Tom"
	c.weight = weight
	c.number = amount
	c.amountFeed = c.weight * float32(c.number) *  catFood
	return c
}

func (c *Cat) getTotalFeed(weight float32, amount int) float32 {
	c = c.newCat(weight, amount)
	fmt.Printf("Cat name: %s, cat weight: %.2f, cats number: %d, cat feed: %.2f\n",
	c.title, c.weight, c.number, c.amountFeed)
	return c.amountFeed
}
