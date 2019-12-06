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
		param3 := pc + 3
		if !isAImmediate {
			param3 = sc.memory[pc+3]
		}
		param2 := sc.memory[pc+2]
		if !isBImmediate {
			param2 = sc.memory[param2]
		}
		param1 := sc.memory[pc+1]
		if !isCImmediate {
			param1 = sc.memory[param1]
		}
		sc.memory[param3] = param1 + param2
		sc.programCounter += 4
		break
	case "02":
		param3 := pc + 3
		if !isAImmediate {
			param3 = sc.memory[pc+3]
		}
		param2 := sc.memory[pc+2]
		if !isBImmediate {
			param2 = sc.memory[param2]
		}
		param1 := sc.memory[pc+1]
		if !isCImmediate {
			param1 = sc.memory[param1]
		}
		sc.memory[param3] = param1 * param2
		sc.programCounter += 4
		break
	case "03":
		intInput := 5
		param1 := pc + 1
		if !isCImmediate {
			param1 = sc.memory[param1]
		}
		sc.memory[param1] = intInput
		sc.programCounter += 2
		break
	case "04":
		param1 := pc + 1
		if !isCImmediate {
			param1 = sc.memory[param1]
		}
		fmt.Printf("Output: %v\n", sc.memory[param1])
		sc.programCounter += 2
		break
	case "05":
		param2 := sc.memory[pc+2]
		if !isBImmediate {
			param2 = sc.memory[param2]
		}
		param1 := sc.memory[pc+1]
		if !isCImmediate {
			param1 = sc.memory[param1]
		}
		if param1 != 0 {
			sc.programCounter = param2
		} else {
			sc.programCounter += 3
		}
		break
	case "06":
		param2 := sc.memory[pc+2]
		if !isBImmediate {
			param2 = sc.memory[param2]
		}
		param1 := sc.memory[pc+1]
		if !isCImmediate {
			param1 = sc.memory[param1]
		}
		if param1 == 0 {
			sc.programCounter = param2
		} else {
			sc.programCounter += 3
		}
		break
	case "07":
		param3 := pc + 3
		if !isAImmediate {
			param3 = sc.memory[pc+3]
		}
		param2 := sc.memory[pc+2]
		if !isBImmediate {
			param2 = sc.memory[param2]
		}
		param1 := sc.memory[pc+1]
		if !isCImmediate {
			param1 = sc.memory[param1]
		}
		result := 0
		if param1 < param2 {
			result = 1
		}
		sc.memory[param3] = result
		sc.programCounter += 4
		break
	case "08":
		param3 := pc + 3
		if !isAImmediate {
			param3 = sc.memory[pc+3]
		}
		param2 := sc.memory[pc+2]
		if !isBImmediate {
			param2 = sc.memory[param2]
		}
		param1 := sc.memory[pc+1]
		if !isCImmediate {
			param1 = sc.memory[param1]
		}
		result := 0
		if param1 == param2 {
			result = 1
		}
		sc.memory[param3] = result
		sc.programCounter += 4
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
