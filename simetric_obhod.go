package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	PrintFilesWithFuncFilter(".", containsDot)
}

func PrintAllFiles(path string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}
	//  проходим по списку
	for _, f := range files {
		// получаем имя элемента
		// filepath.Join — функция, которая собирает путь к элементу с разделителями
		filename := filepath.Join(path, f.Name())
		// печатаем имя элемента
		fmt.Println(filename)
		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			PrintAllFiles(filename)
		}
	}
}

func PrintAllFilesWithFilter(path string, filter string) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}

	for _, f := range files {
		fullpath := filepath.Join(path, f.Name())
		fmt.Println(fullpath, filter)
		if f.IsDir() {
			PrintAllFilesWithFilter(fullpath, filter)
		}
	}
}

func PrintFilesWithFuncFilter(path string, predicate func(string) bool) {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}

	for _, f := range files {
		fullpath := filepath.Join(path, f.Name())
		if predicate(fullpath) {
			fmt.Println(fullpath)
		}
		if f.IsDir() {
			PrintFilesWithFuncFilter(fullpath, predicate)
		}
	}
}

func containsDot(s string) bool {
	return strings.Contains(s, ".")
}
