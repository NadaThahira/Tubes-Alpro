package main

import (
	"fmt"
)

const NMAX int = 1000

type Penggalangan struct {
	namaProyek, deskripsi, kategori        string
	targetDana, jumlahDana, jumlahInvestor int
}

type Proyek [NMAX]Penggalangan

var jumlah int

func main() {
	// I.S. : Program dijalankan
	// F.S. : Menampilkan menu dashboard
	Dashboard()
}

func Dashboard() {
	// I.S. : Pengguna menjalankan program
	// F.S. : Menampilkan menu login dan menavigasi pengguna berdasarkan pilihan
	var pilih1 int
	var data Proyek

	menu()
	fmt.Println("Pilih (0/1/2)?")
	fmt.Scan(&pilih1)

	if pilih1 < 0 || pilih1 > 2 {
		fmt.Println("Pilihlah sesuai nomor yang tertera")
	}

	for pilih1 != 0 {
		switch pilih1 {
		case 1:
			pemilik(&data)
		case 2:
			investor(&data)
		default:
			fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")
		}
		menu()
		fmt.Println("Pilih (0/1/2)?")
		fmt.Scan(&pilih1)
	}
	fmt.Print("Terimakasih telah menggunakan aplikasi ini")
	fmt.Println()
}

func menu() {
	// I.S. : Fungsi dipanggil dari Dashboard
	// F.S. : Menampilkan pilihan menu login
	fmt.Println("==============================")
	fmt.Println("|            LOGIN           |")
	fmt.Println("------------------------------")
	fmt.Println("0. Keluar")
	fmt.Println("1. Pemilik Proyek")
	fmt.Println("2. Investor")
	fmt.Println("==============================")
}

//---------------------------------------------------------------------------------------------------------------------------------

func pemilik(data *Proyek) {
	// I.S. : User sebagai pemilik proyek memilih masuk ke menu pemilik proyek
	// F.S. : User dapat membuat, mencari, mengubah, atau melihat proyek berdasarkan pilihannya
	var pilih int
	fmt.Println("==============================")
	fmt.Println("|       Pemilik Proyek       |")
	fmt.Println("------------------------------")
	fmt.Println("0. Kembali ke menu login")
	fmt.Println("1. Membuat Proyek")
	fmt.Println("2. Melihat Progress Proyek")
	fmt.Println("3. Mengubah Proyek")
	fmt.Println("4. Melihat proyek yang telah mancapai target pendanaan")
	fmt.Println("==============================")
	fmt.Println("Pilih (0/1/2/3/4)?")

	fmt.Scan(&pilih)
	if pilih < 0 || pilih > 4 {
		fmt.Println("Pilihlah sesuai nomor yang tertera")
	}

	for pilih != 0 {
		switch pilih {
		case 1:
			buatProyek(data)
		case 2:
			progressProyek(data)
		case 3:
			ubahProyek(data)
		case 4:
			fmt.Println("==============================")
			DisplayProjekTercapai(data)
			fmt.Println("==============================")
			fmt.Println()
		default:
			fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

		}
		fmt.Println("==============================")
		fmt.Println("|      Pemilik Proyek        |")
		fmt.Println("------------------------------")
		fmt.Println("0. Kembali ke menu login")
		fmt.Println("1. Membuat Proyek")
		fmt.Println("2. Melihat Progress Proyek")
		fmt.Println("3. Mengubah Proyek")
		fmt.Println("4. Melihat proyek yang telah mancapai target pendanaan")
		fmt.Println("==============================")
		fmt.Println("Pilih (0/1/2/3/4)?")

		fmt.Scan(&pilih)
		if pilih < 0 || pilih > 4 {
			fmt.Println("Pilihlah sesuai nomor yang tertera")
		}
	}
	fmt.Println()
}

