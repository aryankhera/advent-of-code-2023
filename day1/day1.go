package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInputFile() []string {
	bs, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	words := strings.Split(string(bs), "\n")
	return words
}

func part1(words []string) int {
	sum := 0
	for _, word := range words {
		firstDigit := ""
		lastDigit := ""
		fmt.Println(word)
		for i := 0; i <= len(word)-1; i++ {
			if word[i] <= 57 && word[i] >= 48 {
				firstDigit = string(word[i])
				break
			}
		}
		for i := len(word) - 1; i >= 0; i-- {
			if word[i] <= 57 && word[i] >= 48 {
				lastDigit = string(word[i])
				break
			}
		}
		num, _ := strconv.Atoi(firstDigit + lastDigit)
		fmt.Println(num)
		sum += num
	}
	return sum
}

func part2(words []string) int {
	sum := 0
	var digits = map[string]int{}
	digits["zero"] = 0
	digits["one"] = 1
	digits["two"] = 2
	digits["three"] = 3
	digits["four"] = 4
	digits["five"] = 5
	digits["six"] = 6
	digits["seven"] = 7
	digits["eight"] = 8
	digits["nine"] = 9

	var strDigits = map[string][]string{}
	for key, _ := range digits {
		vals := strDigits[string(key[0])]
		if vals == nil {
			vals = []string{}
			strDigits[string(key[0])] = vals
		}
		strDigits[string(key[0])] = append(vals, key)
	}
	for _, word := range words {
		firstDigit := ""
		lastDigit := ""
		fmt.Println(word)
		for i := 0; i <= len(word)-1; i++ {
			if strDigits[string(word[i])] != nil {
				for _, digit := range strDigits[string(word[i])] {
					if len(digit)+i <= len(word) && digit == word[i:len(digit)+i] {
						firstDigit = strconv.Itoa(digits[digit])
						break
					}
				}
				if firstDigit != "" {
					break
				}
			}
			if word[i] <= 57 && word[i] >= 48 {
				firstDigit = string(word[i])
				break
			}
		}
		for i := len(word) - 1; i >= 0; i-- {
			if strDigits[string(word[i])] != nil {
				for _, digit := range strDigits[string(word[i])] {
					if len(digit)+i <= len(word) && digit == word[i:len(digit)+i] {
						lastDigit = strconv.Itoa(digits[digit])
						break

					}
				}
				if lastDigit != "" {
					break
				}
			}
			if word[i] <= 57 && word[i] >= 48 {
				lastDigit = string(word[i])
				break
			}
		}
		num, _ := strconv.Atoi(firstDigit + lastDigit)
		fmt.Println(num)
		sum += num
	}
	return sum
}

func main() {

	words := readInputFile()
	//fmt.Println("Part1 Sum: ", part1(words))
	fmt.Println("Part2 Sum: ", part2(words))

}
