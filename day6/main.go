package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse() []OperatorList {

	var OperatorLists []OperatorList

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		for index, token := range tokens {
			if len(OperatorLists) <= index {
				OperatorLists = append(OperatorLists, OperatorList{})
			}
			if token == "+" || token == "-" || token == "*" || token == "/" {
				OperatorLists[index].Operator = Operator(token)
			} else {
				number, _ := strconv.Atoi(token)
				OperatorLists[index].Operands = append(OperatorLists[index].Operands, number)
			}
		}
	}

	return OperatorLists
}

func parsePart2() int {

	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	var splits []int
	splits = append(splits, 0)
	for i, _ := range lines[0] {
		isSplit := true
		for j := 0; j < len(lines); j++ {
			if lines[j][i] != ' ' {
				isSplit = false
				break
			}
		}
		if isSplit {
			splits = append(splits, i)
		}
	}
	splits = append(splits, len(lines[0]))
	fmt.Println("Splits:", splits)

	var operators []string
	for _, ch := range lines[len(lines)-1] {
		if ch != ' ' {
			operators = append(operators, string(ch))
		}
	}
	fmt.Println("Operators:", operators)

	total := 0

	for i := 0; i < len(splits)-1; i++ {
		start := splits[i]
		end := splits[i+1]

		columnTotal := 0
		if operators[i] == "*" {
			columnTotal = 1
		}

		for c := start; c < end; c++ {
			number := 0

			for j := 0; j < len(lines)-1; j++ {
				if lines[j][c] != ' ' {
					digit, _ := strconv.Atoi(string(lines[j][c]))
					number = number*10 + digit
				}
			}

			if number > 0 {
				if operators[i] == "+" {
					columnTotal += number
				} else {
					columnTotal *= number
				}
			}
		}

		fmt.Println("Column total:", columnTotal)
		total += columnTotal
	}

	fmt.Println("Total:", total)
	return total

}

func main() {
	//operatorLists := parse()
	//total := 0
	//for _, ol := range operatorLists {
	//	total += ol.Evaluate()
	//}
	//fmt.Println("Evaluation total:", total)

	parsePart2()

}
