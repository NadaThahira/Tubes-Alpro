package main

import "fmt"

type Penggalangan struct {
	namaProyek, deskripsi                  string
	targetDana, jumlahDana, jumlahInvestor int
}

var buat, hasil Penggalangan

func main() {
	var pilih int

	menu()
	fmt.Println("Pilih (1/2)?")
	fmt.Scan(&pilih)

	for pilih <= 2 {
		switch pilih {
		case 1:
			pemilik()
		case 2:
			investor()
		}
	}
}

func menu() {
	fmt.Println("===========================")
	fmt.Println("           LOGIN")
	fmt.Println("---------------------------")
	fmt.Println("1. Pemilik Proyek")
	fmt.Println("2. Investor")
	fmt.Println("===========================")
}

//---------------------------------------------------------------------------------------------------------------------------------
func pemilik() {
	fmt.Println("===========================")
	fmt.Println("      Pemilik Proyek")
	fmt.Println("---------------------------")
	fmt.Println("1. Membuat Proyek")          //nanti nyambung sama procedure buatProyek
	fmt.Println("2. Melihat Progress Proyek") //nyambung sama procedure progressProyek
	fmt.Println("===========================")

}

func buatProyek() {
	fmt.Println("===========================")
	fmt.Println("           Proyek")
	fmt.Println("---------------------------")
	fmt.Println("Nama Proyek : ")
	fmt.Scan(&buat.namaProyek) //di array + lopping

	fmt.Println("Deskripsi Proyek (untuk mengakhiri proyek beri tanda titik)")
	fmt.Scan(&buat.deskripsi) //di array + lopping + dibuat len(?) untuk baca di string di spasi

	fmt.Println("Target Pendanaan : ")
	fmt.Scan(&buat.targetDana) //di array + lopping

	fmt.Println("===========================")

}

func progressProyek() {
	fmt.Println("===========================")
	fmt.Println("      Progress Proyek")
	fmt.Println("---------------------------")
	fmt.Println("1. Jumlah Investor : ", buat.jumlahDana) //nanti nyambung sama procedure buatProyek, bingung dibedain atau masih masuk ke tipe buat
	fmt.Println("2. Jumlah Dana : ", buat.jumlahInvestor) //nyambung sama procedure perkembanganProyek, bingung dibedain atau masih masuk ke tipe buat
	fmt.Println("===========================")
}

//-----------------------------------------------------------------------------------------------------------------------------------------------------------
func investor() {
	fmt.Println("===========================")
	fmt.Println("         Investor")
	fmt.Println("---------------------------")

	//sebanyak yang diisi di pemilik proyek
	fmt.Println("Nama Proyek : ", buat.namaProyek)
	fmt.Println("Deskripsi Proyek : ", buat.deskripsi)
	fmt.Println("Target Pendanaan : ", buat.targetDana)
	fmt.Println("ingin menjadi investor diproyek mana?") // nanti di sambung sama print dana invest terus input
	fmt.Println("===========================")
}
