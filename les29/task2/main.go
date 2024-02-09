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

package main

import (
	"fmt"
	"os"
	"os/signal"

	//"syscall"
	"time"
)

func worker(exitChan, doneChan chan bool) {
	fmt.Println("worker: запускаюсь")
	var num uint
	for {
		select {
		case <-exitChan:
			// Выполнение при получении сигнала Ctrl+C
			fmt.Println("worker: завершаю работу")
			doneChan <- true
			return
		default:
			// Выполнение обычной работы
			fmt.Println("worker: работаю")
			time.Sleep(1 * time.Second)
			fmt.Println("worker: квадрат числа", num, "равен", num*num)
			num++
		}
	}
}

func main() {
	fmt.Println("main: запускаюсь")
	exitChan := make(chan bool)
	doneChan := make(chan bool)

	fmt.Println("main: запускаю worker")
	go worker(exitChan, doneChan)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		os.Interrupt,
		//syscall.SIGHUP,
		//syscall.SIGINT,
		//syscall.SIGTERM,
		//syscall.SIGQUIT,
	)

	fmt.Println("main: ожидаю сигнал о завершении работы")
	<-sigChan // На этом моменте остановка и ожидание сигнала о завершении

	fmt.Println("main: получен сигнал о завершении работы, завершаю worker")
	close(exitChan) // Скажем worker'у чтоб закруглялся

	fmt.Println("main: ожидаю завершения работы worker")
	<-doneChan // Снова остановка и одидание завершения worker

	fmt.Println("main: завершаю работу")
	//time.Sleep(1 * time.Second) // Дополнительное время для завершения работы
}