func buatProyek(insert *Proyek) {
	// I.S. : User memilih untuk menambahkan proyek baru
	// F.S. : Proyek baru berhasil ditambahkan ke dalam array insert
	var n int
	fmt.Println()
	fmt.Println("==============================")
	fmt.Println("|           Proyek           |")
	fmt.Println("------------------------------")
	fmt.Println("Ingin membuat berapa proyek?")
	fmt.Println("==============================")

	fmt.Scan(&n)
	if n > NMAX || n < 0 {
		fmt.Println("Ingin membuat berapa proyek?")
		fmt.Scan(&n)
	}
	for i := 0; i < n; i++ {
		if jumlah > NMAX {
			fmt.Println("Maaf, Proyek telah mencapai maksimum ")
		}
		fmt.Println("Nama Proyek : ")
		fmt.Scan(&insert[i].namaProyek)

		fmt.Println("Deskripsi Proyek :")
		fmt.Println("Note : Setiap kata untuk deskripsi ditulis tanpa spasi")
		fmt.Scan(&insert[i].deskripsi)

		fmt.Println("Target Pendanaan : ")
		fmt.Scan(&insert[i].targetDana)
		if insert[i].targetDana < 1000 {
			fmt.Println("Dana tidak boleh dibawah seribu, silahkan masukan kembali target dana yang ingin di capai")
			fmt.Scan(&insert[i].targetDana)
		}
		fmt.Println("Kategori Proyek")
		fmt.Println("Note: Bila terdapat lebih dari satu kata, penulisan tidak di spasi")
		fmt.Scan(&insert[i].kategori)
		fmt.Println()

		insert[i].jumlahDana = 0
		insert[i].jumlahInvestor = 0
		jumlah++
	}
	fmt.Println("Terima Kasih, Proyek Berhasil Dibuat")
	fmt.Println()
}

func progressProyek(insert *Proyek) {
	// I.S. : User ingin melihat progress proyek berdasarkan kategori tertentu
	// F.S. : Menampilkan proyek yang telah diurutkan sesuai pilihan user
	var pilih int
	if jumlah == 0 {
		fmt.Println("Mohon maaf, Belum ada proyek yang tersedia")
	} else {

		fmt.Println("==============================")
		fmt.Println("|       Progress Proyek      |")
		fmt.Println("------------------------------")
		fmt.Println("Ingin melihat dalam kategori apa?")
		fmt.Println("0. Kembali pada menu pemilik proyek")
		fmt.Println("1. Nama Proyek")
		fmt.Println("2. Jumlah investor")
		fmt.Println("3. Jumlah Dana yang terkumpul")
		fmt.Println("==============================")
		fmt.Println("Pilih (0/1/2/3)?")

		fmt.Scan(&pilih)
		if pilih != 0 {
			if pilih < 0 || pilih > 3 {
				fmt.Println("Pilihlah sesuai nomor yang tertera")
			}

			if pilih == 1 {
				if jumlah == 0 {
					fmt.Println("Mohon maaf untuk saat ini belum ada proyek yang tersedia")
				} else {
					fmt.Println(jumlah)
					insertionSortByNama(insert)
					DisplayProjek(insert)
				}
			} else if pilih == 2 {
				if jumlah == 0 {
					fmt.Println("Mohon maaf untuk saat ini belum ada proyek yang tersedia")
				} else {
					fmt.Println(jumlah)
					selectionSortByJumInvestor(insert)
					DisplayProjek(insert)
				}
			} else {
				if jumlah == 0 {
					fmt.Println("Mohon maaf untuk saat ini belum ada proyek yang tersedia")
				} else {
					fmt.Println(jumlah)
					selectionSortByJumDana(insert)
					DisplayProjek(insert)
				}
			}
		}
	}

	fmt.Println()
}

func ubahProyek(insert *Proyek) {
	// I.S. : User memilih untuk mengubah atau menghapus proyek
	// F.S. : Proyek diubah atau dihapus berdasarkan pilihan
	var pilih int
	var x string

	fmt.Println("==============================")
	fmt.Println("|       Progress Proyek      |")
	fmt.Println("------------------------------")
	fmt.Println("0. Kembali ke Menu Pemilik Proyek")
	fmt.Println("1. Menghapus Proyek")
	fmt.Println("2. Mengedit Proyek")
	fmt.Println("==============================")
	fmt.Println("Pilih (0/1/2)?")
	fmt.Scan(&pilih)

	for pilih != 0 {
		switch pilih {
		case 1:
			fmt.Println("------------------------------")
			DisplayProjek(insert)
			fmt.Println("Masukkan nama proyek yang ingin dihapus:")
			fmt.Println("Note : Jika ingin kembali input angka 0")
			fmt.Println("------------------------------")
			fmt.Scan(&x)

			hapusProyek(insert, x)

		case 2:
			editProyek(insert)
		default:
			fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

		}
		fmt.Println()
		fmt.Println("==============================")
		fmt.Println("|       Progress Proyek      |")
		fmt.Println("------------------------------")
		fmt.Println("0. Kembali ke Menu Pemilik Proyek")
		fmt.Println("1. Menghapus Proyek")
		fmt.Println("2. Mengedit Proyek")
		fmt.Println("=============================")
		fmt.Println("Pilih (0/1/2)?")
		fmt.Scan(&pilih)

		if pilih < 0 || pilih > 2 {
			fmt.Println("Pilihlah sesuai nomor yang tertera")
		}
	}
	fmt.Print()
}

