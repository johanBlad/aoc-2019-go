package day2

import (
	"fmt"

	"johanBlad.aoc-2019/common"
)

func processCode(ints []int, i int) []int {
	opcode := ints[i]
	inputIndex1 := ints[i+1]
	inputIndex2 := ints[i+2]
	outputIndex := ints[i+3]

	if opcode == 1 {
		ints[outputIndex] = ints[inputIndex1] + ints[inputIndex2]
	} else if opcode == 2 {
		ints[outputIndex] = ints[inputIndex1] * ints[inputIndex2]
	} else if opcode == 99 {
		return ints
	} else {
		return ints
	}
	return processCode(ints, i+4)
}

func brute(ints []int, target int) (int, int) {
	var res []int
	var tmp []int
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			tmp = make([]int, len(ints))
			copy(tmp, ints)
			tmp[1] = i
			tmp[2] = j
			res = processCode(tmp, 0)
			if res[0] == target {
				fmt.Println("FOUND!")
				return i, j
			}
		}
	}
	return -1, -1
}

func Run() {
	ints := common.ReadLine("./input/2.in")
	target := 19690720
	// in1 := 12
	// in2 := 2
	// ints[1] = in1
	// ints[2] = in2
	i, j := brute(ints, target)
	fmt.Println(i, j, 100*i+j)

}
