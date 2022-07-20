package main

type Farm interface {
	getTotalFeed(weight float32, number int) float32
}

func getAnimal(animal string) Farm {
	switch animal {
	case "cat":
		return &Cat{}
	case "dog":
		return &Dog{}
	case "cow":
		return &Cow{}
	}
	return nil
}
