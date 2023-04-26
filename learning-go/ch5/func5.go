package main

import (
	"fmt"
	"strconv"
)

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

type onFuncType func(int, int) int

var onMap = map[string]onFuncType{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"two", "+", "three"},
		[]string{"2", "+", "three"},
		[]string{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println(expression, "不正な式です")
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		op := expression[1]
		onFunc, ok := onMap[op]
		if !ok {
			fmt.Print(expression, " -- ", "定義されていない演算子です: ", op, "\n")
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		result := onFunc(p1, p2)
		fmt.Print(expression, " -> ", result, "\n")
	}
}
