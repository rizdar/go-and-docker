package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	// ==========================================
	// KONVERSI TIPE DATA (TYPE CONVERSION) DI GO
	// ==========================================

	// 1. Konversi Antar Angka (Integer <-> Float)
	var nilaiInt int = 42
	var nilaiFloat float64 = float64(nilaiInt) // Mengonversi int ke float64

	var desimal float64 = 9.99
	var bulat int = int(desimal) // Mengonversi float64 ke int (nilai desimal akan dibuang menjadi 9)

	// 2. Konversi Ukuran Tipe Data
	var angkaKecil int8 = 100
	var angkaBesar int64 = int64(angkaKecil)

	// 3. Konversi Integer ke String menggunakan strconv.Itoa()
	var umur int = 25
	// PENTING: string(umur) tidak akan menghasilkan teks "25", melainkan karakter ASCII dengan kode 25.
	// Untuk mendapatkan string "25", kita harus menggunakan package 'strconv'
	var umurString string = strconv.Itoa(umur)

	// 4. Konversi String ke Integer menggunakan strconv.Atoi()
	var stringAngka string = "100"
	// strconv.Atoi mengembalikan 2 nilai: (hasil_angka, error)
	angkaHasil, err := strconv.Atoi(stringAngka)
	var statusKonversi string
	if err != nil {
		statusKonversi = "Gagal mengonversi string ke int"
	} else {
		statusKonversi = fmt.Sprintf("Sukses! Angka: %d", angkaHasil)
	}

	// 5. Konversi String ke Float64 menggunakan strconv.ParseFloat()
	var stringDesimal string = "3.1415"
	floatHasil, _ := strconv.ParseFloat(stringDesimal, 64) // Mengabaikan error menggunakan '_' untuk demonstrasi singkat

	// Menampilkan data tersebut di halaman web dengan styling modern
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Belajar Konversi Tipe Data Go</title>
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
						margin-bottom: 10px;
						font-size: 2.5rem;
					}
					p.subtitle {
						color: #94a3b8;
						text-align: center;
						margin-bottom: 30px;
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
					.code {
						font-family: monospace;
						color: #f43f5e;
						background: #0f172a;
						padding: 2px 4px;
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
					<h1>Konversi Tipe Data di Go</h1>
					<p class="subtitle">Go tidak mendukung konversi implisit. Semua tipe data harus dikonversi secara eksplisit.</p>
					
					<div class="grid">
						<!-- Numeric Conversion -->
						<div class="card">
							<h2 class="card-title">1. Integer &lt;-&gt; Float</h2>
							<table>
								<tr>
									<td class="label">int ke float64 (float64(nilaiInt))</td>
									<td class="value">%.2f</td>
								</tr>
								<tr>
									<td class="label">float64 ke int (int(desimal))</td>
									<td class="value">%d</td>
								</tr>
							</table>
							<p class="desc">Saat mengonversi float ke int, bagian desimal di belakang koma akan **dibuang** (di-truncate), bukan dibulatkan ke atas.</p>
						</div>

						<!-- Size Conversion -->
						<div class="card">
							<h2 class="card-title">2. Konversi Ukuran Integer</h2>
							<table>
								<tr>
									<td class="label">int8 (angkaKecil)</td>
									<td class="value">%d</td>
								</tr>
								<tr>
									<td class="label">int64 (angkaBesar)</td>
									<td class="value">%d</td>
								</tr>
							</table>
							<p class="desc">Mengonversi tipe yang lebih kecil ke besar aman dilakukan. Tapi hati-hati bila sebaliknya, jika nilainya melebihi batas tipe yang lebih kecil akan terjadi *overflow*.</p>
						</div>

						<!-- Integer to String (strconv.Itoa) -->
						<div class="card">
							<h2 class="card-title">3. Integer ke String</h2>
							<table>
								<tr>
									<td class="label">int (umur = 25)</td>
									<td class="value">%d</td>
								</tr>
								<tr>
									<td class="label">string (strconv.Itoa(umur))</td>
									<td class="value">"%s"</td>
								</tr>
							</table>
							<p class="desc">Wajib menggunakan <span class="code">strconv.Itoa()</span> (Integer to ASCII). Jika langsung memakai <span class="code">string(25)</span>, Go akan mengonversinya menjadi karakter ASCII ke-25, bukan teks "25".</p>
						</div>

						<!-- String to Integer (strconv.Atoi) -->
						<div class="card">
							<h2 class="card-title">4. String ke Angka</h2>
							<table>
								<tr>
									<td class="label">String ("100") -> int</td>
									<td class="value">%s</td>
								</tr>
								<tr>
									<td class="label">String ("3.1415") -> float</td>
									<td class="value">%.4f</td>
								</tr>
							</table>
							<p class="desc">Fungsi <span class="code">strconv.Atoi()</span> mengembalikan dua nilai: hasil angka dan error. Kita harus mengoreksi error tersebut sebelum menggunakan hasilnya.</p>
						</div>
					</div>

					<div class="footer">
						Dijalankan di dalam Container Docker 🐳 | Belajar Konversi Tipe Data Go
					</div>
				</div>
			</body>
			</html>
		`, nilaiFloat, bulat,
			angkaKecil, angkaBesar,
			umur, umurString,
			statusKonversi, floatHasil)
	})

	fmt.Println("Server Go berjalan di port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
