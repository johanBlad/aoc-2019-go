package day1

import (
	"fmt"

	"johanBlad.aoc-2019/common"
)

func calculateFuel(weight int) int {
	fuel := (weight / 3) - 2
	if fuel <= 0 {
		return 0
	}
	return fuel + calculateFuel(fuel)
}

func Run() {

	ints := common.ReadInput("./input/1.in")

	fuelRequirements := 0
	for _, e := range ints {
		fuelRequirements += calculateFuel(e)
	}

	fmt.Println(fuelRequirements)
}
