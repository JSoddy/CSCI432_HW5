package main

import "fmt"

func printmatrix(s [][]int){
	for i:= 0; i < 2; i++ {
		for j:= 0; j < 2; j++ {
		
			fmt.Print(s[i][j], " ")
			}
			fmt.Println()
			}
			
	
}

func main(){
	matrix1:= [][]int{
		[]int{2,2},
		[]int{1,1},
		
		}
	matrix2:= [][]int{
		[]int{3,3},
		[]int{2,2},
		}
		
		printmatrix(matrix1)
		printmatrix(matrix2)
		dankalgy1(matrix1, matrix2)
		
 
}
func dankalgy1(matrix1 [][]int, matrix2 [][]int){


	matrix3:= [][]int{
		[]int{0,0},
		[]int{0,0},
		}
		
		for i:= 0; i < 2; i++ {
			for j:= 0; j < 2; j++ {
				for k:= 0; k < 2; k++ {
					matrix3[i][j] = matrix3[i][j] + matrix1[i][k] * matrix2[k][j]
				}
			}
		}
		printmatrix(matrix3)
		
}







