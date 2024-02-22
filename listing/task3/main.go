package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}

/*
Результат работы программы: <nil>
							false

Foo() возвращает nil => переменная err в функции main() будет равна nil. Поэтому выводится <nil>.
Затем проверка err == nil также возвращает false, поскольку возвращаемый тип Foo() - это *os.PathError, а не error.
С точки зрения строгой типизации Go, err == nil считается false, поскольку *os.PathError не является точным nil.
*/
