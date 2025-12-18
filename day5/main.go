package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse() ([]Interval, []int) {
	scanner := bufio.NewScanner(os.Stdin)
	parsingIntervals := true
	var intervals []Interval
	var points []int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parsingIntervals = false
			continue
		}

		if parsingIntervals {
			// split line by -
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])

			intervals = append(intervals, Interval{
				Start: start,
				End:   end,
			})
		} else {
			point, _ := strconv.Atoi(line)
			points = append(points, point)
		}

	}

	return intervals, points
}
func main() {
	intervals, points := parse()

	fresh := 0
	intervalsSet := NewIntervalSet(intervals...)
	fmt.Println("Merged intervals:", intervalsSet.toString())
	for _, point := range points {
		if intervalsSet.Contains(point) {
			fresh++
			fmt.Println("Point", point, "is fresh")
		}
	}

	fmt.Println("Total fresh points:", fresh)
	fmt.Println("Total ranges covered:", intervalsSet.Range())

}
