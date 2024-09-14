package main

import (
	"fmt"
	"strconv"
)

func checkNums(number1, number2 int) string {
	if number1 < 1 || number2 < 1 {
		panic("Ошибка! Неправильный диапозон")

	} else if number1 > 10 || number2 > 10 {
		panic("Ошибка! Неправильный диапозон")
	}

	return "nil"
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func subtractRoman(a, b int) int {
	if a < b {
		panic("Ошибка! Римское число не может быть меньше нуля")
	}
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	if b == 0 {
		panic("Ошибка! Деление на ноль!")

	}
	return a / b
}

func romanToInt(s1 string) int {

	switch {
	case s1 == "I":
		return 1
	case s1 == "II":
		return 2
	case s1 == "III":
		return 3
	case s1 == "IV":
		return 4
	case s1 == "V":
		return 5
	case s1 == "VI":
		return 6
	case s1 == "VII":
		return 7
	case s1 == "VIII":
		return 8
	case s1 == "IX":
		return 9
	case s1 == "X":
		return 10
	default:
		panic("Ошибка!")
	}
}

func main() {
	var num1, num2, operator string

	//fmt.Scanln(&num1, &operator, &num2)

	_, err := fmt.Scanln(&num1, &operator, &num2)
	if err != nil || len(operator) != 1 {
		panic("Ошибка: некорректный ввод")

	}

	var romanNumerals = [10]string{
		"I",
		"II",
		"III",
		"IV",
		"V",
		"VI",
		"VII",
		"VIII",
		"IX",
		"X",
	}

	var romanNumeraCon = map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}

	var arabNumerals = [10]string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
	}

	var check_roman bool
	var check_arab bool
	// код проверки римского числа
	for _, roman := range romanNumerals {
		if roman == num1 {
			check_roman = true
		} else if roman == num2 {
			check_roman = true
		}
	}
	// код проверки арабского числа
	for _, arab := range arabNumerals {
		if arab == num1 {
			check_arab = true
		} else if arab == num2 {
			check_arab = true
		}
	}

	if check_roman != true {
		// код перовода текста в число (араб)
		// АРАБСКОЕ ЧИСЛО
		number1, err := strconv.Atoi(num1)
		number2, err := strconv.Atoi(num2)

		if err != nil {
			fmt.Println("Ошибка преобразования:", err)
			return
		}
		// проверка диапозона от 1 до 10
		checkNums(number1, number2)

		// оператор

		var result int
		switch operator {
		case "+":
			result = add(number1, number2)
		case "-":
			result = subtract(number1, number2)
		case "*":
			result = multiply(number1, number2)
		case "/":
			result = divide(number1, number2)
		default:
			panic("Неправильный оператор!")
		}

		fmt.Println(result)

		// перевод строки в число (римское)
		// РИМСКОЕ ЧИСЛО
	} else if check_arab != true {
		number1 := romanToInt(num1)
		number2 := romanToInt(num2)
		checkNums(number1, number2)
		// ариф. операции
		var result int
		switch operator {
		case "+":
			result = add(number1, number2)
		case "-":
			result = subtractRoman(number1, number2)
		case "*":
			result = multiply(number1, number2)
		case "/":
			result = divide(number1, number2)
		default:
			panic("Неправильный оператор!")
		}

		for i, v := range romanNumeraCon {
			if result == i {
				fmt.Println(v)
			}
		}

	} else {
		panic("Ошибка! Ввод не соответсвует ожидаемому")
	}

}
