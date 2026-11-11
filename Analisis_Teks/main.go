package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan Text: ")
	if scanner.Scan() {
		text := scanner.Text()

		tanpaSpasi := strings.Join(strings.Fields(text), "")
		kata := strings.Fields(text)
		char := []rune(tanpaSpasi)

		jumlahVokal := 0
		jumlahKonsonan := 0

		for _, val := range char {
			if !unicode.IsLetter(val) {
				continue
			}

			if strings.ContainsRune("aiueoAIUEO", val) {
				jumlahVokal++
			} else {
				jumlahKonsonan++
			}
		}

		fmt.Printf("Kalimat           : %s\n", text)
		fmt.Printf("Jumlah karakter   : %d\n", len([]rune(text)))
		fmt.Printf("Tanpa spasi       : %d\n", len(char))
		fmt.Printf("Jumlah vokal      : %d\n", jumlahVokal)
		fmt.Printf("Jumlah konsonan   : %d\n", jumlahKonsonan)
		fmt.Printf("Jumlah kata       : %d\n", len(kata))
	}
}


