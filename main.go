package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ==========================================
	// TIPE DATA PRIMITIF DI GOLANG
	// ==========================================

	// 1. Boolean (bool)
	var booleanBenar bool = true
	var booleanSalah bool = false

	// 2. Integer (Bilangan Bulat)
	// Signed Integer (Bisa bernilai negatif)
	var angkaBulatBiasa int = -42 // Mengikuti arsitektur OS (32-bit atau 64-bit)
	var angkaKecil int8 = 127     // Rentang: -128 s/d 127
	var angkaBesar int64 = 9223372036854775807

	// Unsigned Integer (Hanya positif dan nol)
	var angkaPositif uint = 150
	var angkaPositifKecil uint8 = 255 // Rentang: 0 s/d 255 (Sering disebut 'byte')

	// 3. Floating Point (Bilangan Desimal)
	var desimalBiasa float32 = 3.14
	var desimalPresisi float64 = 3.141592653589793

	// 4. Character / Unicode Code Point
	var karakterByte byte = 'A'  // byte adalah alias dari uint8, nilai ASCII 'A' adalah 65
	var karakterRune rune = '🍎' // rune adalah alias dari int32, digunakan untuk menyimpan karakter Unicode (UTF-8)

	// 5. String (Teks)
	var teks string = "Belajar Golang itu Seru!"

	// Menampilkan data tersebut di halaman web dengan styling modern
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Belajar Tipe Data Primitif Go</title>
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
						width: 100%%;
					}
					h1 {
						color: #38bdf8;
						text-align: center;
						margin-bottom: 30px;
						font-size: 2.5rem;
					}
					.grid {
						display: grid;
						grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
						gap: 20px;
					}
					.card {
						background: #1e293b;
						border-radius: 12px;
						padding: 24px;
						border: 1px solid #334155;
						box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
					}
					.card-title {
						font-size: 1.25rem;
						color: #38bdf8;
						margin-top: 0;
						margin-bottom: 15px;
						border-bottom: 1px solid #334155;
						padding-bottom: 10px;
					}
					table {
						width: 100%%;
						border-collapse: collapse;
					}
					td {
						padding: 8px 0;
					}
					.label {
						color: #94a3b8;
						font-weight: 500;
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
					.desc {
						font-size: 0.85rem;
						color: #64748b;
						margin-top: 15px;
						line-height: 1.4;
					}
					.footer {
						text-align: center;
						margin-top: 40px;
						color: #64748b;
						font-size: 0.9rem;
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h1>Tipe Data Primitif di Go</h1>
					
					<div class="grid">
						<!-- Boolean Card -->
						<div class="card">
							<h2 class="card-title">1. Boolean (bool)</h2>
							<table>
								<tr>
									<td class="label">booleanBenar</td>
									<td class="value">%t</td>
								</tr>
								<tr>
									<td class="label">booleanSalah</td>
									<td class="value">%t</td>
								</tr>
							</table>
							<p class="desc">Hanya menampung nilai kebenaran: true (benar) atau false (salah).</p>
						</div>

						<!-- Integer Card -->
						<div class="card">
							<h2 class="card-title">2. Integer (Bilangan Bulat)</h2>
							<table>
								<tr>
									<td class="label">int (angkaBulatBiasa)</td>
									<td class="value">%d</td>
								</tr>
								<tr>
									<td class="label">int8 (angkaKecil)</td>
									<td class="value">%d</td>
								</tr>
								<tr>
									<td class="label">int64 (angkaBesar)</td>
									<td class="value">%d</td>
								</tr>
								<tr>
									<td class="label">uint (angkaPositif)</td>
									<td class="value">%d</td>
								</tr>
								<tr>
									<td class="label">uint8 (angkaPositifKecil)</td>
									<td class="value">%d</td>
								</tr>
							</table>
							<p class="desc"><strong>int/uint</strong> otomatis menyesuaikan dengan OS (32/64 bit). Versi spesifik (int8 s/d int64) membatasi ukuran penyimpanan memori.</p>
						</div>

						<!-- Float Card -->
						<div class="card">
							<h2 class="card-title">3. Floating Point (Desimal)</h2>
							<table>
								<tr>
									<td class="label">float32 (desimalBiasa)</td>
									<td class="value">%.2f</td>
								</tr>
								<tr>
									<td class="label">float64 (desimalPresisi)</td>
									<td class="value">%.15f</td>
								</tr>
							</table>
							<p class="desc">Digunakan untuk bilangan berkoma. float64 memiliki tingkat presisi angka di belakang koma yang jauh lebih tinggi daripada float32.</p>
						</div>

						<!-- Byte & Rune Card -->
						<div class="card">
							<h2 class="card-title">4. Byte & Rune (Karakter)</h2>
							<table>
								<tr>
									<td class="label">byte (karakterByte - 'A')</td>
									<td class="value">ASCII: %d (Karakter: %c)</td>
								</tr>
								<tr>
									<td class="label">rune (karakterRune - '🍎')</td>
									<td class="value">Unicode: %d (Karakter: %c)</td>
								</tr>
							</table>
							<p class="desc"><strong>byte</strong> adalah alias dari uint8 (menyimpan ASCII). <strong>rune</strong> adalah alias dari int32 (menyimpan karakter Unicode seperti emoji).</p>
						</div>

						<!-- String Card -->
						<div class="card" style="grid-column: 1 / -1;">
							<h2 class="card-title">5. String (Teks)</h2>
							<table>
								<tr>
									<td class="label">teks (string)</td>
									<td class="value">"%s"</td>
								</tr>
							</table>
							<p class="desc">Menyimpan sekumpulan karakter Unicode UTF-8 secara read-only. Dibuat dengan tanda kutip ganda.</p>
						</div>
					</div>

					<div class="footer">
						Dijalankan di dalam Container Docker 🐳 | Selamat Belajar Golang!
					</div>
				</div>
			</body>
			</html>
		`, booleanBenar, booleanSalah,
			angkaBulatBiasa, angkaKecil, angkaBesar, angkaPositif, angkaPositifKecil,
			desimalBiasa, desimalPresisi,
			karakterByte, karakterByte, karakterRune, karakterRune,
			teks)
	})

	fmt.Println("Server Go berjalan di port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
