package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Clock struct {
	CurrentState int
	TouchedZero  int
}

func (c *Clock) Rotate(degrees int) {

	c.TouchedZero += int(math.Abs(float64(degrees)) / 100)
	degrees = degrees % 100

	newState := c.CurrentState + degrees
	if newState < 0 {
		newState += 100
	} else if newState >= 100 {
		newState -= 100
	}

	if c.CurrentState != 0 &&
		(newState == 0 ||
			(degrees < 0 && newState > c.CurrentState) ||
			(degrees > 0 && newState < c.CurrentState)) {
		c.TouchedZero++
	}
	println("Rotated from", c.CurrentState, "by:", degrees, "to:", newState, "TouchedZero:", c.TouchedZero)
	c.CurrentState = newState
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Error closing file: %v\n", err)
		}
	}()

	clock := &Clock{CurrentState: 50}
	numZeros := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}
		direction := 1
		if line[0] == 'L' {
			direction = -1
		}
		degrees, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing degrees from line: %s\n", line)
			continue
		}
		clock.Rotate(direction * degrees)
		if clock.CurrentState == 0 {
			numZeros++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Number of times clock hit 0:", numZeros)
	fmt.Println("Number of times clock touched 0:", clock.TouchedZero)
}
