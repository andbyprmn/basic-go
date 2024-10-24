# 5. Konstanta

Dalam pembuatan atau pengembangan aplikasi, biasanya kita perlu yang namanya variabel konstanta. *`Tujuannya apa?`* Tujuannya adalah supaya nilai dalam suatu variabel itu tidak berubah-ubah atau variabel tersebut tidak bisa diganti nilainya(tetap) dan kalau kita paksa ganti nilai di satu variabel konstanta itu maka yang terjadi adalah error.

Misalkan, kita sedang membuat aplikasi matematika dimana salah satu fiturnya bisa menghitung keliling lingkaran. Seperti yang kita ingat bahwa dalam menghitung keliling lingkaran ada rumus yang familiar bernama *`π(pi)`* yang salah satu nilai konstanta/tetapnya adalah ![\small \pi=\frac{22}{7}](https://latex.codecogs.com/svg.latex?\small&space;\pi=\frac{22}{7}).

> *Jadi, konstanta adalah suatu jenis variabel yang nilainya tidak dapat diubah-ubah lagi atau kita ganti setelah kita deklarasikan atau pada saat kita inisialisasi di awal.*

## 5.A. Penggunaan Konstanta
Variabel konstanta di Go dideklarasikan menggunakan keyword *`const`* sedangkan variabel biasa atau variabel yang bisa diubah-ubah nilainya itu selalu dideklarasikan dengan menggunakan keyword *`var`*.

Kalau kita menggunakan perumpamaan diatas tentang aplikasi matematika dengan fitur menghitung keliling lingkaran maka kita bisa mendeklarasikan salah satu variabel konstantanya yaitu *`π(pi)`* seperti dibawah ini.
```go
const Pi float32 = 22.0/7.0
```
Kita juga dapat menggunakan metode ***`type inference`*** pada variabel konstanta yang kita deklarasikan dimana dengan menggunakan deklarasi variable dengan metode ini kita tidak perlu pusing memikirkan tipe data yang spesifik karena Go lah yang akan secara otomatis tipe data dari nilai yang kita deklarasikan.
```go
const Pi = 22.0/7.0
```

### 5.A.1. Deklarasi Multi Konstanta
Seperti yang sudah kita bahasa di folder *`3. variables`* tentang deklarasi multiple variabel, variabel konstanta pun mendukung hal yang sama. Konstanta bisa dideklarasikan dengan banyak deklarasi variabel baik yg bertipe data sama maupun berbeda secara bersamaan (langsung).

> ⚠ *Untuk lebih jelas mari kita buka file `main.go` pada folder *`5. konstanta`* agar kita lebih paham.*






