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

func divideRoman(a, b int) int {
	if a < b {
		panic("Ошибка! Римское число не может быть меньше нуля")
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
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		22:  "XXII",
		23:  "XXIII",
		24:  "XXIV",
		25:  "XXV",
		26:  "XXVI",
		27:  "XXVII",
		28:  "XXVIII",
		29:  "XXIX",
		30:  "XXX",
		31:  "XXXI",
		32:  "XXXII",
		33:  "XXXIII",
		34:  "XXXIV",
		35:  "XXXV",
		36:  "XXXVI",
		37:  "XXXVII",
		38:  "XXXVIII",
		39:  "XXXIX",
		40:  "XL",
		41:  "XLI",
		42:  "XLII",
		43:  "XLIII",
		44:  "XLIV",
		45:  "XLV",
		46:  "XLVI",
		47:  "XLVII",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		51:  "LI",
		52:  "LII",
		53:  "LIII",
		54:  "LIV",
		55:  "LV",
		56:  "LVI",
		57:  "LVII",
		58:  "LVIII",
		59:  "LIX",
		60:  "LX",
		61:  "LXI",
		62:  "LXII",
		63:  "LXIII",
		64:  "LXIV",
		65:  "LXV",
		66:  "LXVI",
		67:  "LXVII",
		68:  "LXVIII",
		69:  "LXIX",
		70:  "LXX",
		71:  "LXXI",
		72:  "LXXII",
		73:  "LXXIII",
		74:  "LXXIV",
		75:  "LXXV",
		76:  "LXXVI",
		77:  "LXXVII",
		78:  "LXXVIII",
		79:  "LXXIX",
		80:  "LXXX",
		81:  "LXXXI",
		82:  "LXXXII",
		83:  "LXXXIII",
		84:  "LXXXIV",
		85:  "LXXXV",
		86:  "LXXXVI",
		87:  "LXXXVII",
		88:  "LXXXVIII",
		89:  "LXXXIX",
		90:  "XC",
		91:  "XCI",
		92:  "XCII",
		93:  "XCIII",
		94:  "XCIV",
		95:  "XCV",
		96:  "XCVI",
		97:  "XCVII",
		98:  "XCVIII",
		99:  "XCIX",
		100: "C",
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
			result = divideRoman(number1, number2)
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
