package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	//fmt.Printf("Solve A: %d\n", SolveA(lines))
	fmt.Printf("Solve B: %d\n", SolveB(lines))
}

type ConversionRange struct {
	sourceStart int64
	destStart   int64
	length      int64
}

func SolveA(inputs []string) int64 {
	seedStrings := strings.Split(strings.Split(inputs[0], ":")[1], " ")
	var seeds []int64
	for _, seedStr := range seedStrings {
		if seedStr != "" {
			seed, err := strconv.ParseInt(seedStr, 10, 64)
			if err != nil {
				panic(err)
			}
			seeds = append(seeds, seed)
		}
	}
	fmt.Println("got the seeds!")
	var seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation []ConversionRange

	for i, line := range inputs {
		if strings.Contains(line, "seed-to-soil") {
			ReadConversionRanges(inputs[i+1:], &seedToSoil)
			fmt.Println("got the seed to soil!")
		} else if strings.Contains(line, "soil-to-fertilizer") {
			ReadConversionRanges(inputs[i+1:], &soilToFert)
			fmt.Println("Got the soil to fertilizer!")
		} else if strings.Contains(line, "fertilizer-to-water") {
			ReadConversionRanges(inputs[i+1:], &fertToWater)
			fmt.Println("Got the fertilizer-to-water!")
		} else if strings.Contains(line, "water-to-light") {
			ReadConversionRanges(inputs[i+1:], &waterToLight)
			fmt.Println("Got the water-to-light!")
		} else if strings.Contains(line, "light-to-temp") {
			ReadConversionRanges(inputs[i+1:], &lightToTemp)
			fmt.Println("Got the light-to-temp!")
		} else if strings.Contains(line, "temperature-to-humidity") {
			ReadConversionRanges(inputs[i+1:], &tempToHumidity)
			fmt.Println("Got the temp-to-humidity!")
		} else if strings.Contains(line, "humidity-to-location") {
			ReadConversionRanges(inputs[i+1:], &humidityToLocation)
			fmt.Println("Got the humidity-to-location!")
		} else {
		}
	}

	minLocation := int64(math.MaxInt64)

	for _, seed := range seeds {
		fmt.Println("seed -> soil")
		soil := GetDestValue(seed, seedToSoil)
		fmt.Println("soil -> fert")
		fert := GetDestValue(soil, soilToFert)
		fmt.Println("fert -> water")
		water := GetDestValue(fert, fertToWater)
		fmt.Println("water -> light")
		light := GetDestValue(water, waterToLight)
		fmt.Println("light -> temp")
		temp := GetDestValue(light, lightToTemp)
		fmt.Println("temp -> humidity")
		humidity := GetDestValue(temp, tempToHumidity)
		fmt.Println("humidity -> location")
		location := GetDestValue(humidity, humidityToLocation)
		if location < minLocation {
			minLocation = location
		}
	}

	if minLocation == math.MaxInt64 {
		log.Fatalf("no min location")
	}

	return minLocation
}

type Range struct {
	start int64
	end   int64
}

