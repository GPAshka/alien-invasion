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

//Generate aliens
func CreateAliens(aliensNumber int) Aliens {
	aliens := make(Aliens, aliensNumber)
	for i := 0; i < aliensNumber; i++ {
		aliens[i] = &Alien{}
	}

	return aliens
}

//Remove aliens with specified Ids
func (aliens Aliens) RemoveAliens(ids []int) {
	for _, id := range ids {
		delete(aliens, id)
	}
}

//Check if all aliens have made needed amount of moves
func (aliens Aliens) EachAlienMadeNeededNumberOfMoves() bool {
	for _, al := range aliens {
		if al.MovesNumber < movesNumberToFinish {
			return false
		}
	}

	return true
}

//Group aliens by their current city. For each city write Id of all aliens in that city
func (aliens Aliens) GroupByCity() map[world.City][]int {
	cityGroup := make(map[world.City][]int)

	for id, al := range aliens {
		cityGroup[al.CurrentCity] = append(cityGroup[al.CurrentCity], id)
	}

	return cityGroup
}
