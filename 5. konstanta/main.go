package main

import "fmt"

func main() {
	/* ------------------------ START | DEKLARASI SINGLE KONSTANTA ------------------------ */

	//Deklarasi variabel konstanta dengan metode Manifest Typing
	const piManifestTyping float32 = 22.0 / 7.0

	fmt.Print("cetak : ", piManifestTyping, "\n")

	//Deklarasi variabel konstanta dengan metode Type Inference
	const piTypeInference = 22.0 / 7.0

	fmt.Print("cetak : ", piTypeInference, "\n")

	/* ------------------------ END | DEKLARASI SINGLE KONSTANTA ------------------------ */

	// ============================================================================================================================= //

	/* ------------------------ START | DEKLARASI MULTI KONSTANTA ------------------------ */

	// Style Penulisan Pertama
	const (
		fullname       = "Xaipul Gopur" // Metode Type Inference
		isMale    bool = true           // Metode Manifest Typing
		numberOne uint = 1              // Metode Manifest Typing
		Pi             = 22.0 / 7.0     // Metode Type Inference
	)

	fmt.Print("cetak : ", fullname, ", ", isMale, ", ", numberOne, ", ", Pi, "\n")

	// Style Penulisan Kedua
	const (
		varConstA = "Constanta"
		varConstB
		/* Ketentuan di Go adalah kalau salah satu variabel multi konstanta tidak dideklarasikan nilainya
		 * maka nilainya akan mengikuti nilai yang sudah dideklarasikan sebelumnya/diatasnya.
		 * Dalam hal ini nilai dari varConstB jika kita cetak outputnya adalah nilai yang sama dengan varConstA
		 */
	)

	fmt.Print("cetak : ", varConstA, ", ", varConstB, "\n")

	// Style Penulisan Ketiga | Acak
	const (
		today = "Jumat"
		hariIni
		isJumat = true
	)

	fmt.Print("cetak : ", today, ", ", hariIni, ", ", isJumat, "\n")

	// Style Penulisan Keempat | Satu Baris
	const first, second = 1, 2
	fmt.Print("cetak konstanta dengan tipe data integer : ", first, ", ", second, "\n")

	const third, fourth string = "tiga", "empat"
	fmt.Print("cetak konstanta dengan tipe data string : ", third, ", ", fourth, "\n")

	/* ------------------------ END | DEKLARASI MULTI KONSTANTA ------------------------ */

}
