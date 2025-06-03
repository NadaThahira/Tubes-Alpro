# crowfunding_go
Aplikasi ini merupakan platform yang menyediakan penggalangan dana, berdasarkan permintaan dan kebutuhan pelanggan. Aplikasi ini mencakup berbagai layanan, mulai dari penggalangan dana. Tujuannya adalah untuk menyederhanakan proses penggalangan dan mencari investor.
## Installation

Clone project ini

```bash
https://github.com/NadaThahira/Tubes-Alpro/blob/a7918e763220248f5704e0542f9a2aed987cf362/tubes.go
```

Menjalankan project

```bash
go run tubes.go
## Fitur Utama
```
### 1. **Login Menu**

Pengguna dapat memilih peran:

* **Pemilik Proyek**
* **Investor**

### 2. **Pemilik Proyek**

Sebagai pemilik proyek, pengguna dapat:

* Membuat proyek penggalangan dana.
* Melihat semua proyek dan progresnya berdasarkan:

  * Nama proyek
  * Jumlah dana terkumpul
  * Jumlah investor
* Mengedit atau menghapus proyek.
* Melihat proyek yang telah mencapai target dana.

### 3. **Investor**

Sebagai investor, pengguna dapat:

* Mencari proyek berdasarkan nama atau kategori.
* Menambahkan dana ke proyek pilihan.
* Mengubah nominal dana yang telah diinvestasikan.
* Melihat progres proyek-proyek yang tersedia.

### 4. **Sorting dan Searching**

Aplikasi menggunakan:

* **Insertion Sort** untuk pengurutan berdasarkan nama dan jumlah investor.
* **Selection Sort** untuk pengurutan berdasarkan jumlah dana.
* **Binary Search** untuk pencarian berdasarkan nama.
* **Sequential Search** untuk pencarian berdasarkan kategori.

## Struktur Data

* Menggunakan array statis `Proyek` berisi maksimal 1000 entri `Penggalangan`.
* Tipe data `Penggalangan` menyimpan informasi proyek:

  * Nama, deskripsi, kategori
  * Target dana, dana terkumpul, jumlah investor

## Petunjuk Penggunaan

1. Jalankan aplikasi dengan `go run tubes.go`
2. Pilih peran: Pemilik Proyek atau Investor
3. Gunakan menu interaktif untuk menavigasi dan mengelola data proyek

## Catatan Teknis

* Program ini bersifat **terminal-based**, tidak menggunakan database atau penyimpanan file.
* Beberapa validasi input sederhana telah disediakan untuk menjaga integritas data.
* Nilai input harus sesuai format (contoh: tanpa spasi untuk deskripsi atau kategori).
* Program dapat dikembangkan lebih lanjut dengan integrasi database atau GUI.

## Contoh Screenshot Menu (CLI)

```
===========================
           LOGIN
---------------------------
0. Keluar
1. Pemilik Proyek
2. Investor
===========================
Pilih (0/1/2)?
```

## Kontributor

* Mahasiswa Proyek Akhir (Semester 2)
* Bahasa: Go (Golang)
