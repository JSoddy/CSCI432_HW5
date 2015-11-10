package main

import (
	"math"
	"fmt"
	"math/rand"
	"time"
	)

const cellmax = 11

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
	// matrix_mult_recursive(factor1, factor2, product)
	matrix_mult_recursive(factor1, factor2)

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
// gr33n I redid this func inputs for testing
// matrix [][]int){
// 	matrix = make([][]int,length)
func matrix_mult_recursive(factor1 [][]int, factor2[][]int) (product [][]int){
// func matrix_mult_recursive(factor1[][]int, factor2[][]int){
	lenF2 := len(factor2)
	lenF1 := len(factor1)
	product = make([][]int,lenF2)
	product[0] = make([]int,lenF2)
	//exit case
	if lenF1==1 {
		product[0][0] = factor1[0][0]*factor2[0][0]
		return product
	}
	for i := 1; i < lenF1; i++ {
		product[i] = make([]int,lenF2)
	}

	// In comments, factor1 = A, factor2 = B, product = C
	//Calculate bounds for 4 submatrices each for A and B
		//A11 = A[0:n/2][0:n/2]
		//A12 = A[n/2:n][0:n/2]
		//A21 = A[0:n/2][n/2:n]
		//A22 = A[n/2:n][n/2:n]
	a11:=[][]int{}
	for i := 0; i < lenF1/2; i++ {
			a11 = append(a11, factor1[i][:lenF1/2])
	}
	// printmatrix(a11)

	a12:=[][]int{}
	for i := 0; i < lenF1/2; i++ {
			a12 = append(a12, factor1[i][lenF2/2:])
	}
	// printmatrix(a12)

	a21:=[][]int{}
	for i := lenF1/2; i < lenF1; i++ {
			a21 = append(a21, factor1[i][:lenF2/2])
	}
	// printmatrix(a21)

	a22:=[][]int{}
	for i := lenF1/2; i < lenF1; i++ {
			a22 = append(a22, factor1[i][lenF2/2:])
	}
	// printmatrix(a22)

		//B11 = B[0:n/2][0:n/2]
		//B12 = B[n/2:n][0:n/2]
		//B21 = B[0:n/2][n/2:n]
		//B22 = B[n/2:n][n/2:n]

		b11:=[][]int{}
		for i := 0; i < lenF2/2; i++ {
				b11 = append(b11, factor2[i][:lenF2/2])
		}
		// printmatrix(b11)

		b12:=[][]int{}
		for i := 0; i < lenF2/2; i++ {
				b12 = append(b12, factor2[i][lenF2/2:])
		}
		// printmatrix(b12)

		b21:=[][]int{}
		for i := lenF2/2; i < lenF2; i++ {
				b21 = append(b21, factor2[i][:lenF2/2])
		}
		// printmatrix(b21)

		b22:=[][]int{}
		for i := lenF2/2; i < lenF2; i++ {
				b22 = append(b22, factor2[i][lenF2/2:])
		}
		// printmatrix(b22)

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

	//Compute seven matrix products P1-P7
		// P1 := A11*S1
		p1 := matrix_mult_recursive(a11, s1)
		// p1 := multiplyMatrix(a11, s1)

		// P2 = S2*B22
		p2 := matrix_mult_recursive(s2, b22)
		// p2 := multiplyMatrix(s2, b22)

		// P3 = S3*B11
		p3 := matrix_mult_recursive(s3,b11)
		// p3 := multiplyMatrix(s3,b11)

		// P4 = A22*S4
		p4 := matrix_mult_recursive(a22, s4)
		// p4 := multiplyMatrix(a22, s4)

		// P5 = S5*S6
		p5 := matrix_mult_recursive(s5, s6)
		// p5 := multiplyMatrix(s5, s6)

		// P6 = S7*S8
		p6 := matrix_mult_recursive(s7, s8)
		// p6 := multiplyMatrix(s7, s8)

		// P7 = S9*S10
		p7 := matrix_mult_recursive(s9, s10)
		// p7 := multiplyMatrix(s9, s10)


	//Compute product via sums or differences of P1-P7
		//C11 = P5 + P4 - P2 + P6
		c11 := addMatrix(subMatrix(addMatrix(p5,p4),p2),p6)
		//C12 = P1 + P2
		c12 := addMatrix(p1,p2)
		//C21 = P3 + P4
		c21 := addMatrix(p3,p4)
		//C22 = P5 + P1 - P3 - P7
		c22 := subMatrix(subMatrix(addMatrix(p1,p5),p3),p7)
	//And we're done
	//After we build the return 2d slice from c11-c22
	for i := 0; i < lenF2/2; i++ {
		for j := 0; j < lenF2/2; j++ {
			//quadrant 11
			product[i][j]=c11[i][j]
			//quadrant 12
			product[i][j+lenF2/2]=c12[i][j]
			//quadrant 21
			product[i+lenF2/2][j]=c21[i][j]
			//quadrant 22
			product[i+lenF2/2][j+lenF2/2]=c22[i][j]
		}
	}
	return product
}

