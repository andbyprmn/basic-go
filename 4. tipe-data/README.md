# 4. Tipe Data

Seperti selayaknya bahasa pemrograman compiler seperti Java, dll, Go juga mengenal beberapa jenis tipe data, yaitu :

- Tipe Data Numerik 
  - Non-Desimal (`uint`, `int`) 
  - Desimal (`float32`, `float64`) 
- Tipe Data String (`string`)
- Tipe Data Boolean (`bool`)
  
## 4.A. Tipe Data Numerik Non-Desimal
Untuk tipe data non-desimal atau biasa disebut *non floating point* seperti `uint` dan `int`, dalam deklarasi atau inisialisasinya nanti akan dibagi lagi menjadi beberapa jenis. Pembagian tersebut berdasarkan dari besaran atau lebar dari cakupan nilainya, atau pada saat kita ingin deklarasi variable dengan tipe data numerik non-desimal kita bisa tentukan maksimal si `uint` ini mau diisi dari *`0`* sampai *`255`* misalnya, maka pada saat deklarasi variable kita bisa gunakan `uint8` yang dimana value atau nilai dari `uint8` tidak boleh melebihi *`255`*.

Jika kita tuliskan di *`main.go`*, maka akan jadi seperti berikut :
```go
var positiveNumberNonDecimal uint8 = 250
```
Maksud dari kode pemrograman diatas adalah saya coba mencontohkan untuk mengisi nilai variabel *positiveNumberNonDecimal* dengan *250*. Nah sekarang coba ada ganti nilai tersebut menjadi *260*, maka yg terjadi pada saat kita ketikan perintah *`go run main.go`*, outputnya adalah error. 

## *4.A.1. Mengapa demikian?*
Karena `uint8` cakupan lebarnya hanya sebanyak `255`, jika lebih dari itu maka akan error. Dan juga pada level ***middle to advance*** kita tidak bisa sembarangan mendeklarasikan setiap tipe data khususnya tipe data numerik baik desimal atau non-desimal sesuai dengan yang kita inginkan karena itu erat kaitannya dengan alokasi memori variabel dimana hal tersebut akan mempengaruhi terjadinya pemakaian memori yang tidak maksimal. Dengan kita paham secara mendetail akan kebutuhan kita, maka akan membuat memori bekerja lebih optimal.


## *4.A.2. Lalu, bagaimana solusinya jika kita ingin mendeklarasikan nilai pada tipe data numerik non-desimal yang cakupannya sesuai dengan yang kita butuhkan?*

Berikut adalah table yang bisa kita jadikan patokan untuk mendeklarasikan variable dengan tipe data numerik non-desimal.


| Tipe Data     | Cakupan Bilangan |
| -----------   | ----------- |
| `uint8`       | 0 ↔ 255 |
| `uint16`      | 0 ↔ 65535 |
| `uint32`      | 0 ↔ 4294967295 |
| `uint64`      | 0 ↔ 18446744073709551615 |
| `uint`        | sama dengan `uint32` atau `uint64` (tergantung nilai) |
| `byte`        | sama dengan uint8 |
| `int8`        | -128 ↔ 127 |
| `int16`       | -32768 ↔ 32767 |
| `int32`       | -2147483648 ↔ 2147483647 |
| `int64`       | -9223372036854775808 ↔ 9223372036854775807 |
| `int`         | sama dengan `int32` atau `int64` (tergantung nilai) |
| `rune`        | sama dengan `int32` |
  
## 4.B. Tipe Data Numerik Desimal
Pada tipe data numerik desimal, kita hanya perlu tahu 2 saja, yaitu `float32` dan `float64`. Untuk perbedaan tipe data tersebut ada pada besaran cakupan bilangannya (*sudah dijelaskan pada file `main.go`*)

## 4.C. Tipe Data String
Tipe data ini sebenarnya sudah pernah kita gunakan di project `2. hello-world`. Dan penjelasan dari cara atau deklarasi penulisan nilai pada tipe data `string` ini sudah dijelaskan didalam file `main.go`.

## 4.C. Tipe Data Boolean
Tipe data ini di Go dideklarasikan menggunakan `bool`. Dan penjelasan mulai dari cara atau deklarasi penulisan nilai sampai cetak nilai pada variabel `bool` pada tipe data ini sudah dijelaskan didalam file `main.go`.



> ⚠ *Mari kita lihat file `main.go` pada folder *`4. tipe-data`* agar lebih memahami apa yang saya maksud.*