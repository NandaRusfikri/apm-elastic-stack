# Langkah 1: Tentukan gambar dasar yang ingin Anda gunakan
FROM golang:1.21

# Langkah 4: Buat direktori kerja (misalnya /app) di dalam container
WORKDIR /app

# Langkah 5: Salin kode sumber aplikasi Golang Anda ke dalam container
COPY . .

# Langkah 6: Kompilasi aplikasi Golang
RUN go build -o main app/main.go

EXPOSE 9090

CMD ["./main"]