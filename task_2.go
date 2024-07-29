package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func genPass(password string, level string) string {
	const letters string = "abcdefghijklmnopqrstuvwxyz"
	const upperLetters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const digits string = "0123456789"
	const specialsChar string = "!@#$%^&*()_+-=[]{}|;:',.<>?/"

	// Variabel untuk menyimpan kumpulan karakter yang dipilih dan panjang password yang diizinkan
	var charSet string
	var minLength, maxLength int

	// Menentukan kumpulan karakter dan panjang password berdasarkan level
	switch level {
	case "low":
		charSet = letters + digits
		minLength = 6
		maxLength = 10
	case "med":
		charSet = letters + upperLetters + digits
		minLength = 11
		maxLength = 15
	case "strong":
		charSet = letters + upperLetters + digits + specialsChar
		minLength = 16
		maxLength = 20
	default:
		fmt.Println("Invalid level, use 'low', 'med', or 'strong'")
		return ""
	}

	// Mengatur panjang password yang dihasilkan
	passwordLength := minLength + rand.Intn(maxLength-minLength+1)

	// Mengonversi password yang diberikan menjadi slice of rune
	passwordSlice := []rune(password)

	// Inisialisasi generator angka acak
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// Menambahkan setidaknya satu karakter dari setiap jenis yang diperlukan
	if strings.Contains(charSet, digits) {
		passwordSlice = append(passwordSlice, rune(digits[r.Intn(len(digits))]))
	}
	if strings.Contains(charSet, upperLetters) {
		passwordSlice = append(passwordSlice, rune(upperLetters[r.Intn(len(upperLetters))]))
	}
	if strings.Contains(charSet, letters) {
		passwordSlice = append(passwordSlice, rune(letters[r.Intn(len(letters))]))
	}
	if strings.Contains(charSet, specialsChar) {
		passwordSlice = append(passwordSlice, rune(specialsChar[r.Intn(len(specialsChar))]))
	}

	// Menambahkan karakter acak dari kumpulan karakter sampai mencapai panjang password yang diinginkan
	for len(passwordSlice) < passwordLength {
		char := charSet[r.Intn(len(charSet))]
		passwordSlice = append(passwordSlice, rune(char))
	}

	// Mengacak urutan karakter dalam password
	for i := range passwordSlice {
		j := r.Intn(len(passwordSlice))
		passwordSlice[i], passwordSlice[j] = passwordSlice[j], passwordSlice[i]
	}

	// Mengonversi slice of rune kembali menjadi string dan mengembalikannya
	return string(passwordSlice)
}
