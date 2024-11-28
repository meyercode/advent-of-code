package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := string(b)
	//fmt.Println("Solve A:", SolveA(input))
	fmt.Println("Solve B:", SolveB(input))
}

func SolveA(input string) int {
	lines := strings.Split(input, "\n")
	var xVals []int

	for _, line := range lines {
		if line == "" {
			break
		}
		var stack [][]int
		vals := strings.Split(line, " ")
		var firstLine []int
		for _, x := range vals {
			y, err := strconv.Atoi(x)
			if err != nil {
				panic(err)
			}
			firstLine = append(firstLine, y)
		}
		stack = append(stack, firstLine)
		for !IsZeroes(stack[len(stack)-1]) {
			var newLine []int
			for i := 0; i < len(stack[len(stack)-1])-1; i++ {
				newLine = append(newLine, stack[len(stack)-1][i+1]-stack[len(stack)-1][i])
			}
			stack = append(stack, newLine)
		}

		xVal := 0
		for len(stack) > 0 {
			l := stack[len(stack)-1:][0]
			stack = stack[:len(stack)-1]
			if !IsZeroes(l) {
				xVal = l[len(l)-1] + xVal
			}
		}

		xVals = append(xVals, xVal)
	}

	fmt.Println("extrapolated values:", xVals)

	sum := 0
	for _, x := range xVals {
		sum += x
	}

	return sum
}

func SolveB(input string) int {
	lines := strings.Split(input, "\n")
	var xVals []int

	for _, line := range lines {
		if line == "" {
			break
		}
		var stack [][]int
		vals := strings.Split(line, " ")
		var firstLine []int
		for _, x := range vals {
			y, err := strconv.Atoi(x)
			if err != nil {
				panic(err)
			}
			firstLine = append(firstLine, y)
		}
		stack = append(stack, firstLine)
		for !IsZeroes(stack[len(stack)-1]) {
			var newLine []int
			for i := 0; i < len(stack[len(stack)-1])-1; i++ {
				newLine = append(newLine, stack[len(stack)-1][i+1]-stack[len(stack)-1][i])
			}
			stack = append(stack, newLine)
		}

		xVal := 0
		for len(stack) > 0 {
			l := stack[len(stack)-1:][0]
			stack = stack[:len(stack)-1]
			if !IsZeroes(l) {
				xVal = l[0] - xVal
			}
		}

		xVals = append(xVals, xVal)
	}

	fmt.Println("extrapolated values:", xVals)

	sum := 0
	for _, x := range xVals {
		sum += x
	}

	return sum
}

func IsZeroes(line []int) bool {
	for _, x := range line {
		if x != 0 {
			return false
		}
	}
	return true
}
