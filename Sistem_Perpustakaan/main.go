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
				fmt.Println("Masukan Informasi Buku yang ingin ditambah dengan format [judul , pencipta , stok]")
				if scanner.Scan() {
					input := strings.TrimSpace(scanner.Text())
					data, err := ParseBukuInput(input)
					if err != nil {
						fmt.Println(err)
						break
					}
					datasebelum, ditemukan := Cari_judul(buku, data.Judul)
					if ditemukan {
						fmt.Printf("Gagal, Buku dengan judul %s sudah tersedia\n", datasebelum.Judul)
						break
					}

					buku = append(buku,
						Buku{ID: len(buku) + 1, Judul: data.Judul, Penulis: data.Penulis, Stok: data.Stok},
					)
					fmt.Printf("Buku dengan judul '%s' berhasil ditambah\n", data.Judul)
				}
			case "3":
				fmt.Println("Masukan Informasi Buku yang ingin diubah dengan format [Judul, pencipta, stok]")
				if scanner.Scan() {
					input := strings.TrimSpace(scanner.Text())
					data, err := ParseBukuInput(input)
					if err != nil {
						fmt.Println(err)
						break
					}
					buku, ditemukan := Cari_judul(buku, data.Judul)
					if !ditemukan {
						fmt.Printf("Gagal, Buku dengan judul %s tidak tersedia\n", data.Judul)
						break
					}
					buku.Penulis = data.Penulis
					buku.Stok = data.Stok
					fmt.Printf("Buku dengan judul %s berhasil diubah\n", buku.Judul)
				}
			case "4" :
				fmt.Println("Masukan Judul Buku yang ingin dihapus")
			}

		}
	}
}

func Cari_judul(buku []Buku, judul string) (*Buku, bool) {
	for i, data := range buku {
		if strings.EqualFold(data.Judul, judul) {
			return &buku[i], true
		}
	}
	return nil, false
}

func ParseBukuInput(input string) (Buku, error) {
	parts := strings.Split(input, ",")
	if len(parts) != 3 {
		return Buku{}, fmt.Errorf("format salah, gunakan format: judul, pencipta, stok")
	}

	inputJudul := strings.TrimSpace(parts[0])
	inputPencipta := strings.TrimSpace(parts[1])
	inputStok := strings.TrimSpace(parts[2])

	if inputJudul == "" || inputPencipta == "" {
		return Buku{}, fmt.Errorf("judul dan pencipta tidak boleh kosong")
	}

	stok, err := strconv.Atoi(inputStok)
	if err != nil {
		return Buku{}, fmt.Errorf("format salah, stok harus angka")
	}

	hasil := Buku{

		Judul:   inputJudul,
		Penulis: inputPencipta,
		Stok:    stok,
	}

	return hasil, nil
}
