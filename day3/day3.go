package day3

import (
	"fmt"
	"strconv"
	"strings"

	"johanBlad.aoc-2019/common"
)

type coord struct {
	x int
	y int
}

type crossIndex struct {
	leader   int
	follower int
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func move(start coord, instruction string) coord {
	direction := string(instruction[0])
	distance, err := strconv.Atoi(instruction[1:])

	if err != nil {
		panic(err)
	}

	newX := start.x
	newY := start.y
	if direction == "U" {
		newY += distance
	} else if direction == "R" {
		newX += distance
	} else if direction == "D" {
		newY -= distance
	} else if direction == "L" {
		newX -= distance
	} else {
		panic("wierd state!")
	}
	// take previous position and a move instruction
	// return new position
	return coord{newX, newY}
}

func vertical_catch_up(leaderA coord, leaderB coord, followerA coord, followerB coord) bool {
	// follower is catching up to leader vertically
	return leaderA.x == leaderB.x &&
		followerA.x == leaderA.x &&
		Min(followerA.y, followerB.y) < Min(leaderA.y, leaderB.y) &&
		Max(followerA.y, followerB.y) > Min(leaderA.y, leaderB.y)
}

func horizontal_catch_up(leaderA coord, leaderB coord, followerA coord, followerB coord) bool {
	// follower is catching up to leader horizontally
	return leaderA.y == leaderB.y &&
		followerA.y == leaderA.y &&
		Min(followerA.x, followerB.x) < Min(leaderA.x, leaderB.x) &&
		Max(followerA.x, followerB.x) > Min(leaderA.x, leaderB.x)
}

func vertical_follower_horizontal_leader(leaderA coord, leaderB coord, followerA coord, followerB coord) bool {
	// follower is crossing a horizontal line of leader
	return leaderA.y == leaderB.y &&
		followerA.x < Max(leaderA.x, leaderB.x) &&
		followerA.x > Min(leaderA.x, leaderB.x) &&
		Min(followerA.y, followerB.y) < leaderA.y &&
		Max(followerA.y, followerB.y) > leaderA.y
}

func horizontal_follower_vertical_leader(leaderA coord, leaderB coord, followerA coord, followerB coord) bool {
	// follower is crossing a vertical line of leader
	return leaderA.x == leaderB.x &&
		followerA.y < Max(leaderA.y, leaderB.y) &&
		followerA.y > Min(leaderA.y, leaderB.y) &&
		Min(followerA.x, followerB.x) < leaderA.x &&
		Max(followerA.x, followerB.x) > leaderA.x
}

func splitPath(path string) []string {
	return strings.Split(path, ",")
}

func calcPath(
	leaderIndex int,
	followerIndex int,
	leader []coord,
	follower []coord,
	leaderTopUp int,
	followerTopUp int,
) int {
	leaderPathSum := 0
	followerPathSum := 0
	for i := 1; i <= leaderIndex; i++ {
		leaderPathSum += Abs((leader[i].x - leader[i-1].x) + (leader[i].y - leader[i-1].y))
	}
	leaderPathSum += leaderTopUp

	for j := 1; j <= followerIndex; j++ {
		followerPathSum += Abs((follower[j].x - follower[j-1].x) + (follower[j].y - follower[j-1].y))
	}
	followerPathSum += followerTopUp
	return leaderPathSum + followerPathSum
}

func Run() {
	input := common.Read2Lines("./input/3.in")
	i1 := splitPath(input[0])
	i2 := splitPath(input[1])

	leader := make([]coord, len(i1)+1)
	follower := make([]coord, len(i2)+1)
	leader[0] = coord{x: 0, y: 0}
	follower[0] = coord{x: 0, y: 0}

	for i := 0; i < Max(len(i1), len(i2)); i++ {
		if i < len(i1) {
			leader[i+1] = move(leader[i], i1[i])
		}

		if i < len(i2) {
			follower[i+1] = move(follower[i], i2[i])
		}
	}
	crossings := make([]coord, 0)
	crossPaths := make([]int, 0)

	for i := 0; i < len(follower)-1; i++ {
		for j := 0; j < len(leader)-1; j++ {
			followerA := follower[i]
			followerB := follower[i+1]

			leaderA := leader[j]
			leaderB := leader[j+1]

			if followerA.x == followerB.x {
				// follower is vertical
				if vertical_catch_up(leaderA, leaderB, followerA, followerB) {
					crossings = append(crossings, coord{followerA.x, Min(leaderA.y, leaderB.y)})
				} else if vertical_follower_horizontal_leader(leaderA, leaderB, followerA, followerB) {
					crossings = append(crossings, coord{followerA.x, leaderA.y})
					followerTopUp := Abs(followerA.y - leaderA.y)
					leaderTopUp := Abs(leaderA.x - followerA.x)
					crossPaths = append(crossPaths, calcPath(j, i, leader, follower, leaderTopUp, followerTopUp))

				}
			} else if followerA.y == followerB.y {
				// follower is horizonal
				if horizontal_catch_up(leaderA, leaderB, followerA, followerB) {
					crossings = append(crossings, coord{Min(leaderA.x, leaderB.x), followerA.y})
				} else if horizontal_follower_vertical_leader(leaderA, leaderB, followerA, followerB) {
					crossings = append(crossings, coord{leaderA.x, followerA.y})
					followerTopUp := Abs(followerA.x - leaderA.x)
					leaderTopUp := Abs(leaderA.y - followerA.y)
					crossPaths = append(crossPaths, calcPath(j, i, leader, follower, leaderTopUp, followerTopUp))
				}
			} else {
				fmt.Println("invalid move")
			}
		}
	}
	minDist := -1
	for _, crossing := range crossings {
		dist := Abs(crossing.x) + Abs(crossing.y)
		if minDist == -1 || dist < minDist {
			minDist = dist
		}
	}
	fmt.Println("minDist: ", minDist)

	minPath := -1
	for _, crossPath := range crossPaths {
		if minPath == -1 || crossPath < minPath {
			minPath = crossPath
		}
	}
	fmt.Println("minPath: ", minPath)

}
