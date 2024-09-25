package main

import "fmt"

// func filling([][]int) [][]int{
// 	for
// } 
func main() {
	matrix := make([][]int, 10)
for i:=0; i<10; i++ {
	matrix[i]=make([]int, 10)
	// for j:=0; j<10; j++{
		
	// }
	matrix[i][i]=i
	fmt.Println(matrix[i])
}
}