# 8. Perulangan
Perulangan (looping) adalah salah satu fitur penting dalam pemrograman yang memungkinkan eksekusi berulang dari sekumpulan instruksi. Di Golang, perulangan dikendalikan menggunakan struktur *`for`*. Golang hanya memiliki satu jenis perulangan, yaitu *`for loop`*, yang dapat diadaptasi untuk berbagai kebutuhan perulangan.

## 8.A. *`for`* Loop Dasar
**Struktur :**
```go
    for init; kondisi; post {
        // blok kode yang diulang
    }
```
- *`init`*: Inisialisasi variabel, biasanya variabel counter (misalnya *`i := 0`*).
- *`kondisi`*: Kondisi yang harus dipenuhi untuk tetap menjalankan loop (misalnya *`i < 10`*).
- *`post`*: Ekspresi yang dieksekusi di akhir setiap iterasi (misalnya *`i++`* untuk menambah nilai counter).

**Contoh :**
```go
    package main

    import "fmt"

    func main() {
        for i := 0; i < 5; i++ {
            fmt.Println(i)  // Output: 0 1 2 3 4
        }
    }
```
**Penjelasan :**
- Loop akan dimulai dari *`i = 0`*.
- Kondisi *`i < 5`* memastikan loop terus berjalan selama *`i`* kurang dari 5.
- Setiap iterasi, *`i`* bertambah 1 (*`i++`*).

## 8.B. *`for`* Tanpa Bagian Inisialisasi dan Post
Dalam Golang, kita bisa menghilangkan bagian init dan post jika tidak diperlukan. Hanya bagian kondisi yang wajib.

**Struktur :**
```go
    for kondisi {
        // blok kode yang diulang
    }
```

**Contoh :**
```go
    package main

    import "fmt"

    func main() {
        i := 0
        for i < 5 {
            fmt.Println(i)
            i++
        }
    }
```

**Penjelasan :**
- Variabel *`i`* diinisialisasi di luar loop.
- Loop akan berjalan selama *`i < 5`*.
- Di dalam loop, *`i`* terus bertambah hingga mencapai 5.


## 8.C. Infinite *`for`* Loop
Jika kita menghilangkan semua bagian (*`init`*, *`kondisi`*, dan *`post`*), loop akan terus berjalan selamanya, kecuali ada mekanisme untuk keluar dari loop.

**Struktur :**
```go
    for {
        // blok kode yang diulang tanpa batas
    }
```

**Contoh :**
```go
    package main

    import "fmt"

    func main() {
        i := 0
        for {
            fmt.Println(i)
            i++
            if i >= 5 {
                break  // Keluar dari loop jika i >= 5
            }
        }
    }
```

**Penjelasan :**
- Tanpa syarat *`kondisi`*, loop berjalan tanpa henti.
- Keyword *`break`* digunakan untuk menghentikan perulangan.


## 8.D. *`for`* Sebagai Loop *`while`*
Golang tidak memiliki keyword khusus untuk *`while`*, tetapi kita bisa meniru perilaku *`while`* dengan menggunakan *`for`* loop yang hanya mengecek kondisi.

**Struktur :**
```go
    for kondisi {
        // blok kode yang diulang selama kondisi terpenuhi
    }
```

**Contoh :**
```go
    package main

    import "fmt"

    func main() {
        i := 0
        for i < 5 {
            fmt.Println(i)
            i++
        }
    }
```

**Penjelasan :**
- Loop akan terus berjalan selama *`i < 5`*, seperti halnya *`while`* pada bahasa lain.


## 8.E. Perulangan Menggunakan Keyword *`range`*
Keyword *`range`* dalam Golang digunakan untuk iterasi melalui berbagai tipe data seperti *`array`*, *`slice`*, *`map`*, atau *`channel`*. Dengan *`range`*, kita bisa mengambil indeks dan nilai dari elemen yang diiterasi.

**Struktur :**
```go
    for indeks, nilai := range koleksi {
        // blok kode untuk memproses indeks dan nilai
    }
```

**Contoh : 1. Iterasi *`array`* atau *`slice`***
```go
    package main

    import "fmt"

    func main() {
        arr := []int{1, 2, 3, 4, 5}
        for idx, val := range arr {
            fmt.Printf("Index: %d, Value: %d\n", idx, val)
        }
    }
```

