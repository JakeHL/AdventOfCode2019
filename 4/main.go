package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/JakeHL/AdventOfCode2019/utils"
)

func getRangeFromFile(filename string) (bottomRange, topRange int) {
	inputFile, err := utils.ReadTextFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	inputStrings := strings.Split(inputFile, "-")
	bottomRange, err = strconv.Atoi(inputStrings[0])
	if err != nil {
		log.Fatal(err)
	}
	topRange, err = strconv.Atoi(inputStrings[1])
	if err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	bottomRange, topRange := getRangeFromFile("input.txt")
	viablePasswords := 0
	for i := bottomRange; i < topRange; i++ {
		isIncremental, hasDoubleChar := false, false

		numberAsString := strconv.FormatInt(int64(i), 10)
		fmt.Printf("trying: %v", numberAsString)

		for charIndex, currentChar := range numberAsString {
			if charIndex != 0 {
				lastChar := numberAsString[charIndex-1]
				lastCharAsInt, _ := strconv.Atoi(string(lastChar))
				currentCharAsInt, _ := strconv.Atoi(string(currentChar))
				if !(lastCharAsInt <= currentCharAsInt) {
					isIncremental = false
					break
				}
				isIncremental = true
			}

			doubleChar := string([]rune{currentChar, currentChar})
			doubleCharExists := strings.Contains(numberAsString, doubleChar)
			if doubleCharExists {
				trippleChar := string([]rune{currentChar, currentChar, currentChar})
				doubleIsTripple := strings.Contains(numberAsString, trippleChar)
				quadChar := string([]rune{currentChar, currentChar, currentChar, currentChar})
				doubleIsQuad := strings.Contains(numberAsString, quadChar)
				doubleCharExists = !doubleIsTripple && !doubleIsQuad
			}
			hasDoubleChar = hasDoubleChar || doubleCharExists
		}

		if isIncremental && hasDoubleChar {
			viablePasswords++
			fmt.Print(" SUCCESS \n")
		} else {
			fmt.Print("\n")
		}
	}
	fmt.Print(viablePasswords)
}
