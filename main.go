package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var choise string
	for choise != "x" {
		fmt.Println("Input pilihan:")
		fmt.Println("1. Penjumlahan")
		fmt.Println("2. Pengurangan")
		fmt.Print("Masukkan pilihan anda: ")
		choise = readString()
		pilih(choise)
		fmt.Println()
	}
}

func pilih(choise string) {
	switch choise {
	case "1":
		fmt.Println("======Penjumlahan======")
		break
	case "2":
		fmt.Println("======Pengurangan======")
		break
	default:
		fmt.Println("======Salah pilih bos======")
	}
}

func readString() string {
	input := bufio.NewReader(os.Stdin)
	str, err := input.ReadString('\n')
	if err != nil {
		fmt.Println("error input", err)
	}

	strNew := strings.Split(str, "\n")
	return strNew[0]
}
