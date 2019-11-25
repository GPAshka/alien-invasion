package alien

import (
	"alien-invasion/world"
	"testing"
)

func TestAliens_RemoveAliens(t *testing.T) {
	var testCases = []struct {
		inputAliens    Aliens
		expectedAliens Aliens
		aliensToRemove []int
	}{
		{
			inputAliens: Aliens{
				1: &Alien{},
				2: &Alien{},
				3: &Alien{},
				4: &Alien{},
			},
			expectedAliens: Aliens{
				1: &Alien{},
				3: &Alien{},
			},
			aliensToRemove: []int{2, 4},
		},
	}

	for i, input := range testCases {
		input.inputAliens.RemoveAliens(input.aliensToRemove)

		if len(input.inputAliens) != len(input.expectedAliens) {
			t.Errorf("FAIL: input %v, expected aliens len %v, but got %v", i, len(input.expectedAliens), len(input.inputAliens))
		}

		for j, v := range input.expectedAliens {
			if *input.inputAliens[j] != *v {
				t.Errorf("FAIL: input %v, expected alien %v, but got %v", i, v, input.inputAliens[j])
			}
		}
	}
}

func TestAliens_GroupByCity(t *testing.T) {
	var testCases = []struct {
		inputAliens    Aliens
		expectedResult map[world.City][]int
	}{
		{
			inputAliens: Aliens{
				1: &Alien{CurrentCity: world.City("Foo")},
				2: &Alien{CurrentCity: world.City("Foo")},
				3: &Alien{CurrentCity: world.City("Bar")},
				4: &Alien{CurrentCity: world.City("Foo")},
			},
			expectedResult: map[world.City][]int{
				world.City("Foo"): {1, 2, 4},
				world.City("Bar"): {3},
			},
		},
	}

	for i, input := range testCases {
		actual := input.inputAliens.GroupByCity()

		if len(actual) != len(input.expectedResult) {
			t.Errorf("FAIL: input %v, expected result map length %v, but got %v", i, len(input.expectedResult), len(actual))
		}

		for key, expectedValue := range input.expectedResult {
			actualValue := actual[key]

			for j, v := range actualValue {
				if v != expectedValue[j] {
					t.Errorf("FAIL: key %v, expected element %v, but got %v", key, expectedValue[j], v)
				}
			}
		}
	}
}