func hapusProyek(insert *Proyek, namaCari string) {
	// I.S. : User memilih untuk menghapus proyek
	// F.S. : Proyek dengan nama tertentu dihapus dari array jika ditemukan
	var i, low, high, mid, pos int
	var found bool

	if namaCari == "0" {
		pemilik(insert)
	}
	fmt.Println("==============================")
	insertionSortByNama(insert)
	fmt.Println("------------------------------")

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

func insertionSortByNama(insert *Proyek) {
	// I.S. : Data proyek belum terurut berdasarkan nama
	// F.S. : Data proyek terurut berdasarkan nama secara ascending
	var pass, i int

	for pass = 1; pass <= jumlah-1; pass++ {
		i = pass
		temp := insert[pass]
		for i > 0 && temp.namaProyek < insert[i-1].namaProyek {
			insert[i] = insert[i-1]
			i = i - 1
		}
		insert[i] = temp

	}

}

func editProyek(insert *Proyek) {
	// I.S. : User memilih proyek dan bagian mana yang ingin diubah
	// F.S. : Nilai atribut proyek berubah sesuai input user
	var pilih int
	var baru, targetBaru int
	var namaBaru, deskripsiBaru string

	fmt.Println("------------------------------")
	insertionSortByNama(insert)
	DisplayProjek(insert)
	fmt.Println("------------------------------")

	if jumlah == 0 {
		pemilik(insert)
	}
	fmt.Println("Mau edit proyek ke berapa?")
	fmt.Println("Note : Jika ingin kembali input angka 0")
	fmt.Println("==============================")
	fmt.Scan(&pilih)
	if pilih == 0 {
		pemilik(insert)
	} else if pilih < 1 || pilih > jumlah {
		fmt.Println("Pilihlah sesuai nomor yang tertera")
	}
	fmt.Println("Mau edit nomor berapa?")
	fmt.Println("1. Nama Proyek")
	fmt.Println("2. Deskripsi Proyek")
	fmt.Println("3. Target Dana Proyek")
	fmt.Println("------------------------------")

	fmt.Scan(&baru)
	switch baru {
	case 1:
		fmt.Print("Tulislah nama baru proyek anda: ")
		fmt.Scan(&namaBaru)
		insert[pilih-1].namaProyek = namaBaru

	case 2:
		fmt.Print("Tulislah deskripsi baru proyek anda (tanpa spasi): ")
		fmt.Scan(&deskripsiBaru)
		insert[pilih-1].deskripsi = deskripsiBaru

	case 3:
		fmt.Print("Tulislah Target baru proyek anda: ")
		fmt.Scan(&targetBaru)
		for targetBaru < 1000 {
			fmt.Println("Dana tidak boleh dibawah seribu. Masukkan lagi: ")
			fmt.Scan(&targetBaru)
		}
		insert[pilih-1].targetDana = targetBaru

	default:
		fmt.Println("Pilihan tidak valid.")
	}
	fmt.Println("Terimakasih proyek telah berhasil di edit")
	fmt.Println("==============================")
	fmt.Println()

}

func selectionSortByJumInvestor(insert *Proyek) {
	// I.S. : Data proyek belum terurut berdasarkan jumlah investor
	// F.S. : Data proyek terurut berdasarkan jumlah investor secara ascending
	var i, j, idx_min int
	var t Penggalangan

	for i = 1; i <= jumlah-1; i++ {
		idx_min = i - 1
		j = i
		for j < jumlah {
			if insert[idx_min].jumlahInvestor > insert[j].jumlahInvestor {
				idx_min = j
			}
			j++
		}
		t = insert[idx_min]
		insert[idx_min] = insert[i-1]
		insert[i-1] = t
	}

}

//-----------------------------------------------------------------------------------------------------------------------------------------------------------

func cariProyek(data *Proyek) {
	// I.S. : User ingin mencari proyek berdasarkan nama atau kategori
	// F.S. : Menampilkan proyek yang sesuai dengan kategori atau nama jika ditemukan
	var pilih int
	var x string

	fmt.Println("==============================")
	fmt.Println("|           Investor         |")
	fmt.Println("------------------------------")

	if jumlah == 0 {
		fmt.Println("Maaf, Belum ada proyek yang tersedia")
	} else {

		fmt.Println("Ingin mencari Proyek berdasarkan apa?")
		fmt.Println("1. Kategori")
		fmt.Println("2. Nama")
		fmt.Println("==============================")
		fmt.Println("Pilih (1/2)?")

		fmt.Scan(&pilih)
		for pilih != 0 {
			switch pilih {
			case 1:
				fmt.Println("Masukkan kategori yang ingin dicari: ")
				fmt.Println("Note: Bila terdapat lebih dari satu kata, penulisan tidak di spasi")
				fmt.Scan(&x)
				cariProyekByKategori(data, x)
				pilih = 0

			case 2:
				fmt.Println("Masukkan nama yang ingin dicari: ")
				fmt.Scan(&x)
				cariProyekByNama(data, x)
				pilih = 0
			default:
				fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

			}
		}
	}
	fmt.Println()
}

func cariProyekByNama(data *Proyek, nama string) {
	// I.S. : User memasukkan nama proyek yang ingin dicari
	// F.S. : Menampilkan informasi proyek jika ditemukan, atau notifikasi jika tidak ditemukan

	index := seqSeachNama(*data, nama)
	if index != -1 {
		for i := index; i >= 0 && data[i].namaProyek == nama; i-- {
			fmt.Println("Proyek ditemukan:")
			fmt.Println("Nama Proyek : ", data[index].namaProyek)
			fmt.Println("Deskripsi Proyek : ", data[index].deskripsi)
			fmt.Println("Target Pendanaan : ", data[index].targetDana)
			fmt.Println("Dana Terkumpul :", data[i].jumlahDana)
			fmt.Println("Jumlah Investor   :", data[i].jumlahInvestor)
		}
	} else {
		fmt.Println("Proyek tidak ditemukan.")
	}
}
func seqSeachNama(insert Proyek, kategori string) int {
	// I.S. : Pencarian belum dilakukan
	// F.S. : Mengembalikan indeks proyek dengan kategori yang dicari, -1 jika tidak ditemukan
	for i := 0; i < jumlah; i++ {
		if insert[i].kategori == kategori {
			return i
		}
	}
	return -1
}

func cariProyekByKategori(data *Proyek, kategori string) {
	// I.S. : User memasukkan kategori yang ingin dicari
	// F.S. : Menampilkan proyek dengan kategori tersebut jika ditemukan

	index := seqSeachKategori(*data, kategori)
	if index != -1 {
		for i := index; i >= 0 && data[i].namaProyek == kategori; i-- {
			fmt.Println("Proyek ditemukan:")
			fmt.Println("Nama Proyek : ", data[index].namaProyek)
			fmt.Println("Deskripsi Proyek : ", data[index].deskripsi)
			fmt.Println("Target Pendanaan : ", data[index].targetDana)
			fmt.Println("Dana Terkumpul :", data[index].jumlahDana)
			fmt.Println("Jumlah Investor   :", data[index].jumlahInvestor)
		}
	} else {
		fmt.Println("Tidak ada proyek dengan nama tersebut.")
	}
}

func seqSeachKategori(insert Proyek, kategori string) int {
	// I.S. : Pencarian belum dilakukan
	// F.S. : Mengembalikan indeks proyek dengan kategori yang dicari, -1 jika tidak ditemukan
	for i := 0; i < jumlah; i++ {
		if insert[i].kategori == kategori {
			return i
		}
	}
	return -1
}

func investor(data *Proyek) {
	// I.S. : User memilih login sebagai investor
	// F.S. : Menampilkan proyek dan memberikan opsi untuk menambah atau mengubah dana
	var pilih int

	cariProyek(data)
	if jumlah > 0 {
		fmt.Println("------------------------------")
		fmt.Println("0. Kembali ke menu Utama")
		fmt.Println("1. Mengubah Dana yang di investasikan")
		fmt.Println("2. Menambah Dana yang ingin di investasikan")
		fmt.Println("3. Melihat Progress Proyek")
		fmt.Println("==============================")
		fmt.Println("Pilih (0/1/2/3)?")
		fmt.Scan(&pilih)

		for pilih != 0 {
			switch pilih {
			case 0:
				Dashboard()
			case 1:
				ubahDanaInvestor(data)
			case 2:
				tambahDanaInvestor(data)
			case 3:
				progressProyek(data)
			case 4:
				cariProyek(data)
			default:
				fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

			}
			fmt.Println("==============================")
			fmt.Println("|          Investor          |")
			fmt.Println("------------------------------")
			fmt.Println("0. Kembali ke menu Utama")
			fmt.Println("1. Mengubah Dana yang di investasikan")
			fmt.Println("2. Menambah Dana yang ingin di investasikan")
			fmt.Println("3. Melihat Progress Proyek")
			fmt.Println("4. Mencari Proyek")
			fmt.Println("==============================")
			fmt.Println("Pilih (0/1/2/3/4)?")

			fmt.Scan(&pilih)
		}
	}
}

func ubahDanaInvestor(insert *Proyek) {
	// I.S. : Investor ingin mengubah jumlah dana pada proyek tertentu
	// F.S. : Dana proyek berubah sesuai input user jika proyek ditemukan
	var namaCari string
	var danaBaru int
	var ditemukan bool
	var idx int

	selectionSortByJumDana(insert)
	DisplayProjek(insert)

	fmt.Println("Masukkan nama proyek yang ingin Anda ubah dananya:")
	fmt.Scan(&namaCari)
	
	ditemukan = false
	idx = 0

	selectionSortByJumDana(insert)

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

func tambahDanaInvestor(insert *Proyek) {
	// I.S. : Investor memilih proyek dan ingin menambah dana
	// F.S. : Jumlah dana dan jumlah investor proyek bertambah
	var pilih int
	var danaInvest int

	selectionSortByJumDana(insert)
	DisplayProjek(insert)

	if jumlah == 0 {
		fmt.Println("Maaf, Belum ada Proyek yang Tersedia")
	}

	fmt.Println("Silahkan pilih Proyek keberapa yang ingin Anda investasikan") 
	fmt.Scan(&pilih)
	if pilih < 1 || pilih > jumlah {
		fmt.Println("Pilihlah sesuai nomor yang tertera")
	}

	fmt.Println("Masukan jumlah dana yang ingin anda investasikann : ")
	fmt.Println("Note : masukan dalam bilangan bulat")
	fmt.Scan(&danaInvest)
	fmt.Println("Terima kasih telah berinvestasi pada proyek ini")

	insert[pilih-1].jumlahDana += danaInvest
	insert[pilih-1].jumlahInvestor++

}

func selectionSortByJumDana(insert *Proyek) {
	// I.S. : Data proyek belum diurutkan berdasarkan jumlah dana
	// F.S. : Data proyek terurut secara descending berdasarkan jumlah dana
	var i, j, idx_max int
	var t Penggalangan
	for i = 1; i <= jumlah-1; i++ {
		idx_max = i - 1
		j = i
		for j < jumlah {
			if insert[idx_max].jumlahDana < insert[j].jumlahDana {
				idx_max = j
			}
			j++
		}
		t = insert[idx_max]
		insert[idx_max] = insert[i-1]
		insert[i-1] = t
	}

}

//---------------------------------------------------------------------------------------------------------------------------------
//tampilan

func DisplayProjek(insert *Proyek) {
	// I.S. : Dipanggil untuk menampilkan semua proyek
	// F.S. : Menampilkan informasi proyek satu per satu

	var i int
	if jumlah == 0 {
		fmt.Println("==============================")
		fmt.Println("Mohon maaf untuk saat ini belum ada proyek yang tersedia")
		fmt.Println("==============================")
	} else {
		fmt.Println("-----------------------------")
		fmt.Println("Berikut adalah proyek yang tersedia :")

		for i = 0; i < jumlah; i++ {
			fmt.Println("Proyek ", i+1)
			fmt.Println("Nama Proyek :", insert[i].namaProyek)
			fmt.Println("Deskripsi proyek : ", insert[i].deskripsi)
			fmt.Println("Target Dana :", insert[i].targetDana)
			fmt.Println("Dana Terkumpul :", insert[i].jumlahDana)
			fmt.Println("Jumlah Investor   :", insert[i].jumlahInvestor)
			fmt.Println()
		}
	}

}

func DisplayProjekTercapai(insert *Proyek) {
	// I.S. : Dipanggil untuk menampilkan proyek yang mencapai target pendanaan
	// F.S. : Menampilkan proyek-proyek yang telah mencapai target
	
	ditemukan := false
	for i := 0; i < jumlah; i++ {
		if insert[i].jumlahDana >= insert[i].targetDana {
			if !ditemukan {
				fmt.Println("------------------------------")
				fmt.Println("Berikut adalah proyek yang telah mencapai target pendanaan:")
			}
			fmt.Println("Nama Proyek       :", insert[i].namaProyek)
			fmt.Println("Deskripsi Proyek  :", insert[i].deskripsi)
			fmt.Println("Target Dana       :", insert[i].targetDana)
			fmt.Println("Dana Terkumpul    :", insert[i].jumlahDana)
			fmt.Println()
			ditemukan = true
		}
	}
	fmt.Println("------------------------------")

	if !ditemukan {
		fmt.Println("Mohon maaf, untuk saat ini belum ada proyek yang telah mencapai target.")
		fmt.Println("==============================")
	}
}
