package main

import (
	"bufio"
	"fmt"
	"os"
)

func buildMaxArray(arr []int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	maxArray := make([]int, len(arr))
	maxSoFar := arr[len(arr)-1]
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > maxSoFar {
			maxSoFar = arr[i]
		}
		maxArray[i] = maxSoFar
	}
	return maxArray
}

func findMaxJoltage(arr []int) int {
	maxArray := buildMaxArray(arr)
	maxJoltage := 0
	for i := 0; i < len(arr)-1; i++ {
		if maxJoltage >= arr[i]*10 {
			continue
		}
		joltage := arr[i]*10 + maxArray[i+1]
		if joltage > maxJoltage {
			maxJoltage = joltage
		}
	}

	return maxJoltage
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	totalJoltage := 0

	for scanner.Scan() {
		line := scanner.Text()
		var arr []int
		for _, ch := range line {
			arr = append(arr, int(ch-'0'))
		}
		maxJoltage := findMaxJoltage(arr)
		fmt.Printf("Max joltage for %v is %v\n", line, maxJoltage)
		totalJoltage += maxJoltage
	}
	fmt.Printf("Total joltage: %v\n", totalJoltage)
}
