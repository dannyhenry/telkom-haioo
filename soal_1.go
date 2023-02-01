package main

import (
	"fmt"
	"math"
)

func fractionsMoney(price int) []interface{} {
	money := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 50}
	result := make(map[string]int)
	var res []interface{}
	for _, m := range money {
		if price >= m {
			count := int(math.Floor(float64(price) / float64(m)))
			price = price % m
			result[fmt.Sprintf("Rp. %d", m)] = count
		}
	}

	for k, v := range result {
		res = append(res, map[string]interface{}{k: v})
	}
	return res
}

func main() {
	fmt.Println(fractionsMoney(145000))
	fmt.Println(fractionsMoney(2050))
}

/*
Fungsi countMoney memiliki parameter price yang berisi harga yang ingin dicari pecahannya.
Variabel money berisi daftar nominal uang yang tersedia.
Variabel result berisi hasil akhir yang akan dikembalikan oleh fungsi.
Perulangan for digunakan untuk mengecek apakah price lebih besar dari atau sama dengan nominal uang yang ada di daftar money.
Jika price lebih besar dari atau sama dengan nominal uang, maka ditemukan berapa banyak lembar uang yang dibutuhkan dengan menggunakan fungsi math.Floor untuk membulatkan ke bawah.
Kemudian, sisa dari pembagian tersebut disimpan kembali ke price untuk menentukan pecahan selanjutnya.
Hasil pembagian sebelumnya disimpan ke dalam variabel result dengan format string "Rp. [nominal uang],-".
Setelah selesai melakukan perulangan, fungsi akan mengembalikan hasil dari variabel result.
*/
