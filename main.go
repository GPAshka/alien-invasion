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
	aliensNumberFlag := flag.Int("aliens", 2, "Number of aliens")
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
	g.SetAliens()

	for _, alien := range g.Aliens {
		log.Println(alien.Id, alien.CurrentCity)
	}
}

// Close file and logs error if any
func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Fatalf("error while closing file %v: %v", file.Name(), err)
	}
}
