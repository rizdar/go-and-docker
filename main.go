package main

import (
	"fmt"
	"net/http"
)

// handleHome menampilkan menu dashboard utama
func handleHome(w http.ResponseWriter, r *http.Request) {
	// Menghindari pencocokan wildcard (contoh: r.URL.Path yang tidak valid dialihkan ke 404)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html lang="id">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Dashboard Belajar Go & Docker</title>
			<style>
				body {
					font-family: 'Segoe UI', system-ui, -apple-system, sans-serif;
					background-color: #0f172a;
					color: #e2e8f0;
					margin: 0;
					padding: 40px 20px;
					display: flex;
					flex-direction: column;
					align-items: center;
					min-height: 100vh;
				}
				.container {
					max-width: 800px;
					width: 100%;
				}
				h1 {
					color: #38bdf8;
					text-align: center;
					margin-bottom: 10px;
					font-size: 2.5rem;
				}
				p.subtitle {
					color: #94a3b8;
					text-align: center;
					margin-bottom: 40px;
					font-size: 1.1rem;
				}
				.grid {
					display: grid;
					grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
					gap: 20px;
				}
				.card {
					background: #1e293b;
					border-radius: 16px;
					padding: 24px;
					border: 1px solid #334155;
					box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
					transition: transform 0.2s, border-color 0.2s;
					text-decoration: none;
					color: inherit;
					display: flex;
					flex-direction: column;
					justify-content: space-between;
				}
				.card:hover {
					transform: translateY(-4px);
					border-color: #38bdf8;
				}
				.card-title {
					font-size: 1.25rem;
					color: #38bdf8;
					margin-top: 0;
					margin-bottom: 10px;
					font-weight: 600;
				}
				.card-desc {
					font-size: 0.9rem;
					color: #94a3b8;
					line-height: 1.5;
					margin-bottom: 20px;
				}
				.card-action {
					color: #38bdf8;
					font-weight: bold;
					font-size: 0.95rem;
					display: flex;
					align-items: center;
				}
				.footer {
					text-align: center;
					margin-top: 60px;
					color: #64748b;
					font-size: 0.9rem;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1>Go & Docker Learning Hub 🐳</h1>
				<p class="subtitle">Pilih modul di bawah ini untuk melihat contoh kode dan demonstrasi interaktif.</p>
				
				<div class="grid">
					<!-- Modul 1 -->
					<a href="/variables" class="card">
						<div>
							<h2 class="card-title">1. Variabel & Konstan</h2>
							<p class="card-desc">Mempelajari deklarasi eksplisit, type inference, short declaration, serta konstanta di Go.</p>
						</div>
						<div class="card-action">Buka Modul ➡️</div>
					</a>

					<!-- Modul 2 -->
					<a href="/primitives" class="card">
						<div>
							<h2 class="card-title">2. Tipe Data Primitif</h2>
							<p class="card-desc">Mengenal tipe boolean, integer, unsigned integer, desimal (float), dan karakter (byte & rune).</p>
						</div>
						<div class="card-action">Buka Modul ➡️</div>
					</a>

					<!-- Modul 3 -->
					<a href="/conversion" class="card">
						<div>
							<h2 class="card-title">3. Konversi Tipe Data</h2>
							<p class="card-desc">Mempelajari cara konversi tipe data angka secara eksplisit dan menggunakan package strconv.</p>
						</div>
						<div class="card-action">Buka Modul ➡️</div>
					</a>

					<!-- Modul 4 -->
					<a href="/control-flow" class="card">
						<div>
							<h2 class="card-title">4. Control Flow</h2>
							<p class="card-desc">Mempelajari alur kontrol program menggunakan percabangan if-else, switch, dan loop tunggal for.</p>
						</div>
						<div class="card-action">Buka Modul ➡️</div>
					</a>

					<!-- Modul 5 -->
					<a href="/functions" class="card">
						<div>
							<h2 class="card-title">5. Function (Fungsi)</h2>
							<p class="card-desc">Membuat fungsi dengan parameter, multiple return values, named returns, serta variadic parameters.</p>
						</div>
						<div class="card-action">Buka Modul ➡️</div>
					</a>
				</div>

				<div class="footer">
					Dijalankan di dalam Container Docker 🐳 | Belajar Go-Docker Bersama
				</div>
			</div>
		</body>
		</html>
	`)
}

func main() {
	// Mendaftarkan handler untuk masing-masing modul
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/variables", handleVariables)
	http.HandleFunc("/primitives", handlePrimitives)
	http.HandleFunc("/conversion", handleConversion)
	http.HandleFunc("/control-flow", handleControlFlow)
	http.HandleFunc("/functions", handleFunctions)

	fmt.Println("Server Go berjalan di port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
