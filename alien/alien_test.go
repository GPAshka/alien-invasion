package alien

import "testing"

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
