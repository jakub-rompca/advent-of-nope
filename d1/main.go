package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type ElfWithBag struct {
	number        int
	caloriesCount int
}

func main() {
	file, _ := os.Open("./d1/input.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("oops, could not read the file")
		}
	}(file)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	arrCount := 0
	elfCounter := 1

	var hordeOfElves []ElfWithBag
	var currentElf ElfWithBag
	isNewElf := true

	for fileScanner.Scan() {
		if isNewElf == true {
			currentElf = ElfWithBag{
				number:        elfCounter,
				caloriesCount: 0,
			}
			isNewElf = false
		}
		value, err := strconv.Atoi(fileScanner.Text())
		if err == nil {
			currentElf.caloriesCount += value
		}
		if err != nil {
			elfCounter++
			arrCount++
			isNewElf = true
			hordeOfElves = append(hordeOfElves, currentElf)
		}
	}

	sort.Slice(hordeOfElves, func(i, j int) bool {
		return hordeOfElves[i].caloriesCount > hordeOfElves[j].caloriesCount
	})

	sumOfFirstThree := 0

	for counter := 0; counter < 3; counter++ {
		sumOfFirstThree += hordeOfElves[counter].caloriesCount
	}
}
