/*
Программа устанавливает TCP-соединение с указанным хостом и портом, передает данные из STDIN в сокет и
выводит данные из сокета в STDOUT.
Программа также обрабатывает сигналы SIGINT и SIGTERM для корректного завершения.

Для тестирования: go run main.go --timeout=5s -host "..." -port "..."

На своем устройстве тестировал на заранее запущеном проекте L0 и подключался к NATS с соответсвующей конфигурацией

*/

package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Определение флагов
	hostFlag := flag.String("host", "", "Host (IP or domain name) to connect to")
	portFlag := flag.String("port", "", "Port to connect to")
	timeoutFlag := flag.Duration("timeout", 10*time.Second, "Connection timeout")
	flag.Parse()

	// Проверка обязательных флагов
	if *hostFlag == "" || *portFlag == "" {
		fmt.Println("Usage: go-telnet --timeout=<duration> <host> <port>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Формирование адреса для подключения
	address := fmt.Sprintf("%s:%s", *hostFlag, *portFlag)

	// Устанавливаем соединение
	conn, err := net.DialTimeout("tcp", address, *timeoutFlag)
	if err != nil {
		fmt.Println("Error connecting to the server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Канал для отслеживания сигнала завершения программы
	done := make(chan struct{})
	defer close(done)

	// Запуск горутины для чтения из сокета и вывода в STDOUT
	go func() {
		buffer := make([]byte, 1024)
		for {
			select {
			case <-done:
				return
			default:
				n, err := conn.Read(buffer)
				if err != nil {
					fmt.Println("Error reading from the server:", err)
					return
				}
				fmt.Print(string(buffer[:n]))
			}
		}
	}()

	// Запуск горутины для чтения из STDIN и записи в сокет
	go func() {
		buffer := make([]byte, 1024)
		for {
			select {
			case <-done:
				return
			default:
				n, err := os.Stdin.Read(buffer)
				if err != nil {
					fmt.Println("Error reading from STDIN:", err)
					return
				}
				_, err = conn.Write(buffer[:n])
				if err != nil {
					fmt.Println("Error writing to the server:", err)
					return
				}
			}
		}
	}()

	// Ожидание сигнала завершения программы
	waitForExitSignal()
}

// Ожидание сигнала о завершении
func waitForExitSignal() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals
	fmt.Println("\nReceived interrupt signal. Exiting...")
	os.Exit(0)
}
