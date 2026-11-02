package main

import (
	"fmt"

	"treasury/soal1"
	"treasury/soal2"
	"treasury/soal3"
	"treasury/soal4"
	"treasury/soal5"
)

func main() {
	fmt.Println("Soal 1:")
	fmt.Println(soal1.GO([]int{-1, 1}))
	fmt.Println(soal1.GO([]int{-1, -7, -5}))
	fmt.Println(soal1.GO([]int{1, 2, 1, 6}))

	fmt.Println("\nSoal 2 (hasil, error):")
	n, err := soal2.GO([]int{0, 1, 0}, 1000, 1010)
	fmt.Println(n, err)
	n, err = soal2.GO([]int{0, 1, 0}, 1259, 2601)
	fmt.Println(n, err)

	fmt.Println("\nSoal 3:")
	fmt.Println(soal3.GO([]int{1, 2}, []int{1, 3}))
	fmt.Println(soal3.GO([]int{1, 2, 2}, []int{1, 2, 4}))

	fmt.Println("\nSoal 4:")
	counts := soal4.GO(1223334)
	for digit := 0; digit <= 9; digit++ {
		if count := counts[digit]; count > 0 {
			fmt.Printf("%d -> %d\n", digit, count)
		}
	}

	fmt.Println("\nSoal 5:")
	fmt.Println(soal5.GO(1, 2, 3, 4))
	fmt.Println(soal5.GO(1, 2, 3, 4, 5))
}
