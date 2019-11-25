package game

import (
	"alien-invasion/alien"
	"alien-invasion/world"
	"fmt"
	"io"
	"math/rand"
	"time"
)

const (
	aliensInCityToFight = 2
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

	for _, al := range g.Aliens {
		i := rand.Intn(citiesNumber)
		al.CurrentCity = cities[i]
	}
}

//Move all aliens to the next city
func (g *Game) MoveAliens() {
	//move each alien to the next city from it's current location using the Map
	for _, al := range g.Aliens {
		al.CurrentCity = g.Map.GetNextDestinationFromCity(al.CurrentCity)
		al.MovesNumber++
	}
}

//Check if there are cities with at least aliensInCityToFight aliens and make fight for each case
func (g *Game) FightAliens() {
	//group aliens by their current city
	cityGroup := make(map[world.City][]int)

	for id, al := range g.Aliens {
		cityGroup[al.CurrentCity] = append(cityGroup[al.CurrentCity], id)
	}

	//fight aliens if there are more than aliensInCityToFight aliens in the particular city
	for city, ids := range cityGroup {
		if len(ids) >= aliensInCityToFight {
			//print destroy message
			fmt.Printf("%s has​ ​been​ ​destroyed​ ​by​ ​alien​s %v!\n", city, ids)

			//remove all destroyed aliens
			g.Aliens.RemoveAliens(ids)

			//remove destroyed city from the Map
			g.Map.RemoveCity(city)
		}
	}
}

//Determine whether to continue game
func (g *Game) Continue() bool {
	//continue until all aliens are destroyed or​ ​each​ ​alien​ ​has​ ​moved​ needed amount
	return len(g.Aliens) > 0 && !g.Aliens.EachAlienMadeNeededNumberOfMoves()
}
