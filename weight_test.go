package weight

import (
	"reflect"
	"testing"
)

func TestWeight(t *testing.T) {
	tests := []struct {
		Weight   Weight
		Expected []string
	}{
		{
			Weight: Weight{
				[]*Backend{
					{"A", 5, 0},
					{"B", 3, 0},
					{"C", 1, 0},
				},
			},
			Expected: []string{"A", "B", "A", "C", "A", "B", "A", "B", "A"},
		},
		{
			Weight: Weight{
				[]*Backend{
					{"A", 2, 0},
					{"B", 1, 0},
					{"C", 1, 0},
				},
			},
			Expected: []string{"A", "B", "C", "A"},
		},
		{
			Weight: Weight{
				[]*Backend{
					{"A", 2, 0},
					{"B", 3, 0},
					{"C", 2, 0},
				},
			},
			Expected: []string{"B", "A", "C", "B", "A", "C", "B"},
		},
	}

	for _, tt := range tests {
		t.Run("Test returned server", func(t *testing.T) {
			var got []string

			for i := 0; i < len(tt.Expected); i++ {
				got = append(got, tt.Weight.Next())
			}

			if !reflect.DeepEqual(got, tt.Expected) {
				t.Fatalf("got %v, want %v", got, tt.Expected)
			}
		})
	}
}

func TestAddWeight(t *testing.T) {
	w := NewWeight()
	w.AddWeight("A", 2)
	w.AddWeight("B", 3)
	w.AddWeight("C", 2)

	expected := []string{"B", "A", "C", "B", "A", "C", "B"}
	var got []string

	for i := 0; i < len(expected); i++ {
		got = append(got, w.Next())
	}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("got %v, want %v", got, expected)
	}
}
