package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func countErrors(from, to int) int {
	errors := 0
	for number := from; number <= to; number++ {
		if isFake(number) {
			errors += number
		}
	}
	return errors
}

func countErrorsPart2(from, to int) int {
	errors := 0
	for number := from; number <= to; number++ {
		if isFakePart2(number) {
			errors += number
		}
	}
	return errors
}

func isFake(number int) bool {
	return isFakeCustomParts(number, len(strconv.Itoa(number))/2)
}

func isFakeCustomParts(number, tailLen int) bool {
	numDigits := len(strconv.Itoa(number))

	if tailLen == 0 || tailLen == numDigits || numDigits%tailLen != 0 || numDigits <= 1 {
		return false
	}

	multiplier := int(math.Pow(10, float64(tailLen)))
	part := number % multiplier

	for number != 0 {
		currentPart := number % multiplier
		if currentPart != part {
			return false
		}
		number = number / multiplier
	}

	return true
}

func isFakePart2(number int) bool {
	numDigits := len(strconv.Itoa(number))
	for parts := 1; parts <= numDigits/2; parts++ {
		if isFakeCustomParts(number, parts) {
			fmt.Println("Number", number, "is fake with parts", parts)
			return true
		}
	}

	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var stack int
	buffer := 0
	count := 0
	countPart2 := 0

	for {
		ch, _, err := reader.ReadRune()
		if ch == '-' {
			stack = buffer
			buffer = 0
		} else if ch == ',' || err != nil {
			count += countErrors(stack, buffer)
			countPart2 += countErrorsPart2(stack, buffer)
			buffer = 0
		} else {
			buffer = buffer*10 + int(ch-'0')
		}

		if err != nil {
			break
		}
	}

	println(count)
	println(countPart2)

}
