package day6

import (
	"fmt"
	"strings"

	"johanBlad.aoc-2019/common"
)

func parseInput() []string {
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

		for {
			if current.Parent == nil {
				break
			}
			allTraverses++

			current = *current.Parent
		}
	}

	fmt.Println("all traversed:", allTraverses)
}

func solve2(planetMap map[string]*Planet) {
	santaPlanet := planetMap["SAN"]
	santaTraverses := 0
	santaDistance := make(map[string]int)
	current := *santaPlanet
	for {
		santaDistance[current.Name] = santaTraverses
		if current.Parent == nil {
			break
		}
		santaTraverses++
		current = *current.Parent
	}

	myPlanet := planetMap["YOU"]
	distanceToSanta := 0
	myTraverses := 0
	current = *myPlanet
	for {
		if santaConnection, ok := santaDistance[current.Name]; ok {
			distanceToSanta = santaConnection + myTraverses - 2
			break
		} else if current.Parent == nil {
			break
		}
		myTraverses++
		current = *current.Parent
	}
	fmt.Println("distance to santa", distanceToSanta)
}

func Run() {
	// in := parseInput()
	in := common.ReadInputToString("./input/6.in")
	planetMap := parseOrbits(in)
	solve1(planetMap)
	solve2(planetMap)

}
