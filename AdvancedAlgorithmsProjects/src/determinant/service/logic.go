package service

import (
	"errors"
)

func CheckIfValidMatrix(size int, matrix [][]float64) bool {

	if size != len(matrix) {
		return false
	}

	for _, row := range matrix {

		if size != len(row) {
			return false
		}

	}

	return true

}

func CalculateDeterminant(size int, matrix [][]float64) (float64, error) {

	if size < 1 {
		return -1, errors.New("Matrix size must be bigger than 0")
	} else if size == 1 {
		return matrix[0][0], nil
	} else if size == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0], nil
	} else if size == 3 {

		return CalculateDeterminantForSize3(size, matrix), nil

	} else if size < 11 {

		return CalculateDeterminantForSizeBiggerThan3(size, matrix), nil
	}

	return -1, errors.New("Matrix size must be lower than 11")
}

func CalculateDeterminantForSize3(size int, matrix [][]float64) float64 {

	result := 0.0

	for start_pos := range size {

		part_result := 1.0

		for jump_pos := range size {

			row_pos := jump_pos
			col_pos := start_pos + jump_pos

			if col_pos >= size {
				col_pos -= size
			}

			num := matrix[row_pos][col_pos]

			if num == 0 {
				part_result = 0
				break
			}

			part_result *= num

		}

		result += part_result
		part_result = 1

		for jump_pos := range size {

			row_pos := jump_pos
			col_pos := start_pos - jump_pos

			if col_pos < 0 {
				col_pos += size
			}

			num := matrix[row_pos][col_pos]

			if num == 0 {
				part_result = 0
				break
			}

			part_result *= num

		}

		result -= part_result
	}

	return result
}

func CalculateDeterminantForSizeBiggerThan3(size int, matrix [][]float64) float64 {

	result := 0.0

	row_position := FindRowWithMostZeros(matrix)

	new_matrix := MakeMatrixWithZeros(size - 1)

	for col_position := range size {

		main_number := matrix[row_position][col_position]

		if main_number == 0 {
			continue
		}

		if (row_position+col_position)%2 == 1 {
			main_number *= -1
		}

		new_row := 0

		for matrix_row := 0; matrix_row < size; matrix_row++ {

			if matrix_row == row_position {
				continue
			}

			new_col := 0

			for matrix_col := 0; matrix_col < size; matrix_col++ {

				if matrix_col == col_position {
					continue
				}

				new_matrix[new_row][new_col] = matrix[matrix_row][matrix_col]

				new_col += 1

			}

			new_row += 1

		}

		new_matrix_determinant, _ := CalculateDeterminant(size-1, new_matrix)

		result += main_number * new_matrix_determinant

	}

	return result
}

func MakeMatrixWithZeros(size int) [][]float64 {

	new_matrix := make([][]float64, size)

	for i := 0; i < size; i++ {
		new_matrix[i] = make([]float64, size)
	}

	return new_matrix

}

func FindRowWithMostZeros(matrix [][]float64) int {

	number_of_zeros := 0
	row_with_most_zeros := 0

	for id, row := range matrix {
		zeros_in_row := 0
		for _, num := range row {
			if num == 0 {
				zeros_in_row += 1
			}
		}

		if number_of_zeros < zeros_in_row {
			number_of_zeros = zeros_in_row
			row_with_most_zeros = id
		}
	}

	return row_with_most_zeros
}
