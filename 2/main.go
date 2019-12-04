package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/JakeHL/AdventOfCode2019/utils"
)

func getOperationValues(inputs []int) (val1, val2, position int) {
	val1 = inputs[0]
	val2 = inputs[1]
	position = inputs[2]
	return
}

func getInputsFromFile(filename string) []int {
	var inputsTxt, err = utils.ReadTextFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var inputsStrings = strings.Split(inputsTxt, ",")
	var inputs = make([]int, len(inputsStrings))
	for i, v := range inputsStrings {
		integer, err := strconv.ParseInt(v, 10, 8)
		if err != nil {
			log.Fatal(err)
		}
		inputs[i] = int(integer)
	}
	return inputs
}

func getOutputFromInputs(noun, verb int) int {
	inputs := getInputsFromFile("input.txt")
	inputs[1] = noun
	inputs[2] = verb
	for i := 0; i < len(inputs); i += 4 {
		if inputs[i] == 1 {
			val1, val2, position := getOperationValues(inputs[i+1 : i+4])
			inputs[position] = inputs[val1] + inputs[val2]
		} else if inputs[i] == 2 {
			val1, val2, position := getOperationValues(inputs[i+1 : i+4])
			inputs[position] = inputs[val1] * inputs[val2]
		} else if inputs[i] == 99 {
			break
		}
	}
	return inputs[0]
}

func main() {
	var output, noun, verb = 0, 0, 0
	for ; output != 19690720 && noun <= 99; noun++ {
		verb = 0
		for ; output != 19690720 && verb <= 99; verb++ {
			output = getOutputFromInputs(noun, verb)

			if output == 19690720 {
				fmt.Printf("found: 100 * [noun: %v] + [verb: %v] = [%v]; output = %v\n", noun, verb, 100*noun+verb, output)
			}
		}
	}
}
