package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	b, err := os.ReadFile("input")
	if err != nil {
		log.Fatal("Could not read file: ", err)
	}
	input := string(b)

	fmt.Println("Solve A:", SolveA(input))
}

func SolveA(input string) int {
	sum := 0

	// for all chars
	// if X
	//	check if XMAS in all directions
	//		+sum
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			if lines[i][j] == 'X' {
				if i > 2 && j > 2 {
					// up left
					if lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S' {
						sum++
					}
				}
				if i > 2 {
					// up
					if lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S' {
						sum++
					}
				}
				if i > 2 && j < len(lines[0])-3 { // len(lines) is 10, if X is on 6, M 7, A 8, S 9
					// up right
					if lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' {
						sum++
					}
				}
				if j < len(lines[0])-3 {
					// right
					if lines[i][j+1] == 'M' && lines[i][j+2] == 'A' && lines[i][j+3] == 'S' {
						sum++
					}
				}
				if i < len(lines)-3 && j < len(lines[0])-3 {
					// down right
					if lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
						sum++
					}
				}
				if i < len(lines)-3 {
					// down
					if lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
						sum++
					}
				}
				if i < len(lines)-3 && j > 2 {
					// down left
					if lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
						sum++
					}
				}
				if j > 2 {
					// left
					if lines[i][j-1] == 'M' && lines[i][j-2] == 'A' && lines[i][j-3] == 'S' {
						sum++
					}
				}
			}
		}
	}

	return sum
}
