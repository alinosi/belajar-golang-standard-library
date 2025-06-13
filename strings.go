package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Eko Kurniawan", "Eko")) // apakah string mengandung kata "Eko"
	fmt.Println(strings.Split("Eko Kurniawan", " ")) // memecah setiap kata yang dipisahkan space menjadi slice
	fmt.Println(strings.ToLower("Eko Kurniawan")) // konversi ke huruf kecil
	fmt.Println(strings.ToUpper("Eko Kurniawan")) // konversi ke huruf besar
	fmt.Println(strings.Trim("      Eko Kurniawan      ", " ")) // menghapus karakter pilihan di awal dan akhir kata
	fmt.Println(strings.ReplaceAll("Eko Kurniawan Eko Khannedy", "Eko", "Budi")) // mengganti setiap kata Eko menjadi Budi
}
