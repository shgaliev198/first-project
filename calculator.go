package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArabic = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

var arabicToRoman = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
	"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
	"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
}

func main() {
	fmt.Println("Введите выражение (например, 3 + 5 или IV + II):")
	var input string
	fmt.Scanln(&input)

	tokens := strings.Split(input, " ")

	if len(tokens) != 3 {
		fmt.Println("Ошибка: неверный формат ввода")
		os.Exit(1)
	}

	a, b := tokens[0], tokens[2]
	operator := tokens[1]

	isRomanA := isRoman(a)
	isRomanB := isRoman(b)

	if isRomanA != isRomanB {
		fmt.Println("Ошибка: нельзя смешивать римские и арабские числа")
		os.Exit(1)
	}

	if isRomanA {
		result, err := calculateRoman(a, b, operator)
		if err != nil {
			fmt.Println("Ошибка:", err)
			os.Exit(1)
		}
		fmt.Println(result)
	} else {
		result, err := calculateArabic(a, b, operator)
		if err != nil {
			fmt.Println("Ошибка:", err)
			os.Exit(1)
		}
		fmt.Println(result)
	}
}

func isRoman(s string) bool {
	_, exists := romanToArabic[s]
	return exists
}

func calculateArabic(aStr, bStr, operator string) (int, error) {
	a, err := strconv.Atoi(aStr)
	if err != nil || a < 1 || a > 10 {
		return 0, errors.New("некорректное число")
	}
	b, err := strconv.Atoi(bStr)
	if err != nil || b < 1 || b > 10 {
		return 0, errors.New("некорректное число")
	}

	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, errors.New("неизвестная операция")
	}
}

func calculateRoman(aStr, bStr, operator string) (string, error) {
	a, okA := romanToArabic[aStr]
	b, okB := romanToArabic[bStr]

	if !okA || !okB || a < 1 || a > 10 || b < 1 || b > 10 {
		return "", errors.New("некорректное римское число")
	}

	var result int

	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return "", errors.New("деление на ноль")
		}
		result = a / b
	default:
		return "", errors.New("неизвестная операция")
	}

	if result < 1 {
		return "", errors.New("результат римской операции не может быть меньше I")
	}

	return arabicToRoman[result], nil
}
