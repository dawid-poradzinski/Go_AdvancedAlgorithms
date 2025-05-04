package bisection_logic

import (
	bisection_model "AdvancedAlgorithmsProjects/src/bisection/model"
	"math"
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

func TestBisection(t *testing.T) {

	cases := []struct {
		request  bisection_model.BisectionRequest
		err      bool
		expected float64
	}{
		{request: bisection_model.BisectionRequest{Numbers: []float64{1, 2, 3, 4, 5, 6}, Delta: 0.001, Min: -1, Max: 1}, err: false, expected: -0.670},
		{request: bisection_model.BisectionRequest{Numbers: []float64{1, 2, 3, 4, 5, 6}, Delta: 0.001, Min: 2, Max: 3}, err: true, expected: -1},
		{request: bisection_model.BisectionRequest{Numbers: []float64{1, 2, 3, 4, 5, 6}, Delta: 0.001, Min: -10, Max: 10}, err: false, expected: -0.670},
		{request: bisection_model.BisectionRequest{Numbers: []float64{-5, 2, 5, -2, 3, -2}, Delta: 0.0001, Min: -1, Max: 0}, err: false, expected: -0.8236},
		{request: bisection_model.BisectionRequest{Numbers: []float64{-5, 2, 5, -2, 3, -2}, Delta: 0.0001, Min: 0.5, Max: 1.5}, err: false, expected: 0.8784},
		{request: bisection_model.BisectionRequest{Numbers: []float64{-5, 2, 5, -2, 3, -2}, Delta: 0.0001, Min: 1, Max: 2}, err: false, expected: 1.6857},
		{request: bisection_model.BisectionRequest{Numbers: []float64{-5, 2, 5, -2, 3, -2}, Delta: 0.0001, Min: -0.5, Max: 2}, err: true, expected: -1},
	}

	for id, tc := range cases {
		result, err := Bisection(tc.request)
		if (err != nil) != tc.err {
			t.Errorf("Bisection(id%v) = %v -> want error -> %v", id, err, tc.err)
		} else if math.Abs(result-tc.expected) > tc.request.Delta {
			t.Errorf("Bisection(id%v) = %f;%f. Math.abs -> %f, epsilon, %f", id, result, tc.expected, math.Abs(result-tc.expected), tc.request.Delta)
		}
	}

}
