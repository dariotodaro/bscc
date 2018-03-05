package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	var inputLength = len(input)
	var intArray = make([]int, inputLength)
	var current int
	var sum int
	for i := 0; i < inputLength; i++ {
		c, err := strconv.Atoi(string(input[i%inputLength]))
		if err != nil {
			fmt.Printf("Error in string convertion")
		} else {
			intArray[i] = c
		}
	}
	for i := 0; i < inputLength; i++ {
		current = intArray[i]
		var next = intArray[(i+1)%inputLength]
		if current == next {
			sum += next
		}
	}
	fmt.Printf("\nSum is: %v\n\n", sum)
}
