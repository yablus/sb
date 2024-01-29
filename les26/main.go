package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var args []string
	args = os.Args[1:]
	//args = []string{""}
	//args = []string{"first.txt"}
	//args = []string{"first.txt", "second.txt"}
	//args = []string{"first.txt", "second.txt", "result.txt"}
	//args = []string{"first.txt", "second.txt", "result.txt", "another.txt"}
	//fmt.Println(args)
	//fmt.Println("-----------")

	var content []string
	var err error
	var result string

	switch {
	case len(args) > 3:
		fmt.Println("Введено более 3х файлов")
		break
	case len(args) == 0:
		fmt.Println("Не указано имя файла")
		break
	default:
		content, err = parseContent(content, args[0])
		if err != nil {
			// обработать ошибку или выход
			// в данном примере нет смысла, т.к. программа небольшая
		}
		if len(args) > 1 {
			content, err = parseContent(content, args[1])
			if err != nil {
				// обработать ошибку или выход
				// в данном примере нет смысла, т.к. программа небольшая
			}
		}
		result = strings.Join(content, "\n")
		fmt.Println(result)
		fmt.Println("-----------")
		if len(args) > 2 {
			err = saveContent(args[2], result)
			if err != nil {
				// обработать ошибку или выход
				// в данном примере нет смысла, т.к. программа небольшая
			}
		}
	}
}

func parseContent(content []string, fileName string) ([]string, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println("Ошибка открытия файла", fileName, "для чтения", err)
		//os.Exit(1)
		return content, err
	}
	defer file.Close()
	b, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла", err)
		//os.Exit(1)
		return content, err
	}
	return append(content, string(b)), nil
}

func saveContent(fileName string, content string) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Ошибка открытия файла", fileName, "для записи", err)
		//os.Exit(1)
		return err
	}
	defer file.Close()
	file.WriteString(content)
	fmt.Println("Содержимое файлов записано в", fileName)
	return nil
}

/*
26.5 Практическая работа
Цель задания
Написать программу аналог cat.



Что нужно сделать
Программа должна получать на вход имена двух файлов, необходимо  конкатенировать их содержимое, используя strings.Join.



Что оценивается
При получении одного файла на входе программа должна печатать его содержимое на экран.
При получении двух файлов на входе программа соединяет их и печатает содержимое обоих файлов на экран.
Если программа запущена командой go run firstFile.txt secondFile.txt resultFile.txt, то она должна написать два соединённых файла в результирующий.


Общие условия
Разработка выполняется в среде golang или vs code.

first.txt

контент первого файла

second.txt

контент второго файла

result .txt

контент первого файла

контент второго файла

Input

go run first.txt second.txt result.txt



first.txt

контент первого файла

second.txt

контент второго файла

Input

go run first.txt second.txt


Output

контент первого файла

контент второго файла



Как отправить задание на проверку
Выполните задание в файле Go. Загрузите файл на Google Диск, откройте доступ для всех по ссылке. Отправьте ссылку на файл через форму для сдачи домашнего задания.

Или отправьте файл через онлайн-редактор REPL, или архивом.
*/
