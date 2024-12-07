package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type Guard struct {
	pos       Pos
	distinct  map[Pos]int
	room      [][]string
	direction string
}

type Pos struct {
	x int
	y int
}

const loopThreshold = 4

var (
	mu sync.Mutex
	wg sync.WaitGroup
)

func readLines(path string) Guard {
	var guard Guard
	guard.distinct = make(map[Pos]int)
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var stringArray [][]string
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		stringArray = append(stringArray, []string{})
		for x, c := range scanner.Text() {
			if c == '^' {
				guard.pos = Pos{x: x, y: i}
				guard.distinct[guard.pos] = 1
			}
			stringArray[i] = append(stringArray[i], string(c))
		}
		i++
	}
	guard.room = stringArray

	return guard
}

func (g *Guard) moveUp() {
	if g.room[g.pos.y-1][g.pos.x] == "#" {
		g.direction = "right"
		return
	}
	g.pos.y--
	g.distinct[g.pos] += 1
}

func (g *Guard) moveRight() {
	if g.room[g.pos.y][g.pos.x+1] == "#" {
		g.direction = "down"
		return
	}
	g.pos.x++
	g.distinct[g.pos] += 1
}

func (g *Guard) moveDown() {
	if g.room[g.pos.y+1][g.pos.x] == "#" {
		g.direction = "left"
		return
	}
	g.pos.y++
	g.distinct[g.pos] += 1
}

func (g *Guard) moveLeft() {
	if g.room[g.pos.y][g.pos.x-1] == "#" {
		g.direction = "up"
		return
	}
	g.pos.x--
	g.distinct[g.pos] += 1
}

func (g *Guard) checkEdges() bool {
	// check top row
	if (g.pos.x >= 0 && g.pos.x < len(g.room[0])) && (g.pos.y == 0 || g.pos.y == len(g.room)-1) {
		return true
	}

	// check left and right edges
	if (g.pos.y >= 0 && g.pos.y < len(g.room)) && (g.pos.x == 0 || g.pos.x == len(g.room[0])-1) {
		return true
	}

	return false
}

func deepCopy(src [][]string) [][]string {
	dest := make([][]string, len(src))
	for i := range src {
		dest[i] = make([]string, len(src[i]))
		copy(dest[i], src[i])
	}
	return dest
}

func goMove(g Guard, x, y int) bool {
	g.room[y][x] = "#"
	g.direction = "up"
	for {
		if g.checkEdges() {
			return false
		}
		if g.distinct[Pos{x: g.pos.x, y: g.pos.y}] > loopThreshold {
			return true
		}
		if g.direction == "up" {
			g.moveUp()
		}
		if g.direction == "right" {
			g.moveRight()
		}
		if g.direction == "down" {
			g.moveDown()
		}
		if g.direction == "left" {
			g.moveLeft()
		}
	}
}

func s1(input string) (int, map[Pos]int) {
	g := readLines(input)
	g.direction = "up"

	for {
		if g.checkEdges() {
			break
		}
		if g.direction == "up" {
			g.moveUp()
		}
		if g.direction == "right" {
			g.moveRight()
		}
		if g.direction == "down" {
			g.moveDown()
		}
		if g.direction == "left" {
			g.moveLeft()
		}
	}

	return len(g.distinct), g.distinct
}

func s2(input string, d map[Pos]int) int {
	var count int
	g := readLines(input)

	startingPos := g.pos

	for k := range d {
		x, y := k.x, k.y
		// Make a deep copy of the room for this goroutine
		copyRoom := deepCopy(g.room)

		// Create a copy of g for this goroutine
		localG := g
		localG.room = copyRoom
		localG.distinct = map[Pos]int{startingPos: 1}

		wg.Add(1)
		go func(x, y int, guardCopy Guard) {
			defer wg.Done()
			// Use the copy of g and do not modify the original g
			if goMove(guardCopy, x, y) {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}(x, y, localG)
	}

	wg.Wait()

	return count
}

func main() {
	lngth, distinct := s1("input.txt")
	fmt.Printf("P1: %d \n", lngth)
	fmt.Printf("P2: %d \n", s2("input.txt", distinct))
}