func SolveB(inputs []string) int64 {
	seedStrings := strings.Split(strings.Split(inputs[0], ":")[1], " ")
	var seeds []int64
	for _, seedStr := range seedStrings {
		if seedStr != "" {
			seed, err := strconv.ParseInt(seedStr, 10, 64)
			if err != nil {
				panic(err)
			}
			seeds = append(seeds, seed)
		}
	}
	var seedRanges []Range
	for i := 0; i < len(seeds); i += 2 {
		seedRanges = append(seedRanges, Range{
			start: seeds[i],
			end:   seeds[i] + seeds[i+1],
		})
	}
	var seedToSoil, soilToFert, fertToWater, waterToLight, lightToTemp, tempToHumidity, humidityToLocation []ConversionRange
	var categoryMap [][]ConversionRange

	for i, line := range inputs {
		if strings.Contains(line, "seed-to-soil") {
			ReadConversionRanges(inputs[i+1:], &seedToSoil)
			categoryMap = append(categoryMap, seedToSoil)
		} else if strings.Contains(line, "soil-to-fertilizer") {
			ReadConversionRanges(inputs[i+1:], &soilToFert)
			categoryMap = append(categoryMap, soilToFert)
		} else if strings.Contains(line, "fertilizer-to-water") {
			ReadConversionRanges(inputs[i+1:], &fertToWater)
			categoryMap = append(categoryMap, fertToWater)
		} else if strings.Contains(line, "water-to-light") {
			ReadConversionRanges(inputs[i+1:], &waterToLight)
			categoryMap = append(categoryMap, waterToLight)
		} else if strings.Contains(line, "light-to-temp") {
			ReadConversionRanges(inputs[i+1:], &lightToTemp)
			categoryMap = append(categoryMap, lightToTemp)
		} else if strings.Contains(line, "temperature-to-humidity") {
			ReadConversionRanges(inputs[i+1:], &tempToHumidity)
			categoryMap = append(categoryMap, tempToHumidity)
		} else if strings.Contains(line, "humidity-to-location") {
			ReadConversionRanges(inputs[i+1:], &humidityToLocation)
			categoryMap = append(categoryMap, humidityToLocation)
		} else {
		}
	}

	// for each categorymap
	// while ranges > 0
	// check for intersections with next categorymap
	// add new ranges and missing ranges

	fmt.Printf("Starting ranges: %v\n", seedRanges)
	for _, category := range categoryMap {
		//fmt.Printf("Category ranges: %v\n", category)
		var newRanges []Range
		//fmt.Printf("Using ranges: %v\n", seedRanges)

		for len(seedRanges) > 0 {
			seed := seedRanges[len(seedRanges)-1]
			seedRanges = seedRanges[:len(seedRanges)-1]
			//fmt.Printf("Popped seed %v now have remaining %v\n", seed, seedRanges)
			hasMapping := false
			for _, r := range category {
				overlapStart := max(seed.start, r.sourceStart)
				overlapEnd := min(seed.end, r.sourceStart+r.length)
				if overlapStart < overlapEnd {
					newRanges = append(newRanges, Range{
						start: overlapStart - r.sourceStart + r.destStart,
						end:   overlapEnd - r.sourceStart + r.destStart,
					})
					//fmt.Printf("New ranges: %v\n", newRanges)
					if overlapStart > seed.start {
						seedRanges = append(seedRanges, Range{
							start: seed.start,
							end:   overlapStart,
						})
					}
					if seed.end > overlapEnd {
						seedRanges = append(seedRanges, Range{
							start: overlapEnd,
							end:   seed.end,
						})
					}
					hasMapping = true
				}
			}
			if !hasMapping {
				newRanges = append(newRanges, Range{
					start: seed.start,
					end:   seed.end,
				})
			}
		}
		seedRanges = make([]Range, len(newRanges))
		copy(seedRanges, newRanges)
		//fmt.Printf("copied newRanges to %v\n", seedRanges)
	}

	minLocation := int64(math.MaxInt64)

	for _, r := range seedRanges {
		if r.start < minLocation {
			minLocation = r.start
		}
	}

	return minLocation
}

func ReadConversionRanges(inputs []string, categoryRanges *[]ConversionRange) {
	for _, line := range inputs {
		if line == "" {
			return
		}
		valStrings := strings.Split(line, " ")
		source, err := strconv.ParseInt(valStrings[1], 10, 64)
		dest, err := strconv.ParseInt(valStrings[0], 10, 64)
		length, err := strconv.ParseInt(valStrings[2], 10, 64)
		if err != nil {
			panic(err)
		}

		*categoryRanges = append(*categoryRanges, ConversionRange{
			sourceStart: source,
			destStart:   dest,
			length:      length,
		})
	}
}

func GetDestValue(sourceVal int64, categoryRanges []ConversionRange) int64 {
	for _, r := range categoryRanges {
		fmt.Printf("Check if %d is between %d and %d\n", sourceVal, r.sourceStart, (r.sourceStart + r.length))
		if sourceVal >= r.sourceStart && sourceVal < r.sourceStart+r.length {
			fmt.Printf("using val %d\n", r.destStart+(sourceVal-r.sourceStart))
			return r.destStart + (sourceVal - r.sourceStart)
		}
	}

	return sourceVal
}
