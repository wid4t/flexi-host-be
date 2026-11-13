# Flexi Host Backend

Aplikasi Go sederhana menggunakan Fiber v3. Endpoint `GET /` mengembalikan:

```json
{ "message": "hello world" }
```

Server berjalan pada port `3000`.

## Instalasi

1. Instal Go versi 1.26 atau lebih baru dari [go.dev/dl](https://go.dev/dl/) sesuai sistem operasi. Ikuti petunjuk pemasangannya, buka terminal baru, lalu periksa dengan `go version`.
2. Untuk menggunakan `make live`, instal GNU Make:
   - **Windows:** instal GNU Make, misalnya melalui Chocolatey (`choco install make`), lalu pastikan `make` tersedia di PATH.
   - **Linux:** gunakan pengelola paket distribusi, misalnya `sudo apt install make` pada Ubuntu/Debian.
   - **macOS:** jalankan `xcode-select --install` untuk memasang Command Line Tools yang menyertakan Make.
3. Buka terminal di direktori proyek, lalu unduh dependensi:

   ```sh
   go mod download
   ```

## Menjalankan aplikasi

Pada Windows (PowerShell atau Command Prompt), Linux, maupun macOS:

```sh
go run .
```

Buka `http://localhost:3000/` di browser, atau jalankan `curl http://localhost:3000/`. Hentikan server dengan `Ctrl+C`. Server juga menangani `SIGTERM` dengan graceful shutdown: server berhenti menerima koneksi baru dan menunggu koneksi aktif selesai hingga 10 detik sebelum proses keluar.

## Build dan menjalankan di komputer lokal

Perintah berikut menghasilkan binary sesuai sistem operasi dan arsitektur komputer yang menjalankan Go:

- **Windows:** `go build -o app.exe .`, lalu jalankan `./app.exe` di PowerShell atau `app.exe` di Command Prompt.
- **Linux/macOS:** `go build -o app .`, lalu jalankan `./app`.

## Build Linux ARM64 untuk deployment

Di Windows, Linux, atau macOS, jalankan:

```sh
make live
```

Perintah ini melakukan cross-compile untuk **Linux ARM64** (`GOOS=linux`, `GOARCH=arm64`, `CGO_ENABLED=0`) dan membuat `app.zip` di direktori proyek. File ZIP berisi binary `app` yang dapat dieksekusi; `app.zip` yang sudah ada akan ditimpa. Koneksi internet mungkin diperlukan saat Go pertama kali mengunduh dependensi.

Salin `app.zip` ke mesin Linux ARM64, kemudian jalankan:

```sh
unzip app.zip
chmod +x app
./app
```

Binary ini tidak dapat langsung dijalankan pada Windows, macOS, atau Linux dengan arsitektur selain ARM64. Gunakan build lokal di atas untuk sistem tersebut. Pastikan port `3000` dapat diakses sesuai konfigurasi firewall dan jaringan server.

## Peringatan penyimpanan di SewuCloud

SewuCloud menggunakan pod untuk menjalankan aplikasi. Data yang disimpan di luar folder `upload` dapat hilang saat pod di-restart atau diganti. Simpan file yang perlu dipertahankan di folder `upload` dan pastikan folder tersebut terhubung ke penyimpanan persisten sesuai konfigurasi SewuCloud. Lakukan pencadangan data penting secara berkala.

## Pengujian

```sh
go test ./...
```
