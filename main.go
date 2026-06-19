package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ==========================================
	// CONTROL FLOW DI GOLANG
	// ==========================================

	// 1. If - Else If - Else
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

	// 2. If dengan Temporary Variable (Variabel Sementara)
	// Variabel 'statusUjian' ditentukan berdasarkan 'kriteria' yang dideklarasikan langsung di dalam blok 'if'
	var statusUjian string
	if kriteria := 75; nilai >= kriteria {
		statusUjian = "LULUS (Nilai di atas kriteria)"
	} else {
		statusUjian = "TIDAK LULUS"
	}

	// 3. Switch Case (Tanpa perlu menuliskan 'break')
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

	// 4. For Loop (Satu-satunya keyword perulangan di Go!)
	// A. Loop Standar (For biasa)
	var hasilLoopBiasa string
	for i := 1; i <= 3; i++ {
		hasilLoopBiasa += fmt.Sprintf("Perulangan ke-%d, ", i)
	}

	// B. Loop bergaya While (For dengan satu kondisi saja)
	var counter = 1
	var hasilLoopWhile string
	for counter <= 3 {
		hasilLoopWhile += fmt.Sprintf("Iterasi %d; ", counter)
		counter++
	}

	// C. For Range (Digunakan untuk mengiterasi slice/array/map/string)
	frameworks := []string{"Go", "Docker", "Git"}
	var hasilLoopRange string
	for indeks, item := range frameworks {
		hasilLoopRange += fmt.Sprintf("[%d: %s] ", indeks, item)
	}

	// Menampilkan data tersebut di halaman web dengan styling modern
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Belajar Control Flow Go</title>
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
					<h1>Control Flow di Go</h1>
					<p class="subtitle">Bagaimana mengontrol alur eksekusi program menggunakan Percabangan dan Perulangan.</p>
					
					<div class="grid">
						<!-- IF-ELSE Card -->
						<div class="card">
							<h2 class="card-title">1. If - Else If - Else</h2>
							<table>
								<tr>
									<td class="label">Nilai Ujian</td>
									<td class="value">%d</td>
								</tr>
								<tr>
									<td class="label">Grade Hasil</td>
									<td class="value">%s</td>
								</tr>
								<tr>
									<td class="label">Status (If Temp Var)</td>
									<td class="value">%s</td>
								</tr>
							</table>
							<p class="desc">Go mendukung inisialisasi variabel sementara langsung di dalam statemen <code>if</code> sebelum kondisi dievaluasi. Sangat berguna untuk scope variabel yang sempit.</p>
						</div>

						<!-- SWITCH Card -->
						<div class="card">
							<h2 class="card-title">2. Switch Case</h2>
							<table>
								<tr>
									<td class="label">Angka Hari</td>
									<td class="value">%d</td>
								</tr>
								<tr>
									<td class="label">Nama Hari</td>
									<td class="value">%s</td>
								</tr>
							</table>
							<p class="desc">Di Go, switch case secara default <strong>tidak butuh break</strong>. Eksekusi akan langsung keluar dari switch setelah case terpenuhi, kecuali Anda menggunakan keyword <code>fallthrough</code>.</p>
						</div>

						<!-- FOR Loop Card -->
						<div class="card" style="grid-column: 1 / -1;">
							<h2 class="card-title">3. For Loops (Satu-satunya Keyword Loop)</h2>
							<table>
								<tr>
									<td class="label">For Standard (i:=1; i&lt;=3; i++)</td>
									<td class="value">%s</td>
								</tr>
								<tr>
									<td class="label">For While-Style (counter &lt;= 3)</td>
									<td class="value">%s</td>
								</tr>
								<tr>
									<td class="label">For Range (slice frameworks)</td>
									<td class="value">%s</td>
								</tr>
							</table>
							<p class="desc">Go hanya memiliki <strong>satu</strong> kata kunci untuk perulangan, yaitu <code>for</code>. Namun, <code>for</code> ini sangat fleksibel dan bisa ditulis dengan berbagai variasi gaya.</p>
						</div>
					</div>

					<div class="footer">
						Dijalankan di dalam Container Docker 🐳 | Belajar Control Flow Go
					</div>
				</div>
			</body>
			</html>
		`, nilai, grade, statusUjian,
			hari, namaHari,
			hasilLoopBiasa, hasilLoopWhile, hasilLoopRange)
	})

	fmt.Println("Server Go berjalan di port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
