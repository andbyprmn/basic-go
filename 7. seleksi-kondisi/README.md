# 7. Seleksi Kondisi
Seleksi kondisi adalah mekanisme dalam pemrograman yang memungkinkan program untuk mengambil keputusan berdasarkan kondisi tertentu. Dalam Golang, seleksi kondisi digunakan untuk menjalankan blok kode hanya jika suatu kondisi terpenuhi (*`benar`* atau *`salah`*). Ini adalah salah satu elemen fundamental dari alur kontrol program yang membantu mengarahkan jalannya eksekusi berdasarkan hasil evaluasi kondisi.

Semisal, bayangkan saja kita sedang berkendara dan mendekati lampu lalu lintas. Lampu ini memiliki tiga kondisi: merah, kuning, dan hijau. Kita membuat keputusan berdasarkan kondisi lampu:

- Jika lampu ***hijau***, Anda akan lanjut jalan.
- Jika lampu ***kuning***, Anda mungkin memperlambat atau berhenti.
- Jika lampu ***merah***, Anda harus berhenti.
  
Demikian pula, dalam program, seleksi kondisi membantu memutuskan bagian kode mana yang akan dieksekusi berdasarkan kondisi tertentu.


## 7.A. Macam-Macam Seleksi Kondisi
Ada tiga macam seleksi kondisi utama di Golang:

