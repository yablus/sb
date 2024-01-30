package main

/*
27.8 Практическая работа
Цель задания
Научиться работать с композитными типами данных: структурами и картами

Что нужно сделать
Напишите программу, которая считывает ввод с stdin, создаёт структуру student и записывает указатель на
структуру в хранилище map[studentName] *Student.

type Student struct {
	name string
	age int
	grade int
}

Программа должна получать строки в бесконечном цикле, создать структуру Student через функцию newStudent,
далее сохранить указатель на эту структуру в map, а после получения EOF (ctrl + d) вывести на экран имена всех
студентов из хранилища. Также необходимо реализовать методы put, get.

Общие условия
Разработка выполняется в среде golang или vs code.

Input
-----
go run main.go

Строки
Вася 24 1
Семен 32 2
EOF

Output
-----
Студенты из хранилища:
Вася 24 1
Семен 32 2

Критерии оценки
Зачёт:
при получении одной строки (например, «имяСтудента 24 1») программа создаёт студента и сохраняет его, далее ожидает
следующую строку или сигнал EOF (Сtrl + Z);
при получении сигнала EOF программа должна вывести имена всех студентов из map.
На доработку:
задание не выполнено или выполнено не полностью.
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}

type Data interface {
	// put
	// добавляет студента в карту.
	// возвращает:
	// true, если добавлен;
	// false, если элемент уже присутствует или был совершен некорректный ввод.
	put(std *Student) bool

	// get
	// возвращает:
	// карту студентов.
	get() map[string]*Student
}

type DataTest struct{}

func (test *DataTest) put(std *Student) bool {
	return true
}
func (test *DataTest) get() map[string]*Student {
	return map[string]*Student{
		"Vasya": {"Vasya", 23, 3},
		"Petya": {"Petya", 22, 2},
	}
}

type App struct {
	repository Data
}

func (a *App) Run() {
	for {
		if std, ok := a.inputStudent(); ok {
			a.saveStudent(std)
		} else {
			a.printStudents()
			break
		}
	}
}

func (a *App) printStudents() {
	fmt.Println("Студенты из хранилища:")
	for _, v := range a.repository.get() {
		fmt.Printf("%s %d %d\n", v.name, v.age, v.grade)
	}
}

func (a *App) inputStudent() (*Student, bool) {
	for {
		fmt.Println("Введите через пробел имя студента, его возраст и класс")
		var strName, strAge, strGrade string
		empty := Student{
			name:  "",
			age:   0,
			grade: 0,
		}
		str, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err == io.EOF {
			return &empty, false
		}
		param := strings.Split(strings.TrimSpace(str), " ")
		fmt.Println(param)
		if len(param) == 3 || len(param) == 4 {
			strGrade = param[len(param)-1]
			strAge = param[len(param)-2]
			if len(param) == 4 {
				strName = param[0] + " " + param[1]
			} else {
				strName = param[0]
			}
		}
		intAge, errAge := strconv.Atoi(strAge)
		intGrade, errGrade := strconv.Atoi(strGrade)
		if intGrade != 0 && intAge != 0 && strName != "" && errAge == nil && errGrade == nil {
			std := Student{
				name:  strName,
				age:   intAge,
				grade: intGrade,
			}
			return &std, true
		}
		fmt.Println("Некорректный ввод")
	}
}

func (a *App) saveStudent(std *Student) {
	var msg string
	if ok := a.repository.put(std); ok {
		msg = "Студент %s успешно добавлен\n"
	} else {
		msg = "Студент %s уже есть в базе\n"
	}
	fmt.Printf(msg, std.name)
}

type Memory struct {
	students map[string]*Student
}

func newStudent() *Memory {
	return &Memory{
		students: make(map[string]*Student),
	}
}

func (m *Memory) put(std *Student) bool {
	if m.contains(std.name) {
		return false
	}
	m.students[std.name] = std
	return true
}

func (m *Memory) contains(stdName string) bool {
	for _, value := range m.students {
		if value.name == stdName {
			return true
		}
	}
	return false
}

func (m *Memory) get() map[string]*Student {
	return m.students
}

func main() {
	//repository := &DataTest{}
	repository := newStudent()
	app := &App{repository}
	app.Run()
}

/*
type Student struct {
	name  string
	age   int
	grade int
}

func main() {
	m := make(map[string]*Student)
	for {
		var student Student
		if _, err := fmt.Scan(&student.name, &student.age, &student.grade); err == io.EOF {
			for _, v := range m {
				fmt.Println(v)
			}
			break
		} else {
			m[student.name] = &student
			fmt.Println(student, "добавлен")
		}
	}
}
*/
