package main

import (
	"fmt"
	"strconv"
)

// Функция возвращающая результат вычислений в римских числах
func Roman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	if number <= 0 {
		fmt.Println("Ошибка! - В Римской системе не бывает отрицательных чисел.\nПроизведите новый расчёт")
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}

// Функция определяющая диапазон римских чисел от I до X
func Arabic(number string) int {
	conversions := []struct {
		value string
		digit int
	}{
		{"X", 10},
		{"IX", 9},
		{"VIII", 8},
		{"VII", 7},
		{"VI", 6},
		{"V", 5},
		{"IV", 4},
		{"III", 3},
		{"II", 2},
		{"I", 1},
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"4", 4},
		{"5", 5},
		{"6", 6},
		{"7", 7},
		{"8", 8},
		{"9", 9},
		{"10", 10},
	}

	less_ten := 0

	for _, conversion := range conversions {
		if number == conversion.value {
			less_ten += conversion.digit
		}
	}
	return less_ten
}

// Задаём переменую объекта map для конвертации
// строковых данных в числовые (string -> int)
var convert = map[string]int{
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
	"1":    1,
	"2":    2,
	"3":    3,
	"4":    4,
	"5":    5,
	"6":    6,
	"7":    7,
	"8":    8,
	"9":    9,
	"10":   10,
}

// Задаём переменную массива в которую сохраняются
// конвертированные числовые данные int
var nums [2]int

// Задаём переменную массива для размещения вводимых данных из строки
var str [2]string

// Задаём функции check_1, check_2 определяющие правило системы счислений
func check_1() bool {
	a := strconv.Itoa(nums[0])
	return a == str[0]
}
func check_2() bool {
	a := strconv.Itoa(nums[1])
	return a == str[1]
}

// Задаём функции математических операций калькулятора
func add() int {
	return (nums[0] + nums[1])
}
func subtract() int {
	return nums[0] - nums[1]
}
func multiply() int {
	return (nums[0] * nums[1])
}
func divide() int {
	res := int(nums[0] / nums[1])
	return res
}

// Объявляем переменные данных ввода
var n1, n2, n3 string

func main() {

	for {
		// Считываем данные из консоли через пробел, где:
		// n1, n2 - числовые данные, n2 - оператор арифметических операций
		fmt.Scanln(&n1, &n2, &n3)
		str[0] = n1 // помещаем считанные данные в массив
		str[1] = n3 // переменной str

		k1, k3 := n1, n3 // промежуточные переменные k1 и k3

		// Определяем переменные c1 и c2 необходимые для условия
		// определяющего установленный интервал ввода чисел от 1 до 10
		c1, err := strconv.Atoi(k1) // преобразуем строку в число
		if err != nil {
			//log.Fatal(err)
			//continue
			str[0] = n1
		}
		c3, err := strconv.Atoi(k3)
		if err != nil {
			//log.Fatal(err)
			//continue
			str[1] = n3

		}
		// Задаём условие диапазона вводимых чисел
		if c1 > 10 || c3 > 10 || Arabic(str[0]) == 0 || Arabic(str[1]) == 0 {
			fmt.Println("Ошибка! - Используйте арабские или римские целочисленные значения от 1 до 10 разделённые пробелом.")
			break
		}

		// Цикл конвертации строковых данных в числовые посредством объекта map переменной convert
		// и размещение их в массиве переменной nums
		for key, value := range convert {
			if key == n1 {
				nums[0] = value
			}
			if key == n3 {
				nums[1] = value
			}
		}

		// Задаём условие расчёта в римской системе счисления
		if check_1() == false && check_2() == false {
			switch {
			case n2 == "+":
				fmt.Println("=", Roman(add()))
			case n2 == "-":
				fmt.Println("=", Roman(subtract()))
			case n2 == "*":
				fmt.Println("=", Roman(multiply()))
			case n2 == "/":
				fmt.Println("=", Roman(divide()), "( остаток от деления =", nums[0]%nums[1], ")")
			default:
				fmt.Println("Ошибка ввода! Используйте операнды: \"+\", \"-\", \"*\", \"/\"")
				break
			}
		}

		// Задаём условие расчета в арабской системе счисления
		if check_1() == true && check_2() == true {
			switch {
			case n2 == "+":
				fmt.Println("=", add())
			case n2 == "-":
				fmt.Println("=", subtract())
			case n2 == "*":
				fmt.Println("=", multiply())
			case n2 == "/":
				fmt.Println("=", divide(), "( остаток от деления =", nums[0]%nums[1], ")")
			default:
				fmt.Println("Ошибка ввода! Используйте операнды: \"+\", \"-\", \"*\", \"/\"")
				break
			}

		}

		// Задаём условие проверки корректно вводимых данных по предложенным системам счислений
		if check_1() == true && check_2() == false || check_1() == false && check_2() == true {
			fmt.Println("Ошибка! - Используйте одинаковые системы счислений, либо римскую, либо арабскую.")
			break
		}

	}

}
