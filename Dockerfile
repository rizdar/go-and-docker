# Stage 1: Builder (Tahap Kompilasi)
# Menggunakan versi Go yang modern dan ringan (alpine)
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Menyalin file konfigurasi modul Go (go.sum dibuat opsional menggunakan wildcard *)
COPY go.mod go.sum* ./

# Mengunduh dependensi Go yang dibutuhkan
RUN go mod download

# MENYALIN seluruh source code Go ke dalam container
COPY . .

# MENJALANKAN kompilasi (go build) untuk menghasilkan binary statis
RUN CGO_ENABLED=0 GOOS=linux go build -o my-golang-app .

# Stage 2: Final Stage (Tahap Menjalankan Aplikasi)
# Menggunakan image alpine yang sangat ringan untuk menjalankan aplikasi
FROM alpine:latest 

WORKDIR /app

# Menyalin file binary hasil kompilasi dari stage 'builder'
COPY --from=builder /app/my-golang-app .

# Membuka port yang digunakan oleh aplikasi Go 
EXPOSE 8080 

# Perintah default untuk menjalankan aplikasi saat container aktif
CMD ["./my-golang-app"]

