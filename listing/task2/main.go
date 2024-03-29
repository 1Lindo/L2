package main

import (
	"fmt"
)

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}

/*
Программа выведет: 2
				   1

test() :
1) Значение x объявляется в сигнатуре функции, далее x устанавливается равным 1;
2) После того, как возвращаемое значение уже было установлено, выполняется defer;
3) Далее выполнение функции доходит до return и возвращается текущее значение x 2.

anotherTest() :
1) В этой функции инициализируется переменная x nil;
2) Затем выполняется defer, которая увеличивает x на 1, но это происходит только в рамках этой фунции и не меняет переменную;
3) Присваивается значение x = 1, которое выводится далее. Поэтому результат - 1.

*/
