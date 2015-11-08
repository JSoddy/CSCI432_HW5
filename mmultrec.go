package main

import (
	"math"
	)

// Function to calculate the product of two integer matrices
//  Takes two matrices represented as 2D slices of integers
//  Returns the product of the two matrices if possible
//  and nil if the product is not computable
func mm_rec_init(factor1 [][]int, factor2 [][]int) (product [][]int){
	// If any dimension is 0, then we can't multiply
	if (len(factor1)    == 0 || len(factor2)    == 0 || 
		len(factor1[0]) == 0 || len(factor2[0]) == 0 ){
		return nil
	}
	// Save the side lengths for later
	m, n := len(factor1), len(factor2[0])
	// Normalize the matrices to squares which have sides with a power of 2
		// First find the longest side
	maxside := math.Max(float64(len(factor1)), 
		       math.Max(float64(len(factor2)), 
		       math.Max(float64(len(factor1[0])), 
		                float64(len(factor2[0])))))
		// Increase the longest side value to the next power of 2
	maxside = math.Pow(2, math.Ceil(math.Log2(maxside)))
		// Then expand each matrix to a square of that size
	factor1 = expand_matrix(factor1, int(maxside))
	factor2 = expand_matrix(factor2, int(maxside))
	// Initialize the result matrix with the proper size
	product = make([][]int, int64(maxside))
	for i := range product {
		product[i] = make([]int, int(maxside))
	}
	// Now call the recursive function on the normalized matrices
	matrix_mult_recursive(factor1, factor2, product)
	// Now trim off all trailing rows or columns which are
	//  entirely zeroes, to return the matrix to the proper size
	//  this is the same as returning it to an m x n matrix
	product = trim_matrix(product, m, n)
	// And we're done
	return // product
}

// !!! Stub
// Recursive parrallelized function to compute the product of
//  two n x n matrices. Takes as argument three matrices, represented
//  by 2D slices. The first two arguments are the matrices 
//  to be multiplied, and the third will hold the product.
// Implement's Strassen's method
func matrix_mult_recursive(factor1 [][]int, factor2[][]int, product[][]int){
	n := len(factor1)
	// In comments, factor1 = A, factor2 = B, product = C
	//Calculate bounds for 4 submatrices each for A and B
		//A11 = A[0:n/2][0:n/2]
		//A12 = A[n/2:n][0:n/2]
		//A21 = A[0:n/2][n/2:n]
		//A22 = A[n/2:n][n/2:n]
		//B11 = B[0:n/2][0:n/2]
		//B12 = B[n/2:n][0:n/2]
		//B21 = B[0:n/2][n/2:n]
		//B22 = B[n/2:n][n/2:n]
	//Create ten sub-matrices, S1-S10 through addition or subtraction
		//S1 = B12-B22
		//S2 = A11+A12
		//S3 = A21+A22
		//S4 = B21-B11
		//S5 = A11+A22
		//S6 = B11+B22
		//S7 = A12-A22
		//S8 = B21+B22
		//S9 = A11-A21
		//S10 = B11+B12
	//Compute seven matrix products P1-P7
		//P1 = A11*S1
		//P2 = S2*B22
		//P3 = S3*B11
		//P4 = A22*S4
		//P5 = S5*S6
		//P6 = S7*S8
		//P7 = S9*S10
	//Compute product via sums or differences of P1-P7
		//C11 = P5 + P4 - P2 + P6
		//C12 = P1 + P2
		//C21 = P3 + P4
		//C22 = P5 + P1 - P3 - P7
	//And we're done
}

// !!! Stub
// Function that returns one of the four n/2 x n/2 divisions of
//  an n x n matrix
// Takes as arguments a matrix, represented by a 2D slice of ints,
//  and an integer one through 4, specifying which quadrant of the
//  matrix to return, with 1 being the top left, 2 being the top right
//  3 being the bottom left and 4 being the bottom right quadrant
// Returns the slice modified to the original bounds
func get_matrix_quadrant(matrix [][]int, quadrant int){
	return matrix
}

// !!! Stub
// Function to return the difference of two matrices
// Takes as arguments two matrices represented by 2D slices,
//  and returns a new 2D slice with their difference
func sub_matrices(minuend [][]int, subtrahend [][]int) (difference [][]int){

	difference = minuend // for test
	return // difference
}

// !!! Stub
// Function to return the sum of two matrices
// Takes as arguments two matrices represented by 2D slices,
//  and returns a new 2D slice with their sum
func add_matrices(matrix1 [][]int, matrix2 [][]int) (sum [][]int){

	sum = matrix1 // for test
	return // sum
}


// Function to increase the width and height of a matrix to a specified
//  size
// Accepts a matrix, represented as a 2D slice of ints
// Returns a new matrix with the same data, with width and
//  height equal to argument length
func expand_matrix(matrix [][]int, length int) (new_matrix [][]int){
	// Make a new slice of the appropriate size
	new_matrix = make([][]int, length)
	for i := range new_matrix {
		new_matrix[i] = make([]int, length)
	}
	// Copy the original data into it
	for i := range matrix {
		copy(new_matrix[i], matrix[i])
	}
	return // new_matrix
}

// Function to reduce a matrix, represented as a 2D slice of ints,
//  in size to width x height
// Accepts a matrix, integer width and integer height
// Returns a new matrix of the proper size
func trim_matrix(matrix [][]int, width int, height int) ([][]int){
	matrix = matrix[:width]
	for i := range matrix {
		matrix[i] = matrix[i][:height]
	}
	return matrix
}