// !!! Stub
// Function that returns one of the four n/2 x n/2 divisions of
//  an n x n matrix
// Takes as arguments a matrix, represented by a 2D slice of ints,
//  and an integer one through 4, specifying which quadrant of the
//  matrix to return, with 1 being the top left, 2 being the top right
//  3 being the bottom left and 4 being the bottom right quadrant
// Returns the slice modified to the original bounds
// func get_matrix_quadrant(matrix [][]int, quadrant int){
// 	return matrix
// }








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

// This function prints a matrix represented by a 2D slice Woot!
func printmatrix(s [][]int){
	// Let's just find out how many spaces we need to allow for each number.
	//  Nevermind why
	numwidth := math.Floor(math.Log10(math.Pow(cellmax-1, 2) * float64(len(s)))) + 1
	fmt.Println()
	for i := range s {
		fmt.Print("| ")
		for j := range s[i] {
			// Pay no attention to this ugliness. It's not really here
			for k := 0; k < int(numwidth -
				math.Max(1,math.Floor(math.Log10(float64(s[i][j]))+1))); k++ {
					fmt.Print(" ")
				}
			// What ugliness? I don't know what you are talking about.
			fmt.Printf("%d ", s[i][j])
		}
		fmt.Println("|")
	}
}

// function to create a randomly initialized matrix
//  accepts two integers w and h
//  returns a w by h matrix represented as a 2D slice
func rndmatrix(h int, w int) (matrix [][]int){
	matrix = make([][]int, h)
	for i := range matrix {
		matrix[i] = make([]int, w)
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(cellmax)
		}
	}
	return // matrix
}

//this assumes that the matrices are nxn
func multiplyMatrix(a [][]int, b[][]int) (product [][]int){
	lenA := len(a)
	var temp int
	product = make([][]int,lenA)
	fmt.Println(product)
	//iterates through rows
	for i := 0; i < lenA; i++ {
		//iterates through colums
		product[i] = make([]int,lenA)
		for j := 0; j < lenA; j++ {
			//iterates through the individual elements of the selected colums
			for q := 0; q < lenA; q++ {
				temp += a[i][q]*b[q][j]
			}
			product[i][j]=temp
			temp=0
		}
	}
	return product
}

// !!! Stub
// Function to return the sum of two matrices
// Takes as arguments two matrices represented by 2D slices,
//  and returns a new 2D slice with their sum
func addMatrix(a [][]int, b [][]int) (matrix [][]int){
	length := len(a)
	matrix = make([][]int,length)
	for i := 0; i < length; i++ {
		matrix[i] = make([]int,length)
		for j := 0; j < length; j++ {
			matrix[i][j]=int(a[i][j]+b[i][j])
		}
	}
	return //matrix
}

// !!! Stub
// Function to return the difference of two matrices
// Takes as arguments two matrices represented by 2D slices,
//  and returns a new 2D slice with their difference
func subMatrix(a [][]int, b [][]int) (matrix [][]int){
	length := len(a)
	matrix = make([][]int,length)
	for i := 0; i < length; i++ {
		matrix[i] = make([]int,length)
		for j := 0; j < length; j++ {
			matrix[i][j]=int(a[i][j]-b[i][j])
		}
	}
	return //matrix
}

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	// l := rand.Intn(9) + 2
	// gr33n took away random for testing
	m := 4
	n := 4

	matrix1 := rndmatrix(m, n)
	matrix2 := rndmatrix(m, n)

	fmt.Println("two randomly generated matrixs")
	printmatrix(matrix1)
	printmatrix(matrix2)
	strassesn:=matrix_mult_recursive(matrix1, matrix2)
	// strassesn:=matrix_mult_recursive(matrix1, matrix2)
	fmt.Println("Strassesn's algorithm on the two above matrixs")
	printmatrix(strassesn)
}
