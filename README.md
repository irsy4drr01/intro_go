# Proyek Introduction Go

Proyek ini terdiri dari beberapa fungsi yang dibagi ke dalam beberapa file untuk mencetak segitiga, menghasilkan kata sandi, dan menemukan kombinasi film berdasarkan durasi penerbangan.

## Struktur Proyek

Proyek ini terdiri dari empat file utama:

1. `main.go` - Berisi fungsi `main()` yang mengatur alur eksekusi program.
2. `1._gen_pass.go` - Berisi fungsi `printSegitiga()` untuk mencetak segitiga.
3. `2._print_segitiga.go` - Berisi fungsi `genPass()` untuk menghasilkan kata sandi berdasarkan level keamanan.
4. `3._find_movies.go` - Berisi fungsi `findMovies()` untuk menemukan kombinasi film berdasarkan durasi penerbangan.

## Fungsi Utama

### main.go

File ini mengandung fungsi `main()` yang mengatur alur eksekusi program:

- Meminta pengguna untuk memasukkan nilai `n` dan mencetak segitiga menggunakan `printSegitiga(n)`.
- Meminta pengguna untuk memasukkan kata sandi dan level keamanan, lalu menghasilkan kata sandi baru menggunakan `genPass(password, level)`.
- Meminta pengguna untuk memasukkan durasi penerbangan dan menemukan kombinasi film menggunakan `findMovies(movieDurations, n)`.

### task_1.go

File ini mengandung fungsi `printSegitiga(n int)` yang digunakan untuk mencetak segitiga berdasarkan nilai `n` yang diberikan.

### task_2.go

File ini mengandung fungsi `genPass(password string, level string) string` yang digunakan untuk menghasilkan kata sandi baru berdasarkan kata sandi yang diberikan dan level keamanan.

### task_3.go

File ini mengandung fungsi `findMovies(durations []int, n int) ([][]int, []int, []int)` yang digunakan untuk menemukan kombinasi film berdasarkan durasi penerbangan yang diberikan:

- Mengecek pasangan film dengan total durasi yang sama dengan durasi penerbangan.
- Jika tidak ditemukan, mengecek pasangan film dengan total durasi kurang dari durasi penerbangan.
- Jika tidak ditemukan, mengecek apakah ada satu film dengan durasi yang sama dengan durasi penerbangan atau satu film dengan durasi kurang dari durasi penerbangan.

## Cara Menjalankan Program

1. Pastikan Anda memiliki Go terinstal di sistem Anda. Anda dapat mengunduh dan menginstalnya dari [situs resmi Go](https://golang.org/dl/).
2. Salin semua file (`main.go`, `task_1.go`, `task_2.go`, `task_3.go`) ke dalam satu direktori.
3. Buka terminal, navigasikan ke direktori tersebut, dan jalankan perintah berikut untuk menjalankan program:

    ```sh
    go run main.go task_1.go task_2.go task_3.go
    ```
    atau dengan perintah yang lebih sederhana
    ```sh
    go run .
    ```

4. Ikuti petunjuk yang diberikan oleh program untuk memasukkan nilai-nilai yang diperlukan.

## Penulis

Proyek ini dikembangkan oleh Irsyad Ridho Romadhon.