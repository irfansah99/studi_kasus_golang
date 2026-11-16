package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	produk := map[string]int{
		"Indomie": 3000,
		"Aqua":    4000,
		"Roti":    5000,
		"Susu":    7000,
		"Kopi":    6000,
	}
	var pembeli string
	pesanan := make(map[string]int)
	var jumlahbayar int
	tagihan := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("=== TOKO Agung ===")

		i := 1
		for nama, harga := range produk {
			fmt.Printf("%d. %s : %d \n", i, nama, harga)
			i++
		}
		fmt.Println("Masukkan Produk yang ingin dibeli dan jumlahnya dengan format [ Indomie 2 ]: ")
		fmt.Println("Jika barang sebelumnya sudah ditambahkan, proses kedua akan mengubah jumlah barang")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(input)

		if len(parts) < 2 {
			fmt.Println("\n[Error: Format input salah! Gunakan format: NamaBarang Jumlah]")
			continue
		}

		namaBarang := parts[0]
		jumlah, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("\n[Error: Jumlah harus berupa angka yang valid!]")
			continue
		}

		nama, _, ditemukam := cek(produk, namaBarang)

		if ditemukam {
			pesanan[nama] = jumlah
			fmt.Printf("\n[Sukses! Menambahkan '%s' dengan jumlah %d ke keranjang]\n", nama, jumlah)
		} else {
			fmt.Printf("\n[Tidak ada barang dengan nama '%s' ]\n", namaBarang)
		}

		fmt.Println("\n----------------------------")
		fmt.Println("Pilih tindakan selanjutnya:")
		fmt.Println("1. Lanjut belanja barang lain")
		fmt.Println("2. Lihat isi keranjang")
		fmt.Println("3. Hapus barang dari keranjang")
		fmt.Println("4. Selesaikan pesanan (Checkout)")
		fmt.Print("Pilih (1-4): ")

		if scanner.Scan() {
			konfirmasi := strings.TrimSpace(scanner.Text())

			if konfirmasi == "4" {
				fmt.Println("\nTerima kasih telah berbelanja!")
				fmt.Println("Berikut Barang yang anda pesan:")

				for nama, jumlah := range pesanan {
					for namaproduk, harga := range produk {
						if nama == namaproduk {
							tagihan[nama] = harga * jumlah
						}
					}
					fmt.Printf("- %s dengan jumlah %d : %d\n", nama, jumlah, tagihan[nama])
				}

				for _, harga := range tagihan {
					jumlahbayar += harga
				}

				if jumlahbayar > 30000 {
					diskon := jumlahbayar * 10 / 100
					jumlahbayar -= diskon
					fmt.Println("\n[Selamat! Anda mendapatkan diskon 10% karena total belanja > 30.000]")
				}
				break

			} else if konfirmasi == "2" {
				fmt.Println("\n=== ISI KERANJANG SEMENTARA ===")
				if len(pesanan) == 0 {
					fmt.Println("Keranjang masih kosong.")
				} else {
					for nama, jumlah := range pesanan {
						harga := produk[nama]
						fmt.Printf("- %s (Jumlah: %d) | Subtotal: %d\n", nama, jumlah, harga*jumlah)
					}
				}

			} else if konfirmasi == "3" {
				fmt.Println("\nMasukkan nama barang yang ingin dihapus dari keranjang:")
				if scanner.Scan() {
					namaBarang := strings.TrimSpace(scanner.Text())
					nama, _, ditemukam := cek(produk, namaBarang)

					if ditemukam {

						_, adaDiPesanan := pesanan[nama]
						if adaDiPesanan {
							delete(pesanan, nama)
							fmt.Printf("\n[Sukses! Barang '%s' berhasil dihapus dari keranjang]\n", nama)
						} else {
							fmt.Printf("\n[Barang '%s' tidak ada di dalam keranjang Anda]\n", nama)
						}
					} else {
						fmt.Printf("\n[Tidak ada produk bernama '%s' di toko]\n", namaBarang)
					}
				}

			} else if konfirmasi == "1" {
				continue
			} else {
				fmt.Println("\nPilihan tidak dikenali, otomatis melanjutkan kembali...")
			}
		}
	}

	fmt.Println("\nMasukkan nama anda beserta uang anda dengan format [agung 50000]:")
	if scanner.Scan() {
		bayar := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(bayar)

		if len(parts) < 2 {
			fmt.Println("[Error: Format input pembayaran salah!]")
			return
		}
		pembeli = parts[0]
		uang, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("\n[Error: Uang harus berupa angka yang valid!]")
			return
		}

		fmt.Printf("\n--- STRUK PEMBELIAN ---\n")
		fmt.Printf("Nama Pembeli : %s\n", pembeli)

		for nama, jumlah := range pesanan {
			fmt.Printf("- %s (Jumlah: %d) : %d \n", nama, jumlah, tagihan[nama])
		}

		fmt.Printf("Total bayar (setelah diskon jika ada) : %d\n", jumlahbayar)
		fmt.Printf("Uang Anda : %d\n", uang)

		if uang < jumlahbayar {
			fmt.Println("Pembelian Anda Gagal")
			fmt.Printf("Uang Anda kurang sebesar : %d\n", jumlahbayar-uang)
		} else {
			fmt.Printf("Kembalian : %d\n", uang-jumlahbayar)
		}
	}
}

func cek(produk map[string]int, nama string) (string, int, bool) {
	for namabarang, jumlah := range produk {
		if strings.EqualFold(namabarang, nama) {
			return namabarang, jumlah, true
		}
	}
	return "", 0, false
}
