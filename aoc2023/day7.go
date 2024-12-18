package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := string(b)

	fmt.Printf("Solve B: %d\n", SolveB(input))
}

type ByHand []string

func (h ByHand) Len() int      { return len(h) }
func (h ByHand) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h ByHand) Less(i, j int) bool {
	if h[i] == "" || h[j] == "" {
		return false
	}
	a := strings.Split(h[i], " ")[0]
	b := strings.Split(h[j], " ")[0]

	ah := readHand(a)
	bh := readHand(b)

	if ah == bh {
		for i, _ := range a {
			if a[i] != b[i] { // continue until suits difer
				return CompareSuits(a[i], b[i])
			} else {
				if i == 4 { // if all cards the same, don't swap hands
					fmt.Printf("Last card, not swapping %c and %c\n", a[i], b[i])
					return false
				}
			}
		}
	}

	return ah < bh
}

type Suit int

const (
	T Suit = iota + 10
	J
	Q
	K
	A
)

func CompareSuits(a, b byte) bool {
	return ReadSuit(a) < ReadSuit(b)
}

func ReadSuit(b byte) int {
	suit := 0
	switch b {
	case 'A':
		suit = int(A)
	case 'K':
		suit = int(K)
	case 'Q':
		suit = int(Q)
	case 'J':
		suit = 1
	case 'T':
		suit = int(T)
	default:
		suit = int(b) - 48
	}

	return suit
}

type Hand int

const (
	HighCard Hand = iota + 100
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func readHand(hand string) int {
	valueMap := make(map[rune]int)
	for _, ch := range hand {
		valueMap[ch]++
	}
	t := HighCard

	var keys []rune
	for k, _ := range valueMap {
		keys = append(keys, k)
	}
	mvs := keys[0]
	if mvs == 'J' && len(keys) > 1 {
		mvs = keys[1]
	}
	for _, k := range keys {
		if valueMap[k] > valueMap[mvs] && k != 'J' {
			mvs = k
		}
	}
	// add joker to highest suit
	if mvs != 'J' {
		valueMap[mvs] += valueMap['J']
	}
	if valueMap[mvs] == 5 {
		t = FiveOfAKind
	} else if valueMap[mvs] == 4 {
		t = FourOfAKind
	} else {
		for _, k := range keys {
			if k != 'J' {
				switch valueMap[k] {
				case 3:
					if t == OnePair {
						t = FullHouse
					} else {
						t = ThreeOfAKind
					}
				case 2:
					if t == ThreeOfAKind {
						t = FullHouse
					} else if t == OnePair {
						t = TwoPairs
					} else {
						t = OnePair
					}
				default:
				}
			}
		}
	}

	return int(t)
}

func SolveB(input string) int {
	hands := strings.Split(input, "\n")
	//fmt.Printf("%v\n", hands)

	sort.Sort(ByHand(hands))

	fmt.Printf("%v\n", hands)

	score := 0
	for i, hand := range hands {
		if hand == "" {
			break
		}
		bid, err := strconv.Atoi(strings.Split(hand, " ")[1])
		if err != nil {
			panic(err)
		}

		score += (i + 1) * bid
	}

	return score
}
