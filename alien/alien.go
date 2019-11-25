package alien

import "alien-invasion/world"

type Alien struct {
	MovesNumber int
	CurrentCity world.City
}

type Aliens map[int]*Alien

func CreateAliens(aliensNumber int) Aliens {
	aliens := make(Aliens, aliensNumber)
	for i := 0; i < aliensNumber; i++ {
		aliens[i] = &Alien{}
	}

	return aliens
}

func (aliens Aliens) RemoveAliens(ids []int) {
	for _, id := range ids {
		delete(aliens, id)
	}
}
