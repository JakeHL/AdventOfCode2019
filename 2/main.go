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

func main() {
	var inputsTxt, err = utils.ReadTextFile("input.txt")
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

	for i := 0; i < len(inputs); i += 4 {
		if inputs[i] == 1 {
			val1, val2, position := getOperationValues(inputs[i+1 : i+4])
			inputs[position] = inputs[val1] + inputs[val2]
		} else if inputs[i] == 2 {
			val1, val2, position := getOperationValues(inputs[i+1 : i+4])
			inputs[position] = inputs[val1] * inputs[val2]
		} else if inputs[i] == 9 {
			break
		}
	}
	fmt.Print(inputs)
}
