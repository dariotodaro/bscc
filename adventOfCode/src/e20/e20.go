package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"regexp"
	"strings"
	"math"
)
type coord struct {
	x int
	y int
	z int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func toCoord(s string) (coord) {
	sub := strings.Trim(s, "<>")
	coords := strings.Split(sub, ",")
	x, ex := strconv.Atoi(coords[0])
	y, ey := strconv.Atoi(coords[1])
	z, ez := strconv.Atoi(coords[2])
	check(ex)
	check(ey)
	check(ez)
	return coord { x: x, y: y, z: z }
}

func parse(input string) (p, v, a coord) {
	re := regexp.MustCompile("<.*?>")
	coords := re.FindAllString(input, -1)

	p = toCoord(coords[0])
	v = toCoord(coords[1])
	a = toCoord(coords[2])

	return
}

func calcManhattanDistance(c coord) int {
	return int(math.Abs(float64(c.x)) + math.Abs(float64(c.y)) + math.Abs(float64(c.z)))
}

func main() {
	inputfile := os.Args[1]
	ticks, err := strconv.Atoi(os.Args[2])
	check(err)

	f, err := os.Open(inputfile)
	check(err)

	reader := bufio.NewReader(f)
	smallestDistance := 10000000000
	smallestDistancePoint := 0

	for i:=0;i<1000;i++ {
		line, err := reader.ReadBytes('\n')
		check(err)
		p, v, a := parse(string(line))
		mV := 10000000000
		for j := 0 ; j < ticks ; j++{
			v = coord{ x: v.x + a.x, y: v.y + a.y, z: v.z + a.z }
			p = coord{ x: p.x + v.x, y: p.y + v.y, z: p.z + v.z }
		}
		mV = calcManhattanDistance(p)
		if( smallestDistance > mV)  {
			smallestDistance = mV
			smallestDistancePoint = i
		}
	}

	fmt.Printf("Smallest Distance is %d, of point No. %d, after %d ticks\n", smallestDistance, smallestDistancePoint, ticks)

	f.Close()
}
