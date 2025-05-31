package main

import "fmt"

const NMAX int = 1000

type Penggalangan struct {
	namaProyek, deskripsi, kategori        string
	targetDana, jumlahDana, jumlahInvestor int
}

type Proyek [NMAX]Penggalangan

//var insert Proyek
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

	if pilih1 < 1 || pilih1 > 2 {
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
	fmt.Println() // masih eror
}

func menu() {
	// I.S. : Fungsi dipanggil dari Dashboard
	// F.S. : Menampilkan pilihan menu login
	fmt.Println("===========================")
	fmt.Println("           LOGIN")
	fmt.Println("---------------------------")
	fmt.Println("0. Keluar")
	fmt.Println("1. Pemilik Proyek")
	fmt.Println("2. Investor")
	fmt.Println("===========================")
}

//---------------------------------------------------------------------------------------------------------------------------------

func pemilik(data *Proyek) {
	// I.S. : User sebagai pemilik proyek memilih masuk ke menu pemilik proyek
	// F.S. : User dapat membuat, mencari, mengubah, atau melihat proyek berdasarkan pilihannya
	var pilih int
	fmt.Println("===========================")
	fmt.Println("      Pemilik Proyek")
	fmt.Println("---------------------------")
	fmt.Println("0. Kembali ke menu login")
	fmt.Println("1. Membuat Proyek")
	fmt.Println("2. Melihat Progress Proyek")
	fmt.Println("3. Mengubah Proyek")
	fmt.Println("4. Melihat proyek yang telah mancapai target pendanaan")
	fmt.Println("===========================")
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
			progressProyek(data) //masih eror
		case 3:
			ubahProyek(data)
		case 4:
			DisplayProjekTercapai(data)
		default:
			fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

		}
		fmt.Println("===========================")
		fmt.Println("      Pemilik Proyek")
		fmt.Println("---------------------------")
		fmt.Println("0. Kembali ke menu login")
		fmt.Println("1. Membuat Proyek")
		fmt.Println("2. Melihat Progress Proyek")
		fmt.Println("3. Mengubah Proyek")
		fmt.Println("4. Melihat proyek yang telah mancapai target pendanaan")
		fmt.Println("===========================")
		fmt.Println("Pilih (0/1/2/3/4)?")

		fmt.Scan(&pilih)
		if pilih < 0 || pilih > 4 {
			fmt.Println("Pilihlah sesuai nomor yang tertera")
		}
	}

}

func buatProyek(insert *Proyek) {
	// I.S. : User memilih untuk menambahkan proyek baru
	// F.S. : Proyek baru berhasil ditambahkan ke dalam array insert
	var n int
	fmt.Println()
	fmt.Println("===========================")
	fmt.Println("           Proyek")
	fmt.Println("---------------------------")
	fmt.Println("Ingin membuat berapa proyek?")

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
		fmt.Println("Note : untuk mengakhiri proyek beri tanda titik")
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
		fmt.Println("Maaf, Belum ada proyek yang tersedia")
	}

	fmt.Println("===========================")
	fmt.Println("      Progress Proyek")
	fmt.Println("---------------------------")
	fmt.Println()
	fmt.Println("Ingin melihat dalam kategori apa?")
	fmt.Println("1. Nama Proyek")
	fmt.Println("2. Jumlah investor")
	fmt.Println("3. Jumlah Dana yang terkumpul")
	fmt.Println("Pilih (1/2/3)?")

	fmt.Scan(&pilih)
	if pilih < 0 || pilih > 3 {
		fmt.Println("Pilihlah sesuai nomor yang tertera")
	}

	if pilih == 1 {
		insertionSortByNama(insert)
	} else if pilih == 2 {
		insertionSortByJumInvestor(insert)
	} else {
		selectionSortByJumDana(insert)
	}
}

func ubahProyek(insert *Proyek) {
	// I.S. : User memilih untuk mengubah atau menghapus proyek
	// F.S. : Proyek diubah atau dihapus berdasarkan pilihan
	var pilih int
	fmt.Println("===========================")
	fmt.Println("      Pemilik Proyek")
	fmt.Println("---------------------------")
	fmt.Println("0. Kembali ke Menu login")
	fmt.Println("1. Menghapus Proyek")
	fmt.Println("2. Mengedit Proyek")
	fmt.Println("===========================")
	fmt.Println("Pilih (0/1/2)?")
	fmt.Scan(&pilih)

	for pilih >= 0 {
		switch pilih {
		case 0:
			Dashboard()
		case 1:
			hapusProyek(insert)
		case 2:
			editProyek(insert)
		default:
			fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

		}

		fmt.Println("===========================")
		fmt.Println("      Pemilik Proyek")
		fmt.Println("---------------------------")
		fmt.Println("0. Kembali ke Menu login")
		fmt.Println("1. Menghapus Proyek")
		fmt.Println("2. Mengedit Proyek")
		fmt.Println("===========================")
		fmt.Println("Pilih (0/1/2)?")
		fmt.Scan(&pilih)

		if pilih < 0 || pilih > 2 {
			fmt.Println("Pilihlah sesuai nomor yang tertera")
		}
	}
	fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

}

