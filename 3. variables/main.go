package main

import "fmt"

func main() {
	/* Schema
	* var <nama-variabel> <tipe-data>
	* var <nama-variabel> <tipe-data> = <nilai>
	*
	* Ketika deklarasi, tipe data yg digunakan harus dituliskan juga. Istilah dari metode deklarasi variabel ini adalah MANIFEST TYPING.
	 */

	// MANIFEST TYPING
	var firstName string = "John"

	var lastname string
	lastname = "Wick"

	fmt.Printf("Hai, %s %s!\n", firstName, lastname)

	/*Schema
	*  var <nama-variabel> = <nilai>
	*  <nama-variabel> := <nilai>
	*
	* Selain manifest typing, Go juga mengadopsi konsep type inference,
	* yaitu metode deklarasi variabel yang tipe data-nya diketahui secara otomatis dari data/nilai variabel
	 */

	// TYPE INFERENCE
	var firstNameTI = "John"
	lastnameTI := "Wick"

	fmt.Printf("Hai, %s %s!\n", firstNameTI, lastnameTI)

	//Deklarasi Multi Variable
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fmt.Printf("Hai, %s %s %s!\n", first, second, third)

	var fourth, fifth, sixth string = "empat", "lima", "enam"
	fmt.Printf("Hai, %s %s %s!\n", fourth, fifth, sixth)

	seventh, eight, ninth := "tujuh", "delapan", "sembilan"
	fmt.Printf("Hai, %s %s %s!\n", seventh, eight, ninth)

	/* Dengan multiple variable dimana jika ada case semisal setiap variable mempunyai tipe data yg berbeda dengan metode delarasi type inference
	 * kita bisa menyederhanakannya menjadi dibawah ini :
	 */
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Printf("Hai, %s %s %s %s!\n", one, isFriday, twoPointTwo, say)

	/* Variable Underscore (_)
	 *
	 * Go memiliki aturan unik yang jarang dimiliki bahasa lain, yaitu tidak boleh ada satupun variabel yang tidak dipakai.
	 * Jika ada variable yg tidak dipakai, maka Go akan memberi tahu kita dengan menyebut si variable itu sebagai error
	 * Maka dari itu Go memiliki fitur _ (Underscore), atau biasa disebut variable pre-defined
	 * _ (Underscore) adalah reserved variable yg bisa dimanfaatkan untuk menampung nilai yg tidak dipakai.
	 */

	_ = "Learn basic Go"
	_ = "Golang is easy going"
	name, _ := "John", "Wick"

	fmt.Printf("Hai, %s!\n", name)

	/*Deklarasi Variablle menggunakan keyword new
	 *
	 * Fungsi new() digunakan untuk membuat variabel pointer dengan tipe data tertentu.
	 * Nilai data default-nya akan menyesuaikan tipe datanya.
	 * Variabel name menampung data bertipe pointer string dimana yg ditampilkan bukanlah nilainya melainkan alamat memori dari nilai variable tsb dalam bentuk hexadecimal
	 * Untuk menampilkan nilai aslinya, variabel tersebut perlu di-dereference terlebih dahulu, caranya dengan menuliskan tanda asterisk (*) sebelum nama variabel.
	 */

	nameNewStr := new(string)

	fmt.Println(nameNewStr)  // 0xc00008e150
	fmt.Println(*nameNewStr) // ""

	/* Deklarasi Variable Menggunakan Keyword make
	 *
	 * Fungsi make() ini hanya bisa digunakan untuk pembuatan beberapa jenis variable saja, yaitu channel, slice dan map
	 * dibahas di poin channel, slice dan map
	 */
}
