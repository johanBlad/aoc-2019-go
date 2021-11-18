package day5

import (
	"fmt"
	"sync"
	"time"

	"johanBlad.aoc-2019/common"
)

type IntCodeComputer struct {
	Instructions []int
	index        int
	base         int
	Input        chan int
	Output       chan int
}

func CreateIntCodeComputer(input []int) *IntCodeComputer {
	intCodeComputer := IntCodeComputer{
		Instructions: input,
		index:        0,
		base:         0,
		Input:        make(chan int, 2),
		Output:       make(chan int, 2),
	}
	return &intCodeComputer
}

func setVal(i, val, mode int) {

}

func (icc *IntCodeComputer) getVal(pos, mode int) int {
	var i int
	if mode == 0 {
		i = icc.Instructions[icc.index+pos]
	} else if mode == 1 {
		i = pos
	}
	return icc.Instructions[i]
}

func (icc *IntCodeComputer) Process(wg *sync.WaitGroup) int {
	for {
		inst := icc.Instructions[icc.index]
		opcode := inst % 100

		mode_p2 := (inst / 1000) - (inst/10000)*10
		mode_p1 := (inst / 100) - (inst/1000)*10

		if opcode < 1 || opcode > 4 {
			fmt.Println("ERROR")
			if wg != nil {
				wg.Done()
			}
			return icc.Instructions[0]
		}
		fmt.Println(icc.index, inst, opcode, mode_p1, mode_p2)
		// fmt.Println(icc.Instructions[icc.index : icc.index+10])
		fmt.Println(icc.Instructions)

		switch opcode {
		case 1:
			// add
			val1 := icc.getVal(1, mode_p1)
			val2 := icc.getVal(2, mode_p2)
			out := icc.Instructions[icc.index+3]
			icc.Instructions[out] = val1 + val2
			fmt.Println(icc.Instructions)
			icc.index += 4
			break
		case 2:
			// muliply
			val1 := icc.getVal(1, mode_p1)
			val2 := icc.getVal(2, mode_p2)
			out := icc.Instructions[icc.index+3]
			icc.Instructions[out] = val1 * val2
			fmt.Println(icc.Instructions)
			icc.index += 4
			break
		case 3:
			// input
			out := icc.Instructions[icc.index+1]
			icc.Instructions[out] = <-icc.Input
			icc.index += 2
			break
		case 4:
			// output
			fmt.Println("output", icc.Instructions[icc.Instructions[icc.index+1]])
			// icc.Output <- icc.Instructions[icc.Instructions[icc.index+1]]
			icc.index += 2
			break
		default:
			fmt.Println("ERR")
			break
		}
		time.Sleep(5 * time.Second)
	}
}

func receiver(in, out chan int) {
	for {
		val := <-in
		if val != 0 {
			out <- val
			return
		}
	}
}

func app() int {
	ints := common.ReadLine("./input/5.in")
	// ints := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	var wg sync.WaitGroup

	icc := CreateIntCodeComputer(ints)

	output := make(chan int)
	defer close(output)

	// go receiver(icc.Output, output)
	wg.Add(1)
	go icc.Process(&wg)

	icc.Input <- 1
	wg.Wait()

	return 1 //<-output
}

func Run() {
	res := app()
	fmt.Println(res)
	// i, j := brute(ints, target)
	// i, j := brute_routine_one(ints, target)

	// fmt.Println(i, j, 100*i+j)

}
