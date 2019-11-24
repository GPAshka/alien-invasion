package game

import (
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

/*type City struct {
	Name string
	Roads []*Road
}*/

type City string

type Road struct {
	Direction   Direction
	Destination City
}

type Map map[City][]*Road

func CreateMap(reader io.Reader) *Map {
	//TODO read map from the file

	m := make(Map, 0)
	return &m
}

func (m *Map) addCityFromInputRecord(record string) {
	cityWithRoads := strings.Split(record, " ")

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