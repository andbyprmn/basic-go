# 6. Operator

> Kali ini saya dapat error menarik, seperti dibawah ini :
> ```go 
>    go reading .\main.go unexpected NUL in input  
> ```
> Ternyata setelah saya *crosscheck*, hal itu dikarenakan saya membuat file main.go itu dengan *powershell* di *Windows 10* menggunakan command `echo "package main" >> main.go`. Sehingga pada saat *create* file *main.go* hal itu menyebabkan kesalahan encoding yg telah di-generate oleh command tersebut. Selengkapnya sila cek [*Stackoverflow*](https://stackoverflow.com/a/48870543).