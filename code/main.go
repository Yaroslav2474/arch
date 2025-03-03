package main

import "fmt"

func main() {
	var students []string
	var name string
	var marks []int
	var choise int
	var mark int

	for {
		fmt.Print("\n1.Добавить студента и оценку \n2. Вывести студентов и их оценки \n3. Удалить студента и оценку \n0. Exit\n\nВведите цифру: ")
		fmt.Scan(&choise)

		if choise == 0 {
			break
		}

		if choise == 1 {
			fmt.Print("Введите имя студента: ")
			fmt.Scan(&name)
			students = append(students, name)
			fmt.Print("Введите оценку студента: ")
			fmt.Scan(&mark)
			marks = append(marks, mark)
		}

		if choise == 2 {
			for i := 0; i < len(students); i++ {
				fmt.Print(students[i], " имеет оценку ", marks[i])
			}
		}

	}
}
