/*
Программа реадизует фукнцию, которая принимает переменное количество done-каналов и возвращает single-канал:

Функция chFunction использует sync.WaitGroup для отслеживания завершения всех горутин слушателей.
Горутины слушателей создаются для каждого done-канала, и они закрывают out-канал при завершении работы.
Главная горутина ждет завершения всех горутин слушателей с помощью WaitGroup и затем закрывает out-канал.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-chFunction(
		//sig(2*time.Hour),
		//sig(5*time.Minute),
		sig(1*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
		//sig(1*time.Hour),
		//sig(1*time.Minute),
	)

	fmt.Printf("done after %v", time.Since(start))
}

func chFunction(channels ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	out := make(chan interface{})

	// Функция, которая закрывает out-канал при завершении работы
	output := func(c <-chan interface{}) {
		defer wg.Done()
		for val := range c {
			out <- val
		}
	}

	// Запуск горутин для слушания каждого done-канала
	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	// Закрываем out-канал при завершении всех горутин слушателей
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
