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

// Recursive parrallelized function to compute the product of
//  two n x n matrices. Takes as argument three matrices, represented
//  by 2D slices. The first two arguments are the matrices to be
//  multiplied, and the third will hold the product
// Implement's Strassen's method
func matrix_mult_recursive(factor1 [][]int, factor2[][]int, product[][]int){

}

// !!! Stub
func expand_matrix(matrix [][]int, length int) (new_matrix [][]int){
	new_matrix = matrix
	return // new_matrix
}

// !!! Stub
func trim_matrix(matrix [][]int, width int, height int) (new_matrix [][]int){
	new_matrix = make([][]int, width)
	for i := range new_matrix {
		new_matrix[i] = matrix[i][:height]
	}
	return // new_matrix
}