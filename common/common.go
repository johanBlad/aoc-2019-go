package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadInput(filename string) []int {
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

func ReadInputToString(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}
	if err := scanner.Err(); err != nil {
		check(err)
	}
	return lines
}

func Read2Lines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	// input := Lines{a: "", b: ""}
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		check(err)
	}
	return lines
}

func ReadLine(filename string) []int {
	data, err := os.ReadFile(filename)
	check(err)
	split := strings.Split(string(data), ",")
	ints := make([]int, len(split))

	for i, e := range split {
		parsedInt, err := strconv.Atoi(e)
		check(err)
		ints[i] = parsedInt
	}
	return ints
}
