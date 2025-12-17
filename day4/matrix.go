package main

import "errors"

type Matrix [][]rune

type Position struct {
	Row int
	Col int
}

func (p Position) Add(other Position) Position {
	return Position{
		Row: p.Row + other.Row,
		Col: p.Col + other.Col,
	}
}

var (
	North     = Position{-1, 0}
	South     = Position{1, 0}
	West      = Position{0, -1}
	East      = Position{0, 1}
	NorthEast = Position{-1, 1}
	NorthWest = Position{-1, -1}
	SouthEast = Position{1, 1}
	SouthWest = Position{1, -1}

	AllDirections = []Position{
		North, South, West, East,
		NorthEast, NorthWest, SouthEast, SouthWest,
	}
)

func NewMatrix(rows, cols int) Matrix {
	m := make(Matrix, rows)
	for i := range m {
		m[i] = make([]rune, cols)
	}
	return m
}

func NewMatrixFromLines(lines []string) Matrix {
	rows := len(lines)
	cols := len(lines[0])
	m := NewMatrix(rows, cols)
	for r, line := range lines {
		for c, ch := range line {
			m[r][c] = ch
		}
	}
	return m
}

func (m Matrix) GetSafe(row, col int) (rune, error) {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
		return 0, errors.New("index out of bounds")
	}
	return m[row][col], nil
}
