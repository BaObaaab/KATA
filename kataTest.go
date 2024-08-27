package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var RomanNumbers = map[string]int{
	"I": 1, "II": 2, "III": 3,
	"IV": 4, "V": 5, "VI": 6,
	"VII": 7, "VIII": 8, "IX": 9,
	"X": 10, "L": 50, "C": 100,
}

func findOp(line string) (string, error) {
	switch {
	case strings.Contains(line, "+"):
		return "+", nil
	case strings.Contains(line, "-"):
		return "-", nil
	case strings.Contains(line, "*"):
		return "*", nil
	case strings.Contains(line, "/"):
		return "/", nil
	default:
		return "", fmt.Errorf("can't find operator")
	}
}

func numericOperation(a, b int, op string) (num int, err error) {
	switch op {
	case "+":
		num = a + b
	case "-":
		num = a - b
	case "*":
		num = a * b
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		num = a / b
	default:
		return 0, errors.New("unsupported operation")
	}

	return num, nil
}

func isItRoman(line string) bool {
	if len(line) == 0 {
		return false
	}
	for _, i := range line {
		if _, ok := RomanNumbers[string(i)]; ok {
			return true
		}
	}
	return false
}

func intToRoman(num int) string {
	roman := ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return roman
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите выражение: ")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	line = strings.ReplaceAll(line, " ", "")

	op, err := findOp(line)
	if err != nil {
		fmt.Println(err)
		return
	}

	arguments := strings.Split(line, op)
	if len(arguments) != 2 {
		fmt.Println("incorrect format of expression!")
		return
	}

	firstArgument := arguments[0]
	secondArgument := arguments[1]
	
	var a, b int
	var isRoman bool

	if isItRoman(firstArgument) && isItRoman(secondArgument) {
		a = RomanNumbers[firstArgument]
		b = RomanNumbers[secondArgument]
		isRoman = true
	} else {
		a, err = strconv.Atoi(firstArgument)
		if err != nil {
			fmt.Println("conflicting arguments")
			return
		}

		b, err = strconv.Atoi(secondArgument)
		if err != nil {
			fmt.Println("conflicting arguments")
			return
		}
	}

	if a < 1 || a > 10 || b < 0 || b > 10 {
		fmt.Println("Numbers cannot be less 0 or more than 10")
		return
	}

	result, err := numericOperation(a, b, op)
	if err != nil {
		fmt.Println(err)
		return
	}

	if isRoman {
		if result <= 0 {
			fmt.Println("in Roman numerals, the result cannot be less than or equal to zero")
		} else {
			res := intToRoman(result)
			fmt.Printf("Result: %s\n", res)
		}
	} else {
		fmt.Printf("Result: %d\n", result)
	}
}
