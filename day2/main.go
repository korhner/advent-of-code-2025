package main

import (
	"bufio"
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

func isFake(number int) bool {
	fake := false
	numDigits := len(strconv.Itoa(number))
	if numDigits%2 == 0 {
		multiplier := int(math.Pow(10, float64(numDigits/2)))
		leftHalf := number / multiplier
		rightHalf := number % multiplier

		fake = leftHalf == rightHalf
	}

	return fake
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var stack int
	buffer := 0
	count := 0

	for {
		ch, _, err := reader.ReadRune()
		if ch == '-' {
			stack = buffer
			buffer = 0
		} else if ch == ',' || err != nil {
			count += countErrors(stack, buffer)
			buffer = 0
		} else {
			buffer = buffer*10 + int(ch-'0')
		}

		if err != nil {
			break
		}
	}

	println(count)

}
