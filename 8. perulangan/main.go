package main

import "fmt"

func main() {
	i := 0

	// 8.A. for Loop Dasar
	for i := 0; i < 5; i++ {
		fmt.Println(i) // Output: 0 1 2 3 4
	}

	fmt.Println("")

	// 8.B. for Tanpa Bagian Inisialisasi dan Post
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("")

	// 8.C. Infinite for Loop
	for {
		fmt.Println(i)
		i++
		if i >= 5 {
			break // Keluar dari loop jika i >= 5
		}
	}

	fmt.Println("")

	// 8.D. for Sebagai Loop while
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("")

	// 8.E. Perulangan Menggunakan Keyword range
	// 1. Iterasi array atau slice
	arr := []int{1, 2, 3, 4, 5}
	for idx, val := range arr {
		fmt.Printf("Index: %d, Value: %d\n", idx, val)
	}

	fmt.Println("")

	// 2. Iterasi map
	m := map[string]int{"apple": 2, "banana": 5, "orange": 3}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	fmt.Println("")

	// 3. Iterasi string
	str := "hello"
	for idx, char := range str {
		fmt.Printf("Index: %d, Char: %c\n", idx, char)
	}

	fmt.Println("")

	// 8.F. Penggunaan Keyword break, continue dan goto dalam Perulangan
	// 8.F.1. Keyword break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // Keluar dari loop saat i = 5
		}
		fmt.Println(i)
	}

	fmt.Println("")

	// 8.F.2. Keyword continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue // Lewati iterasi saat i = 5
		}
		fmt.Println(i)
	}

	fmt.Println("")

	// 8.F.3. Keyword goto
	for i := 0; i < 5; i++ {
		if i == 3 {
			goto endLoop // Melompat ke label endLoop
		}
		fmt.Println(i)
	}

endLoop:
	fmt.Println("endloop label ended.")

	fmt.Println("")

	// 8.G. Nested Loop (Perulangan Bersarang)
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("matriks [%d][%d]\n", i, j)
		}
	}

	fmt.Println("")

	// 8.H. for Loop dengan Label
outerLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 4 {
				break outerLoop // Keluar dari outerLoop jika kondisi terpenuhi
			}
			fmt.Printf("matriks [%d][%d]\n", i, j)
		}
	}
	fmt.Println("outerloop label ended.")
}
