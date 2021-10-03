package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readInput(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	ints := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		intLine, err := strconv.Atoi(scanner.Text())
		check(err)
		ints = append(ints, intLine)

	}
	if err := scanner.Err(); err != nil {
		check(err)
	}
	return ints

}

func calculateFuel(weight int) int {
	fuel := (weight / 3) - 2
	if fuel <= 0 {
		return 0
	}
	return fuel + calculateFuel(fuel)
}

func main() {

	ints := readInput("1.in")

	fuelRequirements := 0
	for _, e := range ints {
		fuelRequirements += calculateFuel(e)
	}

	fmt.Println(fuelRequirements)
}
