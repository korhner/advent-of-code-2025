package main

import (
	"bufio"
	"math"
	"os"
)

func parsePoints() []Point {
	var points []Point
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		points = append(points, *NewPointFromString(line))
	}
	return points
}

func part1(points []Point) {
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

func part2(points []Point) {
	// create the polygon
	// find the bounding box

	// slowly reduce the bounding box along sides until it fits inside the polygon

	// implement a function to check if a rectangle fits inside the polygon
	// it fits if all four corners are inside the polygon
	// and no edges intersect with the polygon edges
	// a point is inside the polygon if a ray from the point to infinity intersects the polygon edges an odd number of times

	// keep track of the max area
}

func main() {
	points := parsePoints()
	part1(points)
}
