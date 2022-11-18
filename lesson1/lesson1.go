package main

import (
	"errors"
	"fmt"
	"strconv"
)

// Число Фибоначчи
// Реализовать функцию, которая возвращает следующее в ряду число
// Фибоначчи
func fib() func() int {
	first := 0
	second := 1

	// returns the sequence starting with 0 1 1 etc
	return func() int {
		next := first
		first = second
		second = next + second
		return next
	}
}

// Арифметическая прогрессия
// Реализовать функцию, которая принимает на вход правило, так же в виде
// функции. Правило определяет an от an-1
// Возвращает генератор который последовательно возвращает очередной
// член прогрессии
func progression(a0 int, generator func(int) int) func() int {
	aCurrent := a0
	getNext := generator

	return func() int {
		next := aCurrent
		aCurrent = getNext(aCurrent)
		return next
	}
}

func test1() {
	fib1 := fib()

	for i := 0; i < 10; i++ {
		fmt.Println(fib1())
	}

	prog := progression(0, func(a int) int { return a + 5 })

	for i := 0; i < 10; i++ {
		fmt.Println(prog())
	}
}

// Определить есть ли в строке повторяющиеся буквы латинского алфавита,
// игнорируя регистр
// Можно использовать только один цикл range либо for и только одну
// переменную int64 помимо итератора
// Дополнительные пакеты использовать нельзя
func dupLetter(s string) bool {
	var cache int64 = 0

	// convert the letter code to a 0-based index
	norm := func(c rune) int8 {
		if c < 'A' || c > 'z' {
			panic("not an english letter!")
		}
		if c > 'Z' && c < 'a' {
			panic("not an english letter!")
		}

		if c > 'a' {
			return int8(c - 'a')
		}
		return int8(c - 'A')
	}

	for _, c := range s {
		idx := norm(c)

		var bitmask int64 = 1 << idx
		if cache&bitmask != 0 {
			return true
		}
		cache |= bitmask
	}
	return false
}

func test2() {
	str1 := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str2 := str1 + "z"
	str3 := ""
	str4 := "p"
	// str5 := "abcdefj0klmn" // dupLetter(str5) panics

	fmt.Println("false ", dupLetter(str1))
	fmt.Println("true", dupLetter(str2))
	fmt.Println("false ", dupLetter(str3))
	fmt.Println("false ", dupLetter(str4))
}

// Написать функции eval, plus, minus так чтобы
// eval(10, 20, plus, "45", minus) => -15
// eval(10, 2.5, plus) => 12.5
// В случае если вычислить не удается, предусмотреть возвращение error
// Для преобразования пакет strconv:
// strconv.Atoi(“10”) - для целых
// strconv.ParseFloat("1.23", 64) - для float
func plus(l, r float64) float64 {
	return l + r
}
func minus(l, r float64) float64 {
	return l - r
}

func parseArg(arg interface{}) (float64, error) {
	switch v := arg.(type) {
	case int:
		return float64(v), nil
	case float32:
		return float64(v), nil
	case float64:
		return v, nil
	case string:
		if value, err := strconv.Atoi(v); err == nil {
			return float64(value), nil
		} else if value, err := strconv.ParseFloat(v, 64); err == nil {
			return value, nil
		}
		return 0, errors.New("String is not convertible to int or float")
	default:
		fmt.Println("###", arg)
		return 0, errors.New("Unexpected argument type! Need a value")
	}
}

func eval(args ...any) (float64, error) {
	if len(args) < 3 {
		return 0, errors.New("Too few arguments, need at least 3")
	}

	firstArg, err1 := parseArg(args[0])
	secondArg, err2 := parseArg(args[1])

	if err1 != nil {
		return 0., err1
	}
	if err2 != nil {
		return 0., err2
	}

	result := 0.

	switch v := args[2].(type) {
	case func(float64, float64) float64:
		result = v(firstArg, secondArg)
	default:
		return 0, errors.New("Unexpected argument type! Need an operation")
	}

	if len(args) > 3 {
		newArgs := append([]any{result}, args[3:]...)
		//fmt.Println("Going deeper with args=", newArgs)
		return eval(newArgs...)
	}
	return result, nil
}

func test3() {
	// eval(10, 20, plus, "45", minus) => -15
	// eval(10, 2.5, plus) => 12.5
	res1, err := eval(10, 20, plus, "45", minus)
	res2, _ := eval(10, 2.5, plus)
	fmt.Println("eval(10, 20, plus, \"45\", minus) => -15, actual=", res1, err)
	fmt.Println("eval(10, 2.5, plus) => 12.5, actual=", res2)
}

func main() {

	test1()
	test2()
	test3()

}
