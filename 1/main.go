package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/JakeHl/AdventOfCode2019/utils"
)

func getMassesFromFile(filename string) (inputs []int32) {
	var inputsTxt, _ = utils.ReadTextFile(filename)
	var inputsStrings = strings.Split(inputsTxt, "\n")
	inputs = make([]int32, len(inputsStrings))
	for i, v := range inputsStrings {
		integer, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		inputs[i] = int32(integer)
	}
	return
}

func calculateFuelForFuel(fuel int32) (total int32) {
	total += fuel/3 - 2
	if total <= 0 {
		return 0
	}
	total += calculateFuelForFuel(total)
	return
}

func main() {
	inputs := getMassesFromFile("input.txt")

	var total int32 = 0
	for _, v := range inputs {
		fuel := v/3 - 2
		fuel += calculateFuelForFuel(fuel)
		total += fuel
	}

	fmt.Print(total, "\n")
}
