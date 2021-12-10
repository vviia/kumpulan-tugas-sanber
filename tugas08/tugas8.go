package main

import (
	"fmt"
	"strconv"
)

type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type description interface {
	getDescription() string
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	if int(t.jariJari)%7 == 0 {
		return float64(22 / 7 * t.jariJari * t.jariJari * t.tinggi)
	} else {
		return float64(3.14 * t.jariJari * t.jariJari * t.tinggi)
	}
}

func (t tabung) luasPermukaan() float64 {
	if int(t.jariJari)%7 == 0 {
		return float64(2 * 22 / 7 * t.jariJari * (t.jariJari + t.tinggi))
	} else {
		return float64(2 * 3.14 * t.jariJari * (t.jariJari + t.tinggi))
	}
}

func (b balok) volume() float64 {
	return float64(b.panjang) * float64(b.lebar) * float64(b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	p := float64(b.panjang)
	l := float64(b.lebar)
	t := float64(b.tinggi)

	return float64((2 * (p + l)) + (2 * (p + t)) + (2 * (l + t)))
}

func (ph phone) getDescription() string {
	var desc = "name : " + ph.name + "\n" +
		"brand : " + ph.brand + "\n" +
		"year : " + strconv.Itoa(ph.year) + "\n" +
		"colors : "
	for index, item := range ph.colors {
		if index == 0 {
			desc += item
		} else {
			desc += ", " + item
		}
	}
	return desc
}

func luasPersegi(sisi uint, showText bool) interface{} {
	switch {
	case sisi > 0 && showText:
		return "luas persegi dengan sisi " + strconv.Itoa(int(sisi)) + " cm adalah " + strconv.Itoa(int(sisi*sisi)) + " cm"
	case sisi > 0 && !showText:
		return sisi * sisi
	case sisi == 0 && showText:
		return "Maaf anda belum menginput sisi dari persegi"
	default:
		return nil
	}
}

func main() {
	// soal 1
	fmt.Println("soal ke 1")
	var segitiga1 hitungBangunDatar = segitigaSamaSisi{alas: 7, tinggi: 8}
	var persegiPanjang1 hitungBangunDatar = persegiPanjang{panjang: 7, lebar: 8}

	fmt.Println("------bangun datar-------")
	fmt.Println("luas segitiga sama sisi dari alas 7 dan tinggi 8 adalah", segitiga1.luas())
	fmt.Println("keliling segitiga sama sisi dari alas 7 dan tinggi 8 adalah", segitiga1.keliling())

	fmt.Println("luas persegi panjang sama dari panjang 7 dan lebar 8 adalah", persegiPanjang1.luas())
	fmt.Println("keliling persegi panjang sama dari panjang 7 dan lebar adalah", persegiPanjang1.keliling())

	fmt.Println("------bangun ruang-------")

	var tabung1 hitungBangunRuang = tabung{jariJari: 7, tinggi: 10}
	var balok1 hitungBangunRuang = balok{panjang: 7, lebar: 5, tinggi: 6}

	fmt.Println("volume tabung dari jari-jari 7 dan tinggi 10 adalah", tabung1.volume())
	fmt.Println("luas permukaan tabung dari jari-jari 7 dan tinggi 10 adalah", tabung1.luasPermukaan())

	fmt.Println("volume balok dari panjang 7, lebar 5 dan tinggi 6 adalah", balok1.volume())
	fmt.Println("luas permukaan balok dari panjang 7, lebar 5 dan tinggi 6 adalah", balok1.luasPermukaan())

	// soal 2
	fmt.Println("soal ke 2")
	var samsung description = phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Println(samsung.getDescription())

	// soal 3
	fmt.Println("soal ke 3")
	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))

	// soal 4
	fmt.Println("soal ke 4")
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	var jumlah int

	kumpulanAngka := kumpulanAngkaPertama.([]int)

	kalimat := prefix.(string)

	angkaKedua := kumpulanAngkaKedua.([]int)
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
