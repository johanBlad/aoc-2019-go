package day6

import (
	"fmt"
	"strings"
	"time"

	"johanBlad.aoc-2019/common"
)

func parseTestInput() []string {
	in := `COM)B
	B)C
	C)D
	D)E
	E)F
	B)G
	G)H
	D)I
	E)J
	J)K
	K)L
	K)YOU
	I)SAN`
	inSplit := strings.Split(in, "\n")
	inTrimmed := make([]string, len(inSplit))
	for i, e := range inSplit {
		inTrimmed[i] = strings.TrimSpace(e)

	}
	return inTrimmed
}

type Planet struct {
	Name   string
	Parent *Planet
}

func adoptOrphans(planetMap map[string]*Planet, orphans []string) map[string]*Planet {
	for len(orphans) > 0 {
		newOrphans := make([]string, 0)
		for _, childOfOrphanName := range orphans {
			childOfOrphanPlanet := planetMap[childOfOrphanName]
			orphan := childOfOrphanPlanet.Parent
			parentOfOrphan := planetMap[orphan.Name].Parent
			orphan.Parent = parentOfOrphan
			if childOfOrphanPlanet.Parent.Parent.Parent == nil {
				newOrphans = append(newOrphans, childOfOrphanPlanet.Parent.Name)
			}
		}
		orphans = newOrphans
		fmt.Println("orphans remaining:", len(orphans))
	}
	return planetMap
}

func parseOrbits(orbits []string) map[string]*Planet {
	planetMap := make(map[string]*Planet)
	childrenOfOrphans := make([]string, 0)

	root := Planet{Name: "COM", Parent: nil}
	planetMap[root.Name] = &root
	for _, orbit := range orbits {
		orbitRelationship := strings.Split(orbit, ")")

		parentName := orbitRelationship[0]
		childName := orbitRelationship[1]

		parentPlanet, parentInMap := planetMap[parentName]
		childPlanet, childInMap := planetMap[childName]

		if !parentInMap {
			parentPlanet = &Planet{Name: parentName, Parent: nil}
		}

		if !childInMap {
			childPlanet = &Planet{Name: childName, Parent: parentPlanet}
		}

		if !parentInMap {
			childrenOfOrphans = append(childrenOfOrphans, childName)
		}
		planetMap[childName] = childPlanet
	}
	planetMap = adoptOrphans(planetMap, childrenOfOrphans)
	return planetMap
}

func solve1(planetMap map[string]*Planet) {
	allTraverses := 0
	for _, planet := range planetMap {
		current := *planet

		for current.Parent != nil {
			allTraverses++
			current = *current.Parent
		}
	}

	fmt.Println("all traversed:", allTraverses)
}

// becomes very inefficient, since the traverseForPlanet method is very efficient (following pointers to memory)
// spawning go routines creates overhead
func solve1Parallel(planetMap map[string]*Planet) {
	allTraverses := 0
	// we have a map of planets and need to traverse to COM for each planet
	//	-> loop over planets and spawn goroutine for each, and traverse.
	// 	-> pass total number of traverses back through a channel
	// 	-> main goroutine collects traverses and sums when finished.
	// planetMap := []*Planet{planetMap["COM"], planetMap["B"]}
	ch := make(chan int)
	for _, planet := range planetMap {
		go traveseForPlanet(planet, ch)
	}
	i := 0
	for e := range ch {
		allTraverses += e
		i++
		if i == len(planetMap) {
			close(ch)
		}
	}

	fmt.Println("all traversed parallel:", allTraverses)
}

func traveseForPlanet(planet *Planet, ch chan int) {
	current := *planet
	traverses := 0
	for current.Parent != nil {
		traverses++
		current = *current.Parent
	}
	ch <- traverses
}

func solve2(planetMap map[string]*Planet) {
	santaPlanet := planetMap["SAN"]
	santaTraverses := 0
	santaDistance := make(map[string]int)
	current := *santaPlanet
	for current.Parent != nil {
		santaDistance[current.Name] = santaTraverses
		santaTraverses++
		current = *current.Parent
	}

	myPlanet := planetMap["YOU"]
	distanceToSanta := 0
	myTraverses := 0
	current = *myPlanet
	for current.Parent != nil {
		if santaConnection, ok := santaDistance[current.Name]; ok {
			distanceToSanta = santaConnection + myTraverses - 2
			break
		}
		myTraverses++
		current = *current.Parent
	}
	fmt.Println("distance to santa", distanceToSanta)
}

func Run() {
	// in := parseTestInput()
	in := common.ReadInputToString("./input/6.in")
	planetMap := parseOrbits(in)
	solve1(planetMap)
	solve1Parallel(planetMap)
	solve2(planetMap)
	// runCallable(solve1, "solve1", planetMap)
	// runCallable(solve1Parallel, "solve1Parallel", planetMap)
	// runCallable(solve2, "solve2", planetMap)

}

type convert func(map[string]*Planet)

func runCallable(fn convert, name string, in map[string]*Planet) {
	start := time.Now()
	fn(in)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(name, elapsed)

}
