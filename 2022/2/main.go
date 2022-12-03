package main

import (
	"bufio"
	"fmt"
	"os"
)

// ROCK     A X 1
// PAPER    B Y 2
// SCISSORS C Z 3

func calculateRoundOne(round string) int {
	scoring := map[string]int{
		"A X": 4, "A Y": 8, "A Z": 3,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 7, "C Y": 2, "C Z": 6,
	}
	return scoring[round]
}

func calculateRoundTwo(round string) int {
	scoring := map[string]int{
		"A X": 3, "A Y": 4, "A Z": 8,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 2, "C Y": 6, "C Z": 7,
	}
	return scoring[round]
}

func main() {
	file, _ := os.Open("input.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("oops")
		}
	}(file)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	totalPointsPartOne := 0
	totalPointsPartTwo := 0

	for fileScanner.Scan() {
		round := fileScanner.Text()
		totalPointsPartOne += calculateRoundOne(round)
		totalPointsPartTwo += calculateRoundTwo(round)
	}

	fmt.Printf("Total score of my games for part one is %d\n", totalPointsPartOne)
	fmt.Printf("Total score of my games for part two is %d\n", totalPointsPartTwo)
}
