package main

import (
	"fmt"
	"math"
)

// Matrix represents a 2D matrix
type Matrix [][]int

// NewMatrix creates a new matrix with the given size
func NewMatrix(size int) Matrix {
	matrix := make(Matrix, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}
	return matrix
}

// Add adds two matrices
func Add(A, B Matrix) Matrix {
	n := len(A)
	C := NewMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	return C
}

// Subtract subtracts matrix B from matrix A
func Subtract(A, B Matrix) Matrix {
	n := len(A)
	C := NewMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] - B[i][j]
		}
	}
	return C
}

// Strassen multiplies two matrices using Strassen's algorithm
func Strassen(A, B Matrix) Matrix {
	n := len(A)

	// Base case: use standard multiplication for small matrices
	// if n <= 2 {
	// 	return standardMultiply(A, B)
	// }

	// Ensure matrix size is a power of 2
	newSize := nextPowerOfTwo(n)
	if newSize != n {
		A = padMatrix(A, newSize)
		B = padMatrix(B, newSize)
		n = newSize
	}

	// Divide matrices into quadrants
	half := n / 2

	A11 := getQuadrant(A, 0, 0, half)
	A12 := getQuadrant(A, 0, half, half)
	A21 := getQuadrant(A, half, 0, half)
	A22 := getQuadrant(A, half, half, half)

	B11 := getQuadrant(B, 0, 0, half)
	B12 := getQuadrant(B, 0, half, half)
	B21 := getQuadrant(B, half, 0, half)
	B22 := getQuadrant(B, half, half, half)

	// Calculate the 7 products using Strassen's formulas
	M1 := Strassen(Add(A11, A22), Add(B11, B22))
	M2 := Strassen(Add(A21, A22), B11)
	M3 := Strassen(A11, Subtract(B12, B22))
	M4 := Strassen(A22, Subtract(B21, B11))
	M5 := Strassen(Add(A11, A12), B22)
	M6 := Strassen(Subtract(A21, A11), Add(B11, B12))
	M7 := Strassen(Subtract(A12, A22), Add(B21, B22))

	// Combine the results into the final matrix
	C11 := Add(Subtract(Add(M1, M4), M5), M7)
	C12 := Add(M3, M5)
	C21 := Add(M2, M4)
	C22 := Add(Subtract(Add(M1, M3), M2), M6)

	// Combine quadrants into result matrix
	C := combineQuadrants(C11, C12, C21, C22)

	return C
}

// standardMultiply performs standard matrix multiplication
func standardMultiply(A, B Matrix) Matrix {
	n := len(A)
	C := NewMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	return C
}

// getQuadrant extracts a quadrant from a matrix
func getQuadrant(M Matrix, row, col, size int) Matrix {
	quadrant := NewMatrix(size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			quadrant[i][j] = M[row+i][col+j]
		}
	}
	return quadrant
}

// combineQuadrants combines four quadrants into a single matrix
func combineQuadrants(C11, C12, C21, C22 Matrix) Matrix {
	half := len(C11)
	n := half * 2
	C := NewMatrix(n)

	for i := 0; i < half; i++ {
		for j := 0; j < half; j++ {
			C[i][j] = C11[i][j]
			C[i][j+half] = C12[i][j]
			C[i+half][j] = C21[i][j]
			C[i+half][j+half] = C22[i][j]
		}
	}
	return C
}

// padMatrix pads a matrix to the given size with zeros
func padMatrix(M Matrix, newSize int) Matrix {
	oldSize := len(M)
	padded := NewMatrix(newSize)
	for i := 0; i < oldSize; i++ {
		for j := 0; j < oldSize; j++ {
			padded[i][j] = M[i][j]
		}
	}
	return padded
}

// nextPowerOfTwo returns the next power of 2 greater than or equal to n
func nextPowerOfTwo(n int) int {
	return int(math.Pow(2, math.Ceil(math.Log2(float64(n)))))
}

// PrintMatrix prints a matrix
func PrintMatrix(M Matrix) {
	for _, row := range M {
		fmt.Println(row)
	}
}

func main() {
	// Example usage
	A := Matrix{
		{0, 1, 1, 2},
		{0, 0, 3, 4},
		{1, 0, 1, 0},
		{0, 0, 1, 0},
	}

	B := Matrix{
		{1, 1, 1, 1},
		{0, 1, 2, 0},
		{4, 2, 1, 1},
		{0, 0, 0, 1},
	}

	fmt.Println("Matrix A:")
	PrintMatrix(A)
	fmt.Println("\nMatrix B:")
	PrintMatrix(B)

	result := Strassen(A, B)
	fmt.Println("\nResult (A × B):")
	PrintMatrix(result)
}
