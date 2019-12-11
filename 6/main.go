package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/JakeHL/AdventOfCode2019/utils"
)

type orbit struct {
	name, orbits string
}

func (o orbit) String() string {
	return fmt.Sprintf("{ name: %v, orbits: %v }\n", o.name, o.orbits)
}

func getOrbitsFromFile(filename string) []orbit {
	inputString, err := utils.ReadTextFile(filename)
	if err != nil {
		log.Panic(err)
	}
	planetFields := strings.Split(inputString, "\n")
	orbits := make([]orbit, len(planetFields))
	for i, v := range planetFields {
		parts := strings.Split(v, ")")
		orbits[i] = orbit{
			name:   parts[1],
			orbits: parts[0],
		}
	}
	return orbits
}

func containsString(array []string, item string) bool {
	for _, v := range array {
		if v == item {
			return true
		}
	}
	return false
}

func getUniquePlanetsFromOrbits(orbits []orbit) (result []string) {
	result = []string{}
	for _, v := range orbits {
		if !containsString(result, v.name) {
			result = append(result, v.name)
		}
	}
	return
}

func getOrbitForPlanet(orbits []orbit, planet string) orbit {
	if planet == "COM" {
		log.Fatal("COM cannot orbit another planet")
	}
	for _, v := range orbits {
		if v.name == planet {
			return v
		}
	}
	log.Fatal("Planet does not exist")
	return orbit{}
}

func main() {
	orbits := getOrbitsFromFile("input.txt")
	planets := getUniquePlanetsFromOrbits(orbits)

	orbitCount := 0
	for _, v := range planets {
		next := v
		for next != "COM" {
			orbit := getOrbitForPlanet(orbits, next)
			next = orbit.orbits
			orbitCount++
		}
	}

	youOrbits := []orbit{}
	youNext := "YOU"
	for youNext != "COM" {
		orbit := getOrbitForPlanet(orbits, youNext)
		youNext = orbit.orbits
		youOrbits = append(youOrbits, orbit)
	}

	sanOrbits := []orbit{}
	sanNext := "SAN"
	for sanNext != "COM" {
		orbit := getOrbitForPlanet(orbits, sanNext)
		sanNext = orbit.orbits
		sanOrbits = append(sanOrbits, orbit)
	}

	jumpsToSanta := 0
	for yi, yv := range youOrbits {
		for si, sv := range sanOrbits {
			if yv.name == sv.name {
				jumpsToSanta = yi + si - 2 // to account for san & you orbits
				break
			}
		}
		if jumpsToSanta != 0 {
			break
		}
	}

	fmt.Printf("total orbits: %v, jumps to santa: %v", orbitCount, jumpsToSanta)
}
