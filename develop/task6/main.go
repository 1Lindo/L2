/*
Программа реализует утилиту аналог консольной команды cut (man cut).
Утилита принимает строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Пооддерживаются следующие ключи:
1) -f - "fields" - выбрать поля (колонки);
2) -d - "delimiter" - использовать другой разделитель;
3) -s - "separated" - только строки с разделителем.

Для тестирования:
echo -e "1\t2\t3\n4\t5\t6" | go run main.go -f 1,3 -d "\t"

Этот код сначала разбивает вводимые строки на колонки с использованием разделителя (по умолчанию, табуляция), а затем
выбирает и выводит указанные поля (колонки). Кроме того, есть возможность использовать другой разделитель и
выводить только строки, содержащие разделитель (если установлен соответствующий флаг).
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Определение флагов
	fields := flag.String("f", "", "выбрать поля (колонки)")
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()

	// Чтение строк из STDIN
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Обработка только строк с разделителем, если флаг установлен
		if *separated && !strings.Contains(line, *delimiter) {
			continue
		}

		// Разбивка строки на колонки с использованием указанного разделителя
		columns := strings.Split(line, *delimiter)

		// Выбор и вывод указанных полей (колонок)
		if *fields != "" {
			selectedFields := parseFields(*fields)
			var selectedColumns []string

			for _, idx := range selectedFields {
				if idx >= 1 && idx <= len(columns) {
					selectedColumns = append(selectedColumns, columns[idx-1])
				} else {
					selectedColumns = append(selectedColumns, "")
				}
			}

			fmt.Println(strings.Join(selectedColumns, *delimiter))
		} else {
			// Если не указаны поля, выводим всю строку
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка при чтении ввода:", err)
		os.Exit(1)
	}
}

// parseFields разбирает строку с номерами полей (колонок) и возвращает их в виде среза int.
func parseFields(fields string) []int {
	var selectedFields []int
	fieldStrs := strings.Split(fields, ",")

	for _, fieldStr := range fieldStrs {
		fieldStr = strings.TrimSpace(fieldStr)
		if fieldIdx, err := parseFieldIndex(fieldStr); err == nil {
			selectedFields = append(selectedFields, fieldIdx)
		}
	}

	return selectedFields
}

// parseFieldIndex преобразует строку с номером поля в int.
func parseFieldIndex(fieldStr string) (int, error) {
	return strconv.Atoi(fieldStr)
}
