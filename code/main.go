package main

import (
	"fmt"
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

// type shape struct {
// 	radius int
// }

// type circle interface {
// 	perim() float64
// 	area() float64
// }

// func (s shape) perim() float64 {
// 	return (2 * math.Pi * float64(s.radius))
// }

// func (s shape) area() float64 {
// 	return (math.Pi * math.Pow(float64(s.radius), 2))
// }

// func printZ(circle circle) {
// 	fmt.Printf("Площадь: %.1f\nПериметр: %.1f\n", circle.area(), circle.perim())
// }

type people struct {
	name              string
	surname           string
	phone_number      int
	post              string
	salary            int
	number_worked_day int
}

func main() {

	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	people1 := people{}
	fmt.Print("Введите имя первого человека: ")
	fmt.Scan(&people1.name)
	fmt.Print("Введите фамилию первого человека: ")
	fmt.Scan(&people1.surname)
	fmt.Print("Введите номер телефона первого человека: ")
	fmt.Scan(&people1.phone_number)
	fmt.Print("Введите должность первого человека: ")
	fmt.Scan(&people1.post)
	fmt.Print("Введите зарпаплату за день первого человека: ")
	fmt.Scan(&people1.salary)
	fmt.Print("Введите колтчество отработаных дней первого человека: ")
	fmt.Scan(&people1.number_worked_day)
	salary1 := people1.salary * people1.number_worked_day
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	people2 := people{}
	fmt.Print("Введите имя второго человека: ")
	fmt.Scan(&people2.name)
	fmt.Print("Введите фамилию второго человека: ")
	fmt.Scan(&people2.surname)
	fmt.Print("Введите номер телефона второго человека: ")
	fmt.Scan(&people2.phone_number)
	fmt.Print("Введите должность второго человека: ")
	fmt.Scan(&people2.post)
	fmt.Print("Введите зарпаплату за день второго человека: ")
	fmt.Scan(&people2.salary)
	fmt.Print("Введите колтчество отработаных дней второго человека: ")
	fmt.Scan(&people2.number_worked_day)
	salary2 := people2.salary * people2.number_worked_day
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")

	if people1.salary > 192000 {
		fmt.Printf("Зарплата %s %s позволяет жить в коттедже\n", people1.name, people1.surname)
	} else if 192000 > people1.salary && people1.salary > 3000 {
		fmt.Printf("Зарплата %s %s позволяет жить в провинциальных квартирах\n", people1.name, people1.surname)
	} else {
		fmt.Printf("Зарплата %s %s позволяет жить на теплотрассе\n", people1.name, people1.surname)
	}
	fmt.Println(" ")

	if people2.salary > 192000 {
		fmt.Printf("Зарплата %s %s позволяет жить в коттедже\n", people2.name, people2.surname)
	} else if 192000 > people2.salary && people2.salary > 3000 {
		fmt.Printf("Зарплата %s %s позволяет жить в провинциальных квартирах\n", people2.name, people2.surname)
	} else {
		fmt.Printf("Зарплата %s %s позволяет жить на теплотрассе\n", people2.name, people2.surname)
	}
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")

	fmt.Printf("%s %s\n з/п в сутки %d\n в месяц %d\n", people1.name, people1.surname, people1.salary, salary1)
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("%s %s\n з/п в сутки %d\n в месяц %d\n", people2.name, people2.surname, people2.salary, salary2)
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")

	// cir := shape{
	// 	radius: 10,
	// }
	// printZ(cir)

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
