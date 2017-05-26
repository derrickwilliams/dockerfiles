package chessboard

import (
	"fmt"
	"strings"
	"math"
)

const boardHeight int = 8
const boardWidth int = 8

var ranks = []string {"8", "7", "6", "5", "4", "3", "2", "1"}
var files = []string {"a", "b", "c", "d", "e", "f", "g", "h"}
var positions = calculatePositions(ranks, files)

func Stats() {
	printStats(ranks, files, positions)
}

func Layout() {
	printBoard(positions)
}

func printStats(ranks, files, positions []string) {
	fmt.Println("STATS");
	fmt.Println("=====");
	fmt.Printf("board dimensions h: %d, w: %d\n", boardHeight, boardWidth)
	fmt.Println("number of ranks", len(ranks))
	fmt.Println("number of files", len(files))
	fmt.Println("number of positions", len(positions))
}

func printBoard(positions []string) {
	fmt.Println("\n=====")
	fmt.Println("BOARD")
	fmt.Println("=====")
	fmt.Println("-----------------------------------------")
	for posIndex, pos := range positions {
		fmt.Printf("| %s ", pos);

		if math.Mod(float64(posIndex + 1), float64(boardWidth)) == 0 {
			fmt.Println("|")
			fmt.Println("-----------------------------------------")

		}
	}
}

func calculatePositions(ranks, files []string) []string {
	positions := []string{}

	for _, r := range ranks {
		for _, f := range files {
			positions = append(positions, strings.Join([]string{f, r}, ""))
		}
	}

	return positions
}