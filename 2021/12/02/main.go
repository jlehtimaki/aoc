package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Path struct {
	Name	string
	Visited	map[string]int
	Revisit	bool
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

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func solveProblem(input string) int {
	var count int
	nodes := map[string][]string{}
	lines := readLines(input)

	for _, line := range lines {
		lineSplit := strings.Split(line,"-")
		nodes[lineSplit[0]] = append(nodes[lineSplit[0]], lineSplit[1])
		nodes[lineSplit[1]] = append(nodes[lineSplit[1]], lineSplit[0])
	}

	start := Path{
		Name:    "start",
		Visited: map[string]int{"start": 0},
		Revisit: false,
	}
	queues := []Path{start}

	for len(queues) > 0 {
		queue := queues[0]
		if queue.Name == "end" {
			count = count + 1
			queues = queues[1:]
			continue
		}
		for _, x := range nodes[queue.Name] {
			if _, ok := queue.Visited[x]; !ok {
				visited := map[string]int{}
				for k,v := range queue.Visited {
					visited[k] = v
				}
				if IsLower(x) {
					visited[x] = 0
				}
				queues = append(queues, Path{
					Name:    x,
					Visited: visited,
					Revisit: false,
				})
				continue
			} else if _, ok := queue.Visited[x]; ok{
				if !queue.Revisit && x != "start" && x != "end" {
					visited := map[string]int{}
					for k,v := range queue.Visited {
						visited[k] = v
					}
					queues = append(queues, Path{
						Name:    x,
						Visited: visited,
						Revisit: true,
					})
				}
			}
		}
		queues = queues[1:]

	}
	return count
}

func main()  {
	fmt.Println(solveProblem("input.txt"))
}