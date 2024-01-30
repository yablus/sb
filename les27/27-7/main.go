package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Storage interface {
	// Add
	// добавляет число в коллекцию.
	// возвращает:
	// true, если элемент добавлен;
	// false, если элемент уже присутствует.
	Add(int) bool

	// Size
	// возвращает:
	// число элементов в коллекции.
	Size() int

	// Values
	// возвращает:
	// срез с числами коллекции.
	Values() []int
}

type StorageTest struct{}

func (st *StorageTest) Add(int) bool {
	return true
}

func (st *StorageTest) Size() int {
	return 3
}

func (st *StorageTest) Values() []int {
	return []int{1, 2, 3}
}

type App struct {
	repository Storage
}

func (a *App) Run() {
	for {
		a.printInfo()
		a.printNumbers()
		if number, ok := a.inputNumber(); ok {
			a.storeNumber(number)
		} else {
			break
		}
	}
}

func (a *App) printInfo() {
	fmt.Printf("Уникальных чисел в коллекции: %v\n", a.repository.Size())
}

func (a *App) printNumbers() {
	fmt.Println("Список введённых значений:")
	fmt.Println(strings.Trim(fmt.Sprint(a.repository.Values()), "[]"))
}

func (a *App) inputNumber() (int, bool) {
	for {
		fmt.Print("Введите цифру или `end` для завершения: ")
		var str string
		fmt.Scanln(&str)
		if value, err := strconv.Atoi(str); err == nil {
			return value, true
		}
		if str == "end" {
			return 0, false
		}
		fmt.Println("Некорректный ввод")
	}
}

func (a *App) storeNumber(number int) {
	var msg string
	if ok := a.repository.Add(number); ok {
		msg = "Число %d успешно добавлено\n"
	} else {
		msg = "Число %d присутствует в коллекции\n"
	}
	fmt.Printf(msg, number)
}

type MemStorage struct {
	numbers []int
}

func NewMemStore() *MemStorage {
	return &MemStorage{
		numbers: make([]int, 0),
	}
}

func (ms *MemStorage) Add(number int) bool {
	if ms.contains(number) {
		return false
	}
	ms.numbers = append(ms.numbers, number)
	return true
}

func (ms *MemStorage) Size() int {
	return len(ms.numbers)
}

func (ms *MemStorage) Values() []int {
	var result []int
	result = append(result, ms.numbers...)
	return result
}

func (ms *MemStorage) contains(number int) bool {
	for _, value := range ms.numbers {
		if value == number {
			return true
		}
	}
	return false
}

func main() {
	//repository := &StorageTest{}
	repository := NewMemStore()
	app := &App{repository}
	app.Run()
}
