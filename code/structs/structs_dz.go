package main

import (
	"fmt"
)

type people struct {
	name              string
	surname           string
	salary            int
	number_worked_day int
}

func (pep people) salary_month() int {
	return pep.salary * pep.number_worked_day
}

func get_info() (name, surname string, salary, number_worked_day int) {
	fmt.Print("Введите имя, фамилию, суточную зарплату, кол-во отработанных дней")
	fmt.Scan(&name, &surname, &salary, &number_worked_day)
	return
}

func (pep people) condition_salary() {
	if pep.salary > 192000 {
		fmt.Printf("Зарплата %s %s позволяет жить в коттедже\n", pep.name, pep.surname)
	} else if 192000 > pep.salary && pep.salary > 3000 {
		fmt.Printf("Зарплата %s %s позволяет жить в провинциальных квартирах\n", pep.name, pep.surname)
	} else {
		fmt.Printf("Зарплата %s %s позволяет жить на теплотрассе\n", pep.name, pep.surname)
	}
	fmt.Printf("%s %s\n з/п в сутки %d\n в месяц %d\n", pep.name, pep.surname, pep.salary, pep.salary_month())
}

func main() {

	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	people1 := people{}
	people1.name, people1.surname, people1.salary, people1.number_worked_day = get_info()
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	people2 := people{}
	people2.name, people2.surname, people2.salary, people2.number_worked_day = get_info()
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	people1.condition_salary()
	people2.condition_salary()
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")

}
