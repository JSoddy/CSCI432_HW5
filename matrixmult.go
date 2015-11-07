package main

import (
	"fmt"
	"math/rand"
	"time"
	)

func printmatrix(s [][]int){
	fmt.Println()
	for i := range s {
		fmt.Print("| ")
		for j := range s[i] {
			fmt.Printf("%7d ", s[i][j])
		}
		fmt.Println("|")
	}
}

func main(){
	rand.Seed(int64(time.Now().Nanosecond()))
	l := rand.Intn(10)
	m := rand.Intn(10)
	n := rand.Intn(10)

	matrix1 := rndmatrix(m + 2, l + 2)
	matrix2 := rndmatrix(n + 2, m + 2)
		
	printmatrix(matrix1)
	printmatrix(matrix2)

	matrix3 := dankalgy1(matrix1, matrix2)

	printmatrix(matrix3)
}

// Function to calculate the product of two integer matrices
//  Takes two matrices represented as 2D slices of integers
//  Returns the product of the two matrices if possible
//  and nil if the product is not computable
func dankalgy1(matrix1 [][]int, matrix2 [][]int) (matrix3 [][]int){
	m, n := 0, 0

	// Make sure that the two matrices have a product
	l := len(matrix1)
	if l > 0 {
		m = len(matrix1[1])
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

// function to create a randomly initialized matrix
//  accepts two integers w and h
//  returns a w by h matrix represented as a 2D slice
func rndmatrix(w int, h int) (matrix [][]int){
	matrix = make([][]int, h)
	for i := range matrix {
		matrix[i] = make([]int, w)
		for j := range matrix[i] {
			matrix[i][j] = rand.Intn(1024)
		}
	}
	return // matrix
}





