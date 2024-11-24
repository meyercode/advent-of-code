package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(SolveA(lines))
	fmt.Println(SolveB(lines))
}

func SolveA(inputs []string) int {
	fmt.Printf("Have %d cards\n", len(inputs))

	var scores []int

	for _, card := range inputs {
		numbers := strings.TrimSpace(strings.Split(card, ":")[1])
		numbersHeld := strings.Split(strings.TrimSpace(strings.Split(numbers, "|")[0]), " ")
		winningNumbers := strings.Split(strings.TrimSpace(strings.Split(numbers, "|")[1]), " ")
		winningNums := make(map[int]int)
		for _, n := range winningNumbers {
			if n == "" {
				continue
			}
			winNum, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			winningNums[winNum] = 0
		}
		score := 0
		for _, n := range numbersHeld {
			if n == "" {
				continue
			}
			num, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			if _, ok := winningNums[num]; ok {
				if score == 0 {
					score += 1
				} else {
					score *= 2
				}
			}
		}
		scores = append(scores, score)
	}

	scoreSum := 0
	for _, score := range scores {
		scoreSum += score
	}

	return scoreSum
}

func SolveB(inputs []string) int {

	cardCopies := make([]int, len(inputs))
	for i := range cardCopies {
		cardCopies[i] = 1
	}

	for i, card := range inputs {
		numbers := strings.TrimSpace(strings.Split(card, ":")[1])
		numbersHeld := strings.Split(strings.TrimSpace(strings.Split(numbers, "|")[0]), " ")
		winningNumbers := strings.Split(strings.TrimSpace(strings.Split(numbers, "|")[1]), " ")
		winningNums := make(map[int]int)
		for _, n := range winningNumbers {
			if n == "" {
				continue
			}
			winNum, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			winningNums[winNum] = 0
		}

		copies := 0
		for _, n := range numbersHeld {
			if n == "" {
				continue
			}
			num, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			if _, ok := winningNums[num]; ok {
				copies++
			}
		}

		// add the copies
		for j := i + 1; j <= i+copies; j++ {
			if i == 199 {
				fmt.Println("here")
			}
			if j < len(inputs) {
				cardCopies[j] += cardCopies[i]
			}
		}
	}

	copySum := 0
	for _, copy := range cardCopies {
		copySum += copy
	}

	return copySum
}
