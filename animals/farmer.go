package main


type Farmer struct {
	farmer Farm
}

func (f *Farmer) setAnimal(animal Farm) {
	f.farmer = animal
}
