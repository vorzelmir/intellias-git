package main
import "fmt"

func main() {
	cat := getAnimal("cat")

	f := Farmer {}
	f.setAnimal(cat)
	feed := f.farmer.getTotalFeed(4.4, 9)

	dog := getAnimal("dog")
	f.setAnimal(dog)
	feed += f.farmer.getTotalFeed(11.1, 99)

	cow := getAnimal("cow")
	f.setAnimal(cow)
	feed += f.farmer.getTotalFeed(101.1, 222)
	fmt.Printf("Total feed: %.2f kg\n", feed)
}