1. [*`if`* - *`else`* **Statement**](https://github.com/andbyprmn/basic-go/tree/main/7.%20seleksi-kondisi#7a1-if---else-statement)
2. [*`switch`* **Statement**](https://github.com/andbyprmn/basic-go/tree/main/7.%20seleksi-kondisi#7a2-switch-statement)
3. [*Nesting Conditional* **Statement** (Seleksi Kondisi Bersarang)](https://github.com/andbyprmn/basic-go/tree/main/7.%20seleksi-kondisi#7a3-nesting-conditional-statement-seleksi-kondisi-bersarang)
   
Mari kita bahas masing-masing secara terperinci.

### 7.A.1. *`if`* - *`else`* **Statement**

Ini adalah bentuk paling dasar dari seleksi kondisi. Jika kondisi dalam *`if`* terpenuhi (*`true`*), maka blok kode di dalamnya akan dieksekusi. Jika tidak terpenuhi (*`false`*), blok kode dalam *`else`* (jika ada) akan dieksekusi.

Berikut adalah cara penulisan kode Go untuk seleksi kondisi menggunakan *`if`* - *`else`* **Statement** :
```go
    if kondisi {
        // Blok kode ini akan dijalankan jika kondisi bernilai true
    } else {
        // Blok kode ini akan dijalankan jika kondisi bernilai false
    }

```
#### 7.A.1.1. *`if`* - *`else if`* - *`else`* **Statement**
Dalam Golang kita juga bisa menggunakan menambahkan keyword *else if* untuk memeriksa beberapa kondisi berurutan serta memungkinkan kita untuk memeriksa kondisi tambahan jika kondisi *`if`* pertama tidak terpenuhi.

Berikut adalah cara penulisan kode Go untuk seleksi kondisi menggunakan *`if`* - *`else if`* - *`else`* :
```go
    if kondisi1 {
        // Eksekusi blok ini jika kondisi1 benar
    } else if kondisi2 {
        // Eksekusi blok ini jika kondisi1 salah, tapi kondisi2 benar
    } else {
        // Eksekusi blok ini jika kondisi1 dan kondisi2 keduanya salah
    }
```
#### 7.A.1.2. Variabel Temporary Pada *`if`* - *`else`*
Pada Golang, kita dapat mendeklarasikan variabel sementara di dalam pernyataan *`if`* yang hanya tersedia di dalam blok *`if`* atau *`else`*. 

Berikut adalah cara untuk mendeklarasikan variabel temporary pada *`if`* - *`else`* :

```go
    if temp := ekspresi; kondisi {
        // Gunakan variabel temp di sini
    } else {
        // Gunakan variabel temp di sini juga
    }
```

Dengan adanya variable temporary pada seleksi kondisi *`if`* - *`else`* ini, kita dapat beberapa keuntungan sebagai berikut :
- **Penghematan memori**: Variabel sementara yang dideklarasikan hanya di dalam blok *`if`* atau *`else`* akan otomatis dibuang setelah keluar dari blok tersebut, mengurangi beban memori.
- **Meningkatkan keamanan**: Karena variabel tersebut hanya tersedia dalam lingkup lokal blok *`if`*-*`else`*, ini mencegah penggunaan variabel tersebut di luar blok yang tidak diperlukan, mengurangi kemungkinan kesalahan.
- **Penggunaan variabel yang efisien**: Variabel sementara ini berguna ketika Anda hanya perlu menggunakan nilai sementara untuk kondisi dan tidak ingin variabel tersebut terus ada di luar blok kondisi. 


### 7.A.2. *`switch`* **Statement**

*`switch`* adalah alternatif dari *`if-else`* yang digunakan ketika ada beberapa kondisi berbeda yang harus dicek. Ini adalah cara yang lebih bersih dan efisien untuk menangani banyak kondisi daripada menggunakan banyak *`if-else`*.

Berikut adalah cara penulisan kode Go untuk seleksi kondisi menggunakan *`switch`* **Statement** :
```go 
    switch ekspresi {
    case nilai1:
        // Blok kode ini dijalankan jika ekspresi cocok dengan nilai1
    case nilai2:
        // Blok kode ini dijalankan jika ekspresi cocok dengan nilai2
    default:
        // Blok kode ini dijalankan jika tidak ada case yang cocok
    }
```
#### 7.A.2.1. Pemanfaatan *`case`* Untuk Banyak Kondisi
Dalam *`switch`* di Go, kita bisa memeriksa beberapa nilai dalam satu *`case`*. Ini berguna ketika kita ingin beberapa kondisi memiliki hasil yang sama.

Berikut adalah cara deklarasi menggunakan keyword *`case`* untuk banyak kondisi :

```go
    switch ekspresi {
    case nilai1, nilai2, nilai3:
        // Eksekusi jika ekspresi cocok dengan salah satu dari nilai1, nilai2, atau nilai3
    default:
        // Eksekusi jika tidak ada case yang cocok
    }
```

#### 7.A.2.2. Kurung Kurawal(`{}`) Pada Keyword *`case`* & *`default`*
Dalam Golang, kita bisa menggunakan kurung kurawal *`{}`* untuk membuat blok kode di dalam setiap *`case`* atau *`default`*. Ini berguna ketika kita ingin mendeklarasikan variabel atau melakukan lebih dari satu perintah.

```go
    switch ekspresi {
    case nilai1: {
        // Blok kode ini dijalankan jika ekspresi cocok dengan nilai1
    }

    case nilai2: {
        // Blok kode ini dijalankan jika ekspresi cocok dengan nilai2
    }

    default:
        // Blok kode ini dijalankan jika tidak ada case yang cocok
    }
```
#### 7.A.2.3. Penggunaan *Style* *`if`*-*`else`* Pada *`switch`* **Statement**
*`switch`* di Golang tidak selalu membutuhkan ekspresi di sebelahnya, sehingga bisa digunakan seperti *`if`*-*`else`*. Hal ini memungkinkan kita untuk mengekspresikan logika kondisi yang lebih kompleks.

Berikut adalah cara mendeklarasikan style *`if`*-*`else`* pada *`switch`* **Statement** :

```go
    switch {
    case kondisi1:
        // Blok kode untuk kondisi1
    case kondisi2:
        // Blok kode untuk kondisi2
    default:
        // Blok default
    }
```

#### 7.A.2.4. Keyword *`fallthrough`* Dalam *`switch`* **Statement**

Dalam *`switch`* **Statement** khususnya pada Golang, kita bisa menggunakan keyword *`fallthrough`* untuk melanjutkan eksekusi ke *`case`* berikutnya meskipun kondisinya tidak terpenuhi. Ini seperti memaksa program untuk terus mengevaluasi case di bawahnya.

Berikut adalah cara penggunaan keyword *`fallthrough`* dalam *`switch`* **Statement** :

```go
    switch ekspresi {
    case nilai1:
        // Blok kode ini dijalankan jika ekspresi cocok dengan nilai1
    case nilai2:
        // Blok kode ini dijalankan jika ekspresi cocok dengan nilai2
        fallthrough 
    case nilai3:
        // Blok kode ini dijalankan jika ekspresi cocok dengan nilai3
    default:
        // Blok kode ini dijalankan jika tidak ada case yang cocok
    }
```

### 7.A.3. *Nesting Conditional* **Statement** (Seleksi Kondisi Bersarang)

Seleksi kondisi bersarang adalah ketika kita meletakkan satu pernyataan kondisi di dalam pernyataan kondisi lainnya. Ini sering digunakan saat ada beberapa kondisi kompleks yang perlu dievaluasi.
Berikut adalah contoh program dengan seleksi kondisi bersarang :

```go
package main
import "fmt"

func main() {
    age := 25
    score := 75
    country := "USA"

    if age >= 18 {
        fmt.Println("You are an adult.")

        // Bersarang dengan switch-case
        switch country {
        case "USA":
            fmt.Println("You are in the USA.")
        case "UK":
            fmt.Println("You are in the UK.")
        case "Canada":
            fmt.Println("You are in Canada.")
        default:
            fmt.Println("Country not recognized.")
        }

    } else {
        fmt.Println("You are a minor.")

        // Bersarang dengan if-else
        if score >= 80 {
            fmt.Println("You passed the exam with a good score.")
        } else if score >= 60 {
            fmt.Println("You passed the exam with an average score.")
        } else {
            fmt.Println("You failed the exam.")
        }
    }
}
```

### Rangkuman

- *`if`*-*`else`*: Digunakan untuk memeriksa satu atau lebih kondisi, dengan opsi *`else if`* untuk kondisi berurutan dan penggunaan variabel sementara.
- *`switch`*: Digunakan untuk memeriksa banyak kasus, dengan dukungan untuk memeriksa banyak nilai dalam satu *`case`*, membuat *`switch`* tanpa ekspresi seperti *`if`*-*`else`*, dan menggunakan *`fallthrough`* untuk melanjutkan ke case berikutnya.
- Seleksi Kondisi Bersarang: Digunakan untuk mengevaluasi kondisi dalam kondisi, ketika diperlukan keputusan berdasarkan lebih dari satu faktor.

