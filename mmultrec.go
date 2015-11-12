package main

import (
	//"math"
	//"time"
	//"fmt"
	)

// Function to calculate the product of two integer matrices
//  Takes two matrices represented as 2D slices of integers
//  Returns the product of the two matrices if possible
//  and nil if the product is not computable
//  assumes all matrix dimensions are equal
func mm_rec_init(factor1 [][]int, factor2 [][]int) (product [][]int){
	// If any dimension is 0, then we can't multiply
	if (len(factor1)    == 0 || len(factor2)    == 0 ||
		len(factor1[0]) == 0 || len(factor2[0]) == 0 ){
		return nil
	}
	/*
	// Save the side lengths for later
	m, n := len(factor1), len(factor2[0])
	// Normalize the matrices to squares which have sides with a power of 2
		// First find the longest side
	maxside := math.Max(float64(len(factor1)),
		       math.Max(float64(len(factor2)),
		       math.Max(float64(len(factor1[0])),
		                float64(len(factor2[0])))))
		// Increase the longest side value to the next power of 2
	/*
	maxside = math.Pow(2, math.Ceil(math.Log2(maxside)))
		// Then expand each matrix to a square of that size
	factor1 = expand_matrix(factor1, int(maxside))
	factor2 = expand_matrix(factor2, int(maxside))
	*/
	// Create a return array
	maxside := len(factor1)
	product = make([][]int, int(maxside))
	for i := range product{
		product[i] = make([]int, int(maxside))
	}

	// Now call the recursive function on the normalized matrices
	//product = 
	matrix_mult_recursive(factor1, factor2, product)
	/*
	// Now trim off all trailing rows or columns which are
	//  entirely zeroes, to return the matrix to the proper size
	//  this is the same as returning it to an m x n matrix
	product = trim_matrix(product, m, n)
	// And we're done
	*/
	return // product
}

// Recursive parrallelized function to compute the product of
//  two n x n matrices. Takes as argument three matrices, represented
//  by 2D slices. The first two arguments are the matrices
//  to be multiplied, and the third will hold the product.
// Implement's Strassen's method
func matrix_mult_recursive(factor1 [][]int, factor2[][]int, product [][]int){
	//start := time.Now()
	lenF1 := len(factor1)
	//exit case
	if lenF1==1 {
		product[0][0] = factor1[0][0]*factor2[0][0]
		return // product
	}
	// In comments, factor1 = A, factor2 = B, product = C
	//Calculate bounds for 4 submatrices each for A and B
		//A11 = A[0:n/2][0:n/2]
		//A12 = A[n/2:n][0:n/2]
		//A21 = A[0:n/2][n/2:n]
		//A22 = A[n/2:n][n/2:n]
	a11:=get_matrix_quadrant(factor1, 1)
	a12:=get_matrix_quadrant(factor1, 2)
	a21:=get_matrix_quadrant(factor1, 3)
	a22:=get_matrix_quadrant(factor1, 4)
		//B11 = B[0:n/2][0:n/2]
		//B12 = B[n/2:n][0:n/2]
		//B21 = B[0:n/2][n/2:n]
		//B22 = B[n/2:n][n/2:n]
	b11:=get_matrix_quadrant(factor2, 1)
	b12:=get_matrix_quadrant(factor2, 2)
	b21:=get_matrix_quadrant(factor2, 3)
	b22:=get_matrix_quadrant(factor2, 4)	
		//C11 = C[0:n/2][0:n/2]
		//C12 = C[n/2:n][0:n/2]
		//C21 = C[0:n/2][n/2:n]
		//C22 = C[n/2:n][n/2:n]
	c11:=get_matrix_quadrant(product, 1)
	c12:=get_matrix_quadrant(product, 2)
	c21:=get_matrix_quadrant(product, 3)
	c22:=get_matrix_quadrant(product, 4)
	//Create ten sub-matrices, S1-S10 through addition or subtraction
		//S1 = B12-B22
		s1:=subMatrix(b12,b22)
		//S2 = A11+A12
		s2:=addMatrix(a11,a12)
		//S3 = A21+A22
		s3:=addMatrix(a21,a22)
		//S4 = B21-B11
		s4:=subMatrix(b21,b11)
		//S5 = A11+A22
		s5:=addMatrix(a11,a22)
		//S6 = B11+B22
		s6:=addMatrix(b11,b22)
		//S7 = A12-A22
		s7:=subMatrix(a12,a22)
		//S8 = B21+B22
		s8:=addMatrix(b21,b22)
		//S9 = A11-A21
		s9:=subMatrix(a11,a21)
		//S10 = B11+B12
		s10:=addMatrix(b11,b12)
	//fmt.Printf("Level %d, done: %s\n", lenF1, time.Since(start))
	//Compute seven matrix products P1-P7
		// P1 := A11*S1
		matrix_mult_recursive(a11, s1, s1)
		// P2 = S2*B22
		matrix_mult_recursive(s2, b22, s2)
		// P3 = S3*B11
		matrix_mult_recursive(s3,b11, s3)
		// P4 = A22*S4
		matrix_mult_recursive(a22, s4, s4)
		// P5 = S5*S6
		matrix_mult_recursive(s5, s6, s5)
		// P6 = S7*S8
		matrix_mult_recursive(s7, s8, s6)
		// P7 = S9*S10
		matrix_mult_recursive(s9, s10, s7)

		//C11
		addMatrix2(s5,s4,c11)
		subMatrix2(c11,s2,c11)
		addMatrix2(c11,s6,c11)
		//C12 = P1 + P2
		addMatrix2(s1,s2,c12)
		//C21 = P3 + P4
		addMatrix2(s3,s4,c21)
		//C22 = P5 + P1 - P3 - P7
		addMatrix2(s1,s5,c22)
		subMatrix2(c22,s3,c22)
		subMatrix2(c22,s7,c22)
	//And we're done
	return
}

