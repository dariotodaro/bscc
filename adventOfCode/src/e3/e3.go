package main

import (
	"fmt"
	"os"
	"strconv"
	//"strings"
)

func main() {
	rawInput := os.Args[1]
	c, err := strconv.Atoi(rawInput)
	if err != nil {
		fmt.Printf("Error in string convertion\n")
	} else {
		var dim = calcDimension(c)
		fmt.Println("Dimension:", dim)
		createMatrix(dim)
		fmt.Println("Done!")
	}
}

func createMatrix(dim int) [][]int {
	var matrix = make([][]int, dim)
	for i := 0; i > dim; i++ {
		matrix[i] = make([]int, dim)
	}
	return matrix
}

func calcDimension(i int) int {
	for dim := 1; ; dim++ {
		if dim*dim > i && dim%2 == 1 {
			return dim
		}
	}
}
