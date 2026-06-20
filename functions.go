package main

import (
	"errors"
	"fmt"
	"net/http"
)

// 1. Fungsi biasa
func tambah(a, b int) int {
	return a + b
}

// 2. Multiple return values
func hitungLuasDanKeliling(panjang, lebar float64) (float64, float64) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}

// 3. Named return values & error return
func bagi(pembilang, penyebut float64) (hasil float64, err error) {
	if penyebut == 0 {
		err = errors.New("tidak bisa membagi dengan nol")
		return
	}
	hasil = pembilang / penyebut
	return
}

// 4. Variadic function
func jumlahkanSemua(angka ...int) int {
	total := 0
	for _, nilai := range angka {
		total += nilai
	}
	return total
}

// handleFunctions menampilkan halaman materi functions
func handleFunctions(w http.ResponseWriter, r *http.Request) {
	hasilTambah := tambah(15, 27)
	luasPersegi, kelilingPersegi := hitungLuasDanKeliling(10.0, 5.0)

	hasilBagiNormal, errNormal := bagi(10, 2)
	var teksBagiNormal string
	if errNormal != nil {
		teksBagiNormal = errNormal.Error()
	} else {
		teksBagiNormal = fmt.Sprintf("%.2f", hasilBagiNormal)
	}

	_, errNol := bagi(10, 0)
	var teksBagiNol string
	if errNol != nil {
		teksBagiNol = fmt.Sprintf("Error: %s", errNol.Error())
	} else {
		teksBagiNol = "Sukses"
	}

	totalVariadic := jumlahkanSemua(10, 20, 30, 40, 50)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="id">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Belajar Function Go</title>
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
				<h1>Fungsi (Function) di Go</h1>
				<div class="grid">
					<div class="card">
						<h3 class="card-title">1. Fungsi Biasa</h3>
						<table>
							<tr><td class="label">tambah(15, 27)</td><td class="value">%d</td></tr>
						</table>
					</div>
					<div class="card">
						<h3 class="card-title">2. Multiple Returns</h3>
						<table>
							<tr><td class="label">Luas (10x5)</td><td class="value">%.1f</td></tr>
							<tr><td class="label">Keliling (10x5)</td><td class="value">%.1f</td></tr>
						</table>
					</div>
					<div class="card">
						<h3 class="card-title">3. Named & Error Returns</h3>
						<table>
							<tr><td class="label">bagi(10, 2)</td><td class="value">%s</td></tr>
							<tr><td class="label">bagi(10, 0)</td><td class="value" style="color: #f43f5e;">%s</td></tr>
						</table>
					</div>
					<div class="card">
						<h3 class="card-title">4. Variadic Function</h3>
						<table>
							<tr><td class="label">jumlahkanSemua(...)</td><td class="value">%d</td></tr>
						</table>
					</div>
				</div>
			</div>
		</body>
		</html>
	`, hasilTambah,
		luasPersegi, kelilingPersegi,
		teksBagiNormal, teksBagiNol,
		totalVariadic)
}
