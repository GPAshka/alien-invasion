package world

import (
	"bufio"
	"io"
	"strings"
)

type Direction string

const (
	North Direction = "north"
	South           = "south"
	East            = "east"
	West            = "west"
)

type City string

type Road struct {
	Direction   Direction
	Destination City
}

type Map map[City][]*Road

func CreateMap(reader io.Reader) *Map {
	m := make(Map, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		m.addCityFromInputRecord(line)
	}

	return &m
}

func (m *Map) GetCities() []City {
	cities := make([]City, len(*m))

	i := 0
	for city := range *m {
		cities[i] = city
		i++
	}

	return cities
}

func (m *Map) addCityFromInputRecord(inputRecord string) {
	cityWithRoads := strings.Split(inputRecord, " ")

	city := City(cityWithRoads[0])
	roads := make([]*Road, 0)

	//skip first item in the cityWithRoads slice because it is name of the city itself
	for i := 1; i < len(cityWithRoads); i++ {
		//split each road
		roadParts := strings.Split(cityWithRoads[i], "=")

		//create new Road
		road := &Road{Direction: Direction(roadParts[0]), Destination: City(roadParts[1])}
		roads = append(roads, road)
	}

	(*m)[city] = roads
}
