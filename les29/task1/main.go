/*
Задание 1. Конвейер
Цели задания
Научиться работать с каналами и горутинами.
Понять, как должно происходить общение между потоками.
Что нужно сделать
Реализуйте паттерн-конвейер:

Программа принимает числа из стандартного ввода в бесконечном цикле и передаёт число в горутину.
Квадрат: горутина высчитывает квадрат этого числа и передаёт в следующую горутину.
Произведение: следующая горутина умножает квадрат числа на 2.
При вводе «стоп» выполнение программы останавливается.
Советы и рекомендации
Воспользуйтесь небуферизированными каналами и waitgroup.

Что оценивается
Ввод : 3

Квадрат : 9

Произведение : 18
*/

package main

import (
	"fmt"
	"strconv"
	"sync"
)

func sq(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	num *= num
	fmt.Println("Квадрат:", num)
	var wgs sync.WaitGroup
	wgs.Add(1)
	go ml(num, &wgs)
	wgs.Wait()
}

func ml(num int, wgs *sync.WaitGroup) {
	defer wgs.Done()
	num *= 2
	fmt.Println("Произведение:", num)
}

func main() {
	fmt.Println("Конвейер")
	fmt.Println("========")

	for {
		fmt.Println("Введите число или \"стоп\" для выхода")
		var str string
		fmt.Scan(&str)
		if str == "стоп" {
			break
		}
		if num, err := strconv.Atoi(str); err == nil {
			fmt.Println("Ввод:", num)
			var wg sync.WaitGroup
			wg.Add(1)
			go sq(num, &wg)
			wg.Wait()
		} else {
			fmt.Println("Некорректный ввод")
		}
	}
	fmt.Println("Программа завершена")
}