func hapusProyek(insert *Proyek) {
	// I.S. : User memilih untuk menghapus proyek
	// F.S. : Proyek dengan nama tertentu dihapus dari array jika ditemukan
	var namaCari string
	var i, low, high, mid, pos int
	var found bool

	fmt.Println("Masukkan nama proyek yang ingin dihapus:")
	fmt.Scan(&namaCari)

	insertionSortByNama(insert)

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
	pass = 1

	for pass <= jumlah-1 {
		i = pass
		temp := insert[pass]
		for i > 0 && temp.namaProyek < insert[i-1].namaProyek {
			insert[i] = insert[i-1]
			i = i - 1
		}
		insert[i] = temp
		pass = pass + 1
	}

}

func editProyek(insert *Proyek) {
	// I.S. : User memilih proyek dan bagian mana yang ingin diubah
	// F.S. : Nilai atribut proyek berubah sesuai input user
	var pilih int
	var baru, targetBaru int
	var namaBaru, deskripsiBaru string

	insertionSortByNama(insert)
	DisplayProjek(insert)

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
}

func insertionSortByJumInvestor(insert *Proyek) {
	// I.S. : Data proyek belum terurut berdasarkan jumlah investor
	// F.S. : Data proyek terurut berdasarkan jumlah investor secara ascending
	var pass, i int
	var total Penggalangan
	pass = 1
	for pass <= total.jumlahInvestor-1 {
		i = pass
		temp := insert[pass]
		for i > 0 && temp.jumlahInvestor < insert[i-1].jumlahInvestor {
			insert[i] = insert[i-1]
			i = i - 1
		}
		insert[i] = temp
		pass = pass + 1
	}
}

//-----------------------------------------------------------------------------------------------------------------------------------------------------------

func cariProyek(data *Proyek) {
	// I.S. : User ingin mencari proyek berdasarkan nama atau kategori
	// F.S. : Menampilkan proyek yang sesuai dengan kategori atau nama jika ditemukan
	var pilih int

	fmt.Println("===========================")
	fmt.Println("          Investor")
	fmt.Println("---------------------------")
	fmt.Println()

	if jumlah == 0 {
		fmt.Println("Maaf, Belum ada proyek yang tersedia")
	}

	fmt.Println("Ingin mencari Proyek berdasarkan apa?")
	fmt.Println("1. Kategori")
	fmt.Println("2. Nama")
	fmt.Println("Pilih (1/2)?")

	fmt.Scan(&pilih)
	for pilih != 0 {
		switch pilih {
		case 1:
			cariProyekByKategori(*data)
			return
		case 2:
			cariProyekByNama(*data)
			return
		default:
			fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

		}
	}

}

func cariProyekByNama(data Proyek) {
	// I.S. : User memasukkan nama proyek yang ingin dicari
	// F.S. : Menampilkan informasi proyek jika ditemukan, atau notifikasi jika tidak ditemukan
	var nama string
	fmt.Println("Masukkan nama yang ingin dicari: ")
	fmt.Scan(&nama)
	index := binarySearch(data, nama)
	if index != -1 {
		fmt.Println("Proyek ditemukan:")
		fmt.Println("Nama Proyek : ", data[index].namaProyek)
		fmt.Println("Deskripsi Proyek : ", data[index].deskripsi)
		fmt.Println("Target Pendanaan : ", data[index].targetDana)
	} else {
		fmt.Println("Proyek tidak ditemukan.")
	}
}

