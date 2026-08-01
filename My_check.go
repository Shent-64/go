package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 1) вставьте определение типа для []error
// 2) определите метод Error для вашего типа, который будет выводить
//    все ошибки слайса
// 3) реализуйте функцию MyCheck

type ErrorList []error

func (e ErrorList) Error() string {
	var parts []string
	for _, err := range e {
		parts = append(parts, err.Error())
	}
	return strings.Join(parts, "\n")
}

func MyCheck(input string) error {
	spaces := 0
	var errors ErrorList
	for i := 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			errors = append(errors, fmt.Errorf("found numbers"))
			break
		}
	}

	if len([]rune(input)) >= 20 {
		errors = append(errors, fmt.Errorf("line is too long"))
	}

	for i := 0; i < len(input); i++ {
		if input[i] == ' ' {
			spaces++
		}
	}
	if spaces != 2 {
		errors = append(errors, fmt.Errorf("no two spaces"))
	}
	if len(errors) == 0 {
		return nil
	}
	return errors
}

func main() {
	for {
		fmt.Printf("Укажите строку (q для выхода): ")
		reader := bufio.NewReader(os.Stdin)
		ret, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		ret = strings.TrimRight(ret, "\n")
		if ret == `q` {
			break
		}
		if err = MyCheck(ret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(`Строка прошла проверку`)
		}
	}
}
