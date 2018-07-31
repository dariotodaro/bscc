package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"regexp"
	"strings"
)
type coord struct {
	x int
	y int
	z int
}

type particle struct {
	p coord
	v coord
	a coord
	collided bool
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

func calcTick(vec particle) particle {
	vec.v = coord{ x: vec.v.x + vec.a.x, y: vec.v.y + vec.a.y, z: vec.v.z + vec.a.z }
	vec.p = coord{ x: vec.p.x + vec.v.x, y: vec.p.y + vec.v.y, z: vec.p.z + vec.v.z }
	return vec
}

func removeCollided(allParticles []particle) []particle{
	for i := 0 ; i < len(allParticles) ; i++ {
		foundCollision := false
		for j := i + 1 ; j < len(allParticles) ; j++ {
			if(allParticles[i] != particle{} &&
				allParticles[i].p.x == allParticles[j].p.x &&
				allParticles[i].p.y == allParticles[j].p.y &&
				allParticles[i].p.z == allParticles[j].p.z) {
				//fmt.Printf("found collision\n")	
				allParticles[j].collided = true
				foundCollision = true
			}
		}
		if(foundCollision){
			allParticles[i].collided = true
		}
	}
	for i := 0 ; i < len(allParticles) ; i++ {
		if(allParticles[i].collided) {
			allParticles = removeParticle(allParticles, i)
			i--
		}
	} 
	return allParticles
}

func removeParticle(particles []particle, index int) []particle{
	copy(particles[index:], particles[index+1:])
	return particles[:len(particles)-1]
}

func main() {
	inputfile := os.Args[1]
	ticks, err := strconv.Atoi(os.Args[2])
	check(err)

	f, err := os.Open(inputfile)
	check(err)

	reader := bufio.NewReader(f)

	allParticles := make([]particle, 1000)
	for i:=0;i<1000;i++ {
		line, err := reader.ReadBytes('\n')
		check(err)
		p, v, a := parse(string(line))
		allParticles[i] = particle{p: p, v: v, a: a}
	}
	for i := 0 ; i < ticks ; i++ {
		for j := 0 ; j < len(allParticles) ; j++{
			allParticles[j] = calcTick(allParticles[j])
		}
		allParticles = removeCollided(allParticles)
	}

	fmt.Printf("After %d ticks, %d particles are left\n", ticks, len(allParticles))
	f.Close()
}
