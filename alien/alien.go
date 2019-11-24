package alien

import "alien-invasion/world"

type Alien struct {
	Id          int
	MovesNumber int
	CurrentCity *world.City
}

type Aliens []*Alien

func CreateAliens(aliensNumber int) Aliens {
	aliens := make(Aliens, aliensNumber, aliensNumber)
	for i := 0; i < aliensNumber; i++ {
		alien := Alien{Id: i}
		aliens = append(aliens, &alien)
	}

	return aliens
}
