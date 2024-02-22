/*
Программа осуществляет примитивную распаковку строки, которая содержит повторяющиеся символы/руны, например:

"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

*/

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	fmt.Println(unicode.IsDigit('e'))
	fmt.Println(unpackString("a4bc2d5e"))  // Output: "aaaabccddddde"
	fmt.Println(unpackString("abcd"))      // Output: "abcd"
	fmt.Println(unpackString("45"))        // Output: ""
	fmt.Println(unpackString("qwe\\4\\5")) // Output: "qwe45" (*)
	fmt.Println(unpackString("qwe\\45"))   // Output: "qwe44444" (*)
	fmt.Println(unpackString("qwe\\"))     // Output:  "qwe\\" (*)
}

func unpackString(s string) string {
	//переменные для поледних прочитанных рун/строк
	var lastRune, lastLetter rune
	//переменная для результата распоковки строки и num для определения сколько раз нужно продублировать руну
	var result, num strings.Builder
	//
	var esc bool
	//сброс значений
	result.Reset()
	num.Reset()
	//заданы значения поумолчанию
	lastRune = 0
	lastLetter = 0

	for i, currentRune := range s {
		//проверка на руны на 0
		if unicode.IsDigit(currentRune) && i == 0 {
			return ""
		}
		//провека руны руны на letter, если letter => конкатинирует к result руну, а lastRune и lastLetter дает новые значения
		if unicode.IsLetter(currentRune) {
			// проверка предыдущей руны, если это число =>
			if unicode.IsDigit(lastRune) {
				numRunes, err := strconv.Atoi(num.String())
				if err != nil {
					log.Fatal(err)
				}
				for j := 0; j < numRunes-1; j++ {
					result.WriteRune(lastLetter)
				}
				num.Reset()
			}
			// если currentRUne идет не после числа
			result.WriteRune(currentRune)
			lastRune = currentRune
			lastLetter = currentRune
		}

		//проверка руны на число, елси число =>
		if unicode.IsDigit(currentRune) {
			//защита от числового значения, если int =>
			if esc {
				result.WriteRune(currentRune)
				lastLetter = currentRune
				lastRune = currentRune
				esc = false
			} else {
				// если предыдущая руна == строка =>скидываем значение у num
				if unicode.IsLetter(lastRune) {
					num.Reset()
				}
				// определяем нвое значение для num, а предыдущей руне задаем значение currentRune
				num.WriteRune(currentRune)
				lastRune = currentRune

				if i == utf8.RuneCountInString(s)-1 {
					numRunes, err := strconv.Atoi(num.String())
					if err != nil {
						log.Fatal(err)
					}
					for j := 0; j < numRunes-1; j++ {
						result.WriteRune(lastLetter)
					}
				}
			}
		}
		if currentRune == '\\' {
			if lastRune == '\\' {
				result.WriteRune(currentRune)
				lastLetter = currentRune
				lastRune = currentRune
				esc = false

			} else {
				esc = true
				lastRune = currentRune
			}
		}
	}
	return result.String()
}
