/*
Программа печатает точное время с использованием NTP -библиотеки.
Программа корректно обрабатываешь ошибки библиотеки в STDERR и возвращает ненулевой код выхода в OS.

Инициализирована как go module. Используется библиотеку github.com/beevik/ntp.
*/

package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	//инициализации переменной с ntp временем
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting NTP time: %v\n", err)
		os.Exit(1)
	}
	//инициализации переменной с обычным временем
	localTime := time.Now()

	//вывод вариантов
	fmt.Printf("Local Time: %s\n", localTime.Format(time.RFC3339))
	fmt.Printf("Exact Time: %s\n", ntpTime.Format(time.RFC3339))
}
