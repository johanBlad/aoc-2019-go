package day2

import (
	"fmt"
	"time"

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

func processCode_ch(ints []int, i int, c chan []int, k int, j int) {
	opcode := ints[i]
	inputIndex1 := ints[i+1]
	inputIndex2 := ints[i+2]
	outputIndex := ints[i+3]

	if len(ints) < outputIndex {
		return
	} else if opcode == 1 {
		ints[outputIndex] = ints[inputIndex1] + ints[inputIndex2]
	} else if opcode == 2 {
		ints[outputIndex] = ints[inputIndex1] * ints[inputIndex2]
	} else if opcode == 99 {
		c <- []int{ints[0], k, j}
		return
	} else {
		return
	}
	processCode_ch(ints, i+4, c, k, j)
}

func brute_routine_two(ints []int, target int) (int, int) {
	var tmp []int
	c := make(chan []int)
	for i := 0; i < 100; i++ {
		go func(iLocal int) {
			for j := 0; j < 100; j++ {
				tmp = make([]int, len(ints))
				copy(tmp, ints)
				tmp[1] = iLocal
				tmp[2] = j
				processCode_ch(tmp, 0, c, iLocal, j)
			}
		}(i)
	}

	for res := range c {
		if res[0] == target {
			return res[1], res[2]
		}
	}
	return -1, -1
}

func brute_routine_one(ints []int, target int) (int, int) {
	var tmp []int
	c := make(chan []int)
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			tmp = make([]int, len(ints))
			copy(tmp, ints)
			tmp[1] = i
			tmp[2] = j
			go processCode_ch(tmp, 0, c, i, j)
		}
	}

	for res := range c {
		if res[0] == target {
			return res[1], res[2]
		}
	}
	return -1, -1
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
				return i, j
			}
		}
	}
	return -1, -1
}

type convert func([]int, int) (int, int)

func runCallable(fn convert, name string, ints []int, target int) {
	start := time.Now()
	// i, j := brute(ints, target)
	// i, j := brute_routine_one(ints, target)
	i, j := fn(ints, target)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(name, elapsed)
	fmt.Println(i, j, 100*i+j)
	fmt.Println()

}

func Run() {
	ints := common.ReadLine("./input/2.in")
	target := 19690720
	runCallable(brute, "brute", ints, target)
	runCallable(brute_routine_one, "brute_routine_one", ints, target)
	runCallable(brute_routine_two, "brute_routine_two", ints, target)

	// i, j := brute(ints, target)
	// i, j := brute_routine_one(ints, target)

	// fmt.Println(i, j, 100*i+j)

}
