package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ==========================================
	// DEKLARASI VARIABEL DI GOLANG
	// ==========================================

	// 1. Deklarasi Eksplisit (var namaTipe tipeData)
	var nama string = "Rizdar"

	// 2. Deklarasi dengan Type Inference (tipe data ditentukan otomatis oleh compiler)
	var umur = 20 // Compiler mendeteksi tipe data integer (int)

	// 3. Deklarasi Singkat (Short Declaration) menggunakan operator ':='
	// Hanya bisa digunakan di dalam body function (fungsi main, dll)
	isBelajar := true // Compiler mendeteksi tipe data boolean (bool)

	// 4. Tipe data desimal (float64)
	var ipk float64 = 3.85

	// 5. Konstanta (Constant) - Nilai yang tidak dapat diubah setelah dideklarasikan
	const versiGo = "1.22"

	// Menampilkan data tersebut di halaman web
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<title>Belajar Variabel & Tipe Data Go</title>
				<style>
					body {
						font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
						background-color: #0f172a;
						color: #e2e8f0;
						display: flex;
						justify-content: center;
						align-items: center;
						min-height: 100vh;
						margin: 0;
					}
					.card {
						background: #1e293b;
						padding: 30px;
						border-radius: 16px;
						box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.3);
						max-width: 450px;
						width: 100%%;
						border: 1px solid #334155;
					}
					h1 {
						color: #38bdf8;
						font-size: 24px;
						margin-top: 0;
						text-align: center;
						border-bottom: 2px solid #334155;
						padding-bottom: 15px;
					}
					ul {
						list-style: none;
						padding: 0;
						margin: 20px 0;
					}
					li {
						padding: 12px 0;
						border-bottom: 1px solid #334155;
						display: flex;
						justify-content: space-between;
						align-items: center;
					}
					li:last-child {
						border-bottom: none;
					}
					.label {
						font-weight: 600;
						color: #94a3b8;
					}
					.value {
						font-family: 'Courier New', Courier, monospace;
						color: #34d399;
						font-weight: bold;
						background: #0f172a;
						padding: 4px 8px;
						border-radius: 6px;
					}
					.footer {
						text-align: center;
						font-size: 12px;
						color: #64748b;
						margin-top: 20px;
					}
				</style>
			</head>
			<body>
				<div class="card">
					<h1>Variabel & Tipe Data Go</h1>
					<ul>
						<li>
							<span class="label">Nama (string)</span>
							<span class="value">%s</span>
						</li>
						<li>
							<span class="label">Umur (int)</span>
							<span class="value">%d tahun</span>
						</li>
						<li>
							<span class="label">Belajar Go (bool)</span>
							<span class="value">%t</span>
						</li>
						<li>
							<span class="label">IPK (float64)</span>
							<span class="value">%.2f</span>
						</li>
						<li>
							<span class="label">Versi Go (const)</span>
							<span class="value">%s</span>
						</li>
					</ul>
					<div class="footer">Dijalankan di dalam Container Docker 🐳</div>
				</div>
			</body>
			</html>
		`, nama, umur, isBelajar, ipk, versiGo)
	})

	fmt.Println("Server Go berjalan di port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