func binarySearch(insert Proyek, x string) int {
	// I.S. : Data proyek belum dicari
	// F.S. : Mengembalikan indeks dari proyek yang dicari jika ditemukan, -1 jika tidak
	insertionSortByNama(&insert)
	low, high := 0, len(insert)-1
	for low <= high {
		mid := (low + high) / 2
		if insert[mid].namaProyek == x {
			return mid
		} else if insert[mid].namaProyek < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1

}

func cariProyekByKategori(data Proyek) {
	// I.S. : User memasukkan kategori yang ingin dicari
	// F.S. : Menampilkan proyek dengan kategori tersebut jika ditemukan
	var kategori string
	fmt.Println("Masukkan kategori yang ingin dicari: ")
	fmt.Println("Note: Bila terdapat lebih dari satu kata, penulisan tidak di spasi")
	fmt.Scan(&kategori)
	index := seqSeach(data, kategori)
	if index != -1 {
		fmt.Println("Proyek ditemukan:")
		fmt.Println("Nama Proyek : ", data[index].namaProyek)
		fmt.Println("Deskripsi Proyek : ", data[index].deskripsi)
		fmt.Println("Target Pendanaan : ", data[index].targetDana)
	} else {
		fmt.Println("Proyek dengan kategori tersebut tidak ditemukan.")
	}
}

func seqSeach(insert Proyek, x string) int {
	// I.S. : Pencarian belum dilakukan
	// F.S. : Mengembalikan indeks proyek dengan kategori yang dicari, -1 jika tidak ditemukan
	for i := 0; i < jumlah; i++ {
		if insert[i].kategori == x {
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
	fmt.Println()
	fmt.Println("---------------------------")
	fmt.Println("0. Kembali ke menu Utama")
	fmt.Println("1. Mengubah Dana yang di investasikan")
	fmt.Println("2. Menambah Dana yang ingin di investasikan")
	fmt.Println("3. Melihat Progress Proyek")
	fmt.Println("Pilih (0/1/2/3)?")
	fmt.Println("===========================")
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
		fmt.Println("===========================")
		fmt.Println("         Investor")
		fmt.Println("---------------------------")
		fmt.Println("0. Kembali ke menu Utama")
		fmt.Println("1. Mengubah Dana yang di investasikan")
		fmt.Println("2. Menambah Dana yang ingin di investasikan")
		fmt.Println("3. Melihat Progress Proyek")
		fmt.Println("4. Mencari Proyek")
		fmt.Println("Pilih (0/1/2/3/4)?")
		fmt.Println("===========================")
		fmt.Scan(&pilih)
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
	for i := 0; i < jumlah; i++ {
		if namaCari != insert[i].namaProyek {
			fmt.Println("Nama proyek tidak valid, masukan kembali nama proyek yang ingin Anda ubah dananya:")
			fmt.Scan(&namaCari)
		}
	}

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
	
	fmt.Println("Silahkan pilih Proyek keberapa yang ingin Anda investasikan") // nanti di sambung sama print dana invest terus input
	fmt.Scan(&pilih)
	if pilih < 1 || pilih > jumlah {
		fmt.Println("Pilihlah sesuai nomor yang tertera")
	}

	fmt.Println("Masukan jumlah dana yang ingin anda investasikann : ")
	fmt.Println("Note : masukan dalam bilangan bulat")
	fmt.Scan(&danaInvest)

	insert[pilih-1].jumlahDana += danaInvest
	insert[pilih-1].jumlahInvestor++

}

func selectionSortByJumDana(insert *Proyek) {
	// I.S. : Data proyek belum diurutkan berdasarkan jumlah dana
	// F.S. : Data proyek terurut secara descending berdasarkan jumlah dana
	var pass, idx, i int

	pass = 1
	for pass <= jumlah-1 {
		idx = pass - 1
		i = pass

		for i < jumlah {
			if insert[idx].jumlahDana < insert[i].jumlahDana {
				idx = i
			}
			i++
		}
		temp := insert[pass-1]
		insert[pass-1] = insert[idx]
		insert[idx] = temp
	}

}

//---------------------------------------------------------------------------------------------------------------------------------
//tampilan

func DisplayProjek(insert *Proyek) {
	// I.S. : Dipanggil untuk menampilkan semua proyek
	// F.S. : Menampilkan informasi proyek satu per satu

	var i int
	fmt.Println("Berikut adalah proyek yang tersedia :")

	for i = 0; i < jumlah; i++ {
		fmt.Println("Proyek ", i+1)
		fmt.Println("Nama Proyek :", insert[i].namaProyek)
		fmt.Println("Deskripsi proyek : ", insert[i].deskripsi)
		fmt.Println("Target Dana :", insert[i].targetDana)
		fmt.Println("Dana Terkumpul :", insert[i].jumlahDana)
	}

}

func DisplayProjekTercapai(insert *Proyek) {
	// I.S. : Dipanggil untuk menampilkan proyek yang mencapai target pendanaan
	// F.S. : Menampilkan proyek-proyek yang telah mencapai target
	var i int
	fmt.Println("Berikut adalah proyek yang telah mencapai target ppendanaan :")

	insertionSortByJumInvestor(insert)
	for i = 0; i < jumlah; i++ {
		if insert[i].jumlahDana >= insert[i].targetDana {
			fmt.Println("Nama Proyek :", insert[i].namaProyek)
			fmt.Println("Deskripsi proyek : ", insert[i].deskripsi)
			fmt.Println("Target Dana :", insert[i].targetDana)
			fmt.Println("Dana Terkumpul :", insert[i].jumlahDana)

		} else {
			fmt.Println("Maaf untuk saat ini belum ada proyek yang telah mencapai target")
		}
	}
}
