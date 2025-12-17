package main

import (
	"bufio"
	"fmt"
	"os"
)

func countNeighbors(matrix Matrix, pos Position) int {
	countNeighbors := 0
	for _, direction := range AllDirections {
		neighborPos := pos.Add(direction)
		ch, err := matrix.GetSafe(neighborPos.Row, neighborPos.Col)
		if err == nil && ch == '@' {
			countNeighbors++
		}
	}

	return countNeighbors
}

func part1(matrix Matrix) {
	total := 0

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {

			pos := Position{Row: r, Col: c}
			if matrix[r][c] != '@' {
				continue
			}
			count := countNeighbors(matrix, pos)
			if count < 4 {
				total++
			}
		}
	}

	fmt.Println("Total '@' with less than 4 '@' neighbors:", total)
}

func part2(matrix Matrix) {

	visited := NewVisitedMatrix(matrix)
	queue := []Position{}
	totalRemoved := 0

	// Initialize the queue with all positions containing '@'
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			if matrix[r][c] == '@' {
				queue = append(queue, Position{Row: r, Col: c})
			}
		}
	}

	for len(queue) > 0 {
		currentPos := queue[0]
		queue = queue[1:]

		count := countNeighbors(matrix, currentPos)
		if visited[currentPos.Row][currentPos.Col] || count > 3 {
			continue
		}

		matrix[currentPos.Row][currentPos.Col] = '.'
		visited[currentPos.Row][currentPos.Col] = true
		totalRemoved++

		for _, direction := range AllDirections {
			val, err := matrix.GetSafe(currentPos.Row+direction.Row, currentPos.Col+direction.Col)
			if err == nil && val == '@' && !visited[currentPos.Row+direction.Row][currentPos.Col+direction.Col] {
				queue = append(queue, Position{Row: currentPos.Row + direction.Row, Col: currentPos.Col + direction.Col})
			}
		}
	}

	fmt.Println("Total '@' with part 2 removed:", totalRemoved)
}

func main() {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	matrix := NewMatrixFromLines(lines)
	part1(matrix)
	part2(matrix)
}
