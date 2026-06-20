package main

import (
	"fmt"
	"net/http"
)

// handleVariables menampilkan halaman materi variabel
func handleVariables(w http.ResponseWriter, r *http.Request) {
	// 1. Deklarasi Eksplisit
	var nama string = "Rizdar"

	// 2. Type Inference
	var umur = 20

	// 3. Short Declaration
	isBelajar := true

	// 4. Tipe desimal
	var ipk float64 = 3.85

	// 5. Konstanta
	const versiGo = "1.22"

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="id">
		<head>
			<meta charset="UTF-8">
			<title>Belajar Variabel & Tipe Data Go</title>
			<style>
				body {
					font-family: 'Segoe UI', system-ui, -apple-system, sans-serif;
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
					box-shadow: 0 10px 25px rgba(0,0,0,0.3);
					max-width: 450px;
					width: 100%%;
					border: 1px solid #334155;
				}
				h1 {
					color: #38bdf8;
					text-align: center;
					border-bottom: 2px solid #334155;
					padding-bottom: 15px;
					margin-top: 0;
				}
				ul {
					list-style: none;
					padding: 0;
				}
				li {
					padding: 12px 0;
					border-bottom: 1px solid #334155;
					display: flex;
					justify-content: space-between;
				}
				.label {
					color: #94a3b8;
					font-weight: bold;
				}
				.value {
					font-family: monospace;
					color: #34d399;
					font-weight: bold;
					background: #0f172a;
					padding: 2px 6px;
					border-radius: 4px;
				}
				.back-link {
					display: block;
					text-align: center;
					margin-top: 25px;
					color: #38bdf8;
					text-decoration: none;
					font-weight: bold;
				}
				.back-link:hover {
					text-decoration: underline;
				}
			</style>
		</head>
		<body>
			<div class="card">
				<h1>Variabel & Tipe Data Go</h1>
				<ul>
					<li><span class="label">Nama (string)</span> <span class="value">%s</span></li>
					<li><span class="label">Umur (int)</span> <span class="value">%d tahun</span></li>
					<li><span class="label">Belajar Go (bool)</span> <span class="value">%t</span></li>
					<li><span class="label">IPK (float64)</span> <span class="value">%.2f</span></li>
					<li><span class="label">Versi Go (const)</span> <span class="value">%s</span></li>
				</ul>
				<a href="/" class="back-link">⬅ Kembali ke Menu Utama</a>
			</div>
		</body>
		</html>
	`, nama, umur, isBelajar, ipk, versiGo)
}
