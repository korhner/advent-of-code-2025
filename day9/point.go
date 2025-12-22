package main

import (
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
