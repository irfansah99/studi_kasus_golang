package main

import "fmt"

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
	for _, data := range karyawan {
		fmt.Printf("%d nama : %s \n Jabatan : %s \n Gaji : %d \n",data.ID, data.Nama, data.Jabatan, data.Gaji)
	}
}
