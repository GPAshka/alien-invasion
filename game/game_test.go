package game

import (
	"alien-invasion/alien"
	"alien-invasion/world"
	"testing"
)

func TestGame_Continue(t *testing.T) {
	var testCase = []struct {
		game           *Game
		expectedResult bool
	}{
		{
			//not enough moves - continue
			game: &Game{alien.Aliens{
				0: &alien.Alien{MovesNumber: 1}},
				&world.Map{}},
			expectedResult: true,
		},
		{
			//enough moves - stop
			game: &Game{alien.Aliens{
				0: &alien.Alien{MovesNumber: 10001}},
				&world.Map{}},
			expectedResult: false,
		},
		{
			//not all aliens enough moves - continue
			game: &Game{alien.Aliens{
				0: &alien.Alien{MovesNumber: 10001},
				1: &alien.Alien{MovesNumber: 1000}},
				&world.Map{}},
			expectedResult: true,
		},
		{
			//all aliens enough moves - stop
			game: &Game{alien.Aliens{
				0: &alien.Alien{MovesNumber: 10001},
				1: &alien.Alien{MovesNumber: 10002}},
				&world.Map{}},
			expectedResult: false,
		},
		{
			//all aliens enough moves - stop
			game: &Game{alien.Aliens{
				0: &alien.Alien{MovesNumber: 10001},
				1: &alien.Alien{MovesNumber: 10002}},
				&world.Map{}},
			expectedResult: false,
		},
		{
			//no more aliens - stop
			game: &Game{alien.Aliens{},
				&world.Map{}},
			expectedResult: false,
		},
	}

	for i, input := range testCase {
		continueGame := input.game.Continue()
		if continueGame != input.expectedResult {
			t.Errorf("FAIL: input %v, expected result %v, but got %v", i, input.expectedResult, continueGame)
		}
	}
}
