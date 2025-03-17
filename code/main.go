package main

import (
	"fmt"
	"math"
)

// func remove(s []string, i int) []string {
// 	s[i] = s[len(s)-1]
// 	return s[:len(s)-1]
// }

// type rect struct {
// 	height int
// 	width  int
// }

// func (r rect) perim() int {
// 	return (r.height + r.width) * 2
// }

// func (r rect) area() int {
// 	return r.height * r.width
// }

type shape struct {
	radius int
}

type circle interface {
	perim() float64
	area() float64
}

func (s shape) perim() float64 {
	return (2 * math.Pi * float64(s.radius))
}

func (s shape) area() float64 {
	return (math.Pi * math.Pow(float64(s.radius), 2))
}

func printZ(circle circle) {
	fmt.Printf("Площадь: %.1f\nПериметр: %.1f\n", circle.area(), circle.perim())
}

func main() {
	cir := shape{
		radius: 10,
	}
	printZ(cir)

	// kvadrat := rect{
	// 	height: 20,
	// 	width: 20,
	// }

	// kvadrat := rect{}
	// kvadrat.height = 20
	// kvadrat.width = 20

	// fmt.Printf("Площадь: %d\nПериметр: %d\n", kvadrat.area(), kvadrat.perim())

	// var students []string
	// var name string
	// var marks []string
	// var choise int
	// var mark string
	// var index int

	// for {
	// 	fmt.Print("\n1.Добавить студента и оценку \n2. Вывести студентов и их оценки \n3. Удалить студента и оценку \n0. Exit\n\nВведите цифру: ")
	// 	fmt.Scan(&choise)
	// 	switch choise {
	// 	case 0:
	// 		break
	// 	case 1:
	// 		fmt.Print("Введите имя студента: ")
	// 		fmt.Scan(&name)
	// 		students = append(students, name)
	// 		fmt.Print("Введите оценку студента: ")
	// 		fmt.Scan(&mark)
	// 		marks = append(marks, mark)
	// 	case 2:
	// 		for i := 0; i < len(students); i++ {
	// 			fmt.Print(i+1, students[i], " имеет оценку ", marks[i])
	// 		}
	// 	case 3:
	// 		fmt.Print("Введите номер студента: ")
	// 		fmt.Scan(&index)
	// 		students = remove(students, index)
	// 		marks = remove(marks, index)
	// 	}

	// }

}
