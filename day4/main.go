package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	matrix := NewMatrixFromLines(lines)
	total := 0

	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[0]); c++ {
			countNeighbors := 0
			pos := Position{Row: r, Col: c}
			if matrix[r][c] != '@' {
				continue
			}

			for _, direction := range AllDirections {
				neighborPos := pos.Add(direction)
				ch, err := matrix.GetSafe(neighborPos.Row, neighborPos.Col)
				if err == nil && ch == '@' {
					countNeighbors++
				}
			}

			if countNeighbors < 4 {
				total++
			}
		}
	}

	fmt.Println("Total '@' with less than 4 '@' neighbors:", total)

}
