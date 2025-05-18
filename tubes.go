package main

import "fmt"

const NMAX int = 1000

type Penggalangan struct {
	namaProyek, deskripsi                  string
	targetDana, jumlahDana, jumlahInvestor int
}

type Proyek [NMAX]Penggalangan

//var insert Proyek
var jumlah int

func main() {
	Dashboard()
}

func Dashboard() {
	var pilih int

	menu()
	fmt.Println("Pilih (0/1/2)?")
	fmt.Scan(&pilih)

	if pilih < 1 || pilih > jumlah {
		fmt.Println("Pilihlah sesuai nomor yang tertera")
	}

	for pilih <= 2 {
		switch pilih {
		case 0:
			fmt.Print("Terimakasih telah menggunakan aplikasi ini")
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
	fmt.Println("0. Keluar")
	fmt.Println("1. Pemilik Proyek")
	fmt.Println("2. Investor")
	fmt.Println("===========================")
}

//---------------------------------------------------------------------------------------------------------------------------------
func pemilik() {
	var pilih int
	var add Proyek
	fmt.Println("===========================")
	fmt.Println("      Pemilik Proyek")
	fmt.Println("---------------------------")
	fmt.Println("0. Kembali ke menu utama")
	fmt.Println("1. Membuat Proyek")          //nanti nyambung sama procedure buatProyek
	fmt.Println("2. Melihat Progress Proyek") //nyambung sama procedure progressProyek
	fmt.Println("3. Mengubah Proyek")
	fmt.Println("===========================")
	fmt.Println("Pilih (0/1/2)?")

	fmt.Scan(&pilih)
	if pilih < 1 || pilih > jumlah {
		fmt.Println("Pilihlah sesuai nomor yang tertera")
	}

	for pilih <= 3 {
		switch pilih {
		case 0:
			Dashboard()
		case 1:
			buatProyek(&add, pilih)
		case 2:
			progressProyek(&add, pilih)
		case 3:
			ubahProyek(&add, pilih)
		default:
			fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

		}
	}
	fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

}

func buatProyek(insert *Proyek, pilih int) {
	var n int
	fmt.Println("===========================")
	fmt.Println("           Proyek")
	fmt.Println("---------------------------")
	fmt.Println("Ingin membuat berapa proyek?")

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		if jumlah > NMAX {
			fmt.Println("Maaf, Proyek telah mencapai maksimum ")
		}
		fmt.Println("Nama Proyek : ")
		fmt.Scan(&insert[i].namaProyek)

		fmt.Println("Deskripsi Proyek")
		fmt.Println("Note : untuk mengakhiri proyek beri tanda titik")
		fmt.Scan(&insert[i].deskripsi)

		fmt.Println("Target Pendanaan : ")
		fmt.Scan(&insert[i].targetDana)
		if insert[i].targetDana < 1000 {
			fmt.Println("Dana tidak boleh dibawah seribu, silahkan masukan kembali target dana yang ingin di capai")
			fmt.Scan(&insert[i].targetDana)
		}

		insert[i].jumlahDana = 0
		insert[i].jumlahInvestor = 0
		jumlah++
	}
	fmt.Println("Terima Kasih, Proyek Berhasil Dibuat") // setelah itu kembali ke menu utama
	fmt.Println("Tekan 0 untuk kembali ke menu utama")
	fmt.Println("===========================")

	fmt.Scan(&pilih)
	if pilih < 1 || pilih > jumlah {
		fmt.Println("Pilihlah sesuai nomor yang tertera")
	}

	for pilih == 0 {
		switch pilih {
		case 1:
			Dashboard()
		default:
			fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

		}
	}
	fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

}

func progressProyek(insert *Proyek, pilih int) {
	if jumlah == 0 {
		fmt.Println("Maaf, Belum ada proyek yang tersedia")
	}

	fmt.Println("===========================")
	fmt.Println("      Progress Proyek")
	fmt.Println("---------------------------")

	for i := 0; i < jumlah; i++ {
		fmt.Printf("%d. %s - %s (target : %d)\n", i+1, insert[i].namaProyek, insert[i].deskripsi, insert[i].targetDana)

		fmt.Println("1. Jumlah Dana : ", insert[i].jumlahDana)         //nanti nyambung sama procedure buatProyek, bingung dibedain atau masih masuk ke tipe buat
		fmt.Println("2. Jumlah Investor : ", insert[i].jumlahInvestor) //nyambung sama procedure perkembanganProyek, bingung dibedain atau masih masuk ke tipe buat
	}
	fmt.Println("0. Kembali ke Menu Utama")
	fmt.Println("===========================")

	fmt.Scan(&pilih)
	if pilih < 1 || pilih > jumlah {
		fmt.Println("Pilihlah sesuai nomor yang tertera")
	}

	for pilih == 0 {
		switch pilih {
		case 1:
			Dashboard()
		default:
			fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

		}
	}
	fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

}

func ubahProyek(insert *Proyek, pilih int) {
	fmt.Println("===========================")
	fmt.Println("      Pemilik Proyek")
	fmt.Println("---------------------------")
	fmt.Println("0. Kembali ke Menu Utama")
	fmt.Println("1. Menghapus Proyek")
	fmt.Println("2. Mengedit Proyek")
	fmt.Println("===========================")

	for pilih == 0 {
		switch pilih {
		case 0:
			Dashboard()
		case 1:
			hapusProyek()
		case 2:
			editProyek()
		default:
			fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

		}
	}
	fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

}

func hapusProyek(insert *Proyek, pilih int) {
	var namaCari string
	var i, low, high, mid, pos int
	var found bool

	fmt.Println("Masukkan nama proyek yang ingin dihapus:")
	fmt.Scan(&namaCari)

	insertionSortByNama()

	low = 0
	high = jumlah - 1
	found = false

	for low <= high && !found {
		mid = (low + high) / 2
		if insert[mid].namaProyek == namaCari {
			found = true
			pos = mid
		} else {
			if insert[mid].namaProyek < namaCari {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}

	if found {
		for i = pos; i < jumlah-1; i++ {
			insert[i] = insert[i+1]
		}
		jumlah--
		fmt.Println("Proyek berhasil dihapus.")
	} else {
		fmt.Println("Proyek tidak ditemukan.")
	}
}


func editProyek(insert *Proyek, pilih int) {
	var baru, targetBaru int
	var namaBaru, deskripsiBaru string
	fmt.Println("Mau edit proyek ke berapa?")
	fmt.Scan(&pilih)
	if pilih < 1 || pilih > jumlah {
		fmt.Println("Pilihlah sesuai nomor yang tertera")
	}
	fmt.Println("Mau edit nomor berapa?")
	fmt.Println("1. Nama Proyek")
	fmt.Println("2. Deskripsi Proyek")
	fmt.Println("3. Target Dana Proyek")

	fmt.Scan(&baru)
	if baru < 1 || baru > jumlah {
		fmt.Println("Pilihlah sesuai nomor yang tertera")
	}
	if baru == 1 {
		fmt.Scan(&namaBaru)
		insert[pilih-1].namaProyek = namaBaru
	} else if baru == 2 {
		fmt.Scan(&deskripsiBaru)
		insert[pilih-1].deskripsi = deskripsiBaru
	} else if baru == 3 {
		fmt.Scan(&targetBaru)
		if targetBaru < 1000 {
			fmt.Println("Dana tidak boleh dibawah seribu, silahkan masukan kembali target dana yang ingin di capai")
			fmt.Scan(&targetBaru)
		}
		insert[pilih-1].targetDana = targetBaru

	}

}

//-----------------------------------------------------------------------------------------------------------------------------------------------------------
func investor() {
	var pilih int

	fmt.Println("===========================")
	fmt.Println("         Investor")
	fmt.Println("---------------------------")
	fmt.Println("0. Kembali ke menu Utama")
	fmt.Println("1. Mengubah Dana yang di investasikan")
	fmt.Println("2. Menambah Dana yang ingin di investasikan")
	fmt.Println("Pilih (0/1/2)?")
	fmt.Println("===========================")
	fmt.Scan(&pilih)

	for pilih == 0 {
		switch pilih {
		case 0:
			Dashboard()
		case 1:
			ubahDanaInvestor()
		case 2:
			tambahDanaInvestor()
		default:
			fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

		}
	}
	fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

}

func ubahDanaInvestor(insert *Proyek, pilih int) {
	var namaCari string
	var danaBaru int
	var ditemukan bool
	var idx int

	fmt.Println("Masukkan nama proyek yang ingin Anda ubah dananya:")
	fmt.Scan(&namaCari)

	ditemukan = false
	idx = 0

	for idx < jumlah && !ditemukan {
		if insert[idx].namaProyek == namaCari {
			ditemukan = true
		} else {
			idx++
		}
	}

	if ditemukan {
		fmt.Println("Proyek ditemukan: ", insert[idx].namaProyek)
		fmt.Println("Jumlah dana saat ini: ", insert[idx].jumlahDana)
		fmt.Println("Masukkan jumlah dana baru:")
		fmt.Scan(&danaBaru)
		insert[idx].jumlahDana = danaBaru
		fmt.Println("Dana berhasil diubah.")
	} else {
		fmt.Println("Proyek tidak ditemukan.")
	}
}


func tambahDanaInvestor(insert *Proyek, pilih int) {
	var danaInvest int

	if jumlah == 0 {
		fmt.Println("Maaf, Belum ada Proyek yang Tersedia")
	}
	//sebanyak yang diisi di pemilik proyek
	for i := 0; i < jumlah; i++ {
		fmt.Println("Proyek ", i+1)
		fmt.Println("Nama Proyek : ", insert[i].namaProyek)
		fmt.Println("Deskripsi Proyek : ", insert[i].deskripsi)
		fmt.Println("Target Pendanaan : ", insert[i].targetDana)

	}

	fmt.Println("Silahkan pilih Proyek yang ingin Anda investasikan") // nanti di sambung sama print dana invest terus input
	fmt.Scan(&pilih)
	if pilih < 1 || pilih > jumlah {
		fmt.Println("Pilihlah sesuai nomor yang tertera")
	}

	fmt.Println("Masukan jumlah dana yang ingin anda investasikann : ")
	fmt.Println("Note : masukan dalam bilangan bulat")
	fmt.Scan(&danaInvest)

	insert[pilih-1].jumlahDana = danaInvest
	insert[pilih-1].jumlahInvestor++

}
