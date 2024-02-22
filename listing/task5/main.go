package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

/*
Программа выведет: "error" т.е. значение println("error")

Это произойдет тк мы присваиваем переменной err типа error значение nil, которое возвращает функция test(), типом которого
является *customError, а не error.

Поэтому выполяется условие err != nil => видим резульатт println("error").
*/
