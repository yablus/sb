package main

import (
	"fmt"
	"strings"
)

func main() {
	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := []rune{'H', 'E', 'L', 'П', 'М', 'S', 'd', 'o', 'w'}
	result := make([][]int, len(sentences)*len(chars))
	result = parseTest(sentences, chars)
	printResult(sentences, chars, result)
	printArray(result)
}

func parseTest(sentences []string, chars []rune) [][]int {
	slice := make([][]int, len(sentences))
	lastWord := make([]string, len(sentences))
	for i := 0; i < len(sentences); i++ {
		tempSliceSentences := strings.Split(sentences[i], " ")
		lastWord[i] = tempSliceSentences[len(tempSliceSentences)-1]
		slice[i] = make([]int, len(chars)) // на этом моменте тупил жестко!
		/*
		   Ума не прилагал, что в Go двухмерные массивы не существуют. Это по сути эквивалент одномерного массива (среза, слайса) в котором в качестве аргументов тоже массивы.
		   Данный факт нашелся в документации.
		*/
		for j := 0; j < len(chars); j++ {
			slice[i][j] = -1
			for index, v := range lastWord[i] {
				if v == chars[j] {
					slice[i][j] = index
				}
			}
		}
	}
	return slice
}

func printResult(sentences []string, chars []rune, slice [][]int) {
	fmt.Println("Поиск символов в нескольких строках")
	fmt.Println("=================")
	for i := 0; i < len(slice); i++ {
		fmt.Println(sentences[i])
		for j, v := range slice[i] {
			if v != -1 {
				fmt.Printf("[%d][%d] = %d: '%c' индекс вхождения в последнее слово: %d\n", i, j, slice[i][j], chars[j], slice[i][j])
			}
		}
		fmt.Println("-----------------")
	}
}

/*
func printArray(slice [][]int) {
	for i, _ := range slice {
		fmt.Println(slice[i])
	}
}
*/

/*
Задание 2. Поиск символов в нескольких строках
Что нужно сделать
Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа rune, а возвращает 2D-массив, где на позиции [i][j] стоит индекс вхождения символа j из chars в последнее слово в предложении i (строку надо разбить на слова и взять последнее). То есть сигнатура следующая:

func parseTest(sentences []string, chars []rune)

Советы и рекомендации
В качестве среды разработки используйте Goland или VScode.
Не забудьте проверить, что вы получили больше чем 0 аргументов.
Подход не важен: можно переписать сортировку пузырьком или отсортировать, а потом перевернуть.
Пример входных данных
sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}

chars := [5]rune{'H','E','L','П','М'}

Пример вывода результата в первом элементе массива

'H' position 0

'E' position 1

'L' position 9

Что оценивается
Правильность размеров исходных и конечной матрицы.
Арифметическая правильность ответа.
*/
