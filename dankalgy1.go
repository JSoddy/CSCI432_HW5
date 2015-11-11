package main

// Function to calculate the product of two integer matrices
// Uses the 'naive' n^3 algorithm
//  Takes two matrices represented as 2D slices of integers
//  Returns the product of the two matrices if possible
//  and nil if the product is not computable
func dankalgy1(matrix1 [][]int, matrix2 [][]int) (matrix3 [][]int){
	m, n := 0, 0

	// Make sure that the two matrices have a product
	l := len(matrix1)
	if l > 0 {
		m = len(matrix1[0])
		if m > 0 && len(matrix2) == m {
			n = len(matrix2[0])
		}
	}
	// If n is 0, then the matrices have no product. Return nil
	if n == 0 {return nil}
	// Initialize the return array
	matrix3 = make([][]int, l)
	for i := range matrix3 {
		matrix3[i] = make([]int, n)
	}
	// Compute all of the necessary products with a triply nested loop
	for i:= 0; i < l; i++ {
		for j:= 0; j < n; j++ {
			for k:= 0; k < m; k++ {
				matrix3[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}
	}
	return //matrix3
}