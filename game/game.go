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

func (g *Game) SetAliens() {
	cities := g.Map.GetCities()

	citiesNumber := len(cities)
	rand.Seed(time.Now().UnixNano())

	for _, alien := range g.Aliens {
		i := rand.Intn(citiesNumber)
		alien.CurrentCity = cities[i]
	}
}
