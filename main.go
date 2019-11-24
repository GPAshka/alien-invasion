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
	aliensNumberFlag := flag.Int("aliens", 10, "Number of aliens")
	flag.Parse()

	//open file with input data
	mapFile, err := os.Open(*inputFileFlag)
	if err != nil {
		log.Fatalln("error while opening input CSV file: ", err)
		return
	}
	defer closeFile(mapFile)

	//create new game
	g := game.NewGame(mapFile, *aliensNumberFlag)
}

// Close file and logs error if any
func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Fatalf("error while closing file %v: %v", file.Name(), err)
	}
}
