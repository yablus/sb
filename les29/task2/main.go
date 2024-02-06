/*
Задание 2. Graceful shutdown
Цель задания
Научиться правильно останавливать приложения.

Что нужно сделать
В работе часто возникает потребность правильно останавливать приложения. Например, когда наш сервер обслуживает
соединения, а нам хочется, чтобы все текущие соединения были обработаны и лишь потом произошло выключение сервиса.
Для этого существует паттерн graceful shutdown.

Напишите приложение, которое выводит квадраты натуральных чисел на экран, а после получения сигнала ^С обрабатывает
этот сигнал, пишет «выхожу из программы» и выходит.

Советы и рекомендации
Для реализации данного паттерна воспользуйтесь каналами и оператором select с default-кейсом.

Что оценивается
Код выводит квадраты натуральных чисел на экран, после получения ^С происходит обработка сигнала и выход из программы.
*/

/*
package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
)

func waitSignal(chanSignal chan os.Signal) {
	<-chanSignal
	fmt.Println("выхожу из программы")
	os.Exit(1) // не знаю, корректно ли... Ведь он по идее убивает все другие действующие горутины...
	// но для этой задачи по-идее подойдет... буду рад рекомендациям.
}

func main() {
	chanSignal := make(chan os.Signal, 1)
	signal.Notify(chanSignal,
		os.Interrupt,
	)

	go waitSignal(chanSignal)

	for {
		fmt.Println("введите число")
		var str string
		fmt.Scan(&str)
		if num, err := strconv.Atoi(str); err == nil {
			fmt.Println(num * num)
		} else if len(chanSignal) == 0 {
			fmt.Println("Некорректный ввод")
		}
	}
}
*/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			select {
			case <-c:
				fmt.Println("выхожу из программы")
				wg.Done()
				return
			default:
				fmt.Println("введите число")
				var str string
				fmt.Scan(&str)
				if num, err := strconv.Atoi(str); err == nil {
					fmt.Println("квадрат числа", num, "равен", num*num)
				} else if len(c) == 0 {
					fmt.Println("Некорректный ввод")
				}
			}
		}
	}()
	wg.Wait()
}
