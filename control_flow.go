package main

import (
	"fmt"
	"net/http"
)

// handleControlFlow menampilkan halaman materi control flow
func handleControlFlow(w http.ResponseWriter, r *http.Request) {
	var nilai int = 85
	var grade string
	if nilai >= 85 {
		grade = "A"
	} else if nilai >= 75 {
		grade = "B"
	} else if nilai >= 60 {
		grade = "C"
	} else {
		grade = "D"
	}

	var statusUjian string
	if kriteria := 75; nilai >= kriteria {
		statusUjian = "LULUS (Nilai di atas kriteria)"
	} else {
		statusUjian = "TIDAK LULUS"
	}

	var hari int = 3
	var namaHari string
	switch hari {
	case 1:
		namaHari = "Senin"
	case 2:
		namaHari = "Selasa"
	case 3:
		namaHari = "Rabu"
	case 4:
		namaHari = "Kamis"
	case 5:
		namaHari = "Jumat"
	default:
		namaHari = "Akhir Pekan"
	}

	var hasilLoopBiasa string
	for i := 1; i <= 3; i++ {
		hasilLoopBiasa += fmt.Sprintf("Perulangan ke-%d, ", i)
	}

	var counter = 1
	var hasilLoopWhile string
	for counter <= 3 {
		hasilLoopWhile += fmt.Sprintf("Iterasi %d; ", counter)
		counter++
	}

	frameworks := []string{"Go", "Docker", "Git"}
	var hasilLoopRange string
	for indeks, item := range frameworks {
		hasilLoopRange += fmt.Sprintf("[%d: %s] ", indeks, item)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="id">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Control Flow Go</title>
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
				<h1>Control Flow di Go</h1>
				<div class="grid">
					<div class="card">
						<h3 class="card-title">1. If - Else If - Else</h3>
						<table>
							<tr><td class="label">Nilai</td><td class="value">%d</td></tr>
							<tr><td class="label">Grade</td><td class="value">%s</td></tr>
							<tr><td class="label">Status (Temp Var)</td><td class="value">%s</td></tr>
						</table>
					</div>
					<div class="card">
						<h3 class="card-title">2. Switch Case</h3>
						<table>
							<tr><td class="label">Hari ke</td><td class="value">%d</td></tr>
							<tr><td class="label">Nama Hari</td><td class="value">%s</td></tr>
						</table>
					</div>
					<div class="card" style="grid-column: 1 / -1;">
						<h3 class="card-title">3. For Loops</h3>
						<table>
							<tr><td class="label">Standard Loop (for i)</td><td class="value">%s</td></tr>
							<tr><td class="label">While-style Loop</td><td class="value">%s</td></tr>
							<tr><td class="label">Range Loop (slice)</td><td class="value">%s</td></tr>
						</table>
					</div>
				</div>
			</div>
		</body>
		</html>
	`, nilai, grade, statusUjian,
		hari, namaHari,
		hasilLoopBiasa, hasilLoopWhile, hasilLoopRange)
}
