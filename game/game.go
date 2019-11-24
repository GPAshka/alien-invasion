package game

import "alien-invasion/alien"

type Game struct {
	alien.Aliens
	Map
}

func NewGame(aliensNumber int) *Game {
	//create aliens
	aliens := alien.CreateAliens(aliensNumber)

	game := &Game{Aliens: aliens}
	return game
}

func (game *Game) CreateAliens() {

}
