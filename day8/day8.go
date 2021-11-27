package day8

import (
	"fmt"
	"strings"

	"johanBlad.aoc-2019/common"
)

const (
	black       = "0"
	white       = "1"
	transparent = "2"
)

func Run() {
	// in := "123456789012" + "764042"
	// layerSize := 3 * 2
	in := common.ReadLineToString("./input/8.in")
	layerSize := 25 * 6

	// get layers as slice of string
	layersStr := stringLayers(in, layerSize)

	// solve1(layersStr)
	solve2(layersStr)
}

func solve2(in []string) {
	finalImage := in[0]
	for j, finalChar := range finalImage {
		if string(finalChar) != transparent {
			continue
		}
		// current pixel of final image in transparent -> swap it
		for i := 1; i < len(in); i++ {
			layerChar := in[i][j]
			if string(layerChar) != transparent {
				finalImage = finalImage[:j] + string(layerChar) + finalImage[j+1:]
				break
			}
		}

	}
	fmt.Println(finalImage)
}

func solve2Parallell() {

}

func solve1(in []string) {
	layerFewestZeros := ""
	maxZerosCount := 999999

	for _, layer := range in {
		zerosInLayer := strings.Count(layer, "0")
		if zerosInLayer < maxZerosCount {
			maxZerosCount = zerosInLayer
			layerFewestZeros = layer
		}
	}

	ones := strings.Count(layerFewestZeros, "1")
	twos := strings.Count(layerFewestZeros, "2")
	fmt.Println(ones * twos)
}

func stringLayers(in string, layerSize int) []string {
	inSize := len(in)
	layers := make([]string, 0)
	for i := 0; i < inSize; i += layerSize {
		layers = append(layers, in[i:i+layerSize])
	}
	return layers
}
