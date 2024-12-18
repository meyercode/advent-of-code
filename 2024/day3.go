package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
	pattern := `mul\(\d{1,3},\d{1,3}\)`
	r := regexp.MustCompile(pattern)

	matches := r.FindAllString(input, -1)
	fmt.Println(matches)

	sum := 0
	for _, m := range matches {
		pattern := `\d{1,3},\d{1,3}`
		r := regexp.MustCompile(pattern)
		stripped := r.FindString(m)
		fStrings := strings.Split(stripped, ",")
		a, err := strconv.Atoi(fStrings[0])
		b, err := strconv.Atoi(fStrings[1])
		if err != nil {
			log.Fatal("problem parsing int: ", err)
		}
		sum += a * b
	}

	return sum
}

func SolveB(input string) int {
	pattern := `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`
	r := regexp.MustCompile(pattern)

	matches := r.FindAllString(input, -1)
	fmt.Println(matches)

	sum := 0
	canMult := true
	for _, m := range matches {
		if m == "do()" {
			canMult = true
			continue
		} else if m == "don't()" {
			canMult = false
			continue
		}
		if canMult {
			pattern := `\d{1,3},\d{1,3}`
			r := regexp.MustCompile(pattern)
			stripped := r.FindString(m)
			fStrings := strings.Split(stripped, ",")
			a, err := strconv.Atoi(fStrings[0])
			b, err := strconv.Atoi(fStrings[1])
			if err != nil {
				log.Fatal("problem parsing int: ", err)
			}
			sum += a * b
		}
	}

	return sum
}
