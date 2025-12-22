package main

import "sort"

type Line struct {
	Start *Point
	End   *Point
}

func (l *Line) IsHorizontal() bool {
	return l.Start.Y == l.End.Y
}

func (l *Line) IsVertical() bool {
	return l.Start.X == l.End.X
}

func (l *Line) Intersects(line *Line) bool {
	if l.IsHorizontal() && line.IsVertical() {
		return line.Start.X >= min(l.Start.X, l.End.X) &&
			line.Start.X <= max(l.Start.X, l.End.X) &&
			l.Start.Y >= min(line.Start.Y, line.End.Y) &&
			l.Start.Y <= max(line.Start.Y, line.End.Y)
	} else if l.IsVertical() && line.IsHorizontal() {
		return l.Start.X >= min(line.Start.X, line.End.X) &&
			l.Start.X <= max(line.Start.X, line.End.X) &&
			line.Start.Y >= min(l.Start.Y, l.End.Y) &&
			line.Start.Y <= max(l.Start.Y, l.End.Y)
	} else if l.IsHorizontal() && line.IsHorizontal() {
		if l.Start.Y != line.Start.Y {
			return false
		}
		return max(l.Start.X, l.End.X) >= min(line.Start.X, line.End.X) &&
			min(l.Start.X, l.End.X) <= max(line.Start.X, line.End.X)
	} else if l.IsVertical() && line.IsVertical() {
		if l.Start.X != line.Start.X {
			return false
		}
		return max(l.Start.Y, l.End.Y) >= min(line.Start.Y, line.End.Y) &&
			min(l.Start.Y, l.End.Y) <= max(line.Start.Y, line.End.Y)
	} else {
		panic("Non-axis-aligned lines are not supported")
	}
}

type Polygon struct {
	Lines   []*Line
	SortedX []*Point
	SortedY []*Point
}

func (p *Polygon) NewPolygon(points []*Point) *Polygon {
	var lines []*Line
	n := len(points)
	for i := 0; i < n; i++ {
		start := points[i]
		end := points[(i+1)%n] // Wrap around to the first point
		lines = append(lines, &Line{Start: start, End: end})
	}

	sortedX := make([]*Point, len(points))
	copy(sortedX, points)
	sort.Slice(sortedX, func(i, j int) bool {
		return sortedX[i].X < sortedX[j].X
	})

	sortedY := make([]*Point, len(points))
	copy(sortedY, points)
	sort.Slice(sortedY, func(i, j int) bool {
		return sortedY[i].Y < sortedY[j].Y
	})

	return &Polygon{
		Lines:   lines,
		SortedX: sortedX,
		SortedY: sortedY,
	}
}

// Calculates if the point is inside the polygon using the ray-casting algorithm
func (p *Polygon) ContainsPoint(point Point) bool {
	//intersections := 0
	//for _, line := range p.Lines {
	//
	//}
	//return intersections%2 == 1
	return false
}
