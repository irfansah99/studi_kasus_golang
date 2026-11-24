package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	sekarang := time.Now()

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		proses("A", 3*time.Second)
	}()
	go func() {
		defer wg.Done()
		proses("B", 2*time.Second)
	}()
	go func() {
		defer wg.Done()
		proses("C", 1*time.Second)

	}()
	wg.Wait()
	selisih := time.Since(sekarang)
	fmt.Printf("Total waktu eksekusi: %v\n", int(math.Round(selisih.Seconds())))
}

func proses(nama string, durasi time.Duration) {
	fmt.Printf("Pekerjaan %s mulai (durasi: %v)\n", nama, durasi)
	time.Sleep(durasi)
	fmt.Printf("Pekerjaan %s selesai\n", nama)
}
