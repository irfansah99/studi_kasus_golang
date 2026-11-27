package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Buku struct {
	ID      int
	Judul   string
	Penulis string
	Stok    int
}
type Pinjam struct {
	ID      int
	Nama    string
	Id_Buku int
}

func main() {
	buku := []Buku{
		{1, "One Piece", "Oda", 20},
		{2, "Naruto", "Kishimoto", 40},
		{3, "Bleach", "Kubo", 5},
		{4, "Attack on Titan", "Isayama", 15},
		{5, "Demon Slayer", "Gotouge", 25},
		{6, "Jujutsu Kaisen", "Akutami", 12},
		{7, "Chainsaw Man", "Fujimoto", 18},
		{8, "Spy x Family", "Endo", 30},
		{9, "Hunter x Hunter", "Togashi", 8},
		{10, "Fullmetal Alchemist", "Arakawa", 14},
	}
	// Pinjam := []Pinjam{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n----------------------------")
		fmt.Println("Pilih tindakan:")
		fmt.Println("1. Tampilkan Daftar Buku")
		fmt.Println("2. Tambah Buku")
		fmt.Println("3. Ubah informasi buku")
		fmt.Println("4. Hapus berdasarkan judul")
		fmt.Println("5. Mencari buku berdasarkan judul")
		fmt.Println("6. Pinjam Buku")
		fmt.Println("7. Kembalikan Buku")
		fmt.Println("8. Daftar Pinjam")
		fmt.Println("9. Selesai")
		fmt.Print("Pilih (1-9): ")

		if scanner.Scan() {
			input := strings.TrimSpace(scanner.Text())

			switch input {
			case "1":
				fmt.Println("Daftar Buku yang tersedia")
				for _, data := range buku {
					fmt.Printf("Id : %d \n Judul : %s \n Pencipta : %s \n Stok : %d \n ", data.ID, data.Judul, data.Penulis, data.Stok)
				}
			case "2":
				fmt.Println("Masukan Informasi Buku yang ingin ditambah dengan format [judul pencipta stok]")
				if scanner.Scan() {
					input := strings.TrimSpace(scanner.Text())
					parts := strings.Fields(input)
					stok, err := strconv.Atoi(parts[2])
					if err != nil {
						fmt.Println("Format salah, Stok harus angka")
						break
					}
					buku = append(buku,
						Buku{ID: len(buku) + 1, Judul: parts[0], Penulis: parts[1], Stok: stok},
					)
					fmt.Printf("Buku dengan judul %s berhasil ditambah", parts[0])
				}
			}

		}
	}
}
