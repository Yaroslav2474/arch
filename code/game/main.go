package main

import (
	"fmt"

	"math/rand/v2"
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

// type people struct {
// 	name              string
// 	surname           string
// 	phone_number      int
// 	post              string
// 	salary            int
// 	number_worked_day int
// }

// type student struct {
// 	name     string
// 	age      int
// 	facultet string
// 	ocenki   []int
// }

// func (std student) print_struct(sred_ocen float64) {
// 	fmt.Printf("\tИмя студента: %s\n", std.name)
// 	fmt.Printf("\tВозраст студента: %d\n", std.age)
// 	fmt.Printf("\tФакультет студента: %s\n", std.facultet)
// 	fmt.Printf("\tОценки студента: ")
// 	for i := 0; i < len(std.ocenki); i++ {
// 		fmt.Print(std.ocenki[i], " ")
// 	}
// 	fmt.Printf("\n\tСредняя оценка студента: %.2f\n", sred_ocen)
// }

// func readInt(prompt string) int {
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		fmt.Print(prompt)
// 		input, _ := reader.ReadString('\n')
// 		input = strings.TrimSpace(input)
// 		value, err := strconv.Atoi(input)
// 		if err == nil {
// 			return value
// 		}
// 		fmt.Println("Ошибка: введите целое число")
// 	}
// }

type Hero struct {
	max_hp int
	name   string
	hp     int
	armor  int
	damage int
	weapon string
	crit   int
}

func (hr Hero) info() {
	fmt.Printf("Имя: %s\nХп: %d\nБроня: %d\nУрон: %d\nОружие: %s\n\n", hr.name, hr.hp, hr.armor, hr.damage, hr.weapon)
	// дополнить вывод макс армор и дамаг
}

func (hr *Hero) Attack(enemy *Hero) {
	fmt.Printf("%s замечает, что вдали сидит %s\n", hr.name, enemy.name)
	enemy.info()
	fmt.Print("Напасть? (1-да/0-нет): ")
	dmg := 0
	c := 1

	def_armor := hr.armor

	for {

		fmt.Scan(&c)
		if c == 1 {
			if rand.IntN(100)+1 <= hr.crit {
				dmg = hr.damage * 3
				fmt.Println("КРИТИЧЕСКИЙ УРОН")
			} else {
				dmg = hr.damage
			}

			enemy.armor -= dmg
			if enemy.armor <= 0 {
				enemy.hp += enemy.armor
				enemy.armor = 0
			}
			if enemy.hp <= 0 {
				enemy.hp = 0
				fmt.Print("Победа\n")
				hr.max_hp += rand.IntN(11) + 5
				hr.damage += rand.IntN(5) + 1
				hr.armor = def_armor
				break
			}
			fmt.Printf("\n%s нанес %s %d урона.\nТеперь у %s %d здоровья и %d брони.\n\n", hr.name, enemy.name, dmg, enemy.name, enemy.hp, enemy.armor)
			if rand.IntN(100)+1 <= enemy.crit {
				dmg = hr.damage * 3
				fmt.Println("КРИТИЧЕСКИЙ УРОН")
			} else {
				dmg = enemy.damage
			}
			hr.armor -= dmg
			if hr.armor <= 0 {
				hr.hp += hr.armor
				hr.armor = 0
			}
			if hr.hp <= 0 {
				fmt.Print("Проигрышь\n")
				break
			}
			fmt.Printf("%s нанес %s %d урона.\nТеперь у %s %d здоровья и %d брони.\n", enemy.name, hr.name, dmg, hr.name, hr.hp, hr.armor)

		} else if c == 0 {
			break
		} else {
			fmt.Print("\nНет такой комнады!!!\n")

		}
		fmt.Print("\nПродолжить? (1-да/0-нет): ")
	}

}

func main() {

	herro := &Hero{100, "Ярик", 100, 50, 20, "ак-47", 50}
	herro.info()
	enemy := &Hero{100, "хохол", 100, 30, 15, "палка", 20}

	herro.Attack(enemy)

	herro.info()
	enemy.info()

	// std1 := student{}
	// reader := bufio.NewReader(os.Stdin)

	// fmt.Print("Введите имя студента: ")
	// std1.name, _ = reader.ReadString('\n')
	// std1.name = strings.TrimSpace(std1.name)

	// std1.age = readInt("Введите возраст студента: ")

	// fmt.Print("Введите факультет студента: ")
	// std1.facultet, _ = reader.ReadString('\n')
	// std1.facultet = strings.TrimSpace(std1.facultet)

	// var summ int
	// fmt.Println("Вводите оценки (0 - стоп):")
	// for {
	// 	ocenka := readInt("Оценка: ")
	// 	if ocenka == 0 {
	// 		break
	// 	}
	// 	if ocenka < 1 || ocenka > 5 {
	// 		fmt.Println("Ошибка: оценка должна быть от 1 до 5")
	// 		continue
	// 	}
	// 	std1.ocenki = append(std1.ocenki, ocenka)
	// 	summ += ocenka
	// }

	// if len(std1.ocenki) == 0 {
	// 	fmt.Println("Ошибка: не введено ни одной оценки")
	// 	return
	// }

	// sred_ocen := float64(summ) / float64(len(std1.ocenki))
	// std1.print_struct(sred_ocen)
	// 	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	// 	people1 := people{}
	// 	fmt.Print("Введите имя первого человека: ")
	// 	fmt.Scan(&people1.name)
	// 	fmt.Print("Введите фамилию первого человека: ")
	// 	fmt.Scan(&people1.surname)
	// 	fmt.Print("Введите номер телефона первого человека: ")
	// 	fmt.Scan(&people1.phone_number)
	// 	fmt.Print("Введите должность первого человека: ")
	// 	fmt.Scan(&people1.post)
	// 	fmt.Print("Введите зарпаплату за день первого человека: ")
	// 	fmt.Scan(&people1.salary)
	// 	fmt.Print("Введите колтчество отработаных дней первого человека: ")
	// 	fmt.Scan(&people1.number_worked_day)
	// 	salary1 := people1.salary * people1.number_worked_day
	// 	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	// 	people2 := people{}
	// 	fmt.Print("Введите имя второго человека: ")
	// 	fmt.Scan(&people2.name)
	// 	fmt.Print("Введите фамилию второго человека: ")
	// 	fmt.Scan(&people2.surname)
	// 	fmt.Print("Введите номер телефона второго человека: ")
	// 	fmt.Scan(&people2.phone_number)
	// 	fmt.Print("Введите должность второго человека: ")
	// 	fmt.Scan(&people2.post)
	// 	fmt.Print("Введите зарпаплату за день второго человека: ")
	// 	fmt.Scan(&people2.salary)
	// 	fmt.Print("Введите колтчество отработаных дней второго человека: ")
	// 	fmt.Scan(&people2.number_worked_day)
	// 	salary2 := people2.salary * people2.number_worked_day
	// 	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")

	// 	if people1.salary > 192000 {
	// 		fmt.Printf("Зарплата %s %s позволяет жить в коттедже\n", people1.name, people1.surname)
	// 	} else if 192000 > people1.salary && people1.salary > 3000 {
	// 		fmt.Printf("Зарплата %s %s позволяет жить в провинциальных квартирах\n", people1.name, people1.surname)
	// 	} else {
	// 		fmt.Printf("Зарплата %s %s позволяет жить на теплотрассе\n", people1.name, people1.surname)
	// 	}
	// 	fmt.Println(" ")

	// 	if people2.salary > 192000 {
	// 		fmt.Printf("Зарплата %s %s позволяет жить в коттедже\n", people2.name, people2.surname)
	// 	} else if 192000 > people2.salary && people2.salary > 3000 {
	// 		fmt.Printf("Зарплата %s %s позволяет жить в провинциальных квартирах\n", people2.name, people2.surname)
	// 	} else {
	// 		fmt.Printf("Зарплата %s %s позволяет жить на теплотрассе\n", people2.name, people2.surname)
	// 	}
	// 	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")

	// 	fmt.Printf("%s %s\n з/п в сутки %d\n в месяц %d\n", people1.name, people1.surname, people1.salary, salary1)
	// 	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")
	// 	fmt.Printf("%s %s\n з/п в сутки %d\n в месяц %d\n", people2.name, people2.surname, people2.salary, salary2)
	// 	fmt.Println("-----------------------------------------------------------------------------------------------------------------------------------------------------------------------")

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

}
