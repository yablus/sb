package main

import (
	"fmt"
)

const n = 10

func main() {
	arr := [n]int{1, 2, 3, 4, 6, 6, 4, 8, 9, 10}
	slEven, slOdd := oddEven(arr)
	fmt.Println("Исходный массив:", arr)
	fmt.Printf("Четные: %v\nНечетные: %v\n", slEven, slOdd)
}

func oddEven(arr [n]int) (even []int, odd []int) {
	for _, v := range arr {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	return
}

/*
Задание 1. Чётные и нечётные
Что нужно сделать
Напишите функцию, которая принимает массив чисел, а возвращает два массива: один из чётных чисел, второй из нечётных.
*/