**Penjelasan :**
- *`range`* pada *`array`*/*`slice`* mengembalikan indeks dan nilai elemen pada setiap iterasi.

**Contoh : 2. Iterasi *`map`***
```go
    package main
        
    import "fmt"

    func main() {
        m := map[string]int{"apple": 2, "banana": 5, "orange": 3}
        for key, value := range m {
            fmt.Printf("Key: %s, Value: %d\n", key, value)
        }
    }
```

**Penjelasan :**
- Pada *`map`*, *`range`* mengembalikan ***key*** dan ***value***.

**Contoh : 3. Iterasi *`string`***
```go
    package main

    import "fmt"

    func main() {
        str := "hello"
        for idx, char := range str {
            fmt.Printf("Index: %d, Char: %c\n", idx, char)
        }
    }
```

**Penjelasan :**
- *`range`* pada nilai dari variabel *`string`* mengiterasi setiap karakter, mengembalikan indeks dan karakter dalam bentuk rune atau karakter Unicode yang diwakili dalam tipe *`int32`*.

## 8.F. Penggunaan Keyword *`break`*, *`continue`* dan *`goto`* dalam Perulangan

### 8.F.1. Keyword *`break`*
*`break`* digunakan untuk menghentikan loop sebelum kondisi selesai.

**Contoh :**
```go
    package main

    import "fmt"

    func main() {
        for i := 0; i < 10; i++ {
            if i == 5 {
                break  // Keluar dari loop saat i = 5
            }
            fmt.Println(i)
        }
    }
```

### 8.F.2. Keyword *`continue`*
*`continue`* digunakan untuk melewati iterasi yang sedang berjalan dan langsung melanjutkan ke iterasi berikutnya.

**Contoh :**
```go
    package main

    import "fmt"

    func main() {
        for i := 0; i < 10; i++ {
            if i == 5 {
                continue  // Lewati iterasi saat i = 5
            }
            fmt.Println(i)
        }
    }

```

**Penjelasan :**
- Pada iterasi ke-5, perintah *`continue`* akan dilewati sehingga tidak mencetak angka 5.

### 8.F.3. Keyword *`goto`*
*`goto`* digunakan untuk melompat ke label tertentu dalam kode. Ini jarang digunakan, tapi bisa berguna untuk keluar dari loop yang dalam atau kondisi khusus.

**Contoh :**
```go
   package main

    import "fmt"

    func main() {
        for i := 0; i < 5; i++ {
            if i == 3 {
                goto endLoop  // Melompat ke label endLoop
            }
            fmt.Println(i)
        }

    endLoop:
        fmt.Println("endloop label ended.")
    }
```


## 8.G. Nested Loop (Perulangan Bersarang)
Golang juga mendukung perulangan bersarang, dimana *`for`* loop diletakkan di dalam *`for`* loop lainnya.

**Contoh :**
```go
    package main

    import "fmt"

    func main() {
        for i := 1; i <= 3; i++ {
            for j := 1; j <= 3; j++ {
                fmt.Printf("matriks [%d][%d]\n", i, j)
            }
        }
    }
```

**Penjelasan :**
- *`for`* loop luar (*`i`*) berjalan sebanyak 3 kali, dan untuk setiap iterasi *`i`*, *`for`* loop dalam (*`j`*) berjalan sebanyak 3 kali. Ini menciptakan 9 output berbeda.


## 8.H. *`for`* Loop dengan Label
Golang juga mendukung loop dengan label atau gampangnya kita bisa meng-inisialisasi `for` loop dengan memberinya sebuah label, sehingga memungkinkan kita pada kondisi tertentu menggunakan keyword *`break`* atau *`continue`* untuk melompat ke loop tertentu dalam nested loop.

**Contoh :**
```go
    package main
    import "fmt"

    func main() {
    outerLoop:
        for i := 1; i <= 3; i++ {
            for j := 1; j <= 3; j++ {
                if i*j > 4 {
                    break outerLoop  // Keluar dari outerLoop jika kondisi terpenuhi
                }
                fmt.Printf("matriks [%d][%d]\n", i, j)
            }
        }
        fmt.Println("outerloop label ended.")
    }
```

**Penjelasan :**

- *`outerLoop`* adalah nama sebuah label yang kita inisialisasi untuk membungkus *`for`* loop yang memungkinkan kita untuk langsung keluar dari *`for`* loop luar saat kondisi terpenuhi di *`for`* loop dalam.

### Rangkuman
- Golang hanya menggunakan *`for`* sebagai satu-satunya cara untuk membuat perulangan.
- Ada berbagai variasi *`for`* loop: *`for`* dengan semua parameter, tanpa inisialisasi(*init*) atau *post*, atau hanya dengan *kondisi*.
- *`range`* digunakan untuk mengiterasi *`array`*, *`slice`*, *`map`* dan *`string`*.
- **Keyword** *`break`*, *`continue`* dan *`goto`* membantu mengendalikan alur dalam perulangan.
- **Nested loop** dan **loop dengan label** memungkinkan pembuatan logika yang lebih kompleks.

Perulangan dalam Golang fleksibel dan cukup sederhana untuk diterapkan dalam berbagai skenario program.