package main

import "fmt"

func main() {
	//-- 9.A.1.1. Deklarasi Sederhana
	var numbers [5]int

	//- 9.A.1.2. Deklarasi Dengan Inisialisasi Nilai
	numbers = [5]int{1, 2, 3, 4, 5}
	// atau
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5

	//- 9.A.1.3. Deklarasi Tanpa Mendefinisikan Ukuran
	autoSizingElementNumbers := [...]int{1, 2, 3, 4, 5}

	fmt.Println(numbers[0])    // Mengakses elemen pertama
	numbers[1] = 25            // Mengubah nilai elemen kedua
	fmt.Println("\n", numbers) // Cetak seluruh isi masing-masing elemen di dalam array

	//- 9.A.3. Ukuran dan Kapasitas Array
	fmt.Printf("\nLength of array: %d\n\n", len(autoSizingElementNumbers)) // Cetak autoSizingElementNumbers

	//- 9.A.4.1. Iterasi Array Menggunakan for Loop
	cars := [5]string{"Toyota", "Honda", "Daihatsu", "Suzuki", "Hyundai"}
	for i := 0; i < len(cars); i++ {
		fmt.Printf("Nilai dari Elemen/Indeks ke-%d : %s\n", i, cars[i])
	}

	fmt.Println("")

	//- 9.A.4.2. Iterasi Array Menggunakan range Loop
	for idx, val := range cars {
		fmt.Printf("Nilai dari Elemen/Indeks ke-%d : %s\n", idx, val)
	}

	//- 9.A.5.1. Deklarasi Array Multidimensi
	var matriks [2][3]int // Array 2D dengan 2 baris dan 3 kolom

	//- 9.A.5.2. Inisialisasi Array Multidimensi

	//Inisialisasi Langsung
	matriks = [2][3]int{
		{1, 2, 3}, // Row atau Baris 1
		{4, 5, 6}, // Row atau Baris 2
	}

	// Inisialisasi Elemen Individu
	matriks[0][0] = 1
	matriks[0][1] = 2
	matriks[0][2] = 3
	matriks[1][0] = 4
	matriks[1][1] = 5
	matriks[1][2] = 6

	//- 9.A.5.3. Mengakses Elemen Array Multidimensi
	fmt.Println("\n", matriks[0][1], "\n") // Cetak row pertama pada kolom kedua

	//- 9.A.5.4. Perulangan Untuk Array Multidimensi
	//- Menggunakan for Loop
	for i := 0; i < len(matriks); i++ { // looping baris
		for j := 0; j < len(matriks[i]); j++ { // looping kolom
			fmt.Printf("Element [%d][%d] = %d\n", i, j, matriks[i][j])
		}
	}

	fmt.Println("")

	//- Menggunakan range Loop
	for i, row := range matriks {
		for j, val := range row {
			fmt.Printf("Element [%d][%d] = %d\n", i, j, val)
		}
	}

	//- 9.A.5.5. Array 3 Dimensi (3D) dan Lebih
	var array3D [2][3][4]int

	//- Inisialisasi Array 3D
	array3D = [2][3][4]int{
		{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		},
		{
			{13, 14, 15, 16},
			{17, 18, 19, 20},
			{21, 22, 23, 24},
		},
	}

	//- Akses Elemen Array 3D
	fmt.Println("\n", array3D[1][2][3]) // Cetak/akses array ke-2, pada baris ke-3, dikolom ke-4
}
