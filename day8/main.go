package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point3D struct {
	X int
	Y int
	Z int
}

func ParsePoint3D(s string) *Point3D {
	tokens := strings.Split(s, ",")

	x, errX := strconv.Atoi(tokens[0])
	y, errY := strconv.Atoi(tokens[1])
	z, errZ := strconv.Atoi(tokens[2])

	if errX != nil || errY != nil || errZ != nil {
		panic("Invalid Point3D format")
	}

	return &Point3D{
		X: x,
		Y: y,
		Z: z,
	}
}

func (p *Point3D) Distance(other *Point3D) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	dz := p.Z - other.Z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

func parse() []*Point3D {
	var points []*Point3D
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		points = append(points, ParsePoint3D(line))
	}
	return points
}

type PairDistance struct {
	Distance  float64
	PointFrom *Point3D
	PointTo   *Point3D
}

func NewPairDistance(from *Point3D, to *Point3D) *PairDistance {
	return &PairDistance{
		Distance:  from.Distance(to),
		PointFrom: from,
		PointTo:   to,
	}
}

type PointDistances struct {
	Distances []*PairDistance
}

func NewPointDistances(points []*Point3D) *PointDistances {
	var distances []*PairDistance
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			distances = append(distances, NewPairDistance(points[i], points[j]))
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	return &PointDistances{
		Distances: distances,
	}
}

func (pd *PointDistances) PopClosestPair() *PairDistance {
	closest := pd.Distances[0]
	pd.Distances = pd.Distances[1:]
	return closest
}

func part1(pointDistances *PointDistances, unionFind *UnionFind[Point3D]) {
	const iterations = 1000

	for i := 0; i < iterations; i++ {
		closestPair := pointDistances.PopClosestPair()
		if !unionFind.Union(closestPair.PointFrom, closestPair.PointTo) {
			continue
		}
	}

	sizes := []int{}
	for point, size := range unionFind.Size {
		if unionFind.Parent[point] == point {
			sizes = append(sizes, size)
		}
	}
	sort.Ints(sizes)
	top3 := sizes[len(sizes)-3:]
	product := 1
	for _, size := range top3 {
		product *= size
	}
	fmt.Println("Product of top 3 cluster sizes:", product)
}

func part2(pointDistances *PointDistances, unionFind *UnionFind[Point3D]) {
	for {
		closestPair := pointDistances.PopClosestPair()
		if !unionFind.Union(closestPair.PointFrom, closestPair.PointTo) {
			continue
		}
		if unionFind.Count == 1 {
			fmt.Println("All points connected. ", closestPair.PointFrom.X*closestPair.PointTo.X)
			break
		} else {
			numRoots := 0
			for point, parent := range unionFind.Parent {
				if point == parent {
					numRoots++
				}
			}
		}

	}
}
func main() {
	points := parse()
	pointDistances := NewPointDistances(points)

	unionFind := NewUnionFind[Point3D]()
	for _, point := range points {
		unionFind.Find(point)
	}

	part2(pointDistances, unionFind)
}
