/*
Программа читает строки из указанного файла, сортирует их с учетом заданных параметров и записывает отсортированные
строки в файл sorted_output.txt.

Программа будет сортировать строки из файла text.txt по числовому значению во второй колонке в
обратном порядке, удаляя повторяющиеся строки.

go run main.go -file text.txt -k 2 -n -r -u
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filePath := flag.String("file", "", "Путь к входному файлу")
	column := flag.Int("k", 0, "Номер колонки для сортировки (по умолчанию 0 - вся строка)")
	numericSort := flag.Bool("n", false, "Сортировать по числовому значению")
	reverseSort := flag.Bool("r", false, "Сортировать в обратном порядке")
	unique := flag.Bool("u", false, "Не выводить повторяющиеся строки")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Укажите путь к входному файлу")
		return
	}

	//Определение файла для чтения
	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	//Инициализация сканера
	scanner := bufio.NewScanner(file)
	var lines []string

	//Скан текста из файла в заранее инициализированный слайс lines
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	sort.SliceStable(lines, func(i, j int) bool {
		var key1, key2 string

		if *column > 0 {
			fields1 := strings.Fields(lines[i])
			fields2 := strings.Fields(lines[j])

			if len(fields1) >= *column {
				key1 = fields1[*column-1]
			}

			if len(fields2) >= *column {
				key2 = fields2[*column-1]
			}
		} else {
			key1 = lines[i]
			key2 = lines[j]
		}

		if *numericSort {
			num1, err1 := strconv.Atoi(key1)
			num2, err2 := strconv.Atoi(key2)

			if err1 == nil && err2 == nil {
				key1 = strconv.Itoa(num1)
				key2 = strconv.Itoa(num2)
			}
		}

		if key1 < key2 {
			return !*reverseSort
		} else if key1 > key2 {
			return *reverseSort
		} else {
			return false
		}
	})

	if *unique {
		var uniqueLines []string
		var prevLine string

		for _, line := range lines {
			if line != prevLine {
				uniqueLines = append(uniqueLines, line)
				prevLine = line
			}
		}

		lines = uniqueLines
	}

	//Ниже представлена создание файла с отсортированными данными  и иницилизация writer, который запишет туда данные
	outputFile, err := os.Create("sorted_output.txt")
	if err != nil {
		fmt.Println("Ошибка создания файла для записи:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()

	fmt.Println("Файл успешно отсортирован. Результат сохранен в sorted_output.txt")
}
