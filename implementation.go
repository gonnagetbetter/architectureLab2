package lab2

import (
	"fmt"
	"strings"
)

func PostfixToPrefix(postfix string) (string, error) {
    stack := make([]string, 0)
    operators := map[string]bool{"+": true, "-": true, "*": true, "/": true, "^": true}

    if postfix == "" {
        return "", fmt.Errorf("invalid expression")
    }

    for _, token := range strings.Split(postfix, " ") {
        if _, isOperator := operators[token]; isOperator {
            if len(stack) < 2 {
                return "", fmt.Errorf("not enough operands")
            }
            operand1, operand2 := stack[len(stack)-1], stack[len(stack)-2]
            stack = stack[:len(stack)-2]
            prefix := fmt.Sprintf("%s %s %s", token, operand2, operand1)
            stack = append(stack, prefix)
        } else {
            stack = append(stack, token)
        }
    }

    if len(stack) != 1 {
        return "", fmt.Errorf("invalid expression")
    }

    return stack[0], nil
}