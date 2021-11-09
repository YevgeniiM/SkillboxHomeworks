package main

import (
	"fmt"
	"io"
	"os"
)

/*
Что нужно сделать
Напишите программу, которая считывает ввод с stdin, создаёт структуру student и записывает указатель на структуру
в хранилище map[studentName] *Student.
Программа должна получать строки в бесконечном цикле, создать структуру Student через функцию newStudent,
далее сохранить указатель на эту структуру в map, а после получения EOF (ctrl + d) вывести на экран имена всех студентов из хранилища.
Также необходимо реализовать методы put, get.
*/

type Student struct {
	name  string
	age   int
	grade int
}

func (p *Student) put(s string, a int, g int) {
	p.name = s
	p.age = a
	p.grade = g
}
func (p *Student) get() {
	fmt.Printf("Имя студента %8s; его возраст %3d; курс %d \n", p.name, p.age, p.grade)
}
func newStudent(m map[string]*Student) error {
	var name string
	var age, grade int

	for {

		fmt.Print("Введите имя, возраст и курс студента:\n")
		_, err := fmt.Fscan(os.Stdin, &name, &age, &grade)
		if err != nil {
			return err
		}
		m[name] = &Student{"", 0, 0}
		m[name].put(name, age, grade)

	}
	return nil
}
func main() {
	m := make(map[string]*Student)
	err := newStudent(m)
	if err == io.EOF {
		fmt.Println("Студенты из хранилища:")
		for v, _ := range m {
			m[v].get()
		}
	} else if err != nil {
		fmt.Println(err)
	}

}
