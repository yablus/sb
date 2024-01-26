package main

import (
	"flag"
	"fmt"
)

func main() {
	var str string
	var substr string
	flag.StringVar(&str, "str", "Строка для поиска", "set string")
	flag.StringVar(&substr, "substr", "подстрока", "set substring")
	flag.Parse()
	fmt.Printf("\"%s\", \"%s\"\n", str, substr)
	fmt.Println(find(str, substr))
}

func find(str string, substr string) bool {
	bStr := []byte(str)
	bSubstr := []byte(substr)
	for i := 0; i < len(bStr); i++ {
		if bStr[i] == bSubstr[0] {
			if len(bSubstr) == 1 {
				return true
			}
			for j := 1; j < len(bSubstr); j++ {
				if bStr[i+j] != bSubstr[j] {
					break
				} else if j == len(bSubstr)-1 {
					return true
				}
			}
		}
	}
	return false
}

/*
25.8 Практическая работа
Цель задания
Написать программу для нахождения подстроки в кириллической подстроке. Программа должна запускаться с помощью команды:

 go run main.go --str "строка для поиска" --substr "поиска"

Для реализации такой работы с флагами воспользуйтесь пакетом flags, а для поиска подстроки в строке вам понадобятся руны.



Что нужно сделать
Спроектировать алгоритм поиска подстроки.
Определить строку и подстроку, используя флаги.
Написать алгоритм реализацию для работы со строками UTF-8 (для этого необходимо воспользоваться рунами).


Что оценивается
Алгоритм может работать с различными символами (кириллица, китайские иероглифы).
Использованы руны.


Как отправить задание на проверку
Выполните задание в файле Go. Загрузите файл на Google Диск, откройте доступ для всех по ссылке. Отправьте ссылку на файл через форму для сдачи домашнего задания.

Или отправьте файл через онлайн-редактор REPL, или архивом.
*/
