package main

import (
	"fmt"
	"log"
	"net/http"
	"tugas14/functions"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/nilai-mahasiswa", functions.GetAllNilai)
	router.POST("/nilai-mahasiswa/create", functions.PostNilai)
	router.PUT("/nilai-mahasiswa/:id/update", functions.UpdateNilai)
	router.DELETE("/nilai-mahasiswa/:id/delete", functions.DeleteNilai)

	fmt.Println("Server Running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
