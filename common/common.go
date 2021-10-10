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
