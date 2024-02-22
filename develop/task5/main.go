/*
Программа реализует утилиту фильтрации по аналогии с консольной утилитой (man grep).

Пример команды для запуска:
go run main.go -i -v -n "место для паттерна" text.txt

Основные шаги программы:

1) Программа парсит флаги во flag;
2) Получаем паттерн для поиска и список файлов для обработки из аргументов командной строки;
3) Обработка ввода: Программа либо читает стандартный ввод, либо обрабатывает каждый указанный файл;
4) Обработка каждой строки: Каждая строка из входного потока (файла или стандартного ввода) проверяется на соответствие
фильтрам и паттерну;
5) Совпадающие строки выводятся в соответствии с параметрами (номер строки, количество строк до
и после совпадения и т. д.);
6) Если указан флаг -c, программа выводит общее количество совпадающих строк.
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Определение флагов
	afterLines := flag.Int("A", 0, "Печатать N строк после совпадения")
	beforeLines := flag.Int("B", 0, "Печатать N строк до совпадения")
	contextLines := flag.Int("C", 0, "Печатать ±N строк вокруг совпадения")
	countOnly := flag.Bool("c", false, "Количество строк")
	ignoreCase := flag.Bool("i", false, "Игнорировать регистр")
	invertMatch := flag.Bool("v", false, "Исключать совпадения")
	fixedString := flag.Bool("F", false, "Точное совпадение со строкой, не паттерн")
	lineNumber := flag.Bool("n", false, "Напечатать номер строки")

	flag.Parse()

	// Получение паттерна поиска
	pattern := flag.Arg(0)
	if pattern == "" {
		fmt.Println("Необходимо указать паттерн для поиска.")
		os.Exit(1)
	}

	// Получение файлов для поиска
	files := flag.Args()[1:]

	if len(files) == 0 {
		// Если файлы не указаны, читаем стандартный ввод
		processInput(os.Stdin, pattern, *afterLines, *beforeLines, *contextLines, *countOnly, *ignoreCase, *invertMatch, *fixedString, *lineNumber)
	} else {
		// Обработка каждого указанного файла
		for _, file := range files {
			file, err := os.Open(file)
			if err != nil {
				fmt.Printf("Ошибка открытия файла %s: %s\n", file, err)
				continue
			}
			defer file.Close()

			processInput(file, pattern, *afterLines, *beforeLines, *contextLines, *countOnly, *ignoreCase, *invertMatch, *fixedString, *lineNumber)
		}
	}
}

func processInput(input *os.File, pattern string, after, before, context int, countOnly, ignoreCase, invertMatch, fixedString, lineNumber bool) {
	scanner := bufio.NewScanner(input)
	lineCount := 0
	matchCount := 0
	linesAfterMatch := 0
	linesBeforeMatch := 0
	matchingLines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		lineCount++

		// Применение фильтров
		matched := matchLine(line, pattern, ignoreCase, fixedString)
		if invertMatch {
			matched = !matched
		}

		if matched {
			matchCount++
			linesAfterMatch = after
			linesBeforeMatch = before
			matchingLines = append(matchingLines, line)
		} else if linesAfterMatch > 0 || linesBeforeMatch > 0 {
			if linesBeforeMatch > 0 {
				fmt.Println(line)
				linesBeforeMatch--
			} else {
				matchingLines = append(matchingLines, line)
				linesAfterMatch--
			}
		} else {
			if len(matchingLines) > 0 {
				printMatchingLines(matchingLines, lineNumber)
				matchingLines = nil
			}
		}
	}

	// Вывод последних совпадающих строк (если есть)
	if len(matchingLines) > 0 {
		printMatchingLines(matchingLines, lineNumber)
	}

	// Вывод статистики (если нужно)
	if countOnly {
		fmt.Printf("Всего совпадений: %d\n", matchCount)
	}
}

func matchLine(line, pattern string, ignoreCase, fixedString bool) bool {
	if ignoreCase {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}

	if fixedString {
		return line == pattern
	} else {
		return strings.Contains(line, pattern)
	}
}

func printMatchingLines(lines []string, lineNumber bool) {
	for i, line := range lines {
		if lineNumber {
			fmt.Printf("%d:%s\n", i+1, line)
		} else {
			fmt.Println(line)
		}
	}
}
