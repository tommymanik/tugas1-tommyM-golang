package main

import "fmt"

func main() {
	// var getMin = func(n []int) int {
	// 	var min int
	// 	for i, e := range n {
	// 		if i == 0 {
	// 			min = e
	// 		} else if e < min {
	// 			min = e
	// 		}
	// 	}
	// 	return min
	// }
	// var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	// var min = getMin(numbers)
	// fmt.Printf("data : %v\nmin", numbers, min)

	func() {
		fmt.Println("Hi")
	}()

	func(a int, b int) {
		fmt.Println(a * b)
	}(2, 4)

	print_func := anonymous_func()
	fmt.Println(print_func("hi there anonymous"))

	a := a()
	a(3, 4)

	b := b()
	mul, min := b(6, 4)
	fmt.Println(mul, min)
}

func anonymous_func() func(string) string {
	return func(msg string) string {
		return msg
	}
}

func a() func(int, int) {
	return func(a int, b int) {
		fmt.Println(a * b)
	}
}

func b() func(int, int) (int, int) {
	return func(a int, b int) (int, int) {
		multiply := a * b
		minus := a - b
		return multiply, minus
	}
}
