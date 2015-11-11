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

	// Now call the recursive function on the normalized matrices
	product = matrix_mult_recursive(factor1, factor2)

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
func matrix_mult_recursive(factor1 [][]int, factor2[][]int) (product [][]int){
	lenF1 := len(factor1)
	product = make([][]int,lenF1)
	for i := range product{
		product[i] = make([]int,lenF1)
	}
	//exit case
	if lenF1==1 {
		product[0][0] = factor1[0][0]*factor2[0][0]
		return product
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

	//Compute seven matrix products P1-P7
		// P1 := A11*S1
		p1 := matrix_mult_recursive(a11, s1) // Make these return to s1-s7!!! Woot ???!!!
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
		addMatrix2(p5,p4,c11)
		subMatrix2(c11,p2,c11)
		addMatrix2(c11,p6,c11)
		//C12 = P1 + P2
		addMatrix2(p1,p2,c12)
		//C21 = P3 + P4
		addMatrix2(p3,p4,c21)
		//C22 = P5 + P1 - P3 - P7
		addMatrix2(p1,p5,c22)
		subMatrix2(c22,p3,c22)
		subMatrix2(c22,p7,c22)
	//And we're done
	//After we build the return 2d slice from c11-c22
		/*
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
		*/
	return
}

// !!! Stub
// Function that returns one of the four n/2 x n/2 divisions of
//  an n x n matrix
// Takes as arguments a matrix, represented by a 2D slice of ints,
//  and an integer one through 4, specifying which quadrant of the
//  matrix to return, with 1 being the top left, 2 being the top right
//  3 being the bottom left and 4 being the bottom right quadrant
func get_matrix_quadrant(input [][]int, quadrant int) (output [][]int){
	size := len(input)
	mid  := size / 2
	var lower, upper, left int
	output = make([][]int, mid)
	switch (quadrant) {
		case 1: left = 0; lower = 0; upper = mid
		case 2: left = 0; lower = mid; upper = size
		case 3: left = mid; lower = 0; upper = mid
		case 4: left = mid; lower = mid; upper = size
	}
	for i := 0; i < mid; i++ {
		output[i] = input[i+left][lower:upper]
	}
 	return // output
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

/*
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
*/

/*
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
*/

/*
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
*/


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

// Function to return the sum of two matrices
// Takes as arguments two matrices represented by 2D slices,
//  and a third matrix to store the result
func addMatrix2(a [][]int, b [][]int, matrix [][]int){
	length := len(a)
	//matrix = make([][]int,length)
	for i := 0; i < length; i++ {
		//matrix[i] = make([]int,length)
		for j := 0; j < length; j++ {
			matrix[i][j]=a[i][j]+b[i][j]
		}
	}
	return 
}

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

// Function to return the difference of two matrices
// Takes as arguments two matrices represented by 2D slices,
//  and returns a new 2D slice with their difference
func subMatrix2(a [][]int, b [][]int, matrix [][]int){
	length := len(a)
	//matrix = make([][]int,length)
	for i := 0; i < length; i++ {
		//matrix[i] = make([]int,length)
		for j := 0; j < length; j++ {
			matrix[i][j]=a[i][j]-b[i][j]
		}
	}
	return 
}

/*
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

*/
