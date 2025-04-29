package service

import (
	"AdvancedAlgorithmsProjects/src/quadruple/model"
	"reflect"
	"testing"
)

func TestMakePrimes(t *testing.T) {
	cases := []struct {
		start    int
		end      int
		expected []int
	}{
		{start: 2, end: 2, expected: []int{2}},
		{start: 2, end: 10, expected: []int{2, 3, 5, 7}},
		{start: 10, end: 20, expected: []int{11, 13, 17, 19}},
	}

	for _, tc := range cases {
		result := MakePrimes(tc.start, tc.end)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("MakePrime(%d,%d) = %v; want %v", tc.start, tc.end, result, tc.expected)
		}
	}
}

func TestCheckIfPair(t *testing.T) {
	cases := []struct {
		first    int
		second   int
		expected bool
	}{
		{
			first:    2,
			second:   3,
			expected: false,
		},
		{
			first:    11,
			second:   13,
			expected: true,
		},
		{
			first:    3,
			second:   5,
			expected: true,
		},
		{
			first:    17,
			second:   19,
			expected: true,
		},
	}

	for _, tc := range cases {
		result := CheckIfPair(tc.first, tc.second)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("CheckIfPair(%d, %d) = %v; want %v", tc.first, tc.second, result, tc.expected)
		}
	}
}

func TestFindPairsAndQuadruplets(t *testing.T) {
	cases := []struct {
		data     []int
		expected model.QuadrupletResponse
	}{
		{data: []int{137, 139, 149, 151}, expected: model.QuadrupletResponse{
			PrimeCount:       4,
			PairsCount:       2,
			Quadruplets:      [][4]int{},
			QuadrupletsCount: 0,
		}},
		{data: []int{2, 3, 5, 7, 11, 13, 17, 19}, expected: model.QuadrupletResponse{
			PrimeCount:       8,
			PairsCount:       4,
			Quadruplets:      [][4]int{{5, 7, 11, 13}, {11, 13, 17, 19}},
			QuadrupletsCount: 2,
		}},
	}

	for _, tc := range cases {
		result := FindPairsAndQuadruplets(tc.data)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("FindPairsAndQuadruplets(%v) = %v; want %v", tc.data, result, tc.expected)
		}
	}
}
