package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"sync"
	"time"
)

func showPhones(wg *sync.WaitGroup, phones ...string) {
	for index, item := range phones {
		fmt.Print(index + 1)
		fmt.Println(".", item)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func getMovies(movieCh chan string, movies ...string) {
	fmt.Println("List Movies:")
	for index, item := range movies {
		number := index + 1
		movieCh <- strconv.Itoa(number) + ". " + item
	}
	close(movieCh)
}

func luasLingkaran(ch chan map[string]float64, radiusList ...float64) {
	for _, item := range radiusList {
		luas := math.Round(math.Pi * item * item)
		ch <- map[string]float64{"jari-jari": item, "luas": luas}
	}
	close(ch)
}

func kelilingLingkaran(ch chan map[string]float64, radiusList ...float64) {
	for _, item := range radiusList {
		keliling := math.Round(2 * math.Pi * item)
		ch <- map[string]float64{"jari-jari": item, "keliling": keliling}
	}
	close(ch)
}

func volumeTabung(ch chan map[string]float64, t float64, radiusList ...float64) {
	for _, item := range radiusList {
		volume := math.Round(math.Pi * item * item * t)
		ch <- map[string]float64{"jari-jari": item, "tinggi": t, "volume": volume}
	}
	close(ch)
}

func luasPersegi(satuan map[string]int, ch chan map[string]int) {
	luas := satuan["panjang"] * satuan["lebar"]
	hasil := map[string]int{
		"panjang": satuan["panjang"],
		"lebar":   satuan["lebar"],
		"luas":    luas,
	}
	ch <- hasil
}

func kelilingPersegi(satuan map[string]int, ch chan map[string]int) {
	keliling := 2 * (satuan["panjang"] + satuan["lebar"])
	hasil := map[string]int{
		"panjang":  satuan["panjang"],
		"lebar":    satuan["lebar"],
		"keliling": keliling,
	}
	ch <- hasil
}

func volumeBalok(satuan map[string]int, ch chan map[string]int) {
	volume := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	hasil := map[string]int{
		"panjang": satuan["panjang"],
		"lebar":   satuan["lebar"],
		"tinggi":  satuan["tinggi"],
		"volume":  volume,
	}
	ch <- hasil
}

func main() {
	//soal 1
	fmt.Println("SOAL ke-1")
	var wg sync.WaitGroup
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	sort.Strings(phones)

	wg.Add(1)
	go showPhones(&wg, phones...)
	wg.Wait()

	//soal 2
	fmt.Println("SOAL Ke-2")

	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}

	//soal 3
	fmt.Println("SOAL KE-3")

	var jariJari = []float64{8, 14, 20}

	luasLingkaranChannel := make(chan map[string]float64)

	go luasLingkaran(luasLingkaranChannel, jariJari...)

	for value := range luasLingkaranChannel {
		fmt.Println(value)
	}

	kelilingLingkaranChannel := make(chan map[string]float64)

	go kelilingLingkaran(kelilingLingkaranChannel, jariJari...)

	for value := range kelilingLingkaranChannel {
		fmt.Println(value)
	}

	volumeTabungChannel := make(chan map[string]float64)

	go volumeTabung(volumeTabungChannel, 10, jariJari...)

	for value := range volumeTabungChannel {
		fmt.Println(value)
	}

	//soal 4
	fmt.Println("SOAL KE-4")

	satuan := map[string]int{
		"panjang": 10,
		"lebar":   7,
		"tinggi":  6,
	}

	luasPersegiChannel := make(chan map[string]int)
	kelilingPersegiChannel := make(chan map[string]int)
	volumeBalokChannel := make(chan map[string]int)

	go luasPersegi(satuan, luasPersegiChannel)
	go kelilingPersegi(satuan, kelilingPersegiChannel)
	go volumeBalok(satuan, volumeBalokChannel)

	for i := 1; i <= 3; i++ {
		select {
		case luas := <-luasPersegiChannel:
			fmt.Println(luas)
		case keliling := <-kelilingPersegiChannel:
			fmt.Println(keliling)
		case volume := <-volumeBalokChannel:
			fmt.Println(volume)
		}
	}

}
