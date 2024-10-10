package main

import "fmt"

func main() {

	/* START | TIPE DATA NUMERIK NON-DESIMAL
	 *
	 * Tipe data numerik non-desimal atau non floating point di Go ada beberapa jenis
	 * In general, ada 2 tipe data numerik non-desimal yg perlu diketahui
	 * uint : tipe data untuk bilangan positif atau bilangan cacah
	 * int 	: tipe data untuk bilangan negatif dan positif atau bilangan bulat
	 *
	 * %d : untuk memformat data numerik non-desimal
	 * \n : enter sejumlah 1 baris
	 *
	 */

	var positiveNumberNonDecimal uint8 = 250 // Isi variabel ini dengan 260 atau lebih dari 255, maka akan error
	var negativeNumberNonDecimal = -1243423644

	fmt.Printf("cetak nilai bilangan positif : %d\n", positiveNumberNonDecimal)
	fmt.Printf("cetak nilai bilangan negatif/minus : %d\n", negativeNumberNonDecimal)

	/* END | TIPE DATA NUMERIK NON-DESIMAL */

	// ============================================================================================================================= //

	/* START | TIPE DATA NUMERIK DESIMAL
	 *
	 * Tipe data numerik desimal atau floating point di Go ada beberapa jenis
	 * Biasanya hanya butuh 2 tipe data numerik desimal yg perlu diketahui
	 * float32 	: tipe data numerik desimal yg besaran cakupan nilainya :
	 			- Range 	: Mendekati 3,4 dengan 38 nol di belakangnya.
				- Presisi 	: Mampu menyimpan sekitar 6 hingga 9 digit angka secara akurat.
	 *
	 * float64 	: tipe data numerik desimal yg besaran cakupan nilainya
	 			- Range 	: Mendekati 1,8 dengan 308 nol di belakangnya.
				- Presisi 	: Mampu menyimpan sekitar 15 hingga 17 digit angka secara akurat.
	 *
	 * %f 	: untuk cetak dengan default 6 digit desimal, dihitung setelah point atau titik.
	 * %.3f : untuk cetak 3 digit desimal, dihitung setelah point atau titik.
	 * %.7f : untuk cetak 7 digit desimal, dihitung setelah point atau titik, dst.
	 * \n : enter sejumlah 1 baris
	 *
	*/

	var decimalNumber = 2.62

	fmt.Printf("cetak bilangan desimal default 6 digit: %f\n", decimalNumber)
	fmt.Printf("cetak bilangan desimal 7 digit: %.7f\n", decimalNumber)

	/* END | TIPE DATA NUMERIK DESIMAL */

	// ============================================================================================================================= //

	/* START | TIPE DATA STRING
	 *
	 * Tipe data string ini bisa diisi dengan segala huruf, angka dan beberapa karakter
	 * Bedanya dengan numerik dan boolean tipe data string :
	 * - Ketika dicetak maupun dideklarasikan nilainya menggunakan tanda quote ( " ) di awal dan akhir nilai
	 * - Bisa juga menggunakan tanda grave accent/backticks ( ` ) yang terletak di sebelah kiri tombol 1 pada keyboard
	 *
	 * %s : untuk cetak variabel yang bernilai string
	 * \n : enter sejumlah 1 baris
	 *
	 */

	var tandaQuoteString string = "Hai, apa kabarnya didewasa ini? Semangat!"

	fmt.Printf("Cetak penggunaan tipe data string dengan tanda quote: %s\n", tandaQuoteString)

	var tandaBackticksString = `Hai, salam kenal!
	Nama saya "Xaipul Gopur".
	Mari kita belajar "Golang" bersama!
	1000% GRATIS!!!`

	fmt.Printf("Cetak penggunaan tipe data string dengan tanda grave accent/backticks: %s\n", tandaBackticksString)

	/* END | TIPE DATA STRING */

	// ============================================================================================================================= //

	/* START | TIPE DATA BOOLEAN
	 *
	 * Tipe data boolean atau bool ini cuma bisa diisi oleh 2 nilai
	 * true 	: tipe data boolean yg menyatakan benar
	 * false 	: tipe data boolean yg menyatakan salah
	 *
	 * %t : untuk cetak variabel yang bernilai boolean(true atau false)
	 * \n : enter sejumlah 1 baris
	 *
	 */

	var aiExist bool = true

	fmt.Printf("Apakah AI makin eksis?: %t\n", aiExist)

	/* END | TIPE DATA BOOLEAN */

	// ============================================================================================================================= //

}
