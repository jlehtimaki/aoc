package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	Y,X int
}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer file.Close()

	var stringArray []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stringArray = append(stringArray, scanner.Text())
	}
	return stringArray
}

func adjacents(p Point, maxY, maxX int) []Point {
	var points []Point
	if p.Y > 0 {
		points = append(points, Point{p.Y-1,p.X})
	}
	if p.Y < maxY-1 {
		points = append(points, Point{p.Y+1, p.X})
	}
	if p.X > 0 {
		points = append(points, Point{p.Y,p.X-1})
	}
	if p.X < maxX-1 {
		points = append(points, Point{p.Y, p.X+1})
	}

	return points
}

func solveProblem(input string) int {
	var count int
	lines := readLines(input)

	maze := [][]int{}
	distance := map[Point]int{}
	visited := map[Point]bool{}

	for n, line := range lines {
		maze = append(maze, []int{})
		numbers := strings.Split(line,"")
		for nn, number := range numbers {
			if n == 0 && nn == 0 {
				distance[Point{n,nn}] = 0
			} else {
				distance[Point{n,nn}] = math.MaxInt64
			}
			numberInt, _ := strconv.Atoi(number)
			maze[n] = append(maze[n], numberInt)
		}
	}

	maxY := len(maze)
	maxX := len(maze[0])

	point := Point{0,0}
	for {
		for _, x := range adjacents(point, maxY, maxX) {
			if visited[x] {
				continue
			}
			nDistance := distance[point] + maze[x.Y][x.X]
			if nDistance < distance[x] {
				distance[x] = nDistance
			}
		}

		visited[point] = true

		if visited[Point{maxY-1,maxX-1}] {
			return distance[Point{maxY-1,maxX-1}]
		}

		mDistance := math.MaxInt64
		point = Point{maxY,maxX}
		for p, v := range distance {
			if !visited[p] && v < mDistance {
				mDistance = v
				point = p
			}
		}
	}
	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}