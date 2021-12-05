package day10

import "fmt"

func getFakeInput() string {
	return (`.#..#
			.....
			#####
			....#
			...##`)
}

func getFakeInput2() string {
	return (`......#.#.
			#..#.#....
			..#######.
			.#.#.###..
			.#..#.....
			..#....#.#
			#..#....#.
			.##.#..###
			##...#..#.
			.#....####`)
}

func getFakeInput3() string {
	return (`.#..##.###...#######
	##.############..##.
	.#.######.########.#
	.###.#######.####.#.
	#####.##.#.##.###.##
	..#####..#.#########
	####################
	#.####....###.#.#.##
	##.#################
	#####.##.###..####..
	..######..##.#######
	####.##.####...##..#
	.#####..#.######.###
	##...#.##########...
	#.##########.#######
	.####.#.###.###.#.##
	....##.##.###..#####
	.#.#.###########.###
	#.#.#.#####.####.###
	###.##.####.##.#..##`)
}

func getRealInput() string {
	return (`#..#....#...#.#..#.......##.#.####
	#......#..#.#..####.....#..#...##.
	.##.......#..#.#....#.#..#.#....#.
	###..#.....###.#....##.....#...#..
	...#.##..#.###.......#....#....###
	.####...##...........##..#..#.##..
	..#...#.#.#.###....#.#...##.....#.
	......#.....#..#...##.#..##.#..###
	...###.#....#..##.#.#.#....#...###
	..#.###.####..###.#.##..#.##.###..
	...##...#.#..##.#............##.##
	....#.##.##.##..#......##.........
	.#..#.#..#.##......##...#.#.#...##
	.##.....#.#.##...#.#.#...#..###...
	#.#.#..##......#...#...#.......#..
	#.......#..#####.###.#..#..#.#.#..
	.#......##......##...#..#..#..###.
	#.#...#..#....##.#....#.##.#....#.
	....#..#....##..#...##..#..#.#.##.
	#.#.#.#.##.#.#..###.......#....###
	...#.#..##....###.####.#..#.#..#..
	#....##..#...##.#.#.........##.#..
	.#....#.#...#.#.........#..#......
	...#..###...#...#.#.#...#.#..##.##
	.####.##.#..#.#.#.#...#.##......#.
	.##....##..#.#.#.......#.....####.
	#.##.##....#...#..#.#..###..#.###.
	...###.#..#.....#.#.#.#....#....#.
	......#...#.........##....#....##.
	.....#.....#..#.##.#.###.#..##....
	.#.....#.#.....#####.....##..#....
	.####.##...#.......####..#....##..
	.#.#.......#......#.##..##.#.#..##
	......##.....##...##.##...##......`)
}

func getTestMap() [][]int {
	return [][]int{{1, 1, 0}, {1, 1, 1}, {0, 0, 1}}
}

func printMap(mmap [][]int) {
	fmt.Println()
	for _, e := range mmap {
		fmt.Println(e)
	}
	fmt.Println()
}
func test_clearMap() {
	asteroidMap := getTestMap()
	printMap(asteroidMap)

	clear1 := []Coord{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}}
	origin1 := Coord{X: 1, Y: 0}
	clearMap(clear1, asteroidMap, origin1)
	printMap(asteroidMap)

	asteroidMap = getTestMap()
	clear2 := []Coord{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}}
	origin2 := Coord{X: 0, Y: 1}
	clearMap(clear2, asteroidMap, origin2)
	printMap(asteroidMap)

	asteroidMap = getTestMap()
	clear3 := []Coord{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}}
	origin3 := Coord{X: 0, Y: 0}
	clearMap(clear3, asteroidMap, origin3)
	printMap(asteroidMap)
}

func test_findAsteroidsInColumn() {
	asteroidMap := getTestMap()
	found1 := findAsteroidsInColumn(0, asteroidMap)
	fmt.Println(found1)

	found2 := findAsteroidsInColumn(1, asteroidMap)
	fmt.Println(found2)

	found3 := findAsteroidsInColumn(2, asteroidMap)
	fmt.Println(found3)
}

func test_findAsteroidsInRow() {
	asteroidMap := getTestMap()
	found1 := findAsteroidsInRow(0, asteroidMap)
	fmt.Println(found1)

	found2 := findAsteroidsInRow(1, asteroidMap)
	fmt.Println(found2)

	found3 := findAsteroidsInRow(2, asteroidMap)
	fmt.Println(found3)
}

func test_countVisibleAsteroids() {
	vertical := []Coord{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}}
	horisontal := []Coord{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}}
	diagonal := []Coord{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}}

	single := []Coord{{X: 1, Y: 1}}
	double := []Coord{{X: 3, Y: 4}, {X: 4, Y: 2}}

	fmt.Println(countVisibleAsteroids(Coord{X: 1, Y: 0}, vertical), 1)
	fmt.Println(countVisibleAsteroids(Coord{X: 1, Y: 1}, vertical), 2)
	fmt.Println(countVisibleAsteroids(Coord{X: 1, Y: 2}, vertical), 1)

	fmt.Println(countVisibleAsteroids(Coord{X: 0, Y: 1}, horisontal), 1)
	fmt.Println(countVisibleAsteroids(Coord{X: 1, Y: 1}, horisontal), 2)
	fmt.Println(countVisibleAsteroids(Coord{X: 2, Y: 2}, horisontal), 1)

	fmt.Println(countVisibleAsteroids(Coord{X: 0, Y: 0}, diagonal), 1)
	fmt.Println(countVisibleAsteroids(Coord{X: 1, Y: 1}, diagonal), 2)
	fmt.Println(countVisibleAsteroids(Coord{X: 2, Y: 2}, diagonal), 1)

	fmt.Println(countVisibleAsteroids(Coord{X: 1, Y: 1}, single), 0)

	fmt.Println(countVisibleAsteroids(Coord{X: 4, Y: 2}, double), 1)

}

func RuntTest() {
	test_clearMap()
	test_findAsteroidsInColumn()
	test_findAsteroidsInRow()
	test_countVisibleAsteroids()
}
