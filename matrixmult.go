package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"os"
	"bufio"
	)

// The maximum value we will put in any of our matrices
const cellmax = 11
const max_power_of_2 = 15

func main(){
	rand.Seed(int64(time.Now().Nanosecond()))

	// Open file or writing results
	file, err := os.Create("Output.csv")
	if err != nil {
		panic(err)
	}
	// close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	// Attach file writer to buffered reader
	r := bufio.NewWriter(file)
	
	// Make sure all of our output gets written on exit
	defer func() {
		r.Flush()
	}()

	r.WriteString(" ,")
	
	var a string
	
	for i := 0; i <= max_power_of_2; i++ {
		a = fmt.Sprintf("%d,", int(math.Pow(2, float64(i))))
		r.WriteString(a)
	}

	pow := 0
	n   := 0

	seconds := 0.0
	r.WriteString("\nnaive,")
	for seconds < 180 {
		n        = int(math.Pow(2, float64(pow)))
		pow++
		fmt.Printf("Computing naive for %d by %d:\n", n, n)
		matrix1 := rndmatrix(n, n)
		matrix2 := rndmatrix(n, n)
		start   := time.Now()
		dankalgy1(matrix1, matrix2)
		elapsed := time.Since(start)
		seconds = elapsed.Seconds()
		fmt.Printf("Took %f seconds.\n", seconds)
		a = fmt.Sprintf("%d,", elapsed.Nanoseconds())
		r.WriteString(a)
	}
	fmt.Println("Done with naive")
	r.WriteString("\nStrassen's,")
	n = 0
	pow = 0
	seconds = 0.0
	for seconds < 180 {
		n        = int(math.Pow(2, float64(pow)))
		pow++
		fmt.Printf("Computing Strassen's for %d by %d:\n", n, n)
		matrix1 := rndmatrix(n, n)
		matrix2 := rndmatrix(n, n)
		start   := time.Now()
		mm_rec_init(matrix1, matrix2)
		elapsed := time.Since(start)
		seconds = elapsed.Seconds()
		fmt.Printf("Took %f seconds.\n", seconds)
		a = fmt.Sprintf("%d,", elapsed.Nanoseconds())
		r.WriteString(a)
	}
	r.WriteString("\n")
}

