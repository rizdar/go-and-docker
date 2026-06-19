package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Menghandle request ke halaman utama "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Halo! Kamu berhasil menjalankan Go di dalam Docker!")
	})

	fmt.Println("Server Go berjalan di port 8080...")
	// Menjalankan server pada port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
