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
    count := 0

    for scanner.Scan() {
        rucksack := scanner.Text()
        compartmentA := rucksack[:len(rucksack)/2]
        compartmentB := rucksack[len(rucksack)/2:]
        var cmpmtAMap = make(map[rune]bool)
        for _, item := range compartmentA {
            cmpmtAMap[item] = true
        }
        for _, item := range compartmentB {
            _, exists := cmpmtAMap[item]
            if(exists) {
                fmt.Println(item, "already exists in other compartment! Prio is", calcPrio(item))
                repeatedItemsSum += calcPrio(item)
                break
            }
        }
        count++
    }
    // int('a')-96
    // int('A')-38
    //fmt.Println(int('L'), "rune", rune('L'), "prio:", calcPrio('L'))
    fmt.Println(count)
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
