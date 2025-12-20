package main

import (
	"errors"
	"fmt"
)

type Cell rune

const (
	EmptyCell    Cell = '.'
	StartCell    Cell = 'S'
	SplitterCell Cell = '^'
	BeamCell     Cell = '|'
)

type Coordinate struct {
	X int
	Y int
}

type Board struct {
	Cells [][]Cell
	Beams []Coordinate
}

func NewBoard(lines []string) *Board {
	cells := make([][]Cell, len(lines))
	var startCell Coordinate

	for y, line := range lines {
		row := make([]Cell, len(line))
		for x, char := range line {
			row[x] = Cell(char)
			if char == 'S' {
				startCell = Coordinate{X: x, Y: y}
				row[x] = BeamCell
			}
		}
		cells[y] = row
	}

	return &Board{
		Cells: cells,
		Beams: []Coordinate{startCell},
	}
}

func (b *Board) String() string {
	result := ""
	for _, row := range b.Cells {
		for _, cell := range row {
			result += string(cell)
		}
		result += "\n"
	}
	return result
}

func (b *Board) SimulateTick() (int, error) {
	nextGenBeams := []Coordinate{}
	splits := 0
	for _, beam := range b.Beams {
		if beam.Y < len(b.Cells)-1 {
			if b.Cells[beam.Y+1][beam.X] == EmptyCell {
				b.Cells[beam.Y+1][beam.X] = BeamCell
				nextGenBeams = append(nextGenBeams, Coordinate{X: beam.X, Y: beam.Y + 1})
			} else if b.Cells[beam.Y+1][beam.X] == SplitterCell {
				splits++
				if beam.X > 0 && b.Cells[beam.Y+1][beam.X-1] == EmptyCell {
					b.Cells[beam.Y+1][beam.X-1] = BeamCell
					nextGenBeams = append(nextGenBeams, Coordinate{X: beam.X - 1, Y: beam.Y + 1})
				}

				if beam.X < len(b.Cells[0])-1 && b.Cells[beam.Y+1][beam.X+1] == EmptyCell {
					b.Cells[beam.Y+1][beam.X+1] = BeamCell
					nextGenBeams = append(nextGenBeams, Coordinate{X: beam.X + 1, Y: beam.Y + 1})
				}

			}
		}
	}

	b.Beams = nextGenBeams
	if len(b.Beams) == 0 {
		return 0, errors.New("no more beams to process")
	}

	return splits, nil
}

func (b *Board) Simulate() {
	splits := 0
	for {
		newSplits, err := b.SimulateTick()
		if err != nil {
			break
		}
		splits += newSplits
		fmt.Println("-------------------------------------")
		fmt.Println(b.String())
	}

	fmt.Println("Splits created:", splits)
}
