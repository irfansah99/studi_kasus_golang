package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Karyawan struct {
	ID      int
	Nama    string
	Jabatan string
	Gaji    int
}

func main() {
	karyawan := []Karyawan{
		{1, "Andi", "Programmer", 7000000},
		{2, "Budi", "Designer", 6000000},
		{3, "Citra", "Manager", 10000000},
		{4, "Doni", "Programmer", 7500000},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n----------------------------")
		fmt.Println("Pilih tindakan:")
		fmt.Println("1. Tampilkan Seluruh data Karyawan")
		fmt.Println("2. Mencari Karyawan berdasar id")
		fmt.Println("3. Menampilkan Karyawan berdasar jabatan")
		fmt.Println("4. Mengubah gaji berdasar Id")
		fmt.Println("5. Menghapus karyawan berdasar Id")
		fmt.Println("6. Selesai")

		fmt.Print("Pilih (1-6): ")

		if scanner.Scan() {

			pilihan := strings.TrimSpace(scanner.Text())

			switch pilihan {
			case "1":
				fmt.Println("Data Seluruh Karyawan")
				for _, data := range karyawan {
					fmt.Printf("%d nama : %s \n Jabatan : %s \n Gaji : %d \n", data.ID, data.Nama, data.Jabatan, data.Gaji)
				}
			case "2":
				fmt.Println("Masukan Id Karyawan : ")
				if scanner.Scan() {
					if err := scanner.Err(); err != nil {
						fmt.Println("Terjadi kesalahan:", err)
					}
					cari := strings.TrimSpace(scanner.Text())
					cariid, err := strconv.Atoi(cari)
					if err != nil {
						fmt.Println("ID harus berupa angka")
						break
					}
					ditemukan, hasil := CekData(karyawan, cariid)

					if !ditemukan {
						fmt.Println("Data tidak ditemukan")
						break
					}

					fmt.Printf("\n--- Ditemukan! ---\n")
					fmt.Printf("ID     : %d\n", hasil.ID)
					fmt.Printf("Nama   : %s\n", hasil.Nama)
					fmt.Printf("Jabatan: %s\n", hasil.Jabatan)
					fmt.Printf("Gaji   : %d\n", hasil.Gaji)

				}
			case "3":
				fmt.Print("Masukkan Jabatan: ")

				if scanner.Scan() {
					if err := scanner.Err(); err != nil {
						fmt.Println("Terjadi kesalahan:", err)
					}
					cari := strings.TrimSpace(scanner.Text())
					ditemukan := false

					for _, data := range karyawan {
						if strings.EqualFold(data.Jabatan, cari) {
							ditemukan = true

							fmt.Printf("\n--- Ditemukan! ---\n")
							fmt.Printf("ID      : %d\n", data.ID)
							fmt.Printf("Nama    : %s\n", data.Nama)
							fmt.Printf("Jabatan : %s\n", data.Jabatan)
							fmt.Printf("Gaji    : %d\n", data.Gaji)
							fmt.Println("----------------------------")
						}
					}

					if !ditemukan {
						fmt.Println("\n--- Jabatan Tidak Ditemukan! ---")
						break
					}
				}
			case "4":
				fmt.Println("Masukkan id dan gaji yang ingin diubah dengan format [id gaji] : ")

				if scanner.Scan() {
					if err := scanner.Err(); err != nil {
						fmt.Println("Terjadi kesalahan:", err)
					}
					input := strings.TrimSpace(scanner.Text())
					data := strings.Fields(input)

					if len(data) != 2 {
						fmt.Println("Format salah! Contoh: 2 8500000")
						break
					}

					id, err := strconv.Atoi(data[0])
					if err != nil {
						fmt.Println("ID harus berupa angka")
						break
					}

					gaji, err := strconv.Atoi(data[1])
					if err != nil || gaji < 0 {
						fmt.Println("Gaji harus berupa angka positif")
						break
					}

					ditemukan, hasil := CekData(karyawan, id)

					if !ditemukan {
						fmt.Println("Data tidak ditemukan")
						break
					}

					hasil.Gaji = gaji

					fmt.Printf("Gaji karyawan dengan ID %d berhasil diubah menjadi %d\n", id, gaji)
				}
			case "5":
				fmt.Println("Masukkan id karyawan yang ingin dihapus : ")
				if scanner.Scan() {
					if err := scanner.Err(); err != nil {
						fmt.Println("Terjadi kesalahan:", err)
					}
					input := strings.TrimSpace(scanner.Text())
					id, err := strconv.Atoi(input)
					if err != nil {
						fmt.Println("ID harus berupa angka")
						break
					}
					var ditemukan bool
					karyawan, ditemukan = HapusData(karyawan, id)

					if !ditemukan {
						fmt.Println("Data tidak ditemukan")
						break
					}

					fmt.Println("Data berhasil dihapus")
				}
			case "6":
				fmt.Println("Program selesai.")
				return

			default:
				fmt.Println("Pilihan tidak tersedia!")
			}
		}

	}
}

func CekData(data []Karyawan, input int) (bool, *Karyawan) {
	for i := range data {
		if data[i].ID == input {
			return true, &data[i]
		}
	}

	return false, nil
}
func HapusData(data []Karyawan, input int) ([]Karyawan, bool) {
	for i := range data {
		if data[i].ID == input {
			data = append(data[:i], data[i+1:]...)
			return data, true
		}
	}

	return data, false
}
