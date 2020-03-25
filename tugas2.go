package main

import "fmt"

func main() {

	// Soal Pertama
	print_func := anonymous_func()
	fmt.Println(print_func("hi there anonymous"))

	// Soal Kedua
	nilaiArgs := []int{45, 23, 56, 89, 900, 26, 78}
	getGanGenap := func(args ...int) ([]int, []int) {
		var ganjil []int
		var genap []int

		for _, args := range args {
			if args%2 == 0 {
				genap = append(genap, args)
			} else {
				ganjil = append(ganjil, args)
			}
		}

		return ganjil, genap
	}

	var ganjil, genap = getGanGenap(nilaiArgs...)
	fmt.Println("Nilai Ganjil : ", ganjil, "\nNilai Genap : ", genap)

}

func anonymous_func() func(string) string {
	return func(msg string) string {
		return msg
	}
}
