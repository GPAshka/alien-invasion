package world

import (
	"io"
	"strings"
	"testing"
)

func TestCreateMap(t *testing.T) {
	var testCases = []struct {
		reader      io.Reader
		expectedMap *Map
	}{
		{reader: strings.NewReader(`Foo north=Bar west=Baz south=Qu-ux`),
			expectedMap: &Map{
				City("Foo"): []*Road{
					{North, City("Bar")},
					{West, City("Baz")},
					{South, City("Qu-ux")}},
			},
		},
	}

	for i, input := range testCases {
		m := CreateMap(input.reader)

		expected := (*input.expectedMap)["Foo"]
		actual := (*m)["Foo"]

		if len(expected) != len(actual) {
			t.Errorf("FAIL: input %v, expected length for 'Foo' %v but got %v", i, len(expected), len(actual))
		}

		for j, v := range actual {
			if *v != *expected[j] {
				t.Errorf("FAIL: input %v, expected Road %v but got %v", i, *expected[j], v)
			}
		}
	}
}

func TestRemoveCityRoads(t *testing.T) {
	var testCases = []struct {
		input          Roads
		cityToRemove   City
		expectedResult Roads
	}{
		{
			//city, which should be removed, exists in the input Roads
			input:          Roads{{Direction: North, Destination: City("Bar")}, {Direction: West, Destination: City("Baz")}},
			cityToRemove:   City("Bar"),
			expectedResult: Roads{{Direction: West, Destination: City("Baz")}},
		},
		{
			//city, which should be removed, does not exists in the input Roads
			input:          Roads{{Direction: North, Destination: City("Bar")}, {Direction: West, Destination: City("Baz")}},
			cityToRemove:   City("Barrr"),
			expectedResult: Roads{{Direction: North, Destination: City("Bar")}, {Direction: West, Destination: City("Baz")}},
		},
	}

	for i, test := range testCases {
		actual := RemoveCityRoads(test.input, test.cityToRemove)

		if len(test.expectedResult) != len(actual) {
			t.Errorf("FAIL: input %v, expected Roads length %v, but got %v", i, len(test.expectedResult), len(actual))
		}

		for j, v := range actual {
			if *test.expectedResult[j] != *v {
				t.Errorf("FAIL: input %v, expected Road %v, but got %v", i, test.expectedResult[j], v)
			}
		}
	}
}

func TestMap_RemoveCity(t *testing.T) {
	var testCases = []struct {
		inputMap     *Map
		cityToRemove City
		expectedMap  *Map
	}{
		{
			inputMap: &Map{
				City("Foo"): []*Road{
					{North, City("Bar")},
					{West, City("Baz")},
					{South, City("Qu-ux")}},
				City("Bar"): []*Road{
					{South, City("Foo")},
					{West, City("Bee")}},
			},
			expectedMap: &Map{
				City("Foo"): []*Road{
					{West, City("Baz")},
					{South, City("Qu-ux")}},
			},
			cityToRemove: City("Bar"),
		},
	}

	for i, input := range testCases {
		input.inputMap.RemoveCity(input.cityToRemove)
		if len(*input.inputMap) != len(*input.expectedMap) {
			t.Errorf("FAIL: input %v, expected Map length %v, but got %v", i, len(*input.expectedMap), len(*input.inputMap))
		}

		//compare City roads
		actualRoads := (*input.inputMap)["Foo"]
		expectedRoads := (*input.expectedMap)["Foo"]

		if len(expectedRoads) != len(actualRoads) {
			t.Errorf("FAIL: input %v, expected Roads length %v, but got %v", i, len(expectedRoads), len(actualRoads))
		}

		for j, v := range actualRoads {
			if *expectedRoads[j] != *v {
				t.Errorf("FAIL: input %v, expected Road %v, but got %v", i, expectedRoads[j], v)
			}
		}
	}
}
