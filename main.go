package main

import (
	"alien-invasion/game"
	"flag"
	"log"
	"os"
)

func main() {
	//read parameters from the command line
	inputFileFlag := flag.String("input", "map.txt", "txt file which contains information about world map")
	aliensNumberFlag := flag.Int("aliens", 20, "Number of aliens")
	flag.Parse()

	//open file with input data
	mapFile, err := os.Open(*inputFileFlag)
	if err != nil {
		log.Fatalln("error while opening input map file: ", err)
		return
	}
	defer closeFile(mapFile)

	//create new game
	g := game.NewGame(mapFile, *aliensNumberFlag)

	//place aliens
	g.SetAliens()

	//main game loop: move aliens and fight them until game ends
	for g.Continue() {
		g.MoveAliens()
		g.FightAliens()
	}

	//output rest of the world
	g.Map.Print(os.Stdout)
}

// Close file and log error if any
func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Fatalf("error while closing file %v: %v", file.Name(), err)
	}
}
