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

	//fmt.Println("Solve A:", SolveA(input))
	fmt.Println("Solve B:", SolveB(input))
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

func SolveB(input string) int {
	sum := 0

	lines := strings.Split(strings.TrimSpace(input), "\n")

	// for each A
	//	check if MMSS in any combination
	//		inc
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if i > 0 && j > 0 && i < len(lines)-1 && j < len(lines[i])-1 {
				if lines[i][j] == 'A' {
					ms := make(map[byte]int)
					ul := lines[i-1][j-1]
					ms[ul]++
					ur := lines[i-1][j+1]
					ms[ur]++
					dr := lines[i+1][j+1]
					ms[dr]++
					dl := lines[i+1][j-1]
					ms[dl]++
					if ms['M'] == 2 && ms['S'] == 2 {
						if (ul == 'M' && dr == 'M') || (ul == 'S' && dr == 'S') { // MM and SS are not adjacent
							continue
						}
						sum++
					}
				}
			}
		}
	}

	return sum
}
