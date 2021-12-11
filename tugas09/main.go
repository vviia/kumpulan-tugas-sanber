package main

import (
	"fmt"
	"strconv"

	"./soal1"
	"./soal2"
	"./soal3"
	"./soal4"
)

func main() {
	// soal 1
	fmt.Println("SOAL KE-1")
	var segitiga1 soal1.HitungBangunDatar = soal1.SegitigaSamaSisi{Alas: 7, Tinggi: 8}
	var persegiPanjang1 soal1.HitungBangunDatar = soal1.PersegiPanjang{Panjang: 7, Lebar: 8}

	fmt.Println("------bangun datar-------")
	fmt.Println("luas segitiga sama sisi dari alas 7 dan tinggi 8 adalah", segitiga1.Luas())
	fmt.Println("keliling segitiga sama sisi dari alas 7 dan tinggi 8 adalah", segitiga1.Keliling())

	fmt.Println("luas persegi panjang sama dari panjang 7 dan lebar 8 adalah", persegiPanjang1.Luas())
	fmt.Println("keliling persegi panjang sama dari panjang 7 dan lebar adalah", persegiPanjang1.Keliling())

	fmt.Println("------bangun ruang-------")

	var tabung1 soal1.HitungBangunRuang = soal1.Tabung{JariJari: 7, Tinggi: 10}
	var balok1 soal1.HitungBangunRuang = soal1.Balok{Panjang: 7, Lebar: 5, Tinggi: 6}

	fmt.Println("volume tabung dari jari-jari 7 dan tinggi 10 adalah", tabung1.Volume())
	fmt.Println("luas permukaan tabung dari jari-jari 7 dan tinggi 10 adalah", tabung1.LuasPermukaan())

	fmt.Println("volume balok dari panjang 7, lebar 5 dan tinggi 6 adalah", balok1.Volume())
	fmt.Println("luas permukaan balok dari panjang 7, lebar 5 dan tinggi 6 adalah", balok1.LuasPermukaan())

	// soal 2
	fmt.Println("SOAL KE-2")
	var samsung soal2.Description = soal2.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Println(samsung.GetDescription())

	// soal 3
	fmt.Println("SOAL KE-3")
	fmt.Println(soal3.LuasPersegi(4, true))

	fmt.Println(soal3.LuasPersegi(8, false))

	fmt.Println(soal3.LuasPersegi(0, true))

	fmt.Println(soal3.LuasPersegi(0, false))

	// // soal 4
	fmt.Println("SOAL KE-4")

	var jumlah int

	kumpulanAngka := soal4.KumpulanAngkaPertama.([]int)

	kalimat := soal4.Prefix.(string)

	angkaKedua := soal4.KumpulanAngkaKedua.([]int)
	kumpulanAngka = append(kumpulanAngka, angkaKedua...)

	for index, item := range kumpulanAngka {
		jumlah += item
		if index == len(kumpulanAngka)-1 {
			kalimat += strconv.Itoa(item) + "=" + strconv.Itoa(jumlah)
		} else {
			kalimat += strconv.Itoa(item) + "+"
		}
	}

	fmt.Println(kalimat)
}
