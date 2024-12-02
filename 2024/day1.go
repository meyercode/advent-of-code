package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("input")
	if err != nil {
		log.Fatal("could not read file", err)
	}

	input := string(b)

	//fmt.Println("Solve A:", SolveA(input))
	fmt.Println("Solve B:", SolveB(input))
}

func SolveA(input string) int {
	lines := strings.Split(input, "\n")

	var a, b []int
	for _, l := range lines {
		if l != "" {
			vals := strings.Split(l, "   ")
			x, err := strconv.Atoi(vals[0])
			y, err := strconv.Atoi(vals[1])
			if err != nil {
				log.Fatal("error parsing int", err)
			}
			a = append(a, x)
			b = append(b, y)
		}
	}

	slices.Sort(a)
	slices.Sort(b)

	sum := 0

	for i := range a {
		sum += Abs(a[i] - b[i])
	}

	return sum
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SolveB(input string) int {
	lines := strings.Split(input, "\n")

	bFreq := make(map[int]int)
	var a []int
	for _, l := range lines {
		if l != "" {
			vals := strings.Split(l, "   ")
			x, err := strconv.Atoi(vals[0])
			y, err := strconv.Atoi(vals[1])
			if err != nil {
				log.Fatal("error parsing int", err)
			}

			a = append(a, x)
			bFreq[y]++
		}
	}

	score := 0

	for _, v := range a {
		score += v * bFreq[v]
	}

	return score
}
