package main

import (
	"fmt"
	"math"
	"math/rand"
	)



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