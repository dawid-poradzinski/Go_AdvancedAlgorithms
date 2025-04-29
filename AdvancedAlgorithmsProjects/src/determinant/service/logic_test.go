package service

import (
	"reflect"
	"testing"
)

func TestCheckIfValidMatrix(t *testing.T) {
	cases := []struct {
		size     int
		matrix   [][]float64
		expected bool
	}{
		{size: 2, matrix: [][]float64{{1, 2}, {1, 2}}, expected: true},
		{size: 2, matrix: [][]float64{{1, 2}}, expected: false},
		{size: 5, matrix: [][]float64{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}, {1, 2, 3, 4}, {1, 2, 3, 4, 5}}, expected: false},
	}

	for id, tc := range cases {
		result := CheckIfValidMatrix(tc.size, tc.matrix)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("checkIfValidMatrix(%d,%d) = %v; want %v", tc.size, id, result, tc.expected)
		}
	}
}

func TestCalculateDeterminant(t *testing.T) {
	cases := []struct {
		size     int
		matrix   [][]float64
		expected float64
		err      bool
	}{
		// {size: 0, matrix: [][]float64{}, err: true, expected: -1},
		// {size: 12, matrix: [][]float64{{1, 2}}, err: true, expected: -1},
		// {size: 1, matrix: [][]float64{{-5}}, err: false, expected: -5},
		// {size: 2, matrix: [][]float64{{-2, 1}, {2, 1}}, err: false, expected: -4},
		// {size: 3, matrix: [][]float64{{-2, 1, 2}, {2, 1, 1}, {2, 4, -2}}, err: false, expected: 30},
		// {size: 4, matrix: [][]float64{{-2, 1, 2, 4}, {2, 1, 1, 0}, {2, 4, -2, 1}, {3, -2, 0, 1}}, err: false, expected: 165},
		{size: 10, matrix: [][]float64{
			{4, 4, 3, 2, 1, 9, 2, 5, 6, 7},
			{9, 0, 2, 2.5, 3, 0, 1, 8, 4, 6},
			{2, 3, 0, 1, 3, 4, 1, 7, 5, 9},
			{0, 2, 1, 0.5, 2, 10, 5, 3, 6, 8},
			{1, 0, 2, 3, 0, 1, 2, 4, 9, 7},
			{3, 0, 1, 1, 2, 1, 2.5, 6, 8, 5},
			{4, 9, 3, 2, 0, 1, 2, 5, 7, 0},
			{5, 6, 8, 9, 3, 7, 4, 2, 1, 0},
			{7, 3, 4, 1, 8, 2, 0, 9, 6, 5},
			{2, 5, 9, 3, 6, 7, 1, 4, 8, 0},
		}, err: false, expected: -1.45684295e+07},
	}

	for id, tc := range cases {
		result, err := CalculateDeterminant(tc.size, tc.matrix)
		if tc.err == (err == nil) {
			t.Errorf("calculateDeterminant(%d,%d) = %v; want have error -> %v", tc.size, id, err, tc.err)
		} else if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("calculateDeterminant(%d,%d) = %v; want %v", tc.size, id, result, tc.expected)
		}
	}
}
