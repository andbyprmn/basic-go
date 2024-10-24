# 6. Operator

Dalam bahasa pemrograman Golang (Go), operator digunakan untuk melakukan operasi terhadap nilai atau variabel. Operator memungkinkan manipulasi data dan pengambilan keputusan dalam program. Operator di Go dapat dibagi menjadi beberapa kategori tergantung pada fungsinya. 

## 6.A. Jenis-Jenis Operator di Go
Secara keseluruhan, operator di Go terbagi menjadi beberapa jenis utama:

1. [Operator Aritmatika](https://github.com/andbyprmn/basic-go/tree/main/6.%20operator#6a1-operator-aritmatika)
2. [Operator Perbandingan](https://github.com/andbyprmn/basic-go/tree/main/6.%20operator#6a2-operator-perbandingan)
3. [Operator Logika](https://github.com/andbyprmn/basic-go/tree/main/6.%20operator#6a3-operator-logika)
4. [Operator Penugasan](https://github.com/andbyprmn/basic-go/tree/main/6.%20operator#6a4-operator-penugasan)
5. [Operator Bitwise](https://github.com/andbyprmn/basic-go/tree/main/6.%20operator#6a5-operator-bitwise)
6. [Operator Pengalamatan & Indireksi](https://github.com/andbyprmn/basic-go/tree/main/6.%20operator#6a6-operator-pengalamatan--indireksi)
7. [Operator Lainnya](https://github.com/andbyprmn/basic-go/tree/main/6.%20operator#6a7-operator-lainnya)

Dengan menggunakan operator-operator ini, kita dapat melakukan berbagai macam operasi dan manipulasi data dalam kode Go.

Berikut adalah penjelasan lengkap mengenai jenis-jenis operator dalam Golang:

## 6.A.1. Operator Aritmatika

Operator aritmatika digunakan untuk melakukan operasi matematika seperti penjumlahan, pengurangan, perkalian, dan sebagainya. 

| Operator      | Description           | Example     |
| -----------   | -----------           | ----------- |
| `+`	        | Penjumlahan           | `a + b`     |
| `-`	        | Pengurangan	        | `a - b`     |
| `*`	        | Perkalian	            | `a * b`     |
| `/`	        | Pembagian	            | `a / b`     |
| `%`	        | Modulus (sisa bagi)   | `a % b`     |


## 6.A.2. Operator Perbandingan
Operator perbandingan digunakan untuk membandingkan dua nilai. Hasil dari operasi ini adalah boolean (*`true`* atau *`false`*).

| Operator      | Description                   | Example       |
| -----------   | -----------                   | -----------   |
| `==`	        | Sama dengan	                | `a == b`      | 
| `!=`	        | Tidak sama dengan	            | `a != b`      |
| `>`	        | Lebih besar dari	            | `a > b`       |
| `<`	        | Lebih kecil dari	            | `a < b`       |
| `>=`	        | Lebih besar atau sama dengan	| `a >= b`      |
| `<=`	        | Lebih kecil atau sama dengan	| `a <= b`      |


## 6.A.3. Operator Logika
Operator logika digunakan untuk menggabungkan dua atau lebih ekspresi logika. Hasilnya juga boolean.

| Operator      | Description                                                                               | Example           |
| -----------   | ------------------------------------------------------------------------------------------| ------------------|
| `&&`	        | AND logika: Menghasilkan *`true`* jika kedua kondisi bernilai *`true`*.	                | `a && b`          |
| `\| \|`	    | OR logika: menghasilkan *`true`* jika salah satu dari kondisi bernilai *`true`*.          | `a \| \| b`       |
| `!`	        | NOT logika/negasi/kebalikan dari nilai:  jika *`true`* menjadi *`false`* dan sebaliknya.	| `!a`              |


## 6.A.4. Operator Penugasan
Operator penugasan digunakan untuk menetapkan nilai ke variabel.

| Operator      | Description                   | Example       |
| -----------   | -----------                   | -----------   |
| `=`	        | Menetapkan nilai	            | `a = 5`       |
| `+=`	        | Penjumlahan dan menetapkan	| `a += 5`      |
| `-=`	        | Pengurangan dan menetapkan	| `a -= 5`      |
| `*=`	        | Perkalian dan menetapkan	    | `a *= 5`      |
| `/=`	        | Pembagian dan menetapkan	    | `a /= 5`      |
| `%=`	        | Modulus dan menetapkan	    | `a %= 5`      |


## 6.A.5. Operator Bitwise
Operator bitwise bekerja pada tingkat bit dari nilai integer.

| Operator      | Description       | Example    |
| -----------   | -----------       | -----------|
| `&`	        | AND bitwise	    | `a & b`    |
| `\|`	        | OR bitwise        | `a \| b`   |
| `^`	        | XOR bitwise	    | `a ^ b`    |
| `&^`	        | AND NOT bitwise	| `a &^ b`   |
| `<<`	        | Left shift	    | `a << b`   |
| `>>`	        | Right shift	    | `a >> b`   |


## 6.A.6. Operator Pengalamatan & Indireksi
Operator ini didalam Go digunakan untuk bekerja dengan pointer (referensi ke lokasi memori variabel).

| Operator      | Description           | Example                       |
| -----------   | -----------           | -----------                   |
| `&`	        | Pengambil alamat	    | `&a` (alamat dari `a`)        |
| `*`	        | Dereferensi pointer	| `*p` (nilai dari pointer `p`) |


## 6.A.7. Operator Lainnya
Operator lain yang sering digunakan di Golang.

| Operator      | Description                       | Example    |
| -----------   | -----------                       | -----------|
| `++`	        | Increment (menambah 1)	        | `a++`      |
| `--`	        | Decrement (mengurangi 1)	        | `a--`      |
| `:=`	        | Deklarasi dan inisialisasi cepat	| `a := 5`   |







> ***Notes*** :
>> Kali ini saya dapat error menarik pada saat saya *running* aplikasi Go ini dan errornya seperti dibawah ini :
>> ```go 
>>    go reading .\main.go unexpected NUL in input  
>> ```
>> Ternyata setelah saya *crosscheck*, hal itu dikarenakan saya membuat file main.go itu dengan *powershell* di *Windows 10* menggunakan command `echo "package main" >> main.go`. Sehingga pada saat *create* file *main.go* hal itu menyebabkan kesalahan encoding yg telah di-generate oleh command tersebut. Selengkapnya sila cek [*Stackoverflow*](https://stackoverflow.com/a/48870543).
>
> Selengkapnya kita bisa bukan file *`main.go`* pada folder *`6. Operator`* agar lebih mudah dalam memahami.
