# basic-go

[Golang](https://go.dev/doc/) (atau biasa disebut dengan Go) adalah bahasa pemrograman yang dikembangkan di Google oleh Robert Griesemer, Rob Pike dan Ken Thompson pada tahun 2007 dan mulai diperkenalkan ke publik tahun 2009. Bahasa Pemrograman Golang basic-nya dibuat dari bahasa C dan C++ maka dari itu banyak sintaks dan gaya penulisan kode yang familiar jika kita biasa menuliskan program dengan bahasa C atau C++.

> ## **Kelebihan Bahasa Go**
1. **Sederhana dan Mudah Dipelajari**
   - **Sintaksis Minimalis**: Go dirancang dengan sintaksis yang sederhana dan bersih, memudahkan pengembang baru untuk memahami dan menulis kode.
    - **Kurva Pembelajaran yang Landai**: Konsep-konsep dasar seperti variabel, kontrol alur, dan fungsi mudah dipelajari, memungkinkan pengembang untuk cepat produktif.

2. **Performa Tinggi**
    - **Kompilasi ke Bahasa Mesin**: Go dikompilasi langsung ke bahasa mesin, memberikan performa yang mendekati bahasa seperti C atau C++.
    - **Efisiensi Eksekusi**: Penggunaan memori yang efisien dan optimisasi kompilator membuat aplikasi Go cepat dan responsif.

3. **Dukungan Konkurensi yang Kuat**
    - **Goroutines**: Fitur ini memungkinkan pembuatan ribuan thread ringan secara efisien, memudahkan pengelolaan tugas-tugas konkuren.
    - **Channels**: Mekanisme komunikasi yang aman antar goroutines, memfasilitasi sinkronisasi dan pertukaran data tanpa risiko kondisi balapan.

4. **Library Standar yang Kuat**
    - **Komprehensif**: Go menyediakan library standar yang luas untuk berbagai kebutuhan, seperti pengelolaan file, jaringan, encoding/decoding, dan lainnya.
    - **Konsistensi**: Library standar yang konsisten memudahkan pengembang untuk menemukan dan menggunakan fungsi yang dibutuhkan tanpa harus bergantung pada library pihak ketiga.

5. **Kompilasi Cepat**
    - **Waktu Kompilasi Singkat**: Go dikenal dengan kecepatan kompilasinya yang luar biasa, memungkinkan iterasi pengembangan yang cepat dan produktif.
    - **Tooling Efisien**: Alat bawaan seperti `go build`, `go test`, dan `go fmt` berfungsi dengan cepat, mendukung alur kerja pengembangan yang efisien.

6. **Cross-Platform**
    - **Dukungan Multi-OS**: Go mendukung berbagai sistem operasi seperti Windows, macOS, dan Linux, serta arsitektur seperti AMD64, ARM, dan lainnya.
    - **Portabilitas**: Aplikasi Go dapat dengan mudah dikompilasi untuk berbagai platform tanpa modifikasi kode yang signifikan.

7. **Alat Bawaan yang Kuat**
    - **Formatter Kode (`gofmt`)**: Menjamin konsistensi gaya penulisan kode di seluruh proyek dan tim.
    - **Pengujian (`go test`)**: Memfasilitasi penulisan dan pelaksanaan unit test secara mudah.
    - **Manajemen Dependensi**: Sistem modul Go (go mod) memudahkan pengelolaan dependensi dan versi paket.

8. **Pengelolaan Memori Otomatis (Garbage Collection)**
    - **Manajemen Memori Otomatis**: Mengurangi beban pengembang dalam mengelola alokasi dan dealokasi memori, mencegah kebocoran memori.
    - **Kinerja Garbage Collector**: Garbage collector Go dirancang untuk minim mengganggu performa aplikasi, menjaga responsivitas.

9. **Tipe Statis dengan Type Inference**
    - **Keamanan Tipe**: Tipe statis membantu menangkap kesalahan pada waktu kompilasi, meningkatkan keamanan dan keandalan kode.
    - **Type Inference**: Memungkinkan penulisan kode yang lebih ringkas tanpa harus menyebutkan tipe secara eksplisit setiap saat.

10. **Skalabilitas**
    - **Desain untuk Sistem Terdistribusi**: Go sangat cocok untuk membangun layanan yang skalabel dan terdistribusi, seperti microservices dan aplikasi cloud-native.
    - **Manajemen Konkurensi yang Efisien**: Memungkinkan aplikasi menangani beban kerja yang besar dengan efisien.


> ## **Kekurangan Bahasa Go**
1. **Pengelolaan Error yang Verbose**
    - **Handling Error Manual**: Go tidak memiliki mekanisme exception handling seperti try-catch. Sebaliknya, pengembang harus secara eksplisit memeriksa dan menangani error setelah setiap operasi yang mungkin gagal.
    - **Kode yang Panjang**: Pendekatan ini dapat menyebabkan kode menjadi lebih panjang dan repetitif, terutama dalam aplikasi besar.

2. **Kurangnya Fitur Lanjutan**
    - **Generik Terbatas**: Meskipun Go telah mengimplementasikan generik sejak versi 1.18, implementasinya masih dianggap kurang fleksibel dibandingkan bahasa lain seperti Rust atau TypeScript.
    - **Kurangnya Macros dan Metaprogramming**: Go tidak mendukung macros atau metaprogramming tingkat tinggi, membatasi kemampuan untuk menulis kode yang sangat generik atau otomatis.

3. **Dukungan GUI yang Terbatas**
    - **Library GUI Terbatas**: Go tidak memiliki library GUI bawaan yang kuat, sehingga pengembangan aplikasi desktop dengan antarmuka pengguna grafis lebih menantang dibandingkan dengan bahasa lain seperti Java atau C#.
    - **Ketergantungan pada Pihak Ketiga**: Untuk pengembangan GUI, pengembang harus mengandalkan library pihak ketiga yang mungkin kurang matang atau memiliki dokumentasi terbatas.

4. **Ukuran Binary yang Cenderung Besar**
    - **Binary Standalone**: Aplikasi Go dikompilasi menjadi binary standalone yang termasuk semua dependensi, yang dapat menghasilkan ukuran file yang lebih besar dibandingkan dengan beberapa bahasa lain.
    - **Penggunaan Memori**: Walaupun performa tinggi, binary besar dapat memakan lebih banyak ruang penyimpanan dan memori saat dijalankan.

5. **Kurangnya Dukungan untuk Beberapa Paradigma Pemrograman**
    - **Pemrograman Fungsional Terbatas**: Go tidak mendukung beberapa fitur pemrograman fungsional tingkat tinggi seperti higher-order functions, pattern matching, atau currying secara eksplisit.
    - **OOP dengan Interface yang Terbatas**: Implementasi Object-Oriented Programming (OOP) di Go menggunakan interfaces yang implisit, yang mungkin membingungkan bagi pengembang yang terbiasa dengan OOP tradisional.

6. **Manajemen Modul yang Bisa Menjadi Rumit**
    - **Dependency Management**: Meskipun `go mod` telah meningkatkan pengelolaan dependensi, masih ada tantangan dalam mengelola versi dan kompatibilitas paket, terutama dalam proyek besar dengan banyak dependensi.
    - **Versi yang Tidak Konsisten**: Beberapa paket mungkin tidak mengikuti standar versi semantik, menyebabkan konflik versi dan kesulitan integrasi.

7. **Kurangnya Dukungan untuk Beberapa Area Spesifik**
    - **Keamanan Tingkat Rendah**: Untuk pengembangan sistem yang membutuhkan kontrol tingkat rendah atas memori atau hardware, Go mungkin tidak seefisien bahasa seperti C atau Rust.
    - **Kurangnya Fitur untuk Pengembangan Mobile dan Web Frontend**: Meskipun Go dapat digunakan untuk backend, dukungan untuk pengembangan aplikasi mobile atau frontend web masih terbatas.
  
8. **Implementasi Generik yang Masih Terbatas**
    - **Fleksibilitas Generik**: Meskipun generik telah diperkenalkan, implementasinya belum sepenuhnya matang dan mungkin tidak memenuhi semua kebutuhan pengembang dalam menulis kode yang sangat generik atau reusable.
    - **Kurangnya Dokumentasi dan Praktik Terbaik**: Penggunaan generik masih relatif baru, sehingga dokumentasi dan praktik terbaiknya belum sebanyak bahasa lain yang telah lebih lama mendukung generik.
  
9. **Kurangnya Penggunaan dalam Beberapa Industri**
    - **Adopsi Terbatas di Beberapa Sektor**: Meskipun populer di industri teknologi dan cloud, Go belum banyak diadopsi di sektor lain seperti pengembangan game atau aplikasi desktop yang kompleks.
    - **Komunitas yang Spesifik**: Komunitas Go mungkin tidak sebesar komunitas bahasa lain seperti Python atau JavaScript, sehingga sumber daya dan dukungan mungkin lebih terbatas.


> ## Instalasi dan Persiapan

Sebenarnya seluruh proses lengkap dari mulai instalasi dan persiapan bahkan sampai pembahasan seluruh kode pemrograman itu telah lengkap di dokumentasikan oleh tim dari [Golang](https://go.dev/doc/). Dan seluruh flow atau alur proses belajar ini diamnbil sebagian besar dari tim Go tersebut. Disini saya akan ambil kesimpulan dari apa yang dibaca didokumentasi tersebut.

> Versi Go yg digunakan di repository ini adalah latest stable Go di versi **`1.23.2`** pada environment **`Windows 10`**.

### - Instalasi pada Environment Windows 10
   1. Download *installer*-nya [Go](https://go.ddev/dl/).
   2. Klik pada bagian ***Microsoft Windows***, lalu klik `go.1.23.2.windows-amd64.msi`.
   3. Selesai proses download, double click pada installer Go tersebut lalu `next` sampai `finish`.
   4. Buka `Command Prompt` atau `cmd`, lalu ketikan perintah :
   
      ```go
      go version
      ```
      > Biasanya setiap windows akan menghasilkan *path/directory Go* yg berbeda, tergantung *default directory* yang di set pada saat kita instalasi file Go. Ada yg *default path*-nya ke *`C:\go`* dan saya sendiri default path-nya di arahkan ke *`C:\Program Files\Go`*, tapi apakah nanti akan jadi masalah pada saat development proyek yang harus butuh banyak dependensi misalnya? Kalau itu terjadi, saya akan update di repository ini.

   5. Kalau outputnya adalah `go version go1.23.2 windows/amd64`, itu tandanya kita sudah berhasil menginstalasi Go kedalam environment kita.
      > Kalau pada saat menjalankan perintah `go version` outputnya adalah error maka biasanya kita bisa coba refresh dengan menutup lalu membuka kembali `cmd` lalu ketikan perintah `go version` kembali. 
