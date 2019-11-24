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
