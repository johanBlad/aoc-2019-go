package day10

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

type Coord struct {
	X int
	Y int
}
type MonitoringStation struct {
	Coordinates      Coord
	VisibleAsteroids int
}

func parseAsteroids(in string) [][]int {
	lineSplits := strings.Split(in, "\n")
	height := len(lineSplits)
	width := len(strings.TrimSpace(lineSplits[0]))
	asteroids := make([][]int, height)
	for y, line := range lineSplits {
		asteroidsForLine := make([]int, width)
		for x, coord := range strings.TrimSpace(line) {
			if string(coord) == "#" {
				asteroidsForLine[x] = 1
			}
		}
		asteroids[y] = asteroidsForLine
	}
	return asteroids
}

func coord(x int, y int) Coord {
	return Coord{X: x, Y: y}
}

func copyMap(asteroidMap [][]int) [][]int {
	duplicate := make([][]int, len(asteroidMap))
	for y, line := range asteroidMap {
		duplicate[y] = make([]int, len(line))
		copy(duplicate[y], asteroidMap[y])
	}
	return duplicate
}

func Abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func getLineParams(origin Coord, current Coord) (float64, float64) {
	k := float64(current.Y-origin.Y) / float64(current.X-origin.X)
	m := float64(current.Y) - k*float64(current.X)
	return k, m
}

func findAsteroidsInDiagonal(origin Coord, target Coord, asteroidMap [][]int) []Coord {
	diagonal := make([]Coord, 0)
	k, m := getLineParams(origin, target)
	for i := 0; i < len(asteroidMap[0]); i++ {
		yRes := k*float64(i) + m
		yRounded := int(math.RoundToEven(yRes))
		if Abs(yRes-float64(math.RoundToEven(yRes))) < 0.000001 &&
			yRounded >= 0 &&
			yRounded < len(asteroidMap) &&
			asteroidMap[yRounded][i] == 1 {
			diagonal = append(diagonal, Coord{X: i, Y: yRounded})
		}

	}
	return diagonal
}

func findAsteroidsInColumn(x int, asteroidMap [][]int) []Coord {
	column := make([]Coord, 0)
	for y := 0; y < len(asteroidMap); y++ {
		if asteroidMap[y][x] == 1 {
			column = append(column, coord(x, y))
		}
	}
	return column
}

func findAsteroidsInRow(y int, asteroidMap [][]int) []Coord {
	row := make([]Coord, 0)
	for x := 0; x < len(asteroidMap[0]); x++ {
		if asteroidMap[y][x] == 1 {
			row = append(row, coord(x, y))
		}
	}
	return row
}

func countVisibleAsteroids(origin Coord, asteroidsInPath []Coord) int {
	xMin := 9999
	xMax := 0
	yMin := 9999
	yMax := 0

	for _, asteroid := range asteroidsInPath {
		if asteroid.X > xMax {
			xMax = asteroid.X
		}
		if asteroid.X < xMin {
			xMin = asteroid.X
		}
		if asteroid.Y > yMax {
			yMax = asteroid.Y
		}
		if asteroid.Y < yMin {
			yMin = asteroid.Y
		}
	}

	numVisible := 0
	if origin.X == xMax && origin.X == xMin {
		if origin.Y < yMax {
			numVisible++
		}
		if origin.Y > yMin {
			numVisible++
		}
	} else {
		if origin.X < xMax {
			numVisible++
		}
		if origin.X > xMin {
			numVisible++
		}
	}

	return numVisible
}

func clearMap(asteroidToClear []Coord, asteroidMap [][]int, origin Coord) [][]int {
	for _, asteroid := range asteroidToClear {
		asteroidMap[asteroid.Y][asteroid.X] = 0
	}
	asteroidMap[origin.Y][origin.X] = 1
	return asteroidMap
}

func traverseMap(origin Coord, asteroidMap [][]int) ([][]int, int) {
	for y, row := range asteroidMap {
		for x, el := range row {
			if el == 1 && !(origin.X == x && origin.Y == y) {
				if y == origin.Y {
					// get visible asteroids from origin on this row, 1 or 2
					row := findAsteroidsInRow(y, asteroidMap)
					numVisible := countVisibleAsteroids(origin, row)
					asteroidMap = clearMap(row, asteroidMap, origin)
					return asteroidMap, numVisible
				}
				if x == origin.X {
					// current asteroid is in same column as origin -> get visible asteroids from origin on this column, 1 or 2
					column := findAsteroidsInColumn(x, asteroidMap)
					numVisible := countVisibleAsteroids(origin, column)
					asteroidMap = clearMap(column, asteroidMap, origin)
					return asteroidMap, numVisible
				} else {
					// we have an asteroid that is not in the same line or column as the origin
					// calculate the line between them. Find all values within the map that falls on that line
					diagonal := findAsteroidsInDiagonal(origin, Coord{X: x, Y: y}, asteroidMap)
					numVisible := countVisibleAsteroids(origin, diagonal)
					asteroidMap = clearMap(diagonal, asteroidMap, origin)
					return asteroidMap, numVisible
				}
			}

		}
	}
	return nil, 0
}

func evaluateMonitoringStation(origin Coord, asteroidMap [][]int) MonitoringStation {
	asteroidMapCopy := copyMap(asteroidMap)
	visibleAsteroids := 0
	var c int
	for asteroidMapCopy != nil {
		asteroidMapCopy, c = traverseMap(origin, asteroidMapCopy)
		// time.Sleep(time.Second * 1)
		// fmt.Println(asteroidMapCopy)
		visibleAsteroids += c
	}
	return MonitoringStation{Coordinates: origin, VisibleAsteroids: visibleAsteroids}
}

func evaluateMonitoringStations(asteroidMap [][]int) []MonitoringStation {
	monitoringStations := make([]MonitoringStation, 0)
	for y, row := range asteroidMap {
		for x, el := range row {
			if el == 1 {
				// monitoring station! evaluate it
				origin := Coord{X: x, Y: y}
				monitoringStations = append(monitoringStations, evaluateMonitoringStation(origin, asteroidMap))

			}
		}
	}
	return monitoringStations
}

func solve1(asteroidMap [][]int) {
	stations := evaluateMonitoringStations(asteroidMap)
	sort.Slice(stations, func(i, j int) bool {
		return stations[i].VisibleAsteroids < stations[j].VisibleAsteroids
	})

	for _, s := range stations {
		fmt.Println(s)
	}
}

func Run() {
	RuntTest()
	in := getRealInput()
	asteroidMap := parseAsteroids(in)
	solve1(asteroidMap)

}
