package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	)

// The maximum value we will put in any of our matrices
const cellmax = 11

func main(){
	rand.Seed(int64(time.Now().Nanosecond()))
	l := rand.Intn(9) + 2
	m := rand.Intn(9) + 2
	n := rand.Intn(9) + 2

	matrix1 := rndmatrix(l, m)
	matrix2 := rndmatrix(m, n)

	printmatrix(matrix1)
	printmatrix(matrix2)

	// matrix3 := dankalgy1(matrix1, matrix2)

	// printmatrix(matrix3)

	// printmatrix(mm_rec_init(matrix1, matrix2))
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
