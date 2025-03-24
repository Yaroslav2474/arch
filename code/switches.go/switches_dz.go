package main

import "fmt"

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func main() {
	var students []string
	var name string
	var marks []string
	var choise int
	var mark string
	var index int

	for {
		fmt.Print("\n1.Добавить студента и оценку \n2. Вывести студентов и их оценки \n3. Удалить студента и оценку \n0. Exit\n\nВведите цифру: ")
		fmt.Scan(&choise)
		switch choise {
		case 0:
			break
		case 1:
			fmt.Print("Введите имя студента: ")
			fmt.Scan(&name)
			students = append(students, name)
			fmt.Print("Введите оценку студента: ")
			fmt.Scan(&mark)
			marks = append(marks, mark)
		case 2:
			for i := 0; i < len(students); i++ {
				fmt.Print(i+1, students[i], " имеет оценку ", marks[i])
			}
		case 3:
			fmt.Print("Введите номер студента: ")
			fmt.Scan(&index)
			students = remove(students, index)
			marks = remove(marks, index)
		}
	}
}
