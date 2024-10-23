package main

import "fmt"

func main() {
	// 6.A.1. Operator Aritmatika
	a := 10
	b := 3

	fmt.Println(a + b) // Penjumlahan: 13
	fmt.Println(a - b) // Pengurangan: 7
	fmt.Println(a * b) // Perkalian: 30
	fmt.Println(a / b) // Pembagian: 3
	fmt.Println(a % b) // Modulus: 1

	// 6.A.2. Operator Perbandingan
	c := 5
	d := 8

	fmt.Println(c == d) // Sama dengan: false
	fmt.Println(c != d) // Tidak sama dengan: true
	fmt.Println(c > d)  // Lebih besar: false
	fmt.Println(c < d)  // Lebih kecil: true
	fmt.Println(c >= d) // Lebih besar atau sama: false
	fmt.Println(c <= d) // Lebih kecil atau sama: true

	// 6.A.3. Operator Logika
	x := true
	y := false

	fmt.Println(x && y) // AND logika: false
	fmt.Println(x || y) // OR logika: true
	fmt.Println(!x)     // NOT logika: false

	// 6.A.4. Operator Penugasan
	e := 10

	e += 5         // Sama dengan: e = e + 5
	fmt.Println(e) // Hasil: 15

	e -= 3         // Sama dengan: e = e - 3
	fmt.Println(e) // Hasil: 12

	e *= 2         // Sama dengan: e = e * 2
	fmt.Println(e) // Hasil: 24

	e /= 4         // Sama dengan: e = e / 4
	fmt.Println(e) // Hasil: 6

	e %= 5         // Sama dengan: e = e % 5
	fmt.Println(e) // Hasil: 1

	// 6.A.5. Operator Bitwise
	f := 10 // 1010 dalam biner
	g := 3  // 0011 dalam biner

	fmt.Println(f & g)  // AND bitwise: 2 (0010)
	fmt.Println(f | g)  // OR bitwise: 11 (1011)
	fmt.Println(f ^ g)  // XOR bitwise: 9 (1001)
	fmt.Println(f &^ g) // AND NOT bitwise: 8 (1000)

	fmt.Println(f << 1) // Left shift: 20 (10100)
	fmt.Println(f >> 1) // Right shift: 5 (0101)

	// 6.A.6. Operator Pengalamatan & Indireksi
	o := 42
	p := &o // Pengambil alamat

	fmt.Println(p)  // Alamat memori variabel o
	fmt.Println(*p) // Nilai dari variabel yang ditunjuk oleh p (42)

	*p = 21        // Mengubah nilai melalui pointer
	fmt.Println(o) // Hasil: 21

	// 6.A.7. Operator Lainnya
	h := 5

	h++            // Increment: h = h + 1
	fmt.Println(h) // Hasil: 6

	h--            // Decrement: h = h - 1
	fmt.Println(h) // Hasil: 5

	i := 10        // Deklarasi dan inisialisasi cepat
	fmt.Println(i) // Hasil: 10
}
