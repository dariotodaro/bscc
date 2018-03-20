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
		for j := 0; j < len(intArr); j++ {
			var d = findDivideable(j, intArr)
			if d != -1 {
				var current = intArr[j]
				if d > current {
					checksum += d / current
					fmt.Printf("d: %v\n", d/current)
				} else {
					checksum += current / d
					fmt.Printf("d: %v\n", current/d)
				}
				break
			}
		}
	}
	fmt.Printf("\nChecksum is: %v\n\n", checksum)
}

func findDivideable(pos int, arr []int) int {
	var d = arr[pos]
	for i := 0; i < len(arr); i++ {
		if i != pos && isDivideable(d, arr[i]) {
			return arr[i]
		}
	}
	return -1
}

func isDivideable(a, b int) bool {
	return a%b == 0
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
