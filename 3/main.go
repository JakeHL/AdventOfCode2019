package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/JakeHL/AdventOfCode2019/utils"
)

type WirePosition struct {
	x int
	y int
}

func (wp WirePosition) add(wirePos WirePosition) WirePosition {
	return WirePosition{
		x: wp.x + wirePos.x,
		y: wp.y + wirePos.y,
	}
}

func (wp WirePosition) distance() int {
	return abs(wp.x) + abs(wp.y)
}

func (wp WirePosition) isEqual(wirePos WirePosition) bool {
	return wp.x == wirePos.x && wp.y == wirePos.y
}

func getWiresFromFile(filename string) (wires [][]string) {
	inputTxt, err := utils.ReadTextFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var wireStrings = strings.Split(inputTxt, "\n")
	var wireCount = len(wireStrings)
	wires = make([][]string, wireCount)
	for i, v := range wireStrings {
		wires[i] = strings.Split(v, ",")
	}
	return
}

var regUp = regexp.MustCompile("^U[0-9]")
var regDown = regexp.MustCompile("^D[0-9]")
var regLeft = regexp.MustCompile("^L[0-9]")
var regRight = regexp.MustCompile("^R[0-9]")
var regDir = regexp.MustCompile("[U,D,L,R]")

func getDistanceFromDirection(direction string) int {
	var distanceStr = regDir.ReplaceAllString(direction, "")
	var distance, err = strconv.Atoi(distanceStr)
	if err != nil {
		log.Fatal(err)
	}
	return distance
}

func main() {
	var wireInstructions = getWiresFromFile("input.txt")

	var wires [][]WirePosition = make([][]WirePosition, len(wireInstructions))
	for instructionsIndex, instructions := range wireInstructions {
		wires[instructionsIndex] = []WirePosition{{x: 0, y: 0}}
		for _, direction := range instructions {
			fmt.Printf("Wire No: %v, Direction: %v\n", instructionsIndex, direction)
			switch {
			case regUp.MatchString(direction):
				var distance = getDistanceFromDirection(direction)
				for i := 0; i < distance; i++ {
					var lastPart = wires[instructionsIndex][len(wires[instructionsIndex])-1]
					wires[instructionsIndex] = append(wires[instructionsIndex], lastPart.add(WirePosition{x: 0, y: 1}))
				}
				break
			case regDown.MatchString(direction):
				var distance = getDistanceFromDirection(direction)
				for i := 0; i < distance; i++ {
					var lastPart = wires[instructionsIndex][len(wires[instructionsIndex])-1]
					wires[instructionsIndex] = append(wires[instructionsIndex], lastPart.add(WirePosition{x: 0, y: -1}))
				}
				break
			case regLeft.MatchString(direction):
				var distance = getDistanceFromDirection(direction)
				for i := 0; i < distance; i++ {
					var lastPart = wires[instructionsIndex][len(wires[instructionsIndex])-1]
					wires[instructionsIndex] = append(wires[instructionsIndex], lastPart.add(WirePosition{x: -1, y: 0}))
				}
				break
			case regRight.MatchString(direction):
				var distance = getDistanceFromDirection(direction)
				for i := 0; i < distance; i++ {
					var lastPart = wires[instructionsIndex][len(wires[instructionsIndex])-1]
					wires[instructionsIndex] = append(wires[instructionsIndex], lastPart.add(WirePosition{x: 1, y: 0}))
				}
				break
			}
		}
	}

	var firstWire = wires[0]
	var seccondWire = wires[1]
	var intersections []WirePosition = []WirePosition{}

	for _, fparts := range firstWire {
		for _, sparts := range seccondWire {
			if fparts.isEqual(sparts) {
				intersections = append(intersections, fparts)
			}
		}
	}

	var lowest = intersections[1].distance()
	for _, i := range intersections {
		var dist = i.distance()
		if dist < lowest && dist > 0 {
			lowest = dist
		}
	}
	fmt.Print(lowest)

}

func abs(val int) int {
	if val < 0 {
		return 0 - val
	}
	return val
}
