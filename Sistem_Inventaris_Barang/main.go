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
	stok := map[string]int{
		"Laptop":   10,
		"Mouse":    25,
		"Keuboard": 15,
		"Monitor":  8,
		"Headset":  20,
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("=== INVENTARIS TOKO ===")
		fmt.Println("1. Info Inventories\n2. Cari barang berdasar nama\n3. Tambah Barang baru\n4. Hapus Barang berdasar nama\n5. Mengubah stok barang berdasar nama\n6. Keluar")
		fmt.Print("Masukkan Perintah: ")

		if !scanner.Scan() {
			break
		}
		pilihan := strings.TrimSpace(scanner.Text())
		switch pilihan {
		case "1":
			hasil := info(stok)
			fmt.Print("\n" + hasil)
		case "2":
			fmt.Print("Masukkan nama barang yang dicari: ")
			if !scanner.Scan() {
				break
			}
			input := strings.TrimSpace(scanner.Text())
			nama, jumlah, ditemukan := cek(stok, input)

			if ditemukan {
				fmt.Printf("\n%s : %d\n", nama, jumlah)
			} else {
				fmt.Println("\n[Barang tidak tersedia di inventaris]")
			}
		case "3":
			fmt.Print("Masukan Nama barang dan stok dengan format 'Handphone 21' ")
			if !scanner.Scan() {
				break
			}
			input := strings.TrimSpace(scanner.Text())
			parts := strings.Fields(input)
			if len(parts) < 2 {
				fmt.Println("\n[Format salah! Gunakan format: NamaBarang JumlahStok (contoh: Handphone 21)]")
				break
			}
			namaBarang := parts[0]

			jumlahStok, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Println("\n[Error: Jumlah stok harus berupa angka yang valid!]")
				break
			}

			stok[namaBarang] = jumlahStok
			fmt.Printf("\n[Sukses! Barang '%s' dengan stok %d berhasil ditambahkan/diperbarui]\n", namaBarang, jumlahStok)
		case "4":
			fmt.Println("Program selesai.")
			return

		default:
			fmt.Println("Pilihan tidak tersedia!")
		}
		fmt.Println("\n----------------------------")
		fmt.Println("Ingin melakukan perintah lain?")
		fmt.Println("1. Lanjut")
		fmt.Println("2. Berhenti")
		fmt.Print("Pilih (1/2): ")

		if scanner.Scan() {
			konfirmasi := strings.TrimSpace(scanner.Text())

			if konfirmasi == "2" {
				fmt.Println("\nTerima kasih telah menggunakan program inventaris. Sampai jumpa!")
				break
			} else if konfirmasi != "1" {
				fmt.Println("\nPilihan tidak dikenali, otomatis melanjutkan kembali...")
			}
		}
	}

}

func info(stok map[string]int) string {
	var semuaNilai []int
	var sb strings.Builder
	total := 0
	fmt.Println("=== INVENTARIS TOKO ===")
	for nama, jumlah := range stok {
		semuaNilai = append(semuaNilai, jumlah)
		total = total + jumlah
		fmt.Printf("%v : %d\n", nama, jumlah)
	}
	sb.WriteString(fmt.Sprintf("Total Barang : %d\n", total))
	sb.WriteString(fmt.Sprintf("Stok Terbanyak : %d\n", slices.Max(semuaNilai)))
	sb.WriteString(fmt.Sprintf("Stok Terendah : %d\n", slices.Min(semuaNilai)))

	return sb.String()
}

func cek(stok map[string]int, input string) (string, int, bool) {
	for nama, jumlah := range stok {
		if strings.EqualFold(nama, input) {
			return nama, jumlah, true
		}
	}

	return "", 0, false
}
