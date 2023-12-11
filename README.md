
# Capstone-Aplikasi 'Ticketing' menggunakan Golang

Proyek ini dibuat untuk untuk memenuhi syarat-syarat menyelesaikan kegiatan Study Independent MSIB batch 5 di mitra MIKTI. Selain itu, proyek ini juga bertujuan untuk memperdalam pemahaman menggunakan Golang.

## Detail Proyek
Tema: 

Aplikasi Ticketing Berbasis Web

Nama Aplikasi:

Depublic

Kelompok : 4

Deskripsi:

Platform ini merupakan tempat jual-beli tiket konser ataupun event. Platform ini membuka dan menyediakan berbagai jenis kategori kebutuhan. User yang mendaftarkan diri pada aplikasi ini dapat berperan buyer. Dalam hal ini, pengguna diharapkan dapat dengan mudah menemukan jadwal konser yang sesuai dengan kebutuhan mereka dan membayar tiket secara online. Selain itu, website ini juga diharapkan dapat menyediakan informasi yang akurat dan terkini tentang event yang sedang berlangsung.

Untuk detail lebih lanjut, dapat dilihat di [sini](https://docs.google.com/presentation/d/1vvOwCKjysLxORL7GOtOJbgiW3XAVuYsRK0ccUa8VOzA/edit#slide=id.g248d5834739_0_47).
## Fitur-fitur

- User registration
- In-App Notification
- Search and filter
- History transaction
- Payment

## Cara Menjalankan Proyek

- Unduh installer [Golang](https://golang.org/dl/) terlebih dahulu
- Setelah terunduh, jalankan installer, klik next hingga proses instalasi selesai. By default jika anda tidak merubah path pada saat instalasi, Go akan ter-install di ```C:\go```. 
- Path tersebut secara otomatis akan didaftarkan dalam ```PATH``` environment variable.
- Buka Command Prompt / CMD, eksekusi perintah berikut untuk mengecek versi Go.
```
go version
```
- Jika output adalah sama dengan versi Go yang ter-install, menandakan proses instalasi berhasil.
- Sering terjadi, command ```go version``` tidak bisa dijalankan meskipun instalasi sukses. Solusinya bisa dengan restart CMD (tutup CMD, kemudian buka lagi). Setelah itu coba jalankan ulang command di atas.
- Unduh dan install [PostgreSQL](https://www.postgresql.org/download/)
* Clone repository proyek ke lokal
```
git clone -b develop https://github.com/trikrama/Depublic.git
```
- Pindah ke direktori repositori lokal dengan command
```
cd Depublic
```
- Jalankan command ```go mod tidy``` untuk memvalidasi dependensi. Jika ada dependensi yang belum terunduh, maka akan otomatis diunduh.
- Unduh dan install [pgAdmin](https://www.pgadmin.org/download/)
* Buat database baru dengan nama ```depublic_db```
* Ganti password di file ```.env``` menggunakan password postgreSQL pribadi
* Migrate data dengan command
```
migrate -database "postgres://postgres:trikrama@localhost:5432/depublic_db?sslmode=disable" -path db/migrations up
```
- Ubah ```trikarma``` menjadi password postgreSQL pribadi
- Jalankan command
```
go run cmd/server/main.go
```


## Dokumentasi

- [Postman](https://documenter.getpostman.com/view/21997905/2s9YkhgjAD#f824225e-ea39-49a0-b149-42a7242a9d7b)