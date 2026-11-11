package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan Nilai: ")

	if scanner.Scan() {
		inputStr := strings.TrimSpace(scanner.Text())

		parts := strings.Fields(inputStr)

		var nilai []int
		total := 0
		lulus := 0
		gagal := 0
		for _, p := range parts {
			angka, err := strconv.Atoi(p)
			if err != nil {
				fmt.Printf("Gagal '%s' bukan angka", p)
				return
			}

			if angka >= 75 {
				lulus++
			} else {
				gagal++
			}

			nilai = append(nilai, angka)
			total += angka
		}
		if len(nilai) == 0 {
			fmt.Println("Belum ada nilai yang dimasukkan")
			return
		}
		fmt.Printf("Daftar Nilai : %v\n", nilai)
		fmt.Printf("Nilai Tertinggi  : %d\n", slices.Max(nilai))
		fmt.Printf("Nilai Terendah  : %d\n", slices.Min(nilai))
		fmt.Printf("Rata-rata  : %.2f\n", float64(total)/float64(len(nilai)))
		fmt.Printf("Jumlah Lulus  : %d\n", lulus)
		fmt.Printf("Tidak Lulus   : %d\n", gagal)
	}
}

// func stdi1(text) {
// 	if scanner.Scan() {
// 		text := scanner.Text()

// 		tanpaSpasi := strings.Join(strings.Fields(text), "")
// 		kata := strings.Fields(text)
// 		char := []rune(tanpaSpasi)

// 		jumlahVokal := 0
// 		jumlahKonsonan := 0

// 		for _, val := range char {
// 			if !unicode.IsLetter(val) {
// 				continue
// 			}

// 			if strings.ContainsRune("aiueoAIUEO", val) {
// 				jumlahVokal++
// 			} else {
// 				jumlahKonsonan++
// 			}
// 		}

// 		fmt.Printf("Kalimat           : %s\n", text)
// 		fmt.Printf("Jumlah karakter   : %d\n", len([]rune(text)))
// 		fmt.Printf("Tanpa spasi       : %d\n", len(char))
// 		fmt.Printf("Jumlah vokal      : %d\n", jumlahVokal)
// 		fmt.Printf("Jumlah konsonan   : %d\n", jumlahKonsonan)
// 		fmt.Printf("Jumlah kata       : %d\n", len(kata))
// 	}
// }
