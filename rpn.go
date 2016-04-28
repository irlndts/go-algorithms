package algorithm

import (
	"fmt"
	"strings"
)

var operators map[string]operator

func init() {
	operators = map[string]operator{
		"+": add,
		"-": subtract,
		"*": multiply,
		"/": divide,
	}
}

// Revers Polish Notation calculator
// input is a string of rpn expression
// output is a result
func Rpn(expr string) interface{} {
	tokens := strings.Split(expr, " ")
	stack := make([]int, len(tokens)/2+1)

	for _, token := range tokens {
		switch token := token.(type) {
		case int:
			fmt.Println("Int", token)
		dafault:
			fmt.Println("Default", token)
		}

	}

	return 0
}
