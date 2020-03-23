package main

import "fmt"

func main() {
	fmt.Println("hello tugas 1")
	multiply, divide, add, minus := m4value(12, 34, 2345, 109)
	fmt.Println("Jawaban No.1 : ", multiply, divide, add, minus)
	nilai1, nilai2 := mValue()
	fmt.Println("Jawaban No.2 : ", nilai1, nilai2)
	nilaiArgs := []int{21, 31, 41}
	fmt.Println("Jawaban No.3 : ", getMultiply(2, nilaiArgs...))
}

func m4value(a int, b int, c int, d int) (int, int, int, int) {
	multiply := a * a
	divide := multiply / 2
	add := b + a
	minus := c - d
	return multiply, divide, add, minus
}

func mValue() (int, int) {
	nilai1 := (16*54 + 23/12) - 130
	nilai2 := 16*54 + 23/(12-130)

	return nilai1, nilai2
}

func getMultiply(number int, args ...int) []int {
	var multiply []int
	for _, args := range args {
		multiply = append(multiply, number*args)
	}

	return multiply
}
