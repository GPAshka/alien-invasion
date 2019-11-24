package alien

import "alien-invasion/world"

type Alien struct {
	Id          int
	MovesNumber int
	CurrentCity world.City
}

type Aliens []*Alien

func CreateAliens(aliensNumber int) Aliens {
	aliens := make(Aliens, aliensNumber)
	for i := 0; i < aliensNumber; i++ {
		aliens[i] = &Alien{Id: i}
	}

	return aliens
}
