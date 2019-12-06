package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/JakeHL/AdventOfCode2019/utils"
)

type ShipComputer struct {
	memory         []int
	programCounter int
}

func (sc *ShipComputer) Start() (exitcode int) {
	for sc.programCounter != -1 {
		sc.Iterate()
	}
	return sc.memory[0]
}

func (sc *ShipComputer) Iterate() {
	pc := sc.programCounter
	intCode := sc.memory[pc]
	intCodeAsString := fmt.Sprintf("%05d", intCode)
	newIntCode := string(intCodeAsString[3:5])
	isAImmediate := string(intCodeAsString[0]) == "1"
	isBImmediate := string(intCodeAsString[1]) == "1"
	isCImmediate := string(intCodeAsString[2]) == "1"

	switch newIntCode {
	case "01":
		position := pc + 3
		if !isAImmediate {
			position = sc.memory[pc+3]
		}
		verb := sc.memory[pc+2]
		if !isBImmediate {
			verb = sc.memory[verb]
		}
		noun := sc.memory[pc+1]
		if !isCImmediate {
			noun = sc.memory[noun]
		}
		sc.memory[position] = noun + verb
		sc.programCounter += 4
		break
	case "02":
		position := pc + 3
		if !isAImmediate {
			position = sc.memory[pc+3]
		}
		verb := sc.memory[pc+2]
		if !isBImmediate {
			verb = sc.memory[verb]
		}
		noun := sc.memory[pc+1]
		if !isCImmediate {
			noun = sc.memory[noun]
		}
		sc.memory[position] = noun * verb
		sc.programCounter += 4
		break
	case "03":
		intInput := 1
		noun := pc + 1
		if !isCImmediate {
			noun = sc.memory[noun]
		}
		sc.memory[noun] = intInput
		sc.programCounter += 2
		break
	case "04":
		noun := pc + 1
		if !isCImmediate {
			noun = sc.memory[noun]
		}
		fmt.Printf("Output: %v\n", sc.memory[noun])
		sc.programCounter += 2
		break
	case "99":
		sc.programCounter = -1
		break
	}
}

func loadRomFromFile(filename string) (rom []int) {
	fileString, err := utils.ReadTextFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	intCodeStrings := strings.Split(fileString, ",")
	rom = make([]int, len(intCodeStrings))
	for i, str := range intCodeStrings {
		intcode, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		rom[i] = intcode
	}
	return
}

func main() {
	rom := (loadRomFromFile("input.txt"))

	computer := ShipComputer{
		memory:         rom,
		programCounter: 0,
	}

	computer.Start()
}
