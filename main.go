package main

import (
	"errors"
	"fmt"
	"net/http"
)

// 1. Fungsi biasa dengan satu return value
// Karena tipe parameter a dan b sama, kita bisa mempersingkatnya menjadi "a, b int"
func tambah(a, b int) int {
	return a + b
}

// 2. Fungsi dengan MULTIPLE RETURN VALUES (Mengembalikan lebih dari satu nilai)
func hitungLuasDanKeliling(panjang, lebar float64) (float64, float64) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}

// 3. Fungsi dengan NAMED RETURN VALUES (Return value yang sudah memiliki nama variabel)
// Fungsi ini juga mendemonstrasikan cara umum penanganan error di Go
func bagi(pembilang, penyebut float64) (hasil float64, err error) {
	if penyebut == 0 {
		err = errors.New("tidak bisa membagi dengan nol")
		return // Naked return: otomatis mengembalikan nilai variabel 'hasil' dan 'err' yang terakhir diset
	}
	hasil = pembilang / penyebut
	return // Naked return
}

// 4. Variadic Function (Fungsi yang bisa menerima parameter tak terbatas jumlahnya)
func jumlahkanSemua(angka ...int) int {
	total := 0
	// Parameter 'angka' dibaca sebagai slice []int di dalam fungsi ini
	for _, nilai := range angka {
		total += nilai
	}
	return total
}

func main() {
	// ==========================================
	// PENGGUNAAN FUNGSI DI GOLANG
	// ==========================================

	// Memanggil Fungsi 1
	hasilTambah := tambah(15, 27)

	// Memanggil Fungsi 2 (Menerima multiple return values)
	luasPersegi, kelilingPersegi := hitungLuasDanKeliling(10.0, 5.0)

	// Memanggil Fungsi 3 (Menguji pembagian normal & pembagian dengan nol)
	hasilBagiNormal, errNormal := bagi(10, 2)
	var teksBagiNormal string
	if errNormal != nil {
		teksBagiNormal = errNormal.Error()
	} else {
		teksBagiNormal = fmt.Sprintf("%.2f", hasilBagiNormal)
	}

	_, errNol := bagi(10, 0) // Menggunakan '_' untuk mengabaikan nilai 'hasil' yang dikembalikan
	var teksBagiNol string
	if errNol != nil {
		teksBagiNol = fmt.Sprintf("Error: %s", errNol.Error())
	} else {
		teksBagiNol = "Sukses"
	}

	// Memanggil Fungsi 4 (Variadic function)
	totalVariadic := jumlahkanSemua(10, 20, 30, 40, 50)

	// Menampilkan data tersebut di halaman web dengan styling modern
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
					.error-val {
						color: #f43f5e;
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
					<h1>Function (Fungsi) di Go</h1>
					<p class="subtitle">Mempelajari cara membuat blok kode modular menggunakan Function, Parameter, Multiple Returns, dan Variadic Arguments.</p>
					
					<div class="grid">
						<!-- Basic Function Card -->
						<div class="card">
							<h2 class="card-title">1. Fungsi Biasa</h2>
							<table>
								<tr>
									<td class="label">Panggilan: tambah(15, 27)</td>
									<td class="value">%d</td>
								</tr>
							</table>
							<p class="desc">Bentuk fungsi umum dengan parameter input dan satu nilai kembalian (return value). Jika tipe parameter bersebelahan sama, bisa disingkat (misal: <code>a, b int</code>).</p>
						</div>

						<!-- Multiple Return Values Card -->
						<div class="card">
							<h2 class="card-title">2. Multiple Return Values</h2>
							<table>
								<tr>
									<td class="label">Luas (panjang=10, lebar=5)</td>
									<td class="value">%.1f</td>
								</tr>
								<tr>
									<td class="label">Keliling (panjang=10, lebar=5)</td>
									<td class="value">%.1f</td>
								</tr>
							</table>
							<p class="desc">Fitur andalan Go yang memungkinkan fungsi mengembalikan lebih dari satu nilai sekaligus. Kita bisa mengabaikan salah satu nilai menggunakan blank identifier (<code>_</code>).</p>
						</div>

						<!-- Named Returns Card -->
						<div class="card">
							<h2 class="card-title">3. Named &amp; Error Returns</h2>
							<table>
								<tr>
									<td class="label">Bagi Normal: bagi(10, 2)</td>
									<td class="value">%s</td>
								</tr>
								<tr>
									<td class="label">Bagi Nol: bagi(10, 0)</td>
									<td class="value error-val">%s</td>
								</tr>
							</table>
							<p class="desc">Mendefinisikan nama variabel di bagian tipe data kembalian. Hal ini memungkinkan <em>Naked Return</em> (cukup tulis <code>return</code> saja). Go juga menggunakan return ganda untuk error handling konvensional.</p>
						</div>

						<!-- Variadic Function Card -->
						<div class="card">
							<h2 class="card-title">4. Variadic Function</h2>
							<table>
								<tr>
									<td class="label">jumlahkanSemua(10,20,30,40,50)</td>
									<td class="value">%d</td>
								</tr>
							</table>
							<p class="desc">Fungsi yang bisa menerima parameter tak terbatas jumlahnya menggunakan operator tiga titik (<code>...</code>). Di dalam fungsi, parameter ini dibaca sebagai slice.</p>
						</div>
					</div>

					<div class="footer">
						Dijalankan di dalam Container Docker 🐳 | Belajar Function Go
					</div>
				</div>
			</body>
			</html>
		`, hasilTambah,
			luasPersegi, kelilingPersegi,
			teksBagiNormal, teksBagiNol,
			totalVariadic)
	})

	fmt.Println("Server Go berjalan di port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
