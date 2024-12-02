package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("input")
	if err != nil {
		log.Fatal("can't read file", err)
	}

	input := string(b)

	//fmt.Println("Solve A:", SolveA(input))
	fmt.Println("Solve B:", SolveB(input))
}

func SolveA(input string) int {
	reports := strings.Split(input, "\n")

	valid := 0

	for _, report := range reports {
		if report == "" {
			continue
		}
		rStrings := strings.Split(report, " ")
		var readings []int
		isValid := true
		increasing := false
		decreasing := false
		for i, rString := range rStrings {
			if rString == "" {
				continue
			}
			r, err := strconv.Atoi(rString)
			if err != nil {
				log.Fatal("error parsing rstring", err)
			}
			readings = append(readings, r)
			if i > 0 {
				diff := Abs(readings[i] - readings[i-1])
				if diff > 3 || diff < 1 {
					isValid = false
					break
				}
				if readings[i] > readings[i-1] {
					increasing = true
					if decreasing {
						isValid = false
						break
					}
				} else if readings[i] < readings[i-1] {
					decreasing = true
					if increasing {
						isValid = false
						break
					}
				} else {
					isValid = false
					break
				}
			}
		}
		if isValid {
			valid++
		}
	}

	return valid
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func SolveB(input string) int {
	reports := strings.Split(input, "\n")

	valid := 0

	for _, report := range reports {
		fmt.Println("report:", report)
		if report == "" {
			continue
		}
		rStrings := strings.Split(report, " ")
		var layers []int
		for _, rString := range rStrings { // parse ints
			if rString == "" {
				continue
			}
			r, err := strconv.Atoi(rString)
			if err != nil {
				log.Fatal("error parsing rstring", err)
			}
			layers = append(layers, r)
		}

		if report == "1 3 2 4 5" {
			fmt.Println("here")
		}
		isValid := false
		if !IsValid(layers) {
			for i := 0; i < len(layers); i++ { // try removing one at a time
				d := make([]int, len(layers))
				copy(d, layers)
				d = append(d[:i], d[i+1:]...)
				if IsValid(d) {
					isValid = true
					break
				}
			}
		} else {
			isValid = true
		}
		fmt.Println("is valid:", isValid)
		if isValid {
			fmt.Println("SAFE")
			valid++
		}
	}

	return valid
}

func IsValid(layers []int) bool {
	isValid := true
	increasing := false
	decreasing := false
	for i := 0; i < len(layers); i++ { // check if sequence is valid
		if i > 0 {
			diff := Abs(layers[i] - layers[i-1])
			if diff > 3 || diff < 1 {
				isValid = false
			}
			if layers[i] > layers[i-1] {
				increasing = true
				if decreasing {
					isValid = false
				}
			} else if layers[i] < layers[i-1] {
				decreasing = true
				if increasing {
					isValid = false
				}
			} else {
				isValid = false
			}
		}
	}

	return isValid
}
