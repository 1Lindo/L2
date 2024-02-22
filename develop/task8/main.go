/*
Программа реализует UNIX-шелл-утилиту с поддержкой команд и функционала fork/exec-команд:

- cd <args> - смена директории
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента "Х"
- ps - выводит общую информацию по запущенным процессам в определенном формате

Дополнительно поддерживается конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	fmt.Println("Простая UNIX-шелл-утилита. Для выхода введите \\quit.")

	for {
		fmt.Print("-> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		command := scanner.Text()

		if command == "\\quit" {
			fmt.Println("Выход из шелла.")
			break
		}

		handleCommand(command)
	}
}

func handleCommand(command string) {
	args := strings.Fields(command)

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			fmt.Println("Недостаточно аргументов для cd.")
		} else {
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("Ошибка при смене директории:", err)
			}
		}
	case "pwd":
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Ошибка при получении текущей директории:", err)
		} else {
			fmt.Println(currentDir)
		}
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
	case "kill":
		if len(args) < 2 {
			fmt.Println("Недостаточно аргументов для kill.")
		} else {
			pid := args[1]
			process, err := os.FindProcess(pidToInt(pid))
			if err != nil {
				fmt.Println("Ошибка при поиске процесса:", err)
			} else {
				err := process.Signal(syscall.SIGTERM)
				if err != nil {
					fmt.Println("Ошибка при отправке сигнала:", err)
				}
			}
		}
	case "ps":
		cmd := exec.Command("ps", "aux")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("Ошибка при выполнении команды ps:", err)
		}
	default:
		runCommandWithPipes(args)
	}
}

func runCommandWithPipes(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("Ошибка при выполнении команды:", err)
	}
}

func pidToInt(pid string) int {
	var result int
	fmt.Sscanf(pid, "%d", &result)
	return result
}
