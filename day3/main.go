package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func findMaxJoltage(arr []int, startIndex, endIndex, length int) int {
	lastAllowedPick := endIndex - length + 1
	if lastAllowedPick < startIndex || lastAllowedPick > endIndex {
		return 0
	}

	maxJoltageLeadingDigit := 0
	for i := startIndex; i <= lastAllowedPick; i++ {
		if arr[i] > maxJoltageLeadingDigit {
			maxJoltageLeadingDigit = arr[i]
		}
	}

	maxJoltageTail := 0
	for i := startIndex; i <= lastAllowedPick; i++ {
		if arr[i] != maxJoltageLeadingDigit {
			continue
		}

		tailJoltage := findMaxJoltage(arr, i+1, endIndex, length-1)
		if tailJoltage > maxJoltageTail {
			maxJoltageTail = tailJoltage
		}
	}

	return maxJoltageLeadingDigit*int(math.Pow(10, float64(length-1))) + maxJoltageTail
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	totalJoltage := 0
	totalJoltagePart2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		var arr []int
		for _, ch := range line {
			arr = append(arr, int(ch-'0'))
		}
		maxJoltage := findMaxJoltage(arr, 0, len(arr)-1, 2)
		maxJoltagePart2 := findMaxJoltage(arr, 0, len(arr)-1, 12)

		fmt.Printf("Max joltage for %v is %v\n", line, maxJoltage)
		fmt.Printf("Max joltage part2 for %v is %v\n", line, maxJoltagePart2)
		totalJoltage += maxJoltage
		totalJoltagePart2 += maxJoltagePart2
	}
	fmt.Printf("Total joltage: %v\n", totalJoltage)             // 17376
	fmt.Printf("Total joltage part 2: %v\n", totalJoltagePart2) // 17376
}
