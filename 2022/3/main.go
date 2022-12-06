package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func prepareScanner() *bufio.Scanner {
	file, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}

func convertLetter(letter rune) int {
	if unicode.IsLower(letter) {
		return int(letter) - 96
	} else {
		return int(letter) - 38
	}
}

func findCommonInStrings(words []string) []rune {
	var res []rune
	main := make(map[rune]int)

	for _, letter := range words[0] {
		main[letter]++
	}

	words = words[1:]

	for _, word := range words {
		temp := make(map[rune]int)

		for _, letter := range word {
			if main[letter] > 0 {
				temp[letter]++
				main[letter]--
			}
		}

		main = temp
	}

	for key, value := range main {
		for i := 0; i < value; i++ {
			res = append(res, key)
		}
	}

	return res
}

func partOne() int {
	fileScanner := prepareScanner()

	var sum int

	for fileScanner.Scan() {
		line := fileScanner.Text()
		rucksack := []string{line[:len(line)/2], line[len(line)/2:]}
		commonLetter := findCommonInStrings([]string{rucksack[0], rucksack[1]})[0]
		sum += convertLetter(commonLetter)
	}
	return sum
}

func partTwo() int {
	fileScanner := prepareScanner()

	var sum int
	var counter int
	var group []string
	for fileScanner.Scan() {
		group = append(group, fileScanner.Text())

		if counter%3 == 2 {
			commonLetter := findCommonInStrings(group)[0]
			sum += convertLetter(commonLetter)
			group = nil
		}
		counter++
	}

	return sum
}

func main() {

	fmt.Printf("Sum of priorities in part one is %d\n", partOne())
	fmt.Printf("Sum of priorities in part two is %d\n", partTwo())
}
