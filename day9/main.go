package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

func NewPointFromString(s string) *Point {
	tokens := strings.Split(s, ",")
	x, _ := strconv.Atoi(tokens[0])
	y, _ := strconv.Atoi(tokens[1])
	return &Point{
		X: x,
		Y: y,
	}
}

func parsePoints() []Point {
	var points []Point
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		points = append(points, *NewPointFromString(line))
	}
	return points
}

func main() {
	points := parsePoints()
	maxArea := 0.0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			area := (math.Abs(float64(points[i].X-points[j].X)) + 1) * (math.Abs(float64(points[i].Y-points[j].Y)) + 1)

			if area > maxArea {
				maxArea = area
			}
		}
	}

	println("Max Area:", int(maxArea))
}
