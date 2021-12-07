package main

import "fmt"

// function soal ke-1
func luasPersegiPanjang(P, L int) int {
	return P * L
}

func kelilingPersegiPanjang(P, L int) int {
	return 2 * (P + L)
}

func volumeBalok(P, L, T int) int {
	return P * L * T
}

//function soal ke-2
func introduce(name, gender, job, age string) (sentence string) {
	if gender == "laki-laki" {
		sentence = "Pak " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	} else if gender == "perempuan" {
		sentence = "Bu " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	}
	return
}

//function soal ke-3
func buahFavorit(nama string, buah ...string) string {
	sentence := "halo nama saya john dan buah favorit saya adalah"
	for index, item := range buah {
		if index == 0 {
			sentence += ` "` + item + `"`
		} else {
			sentence += `, "` + item + `"`
		}
	}

	return sentence
}

func main() {
	// soal ke-1
	fmt.Println("SOAL KE- 1")
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)
	//Soal ke-2
	fmt.Println("SOAL KE- 2")
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
	// Soal ke -3
	fmt.Println("SOAL KE- 3")
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	// Soal ke -4
	fmt.Println("SOAL KE- 4")
	var dataFilm = []map[string]string{}

	//Clusure nomor-4
	tambahDataFilm := func(title, jam, genre, tahun string) {
		item := map[string]string{
			"genre": genre,
			"jam":   jam,
			"tahun": tahun,
			"title": title,
		}
		dataFilm = append(dataFilm, item)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}
