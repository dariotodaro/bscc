package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rawInput := os.Args[1]
	fmt.Printf("rawInput: %v\n\n", rawInput)
	var inputLines = strings.Split(rawInput, "\n")
	var input = make([]string, 0)
	for i := 0; i < len(inputLines); i++ {
		input = append(input, strings.Split(inputLines[i], " ")...)
	}
	var inputLength = len(input)
	var intArray = make([]int, inputLength)
	var current int
	var sum int
	fmt.Printf("inputlength: %v\n\n", inputLength)
	for i := 0; i < inputLength; i++ {
		fmt.Printf("Try to convert: %v\n", input[i])
		c, err := strconv.Atoi(string(strings.Trim(input[i%inputLength], "^M\n ")))
		if err != nil {
			fmt.Printf("Error in string convertion\n")
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
