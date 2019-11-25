package alien

import "alien-invasion/world"

const (
	movesNumberToFinish = 10000
)

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

func (aliens Aliens) EachAlienMadeNeededNumberOfMoves() bool {
	for _, al := range aliens {
		if al.MovesNumber < movesNumberToFinish {
			return false
		}
	}

	return true
}
