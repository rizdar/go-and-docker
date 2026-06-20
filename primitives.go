package main

import (
	"fmt"
	"net/http"
)

// handlePrimitives menampilkan halaman materi tipe data primitif
func handlePrimitives(w http.ResponseWriter, r *http.Request) {
	var booleanBenar bool = true
	var booleanSalah bool = false
	var angkaBulatBiasa int = -42
	var angkaKecil int8 = 127
	var angkaBesar int64 = 9223372036854775807
	var angkaPositif uint = 150
	var angkaPositifKecil uint8 = 255
	var desimalBiasa float32 = 3.14
	var desimalPresisi float64 = 3.141592653589793
	var karakterByte byte = 'A'
	var karakterRune rune = '🍎'
	var teks string = "Belajar Golang itu Seru!"

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="id">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Tipe Data Primitif Go</title>
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
				<h1>Tipe Data Primitif di Go</h1>
				<div class="grid">
					<div class="card">
						<h3 class="card-title">1. Boolean (bool)</h3>
						<table>
							<tr><td class="label">booleanBenar</td><td class="value">%t</td></tr>
							<tr><td class="label">booleanSalah</td><td class="value">%t</td></tr>
						</table>
					</div>
					<div class="card">
						<h3 class="card-title">2. Integer (Bilangan Bulat)</h3>
						<table>
							<tr><td class="label">int (angkaBulatBiasa)</td><td class="value">%d</td></tr>
							<tr><td class="label">int8 (angkaKecil)</td><td class="value">%d</td></tr>
							<tr><td class="label">int64 (angkaBesar)</td><td class="value">%d</td></tr>
							<tr><td class="label">uint (angkaPositif)</td><td class="value">%d</td></tr>
							<tr><td class="label">uint8 (angkaPositifKecil)</td><td class="value">%d</td></tr>
						</table>
					</div>
					<div class="card">
						<h3 class="card-title">3. Float (Desimal)</h3>
						<table>
							<tr><td class="label">float32 (desimalBiasa)</td><td class="value">%.2f</td></tr>
							<tr><td class="label">float64 (desimalPresisi)</td><td class="value">%.15f</td></tr>
						</table>
					</div>
					<div class="card">
						<h3 class="card-title">4. Byte & Rune</h3>
						<table>
							<tr><td class="label">byte (ASCII 'A')</td><td class="value">ASCII: %d (Char: %c)</td></tr>
							<tr><td class="label">rune (Unicode '🍎')</td><td class="value">Unicode: %d (Char: %c)</td></tr>
						</table>
					</div>
					<div class="card" style="grid-column: 1 / -1;">
						<h3 class="card-title">5. String (Teks)</h3>
						<table>
							<tr><td class="label">teks (string)</td><td class="value">"%s"</td></tr>
						</table>
					</div>
				</div>
			</div>
		</body>
		</html>
	`, booleanBenar, booleanSalah,
		angkaBulatBiasa, angkaKecil, angkaBesar, angkaPositif, angkaPositifKecil,
		desimalBiasa, desimalPresisi,
		karakterByte, karakterByte, karakterRune, karakterRune,
		teks)
}
