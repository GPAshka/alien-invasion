package world

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strings"
	"time"
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

type Roads []*Road

type Map map[City]Roads

func (road *Road) String() string {
	return fmt.Sprintf("%s=%s", road.Direction, road.Destination)
}

//String representation of the city roads array
func (roads Roads) String() string {
	str := make([]string, len(roads))
	i := 0
	for _, road := range roads {
		str[i] = road.String()
		i++
	}

	return strings.Join(str, " ")
}

//Get random destination city from all possible direction, which lead out from city
func (m *Map) GetNextDestinationFromCity(city City) City {
	roads := (*m)[city]

	//return current city if there are no more roads from it
	if len(roads) == 0 {
		return city
	}

	rand.Seed(time.Now().UnixNano())
	ind := rand.Intn(len(roads))
	return roads[ind].Destination
}

//Remove city roads
func RemoveCityRoads(roads Roads, city City) Roads {
	res := make(Roads, 0)
	for _, road := range roads {
		if road.Destination != city {
			res = append(res, road)
		}
	}

	return res
}

func CreateMap(reader io.Reader) *Map {
	m := make(Map, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		m.addCityFromInputRecord(line)
	}

	return &m
}

// Get all cities on the map
func (m *Map) GetCities() []City {
	cities := make([]City, len(*m))

	i := 0
	for city := range *m {
		cities[i] = city
		i++
	}

	return cities
}

// Removes city from the map and also any roads​ ​that​ ​lead​ ​into​ ​or​ ​out​ ​of​ ​it.
func (m *Map) RemoveCity(city City) {
	//remove city itself
	delete(*m, city)

	//remove all roads to this city
	for c, cityRoads := range *m {
		(*m)[c] = RemoveCityRoads(cityRoads, city)
	}
}

// Write map to to the specified writer
func (m *Map) Print(writer io.Writer) {
	for city, roads := range *m {
		line := fmt.Sprintf("%s %s\n", city, roads.String())
		if _, err := writer.Write([]byte(line)); err != nil {
			log.Printf("error while writing city %s information to the output file: %v\n", city, err)
		}
	}
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
