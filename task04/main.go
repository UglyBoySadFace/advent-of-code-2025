package main

import (
	"fmt"
)

func main() {
	var line string
	var lines []string
	result := 0
	for true {
		_, err := fmt.Scanf("%s", &line)
		if err != nil {
			break
		}
		lines = append(lines, line)
	}
	grid := parseGrid(lines)

	areas := locateAccassibleAreas(grid)
	for areas > 0 {
		result += areas
		areas = locateAccassibleAreas(grid)
	}
	fmt.Println(result)
}

// . means empty space
// @ means roll of paper
// locate rolls of paper that have less than 4 out of 8 surrounding spaces not a roll of paper
func locateAccassibleAreas(grid [][]rune) int {
	count := 0
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '@' {
				emptyCount := 0
				for _, dir := range directions {
					ni, nj := i+dir[0], j+dir[1]
					if ni >= 0 && ni < len(grid) && nj >= 0 && nj < len(grid[i]) {
						if grid[ni][nj] != '@' {
							emptyCount++
						}
					} else {
						emptyCount++
					}
				}
				if emptyCount > 4 {
					count++
					grid[i][j] = 'X' // mark as counted
				}
			}
		}
	}
	return count
}

func parseGrid(lines []string) [][]rune {
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}
	return grid
}
