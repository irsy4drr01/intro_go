package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int

	fmt.Println("-PRINT SEGITIGA-")
	fmt.Print("Input value n: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Input error: ", err)
		return
	}
	printSegitiga(n)

	var password string
	var level string

	fmt.Println("-GENERATE PASSWORD-")
	fmt.Print("Input Password (minimum 6 characters): ")
	_, err = fmt.Scanln(&password)
	if err != nil {
		fmt.Println("Input error: ", err)
		return
	}
	password = strings.TrimSpace(password)

	if len(password) < 6 {
		fmt.Println("Input password minimum 6 characters.")
		return
	}

	fmt.Print("Input level ('low', 'med', 'strong'): ")
	_, err = fmt.Scanln(&level)
	if err != nil {
		fmt.Println("Input level error: ", err)
		return
	}
	level = strings.TrimSpace(level)

	newPassword := genPass(password, level)
	if newPassword == "" {
		return
	}
	fmt.Println("New Password: ", newPassword)
	fmt.Println()

	fmt.Println("-MOVIE DURATIONS-")
	movieDurations := []int{1, 2, 2, 7, 3, 4, 8, 9}

	fmt.Print("Input duration of flight n: ")
	_, err = fmt.Scanln(&n)
	if err != nil {
		fmt.Println("Input error: ", err)
		return
	}

	moviePairs, singleEqual, singleLess := findMovies(movieDurations, n)
	if len(moviePairs) > 0 {
		fmt.Println("Kombinasi film yang durasinya sama dengan durasi penerbangan:")
		for _, pair := range moviePairs {
			fmt.Printf("Film dengan durasi %d dan %d\n", pair[0], pair[1])
		}
	} else if len(singleEqual) > 0 || len(singleLess) > 0 {
		if len(singleEqual) > 0 {
			fmt.Println("Film dengan durasi sama dengan durasi penerbangan:")
			for _, duration := range singleEqual {
				fmt.Printf("Film dengan durasi %d\n", duration)
			}
		}
		if len(singleLess) > 0 {
			fmt.Println("Film dengan durasi kurang dari durasi penerbangan:")
			for _, duration := range singleLess {
				fmt.Printf("Film dengan durasi %d\n", duration)
			}
		}
	} else {
		fmt.Println("Tidak ada film yang memiliki durasi sesuai dengan kriteria.")
	}
}
