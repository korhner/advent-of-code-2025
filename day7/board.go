package main

import (
	"errors"
	"fmt"
)

type CellType rune

const (
	EmptyCell    CellType = '.'
	StartCell    CellType = 'S'
	SplitterCell CellType = '^'
	BeamCell     CellType = '|'
)

type Cell struct {
	CellType CellType
	Count    int
}

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
			row[x].CellType = CellType(char)
			if CellType(char) == StartCell {
				startCell = Coordinate{X: x, Y: y}
				row[x].CellType = BeamCell
				row[x].Count = 1
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
			if cell.CellType == BeamCell {
				result += fmt.Sprintf("%3d", cell.Count)
			} else {
				result += fmt.Sprintf("%3c", cell.CellType)
			}
		}
		result += "\n"
	}
	return result
}

func (b *Board) SimulateTick() (splits int, err error) {
	nextGenBeams := []Coordinate{}
	for _, beam := range b.Beams {
		if beam.Y < len(b.Cells)-1 {
			if b.Cells[beam.Y+1][beam.X].CellType != SplitterCell {
				b.Cells[beam.Y+1][beam.X].Count += b.Cells[beam.Y][beam.X].Count

				if b.Cells[beam.Y+1][beam.X].CellType == EmptyCell {
					b.Cells[beam.Y+1][beam.X].CellType = BeamCell
					nextGenBeams = append(nextGenBeams, Coordinate{X: beam.X, Y: beam.Y + 1})
				}

			} else {
				splits++
				b.Cells[beam.Y+1][beam.X-1].Count += b.Cells[beam.Y][beam.X].Count
				b.Cells[beam.Y+1][beam.X+1].Count += b.Cells[beam.Y][beam.X].Count

				if b.Cells[beam.Y+1][beam.X-1].CellType == EmptyCell {
					b.Cells[beam.Y+1][beam.X-1].CellType = BeamCell
					nextGenBeams = append(nextGenBeams, Coordinate{X: beam.X - 1, Y: beam.Y + 1})
				}

				if b.Cells[beam.Y+1][beam.X+1].CellType == EmptyCell {
					b.Cells[beam.Y+1][beam.X+1].CellType = BeamCell
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
		//fmt.Print("\033[H\033[2J")
		//fmt.Println(b.String())
		//time.Sleep(300 * time.Millisecond)
	}

	fmt.Println("Splits created:", splits)

	paths := 0
	for _, path := range b.Cells[len(b.Cells)-1] {
		if path.CellType == BeamCell {
			paths += path.Count
		}
	}
	fmt.Println("Paths created:", paths)
}
