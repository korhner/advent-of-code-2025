package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operator string

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

func main() {
	operatorLists := parse()
	total := 0
	for _, ol := range operatorLists {
		total += ol.Evaluate()
	}

	fmt.Println("Evaluation total:", total)

}
