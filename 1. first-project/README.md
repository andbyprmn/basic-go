# Go Toolchain

Toolchain Go merujuk pada seperangkat alat (*tools*) yang disediakan oleh Go untuk membantu kita dalam menulis, mengkompilasi, menguji, dan mengelola kode dalam proyek atau aplikasi Golang yang kita buat. Toolchain ini isinya adalah berbagai perintah (*command*) yang dapat dijalankan melalui baris perintah (***command line interface*** atau ***CLI***). Perintah-perintah ini tidak hanya mempermudah proses pengembangan tetapi juga memastikan bahwa kode yang dihasilkan konsisten, teruji, dan dapat diandalkan.


### Mengapa Perintah dalam Go itu Penting?
Perintah-perintah Go adalah jantung dari proses development dengan bahasa Go. Golang memungkinkan kita untuk:
1. **Mengelola Proyek**: Menginisialisasi proyek baru, mengelola dependensi, dan menjaga struktur proyek tetap terorganisir.
2. **Kompilasi dan Eksekusi**: Mengkompilasi kode sumber menjadi executable, menjalankan aplikasi, dan menguji performa.
3. **Pengujian dan Dokumentasi**: Menulis dan menjalankan unit test, serta menghasilkan dokumentasi otomatis.
4. **Pemeliharaan dan Optimisasi**: Membersihkan build cache, memformat kode secara otomatis, dan memverifikasi integritas dependensi.

Dengan memahami dan memanfaatkan perintah-perintah atau *commands* ini secara efektif, kita dapat meningkatkan produktivitas, memastikan kualitas kode, dan mempercepat development lifecycle.


### Perintah-perintah Golang
1. **`go run`**: perintah untuk menjalankan program Go secara langsung tanpa perlu proses *compile* menjadi binary terlebih dahulu.
   ```go
   go run main.go
   ```
2. **`go build`**: perintah untuk menjalankan proses *compile* paket atau file Go menjadi binary.
    ```go
    go build main.go
    ```
3. **`go install`**: perintah untuk menjalankan proses *compile* dan menginstalasi paket atau binary ke direktori **`$GOPATH/bin`** atau modul Go.
    ```go
    go install main.go
    ```
4. **`go fmt`**: perintah untuk memformat kode sumber Go sesuai dengan standar gaya penulisan yg ditetapkan oleh komunitas Go.
    ```go
    go fmt main.go
    ```
5. **`go test`**: perintah untuk menjalankan unit test dan benchmark dalam paket Go.
   ```go
   go test ./...
   ```
6. **`go get`**: perintah untuk mengambil dan menginstall paket Go dari repository.
    ```go
    go get github.com/gin-gonic/gin
    ```
7. **`go mod`**: perintah untuk mengelola modul Go yang mencakup pembuatan, pembaruan dan pemeliharaan dependensi. Dalam *`go mod`* ada beberapa sub-perintah yang umum digunakan :
   - *`go mod init [nama_modul]`*: Inisialisasi modul baru.
   - *`go mod tidy`*: Menambah dan menghapus dependensi yang diperlukan.
   - *`go mod download`*: Mengunduh semua dependensi yang diperlukan.
   - *`go mod verify`*: Memverifikasi integritas dependensi.
    ```go
    go mod init github.com/username/proyek
    go mod tidy
    ```
8. **`go clean`**: perintah untuk menghapus file-file yang dihasilkan oleh perintah seperti build, test maupun install.
    ```go
    go clean -cache
    ```
9.  **`go env`**: perintah untuk menampilkan atau menagtur ulang environment variable yang digunakan oleh Go.
    ```go
    go env
    go env GOPATH
    ```
10. **`go list`**: perintah untuk menampilkan informasi tentang paket-paket atau libraries di Go.
    ```go
    go list ./...
    go list -m all
    ```
11. **`go generate`**: perintah untuk menjalankan generator kode yang didefinisikan dalam komentar *`//go:generate`* di dalam source code.
    ```go
    //go:generate go run tools/gen.go
    go generate ./...
    ```
12. **`go doc`**: perintah untuk menampilkan dokumentasi paket, fungsi, atau tipe Go.
    ```go
    go doc fmt.Println
    ```
13. **`go mod vendor`**: perintah untuk copy semua dependensi ke dalam direktori vendor di proyek.
    ```go
    go mod vendor
    ```
14. **`go tool`**: perintah untuk memberikan akses ke berbagai developer tools di Go yang lebih mendalam, seperti debugging, profiling, dan analisis.
    ```go
    go tool pprof cpu.prof
    ```
15. **`go version`**: perintah untuk menampilkan versi Go yang terpasang di sistem.
    ```go
    go version
    ```
16. **`go list -json`**: perintah untuk menampilkan informasi paket/library dalam format JSON, memungkinkan integrasi dengan alat lain atau skrip.
    ```go
    go list -json ./...
    ```