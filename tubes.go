package main

import "fmt"

const NMAX int = 1000

type Penggalangan struct {
	namaProyek, deskripsi                  string
	targetDana, jumlahDana, jumlahInvestor int
}

type Proyek [NMAX]Penggalangan

var jumlah int

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

		default:
			fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")
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
	var pilih int
	fmt.Println("===========================")
	fmt.Println("      Pemilik Proyek")
	fmt.Println("---------------------------")
	fmt.Println("1. Membuat Proyek")          //nanti nyambung sama procedure buatProyek
	fmt.Println("2. Melihat Progress Proyek") //nyambung sama procedure progressProyek
	fmt.Println("===========================")
	fmt.Println("Pilih (1/2)?")
	fmt.Scan(&pilih)

	for pilih <= 2 {
		switch pilih {
		case 1:
			buatProyek()
		case 2:
			progressProyek()
		default:
			fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

		}
	}
	fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

}

func buatProyek() {
	var insert Proyek
	var pilih int
	var n int
	fmt.Println("===========================")
	fmt.Println("           Proyek")
	fmt.Println("---------------------------")
	fmt.Println("Ingin membuat berapa Proyek?")

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Println("Nama Proyek : ")
		fmt.Scan(&insert[i].namaProyek) //di array + lopping

		fmt.Println("Deskripsi Proyek (untuk mengakhiri proyek beri tanda titik)")
		fmt.Scan(&insert[i].deskripsi) //di array + lopping + dibuat len(?) untuk baca di string di spasi

		fmt.Println("Target Pendanaan : ")
		fmt.Scan(&insert[i].targetDana) //di array + lopping

		insert[i].jumlahDana = 0
		insert[i].jumlahInvestor = 0
		jumlah++
	}
	fmt.Println("Terima Kasih, Proyek Berhasil Dibuat")
	fmt.Println("1. Return to Menu")
	fmt.Println("===========================")

	for pilih == 1 {
		switch pilih {
		case 1:
			menu()
		default:
			fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

		}
	}
	fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

}

func progressProyek() {
	var progress Proyek
	if jumlah == 0 {
		fmt.Println("Maaf, Belum ada Proyek")
	}

	fmt.Println("===========================")
	fmt.Println("      Progress Proyek")
	fmt.Println("---------------------------")

	for i := 0; i < jumlah; i++ {
		fmt.Printf("%d. %s - %s (target : %d, Terkumpul : %d)\n", i+1, progress[i].namaProyek, progress[i].deskripsi, progress[i].targetDana, progress[i].jumlahDana)

		fmt.Println("1. Jumlah Investor : ", progress[i].jumlahDana) //nanti nyambung sama procedure buatProyek, bingung dibedain atau masih masuk ke tipe buat
		fmt.Println("2. Jumlah Dana : ", progress[i].jumlahInvestor) //nyambung sama procedure perkembanganProyek, bingung dibedain atau masih masuk ke tipe buat
	}
	fmt.Println("===========================")
}

//-----------------------------------------------------------------------------------------------------------------------------------------------------------
func investor() {
	var valid Proyek

	if jumlah == 0 {
		fmt.Println("Maaf, Belum ada Proyek yang Tersedia")
	}
	fmt.Println("===========================")
	fmt.Println("         Investor")
	fmt.Println("---------------------------")

	//sebanyak yang diisi di pemilik proyek
	for i := 0; i < jumlah; i++ {
		fmt.Println("Nama Proyek : ", valid[i].namaProyek)
		fmt.Println("Deskripsi Proyek : ", valid[i].deskripsi)
		fmt.Println("Target Pendanaan : ", valid[i].targetDana)

	}
	fmt.Println("ingin menjadi investor diproyek mana?") // nanti di sambung sama print dana invest terus input
	fmt.Println("===========================")
	fmt.Println("Silahkan pilih Proyek yang ingin Anda investasikan")
}
