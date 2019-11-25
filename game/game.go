package game

import (
	"alien-invasion/alien"
	"alien-invasion/world"
	"io"
	"math/rand"
	"time"
)

type Game struct {
	alien.Aliens
	*world.Map
}

func NewGame(mapReader io.Reader, aliensNumber int) *Game {
	//create aliens
	aliens := alien.CreateAliens(aliensNumber)

	//create Map
	m := world.CreateMap(mapReader)

	game := &Game{Aliens: aliens, Map: m}
	return game
}

//Randomly set aliens to the cities
func (g *Game) SetAliens() {
	cities := g.Map.GetCities()

	citiesNumber := len(cities)
	rand.Seed(time.Now().UnixNano())

	for _, alien := range g.Aliens {
		i := rand.Intn(citiesNumber)
		alien.CurrentCity = cities[i]
	}
}

func (g *Game) MoveAliens() {
	//move each alien to the next city from it's current location using the Map
	for _, al := range g.Aliens {
		al.CurrentCity = g.Map.GetNextDestinationFromCity(al.CurrentCity)
	}
}

func (g *Game) FightAliens() {

}

func (g *Game) Continue() bool {
	return false
}
