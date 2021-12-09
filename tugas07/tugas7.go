package main

import "fmt"

type buah struct {
	nama, warna string
	adaBijinya  bool
	harga       int
}

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type movie struct {
	title, genre   string
	duration, year int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) luas() int {
	return s.alas * s.tinggi / 2
}

func (p persegi) luas() int {
	return p.sisi * p.sisi
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (ph *phone) addColors(colors ...string) {
	for _, color := range colors {
		ph.colors = append(ph.colors, color)
	}
}

func tambahDataFilm(title string,
	duration int,
	genre string,
	year int,
	dataFilm *[]movie) {
	*dataFilm = append(*dataFilm, movie{title, genre, duration, year})
}

func main() {
	// soal 1
	fmt.Println("soal ke-1")

	daftarBuah := []buah{
		{nama: "Nanas", warna: "Kuning", adaBijinya: false, harga: 9000},
		{nama: "Jeruk", warna: "Oranye", adaBijinya: true, harga: 8000},
		{nama: "Semangka", warna: "Hijau & Merah", adaBijinya: true, harga: 10000},
		{nama: "Pisang", warna: "Kuning", adaBijinya: false, harga: 5000},
	}

	for _, item := range daftarBuah {
		fmt.Println(item)
	}

	// soal 2
	fmt.Println("soal ke-2")
	segitiga1 := segitiga{alas: 10, tinggi: 7}
	persegi1 := persegi{sisi: 9}
	persegiPanjang1 := persegiPanjang{panjang: 8, lebar: 6}

	fmt.Println(segitiga1.luas())
	fmt.Println(persegi1.luas())
	fmt.Println(persegiPanjang1.luas())

	// soal 3
	fmt.Println("soal ke-3")
	samsung := phone{name: "Samsung Galaxy Note 20", brand: "Samsung", year: 2020}

	fmt.Println(samsung)
	samsung.addColors("Black", "Bronze", "Silver")
	fmt.Println(samsung)

	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for index, item := range dataFilm {
		fmt.Print(index + 1)
		fmt.Println(". title :", item.title)
		fmt.Println("   duration :", item.duration/60, "jam")
		fmt.Println("   genre :", item.genre)
		fmt.Println("   year :", item.year)
		fmt.Println()
	}

}
