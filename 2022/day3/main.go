package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open("./input")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    repeatedItemsSum := 0
    groupCount := 0
    group := make(map[int]map[rune]bool)
    groupMap := make(map[rune]int)

    for scanner.Scan() {
        rucksack := scanner.Text()
        group[groupCount] = make(map[rune]bool)
        for _, item := range rucksack {
            group[groupCount][item] = true
        }
        groupCount++
        if(groupCount == 3) {
            for _, m := range group {
                for item := range m {
                    groupMap[item] += 1
                    if (groupMap[item] == 3) {
                        repeatedItemsSum += calcPrio(item)
                    }
                }
            }
            groupMap = make(map[rune]int)
            groupCount = 0
        }
    }
    // int('a')-96
    // int('A')-38
    //fmt.Println(int('L'), "rune", rune('L'), "prio:", calcPrio('L'))
    fmt.Println(repeatedItemsSum)

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func calcPrio(item rune) int {
    value := int(item)
    if(value <= 90) {
        return value-38
    } else {
        return value-96
    }
}
