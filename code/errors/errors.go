package main

import (
	"errors"
	"fmt"
)

func triangle_exists(a, b, c int) (string, error) {

	if (((a + b) > c) && c > a && c > b) || (((a + c) > b) && b > a && b > c) || (((c + b) > a) && a > c && a > b) {
		return "Существует", nil
	} else {
		return "нету", errors.New("треугольника не существует")
	}

}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	s, err := triangle_exists(a, b, c)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
