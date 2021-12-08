package main

import "fmt"

func hitungLingkaran(jariJari int, luas *float64, keliling *float64) {
	if jariJari%7 == 0 {
		*luas = float64(22 / 7 * jariJari * jariJari)
		*keliling = float64(2 * 22 / 7 * jariJari)
	} else {
		*luas = float64(3.14 * float64(jariJari) * float64(jariJari))
		*keliling = float64(2 * 3.14 * float64(jariJari))
	}
}

func introduce(sentence *string, name, gender, job, age string) {
	prefix := "Pak"

	if gender == "perempuan" {
		prefix = "Bu"
	}

	*sentence = prefix + " " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
}

func tambahBuah(buah *[]string, namaBuah ...string) {
	for _, nama := range namaBuah {
		*buah = append(*buah, nama)
	}
}

func tambahDataFilm(title, duration, genre, year string, dataFilm *[]map[string]string) {
	data := map[string]string{
		"title":    title,
		"duration": duration,
		"genre":    genre,
		"year":     year,
	}
	*dataFilm = append(*dataFilm, data)
}

func main() {
	// soal 1
	fmt.Println("SOAL KE-1")
	var luasLigkaran float64
	var kelilingLingkaran float64

	hitungLingkaran(7, &luasLigkaran, &kelilingLingkaran)
	fmt.Println(luasLigkaran, kelilingLingkaran)

	hitungLingkaran(10, &luasLigkaran, &kelilingLingkaran)
	fmt.Println(luasLigkaran, kelilingLingkaran)

	// soal 2
	fmt.Println("SOAL KE-2")
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// soal 3
	fmt.Println("SOAL KE-3")
	var buah = []string{}
	tambahBuah(&buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")

	for index, item := range buah {
		fmt.Print(index + 1)
		fmt.Println(".", item)
	}

	//soal 4
	fmt.Println("SOAL KE-4")
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for index, item := range dataFilm {
		fmt.Print(index + 1)
		fmt.Print(". ")
		i := 1
		for key, value := range item {
			if i == 1 {
				fmt.Println(key, ":", value)
			} else {
				fmt.Println("  ", key, ":", value)
			}
			i++
		}
		fmt.Println()
	}

}
