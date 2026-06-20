package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// handleConversion menampilkan halaman materi konversi tipe data
func handleConversion(w http.ResponseWriter, r *http.Request) {
	var nilaiInt int = 42
	var nilaiFloat float64 = float64(nilaiInt)

	var desimal float64 = 9.99
	var bulat int = int(desimal)

	var angkaKecil int8 = 100
	var angkaBesar int64 = int64(angkaKecil)

	var umur int = 25
	var umurString string = strconv.Itoa(umur)

	var stringAngka string = "100"
	angkaHasil, err := strconv.Atoi(stringAngka)
	var statusKonversi string
	if err != nil {
		statusKonversi = "Gagal"
	} else {
		statusKonversi = fmt.Sprintf("Sukses! Angka: %d", angkaHasil)
	}

	var stringDesimal string = "3.1415"
	floatHasil, _ := strconv.ParseFloat(stringDesimal, 64)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="id">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Konversi Tipe Data Go</title>
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
				}
				.container {
					max-width: 800px;
					width: 100%%;
				}
				h1 {
					color: #38bdf8;
					text-align: center;
				}
				.grid {
					display: grid;
					grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
					gap: 20px;
					margin-top: 20px;
				}
				.card {
					background: #1e293b;
					border-radius: 12px;
					padding: 20px;
					border: 1px solid #334155;
				}
				.card-title {
					color: #38bdf8;
					border-bottom: 1px solid #334155;
					padding-bottom: 8px;
					margin-top: 0;
				}
				table {
					width: 100%%;
				}
				td {
					padding: 6px 0;
				}
				.label {
					color: #94a3b8;
				}
				.value {
					text-align: right;
					font-family: monospace;
					color: #34d399;
					font-weight: bold;
					background: #0f172a;
					padding: 2px 6px;
					border-radius: 4px;
				}
				.back-link {
					display: inline-block;
					color: #38bdf8;
					text-decoration: none;
					font-weight: bold;
					margin-bottom: 20px;
				}
				.back-link:hover {
					text-decoration: underline;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<a href="/" class="back-link">⬅ Kembali ke Menu Utama</a>
				<h1>Konversi Tipe Data di Go</h1>
				<div class="grid">
					<div class="card">
						<h3 class="card-title">1. Integer &lt;-&gt; Float</h3>
						<table>
							<tr><td class="label">int ke float64 (float64(nilaiInt))</td><td class="value">%.2f</td></tr>
							<tr><td class="label">float64 ke int (int(desimal))</td><td class="value">%d</td></tr>
						</table>
					</div>
					<div class="card">
						<h3 class="card-title">2. Ukuran Integer</h3>
						<table>
							<tr><td class="label">int8 (angkaKecil)</td><td class="value">%d</td></tr>
							<tr><td class="label">int64 (angkaBesar)</td><td class="value">%d</td></tr>
						</table>
					</div>
					<div class="card">
						<h3 class="card-title">3. Integer ke String</h3>
						<table>
							<tr><td class="label">int (umur = 25)</td><td class="value">%d</td></tr>
							<tr><td class="label">string (strconv.Itoa(umur))</td><td class="value">"%s"</td></tr>
						</table>
					</div>
					<div class="card">
						<h3 class="card-title">4. String ke Angka</h3>
						<table>
							<tr><td class="label">String ("100") -> int</td><td class="value">%s</td></tr>
							<tr><td class="label">String ("3.1415") -> float</td><td class="value">%.4f</td></tr>
						</table>
					</div>
				</div>
			</div>
		</body>
		</html>
	`, nilaiFloat, bulat,
		angkaKecil, angkaBesar,
		umur, umurString,
		statusKonversi, floatHasil)
}
