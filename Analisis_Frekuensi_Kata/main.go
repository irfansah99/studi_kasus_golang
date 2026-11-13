package main

import (
	"fmt"
	"strings"
)

func main() {
	kalimat := "go belajar go programming belajar go"
	parts := strings.Fields(kalimat)
	frek := make(map[string]int)

	for _ , kata := range parts {
		frek[kata]++
	}
	var jumlahmax int
	var namamax string
	var jumlahmin int
	var namamin string
	first := true

	for nama, jumlah := range frek {
		if first {
			jumlahmax = jumlah
			jumlahmin = jumlah
			namamax = nama
			namamin = nama
			first = false
		}
		if jumlahmax > jumlah {
			namamax = nama
			jumlahmax = jumlah
		}
		if jumlahmin < jumlah  {
			namamin = nama
			jumlahmin = jumlah
		}
	} 
	fmt.Printf("Kata : %s \n", kalimat)
	for nama, jumlah := range frek {
		fmt.Printf("%s : %d \n",nama, jumlah)
	} 
	fmt.Printf("Kata terbanyak adalah %s dengan jumlah %d \n", namamax, jumlahmax)
	fmt.Printf("Kata terkecil adalah %s dengan jumlah %d \n", namamin, jumlahmin)
}
