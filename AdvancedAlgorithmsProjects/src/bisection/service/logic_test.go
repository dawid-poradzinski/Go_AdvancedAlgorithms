package bisection_logic

import (
	"reflect"
	"testing"
)

func TestValidateRange(t *testing.T) {

	cases := []struct {
		delta    float64
		min      float64
		max      float64
		expected bool
	}{
		{delta: 1, min: -10, max: 10, expected: false},
		{delta: 30, min: -10, max: 10, expected: true},
		{delta: 1, min: 10, max: -10, expected: true},
	}

	for _, tc := range cases {
		err := ValidateRange(tc.delta, tc.min, tc.max)
		result := err != nil
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("ValidateRange(%f,%f,%f) = %v; want error -> %v", tc.delta, tc.min, tc.max, err, tc.expected)
		}
	}

}

func TestCalculateValueInPoint(t *testing.T) {

	cases := []struct {
		numbers  []float64
		x        float64
		expected float64
	}{
		{numbers: []float64{2, 5, 3, -2, 3, 1}, x: 0, expected: 2},
		{numbers: []float64{20, -6, -5, 1, -2, 5}, x: 5, expected: 14365},
		{numbers: []float64{40, 5, 12, 7, 4, 12}, x: 1, expected: 80},
		{numbers: []float64{40, 5, 12, 7, 4, 12}, x: -1, expected: 32},
	}

	for id, tc := range cases {
		result := CalculateValueInPoint(tc.numbers, tc.x)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("CalculateValueInPoint(id%v,%f) = %v; want error -> %v", id, tc.x, result, tc.expected)
		}
	}

}
