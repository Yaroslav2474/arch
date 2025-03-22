package main

import "fmt"

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
}
