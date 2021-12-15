package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Output struct {
	Pair1, Pair2 string
	NewElement   string
}

type Step struct {
	Pairs map[string]int
	Count map[string]int
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


func solveProblem(input string) int {
	data := readLines(input)

	rules := map[string]Output{}

	for _, rule := range data[2:] {
		var in, out string
		fmt.Sscanf(rule, "%s -> %s", &in, &out)

		rules[in] = Output{
			Pair1:      string(in[0]) + out,
			Pair2:      out + string(in[1]),
			NewElement: out,
		}
	}

	// 2. convert input polymer to internal format
	s := NewStep(data[0])
	fmt.Println(s)

	// 3. execute 40 steps
	for i := 0; i < 40; i++ {
		s = next(s, rules)
	}

	// 4. find the min and max counts
	min := math.MaxInt64
	max := 0
	for _, v := range s.Count {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	// 5. print out your answer
	return max - min
}
func main() {
	fmt.Println(solveProblem("input.txt"))
}

func next(in *Step, rules map[string]Output) *Step {
	// 0. init a new struct to hold the next step
	out := &Step{
		Pairs: map[string]int{},
		Count: map[string]int{},
	}

	// 1. copy all existing counts over
	for k, v := range in.Count {
		out.Count[k] = v
	}

	// 2. pairs don't get copied over -> they are expanded into different pairs
	// for each pair, increment the count of the new pairs that it becomes,
	// and also increment the count for the new element that gets inserted
	for p, count := range in.Pairs {
		t := rules[p]
		out.Pairs[t.Pair1] += count
		out.Pairs[t.Pair2] += count
		out.Count[t.NewElement] += count
	}

	return out
}

func NewStep(in string) *Step {
	s := &Step{
		Pairs: map[string]int{}, // pairs in the polymer and how many of each
		Count: map[string]int{}, // elements in the polymer and how many of each
	}

	// 1. take every pair in the polymer and add it to Pairs
	for i := 0; i < len(in)-1; i++ {
		pair := in[i : i+2]
		s.Pairs[pair]++
	}

	// 2. count every element in the polymer and add it to Counts
	// (you could do this with the previous loop)
	for i := 0; i < len(in); i++ {
		s.Count[string(in[i])]++
	}

	return s
}