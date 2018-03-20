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
	var checksum int
	for i := 0; i < len(inputLines); i++ {
		var input = strings.Split(inputLines[i], " ")
		var intArr = stringArrToIntArr(input)
		var max, min = 0, 100000000
		for j := 0; j < len(intArr); j++ {
			var current = intArr[j]
			if current < min {
				min = current
			}
			if current > max {
				max = current
			}
		}
		//		fmt.Printf("max: %v, min: %v\n", max, min)
		checksum += (max - min)
		max, min = 0, 0
	}
	fmt.Printf("\nChecksum is: %v\n\n", checksum)
}
func stringArrToIntArr(input []string) []int {
	var intArray = make([]int, len(input))
	for i := 0; i < len(input); i++ {
		c, err := strconv.Atoi(string(strings.Trim(input[i], "^M\n ")))
		if err != nil {
			fmt.Printf("Error in string convertion\n")
		} else {
			intArray[i] = c
		}
	}
	return intArray
}
