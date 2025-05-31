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
	Dashboard()
}

func Dashboard() {
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
			fmt.Print()
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
	fmt.Print("Terimakasih telah menggunakan aplikasi ini") // masih eror
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
func pemilik(data *Proyek) {
	var pilih int
	var add Proyek
	fmt.Println("===========================")
	fmt.Println("      Pemilik Proyek")
	fmt.Println("---------------------------")
	fmt.Println("0. Kembali ke menu login")
	fmt.Println("1. Membuat Proyek")
	fmt.Println("2. Mencari Proyek")
	fmt.Println("3. Mengubah Proyek")
	fmt.Println("4. Melihat proyek yang telah mancapai target pendanaan")
	fmt.Println("===========================")
	fmt.Println("Pilih (0/1/2)?")

	fmt.Scan(&pilih)
	if pilih < 0 || pilih > 4 {
		fmt.Println("Pilihlah sesuai nomor yang tertera")
	}
	// fmt.Print(pilih)

	for pilih != 0 {
		fmt.Print("aa")
		switch pilih {
		// case 0:
		// 	Dashboard()
		case 1:
			buatProyek(data)
		case 2:
			progressProyek(data)
		case 3:
			ubahProyek(data, pilih)
		case 4:
			DisplayProjekTercapai(&add)
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
		fmt.Println("Pilih (0/1/2)?")

		fmt.Scan(&pilih)
		if pilih < 0 || pilih > 4 {
			fmt.Println("Pilihlah sesuai nomor yang tertera")
		}
	}

}

func buatProyek(insert *Proyek) {
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
			fmt.Scan(&insert[i].targetDana) //ga keluar minta inputan
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
	if pilih < 1 || pilih > 3 {
		fmt.Println("Pilihlah sesuai nomor yang tertera")
	}

	if pilih == 1 {
		insertionSortByNama(insert)
	} else if pilih == 2 {
		insertionSortByJumInvestor(insert)
	} else {
		selectionSortByJumDana(insert)
	}

	// for i := 0; i < jumlah; i++ {
	// 	fmt.Printf("%d. %s - %s (target : %d)\n", i+1, insert[i].namaProyek, insert[i].deskripsi, insert[i].targetDana)

	// 	fmt.Println("1. Jumlah Dana : ", insert[i].jumlahDana)
	// 	fmt.Println("2. Jumlah Investor : ", insert[i].jumlahInvestor)
	// 	fmt.Println("0. Kembali ke Menu Utama")
	// 	fmt.Println("===========================")

	// fmt.Scan(&pilih)
	// if pilih < 1 || pilih > jumlah {
	// 	fmt.Println("Pilihlah sesuai nomor yang tertera")
	// }

	// if pilih == 0 {
	// 	Dashboard()
	// } else {
	// 	fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

	// }
	// fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

	//}
}

func ubahProyek(insert *Proyek, pilih int) {
	var n int

	fmt.Println("===========================")
	fmt.Println("      Pemilik Proyek")
	fmt.Println("---------------------------")
	fmt.Println("0. Kembali ke Menu login")
	fmt.Println("1. Menghapus Proyek")
	fmt.Println("2. Mengedit Proyek")
	fmt.Println("===========================")

	for pilih == 0 {
		switch pilih {
		case 0:
			Dashboard()
		case 1:
			hapusProyek(insert, n)
		case 2:
			editProyek(insert, n)
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

func editProyek(insert *Proyek, pilih int) {
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
func insertionSortByJumInvestor(insert *Proyek) {
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
		case 2:
			cariProyekByNama(*data)
		default:
			fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

		}

		// if pilih < 1 || pilih > 3 {
		// 	fmt.Println("Pilihlah sesuai nomor yang tertera")
		// }

		// if pilih == 1 {
		// 	insertionSortByNama(insert)
		// } else if pilih == 2 {
		// 	insertionSortByJumInvestor(insert)
		// } else {
		// 	selectionSortByJumDana(insert)
		// }
	}

}

// func bacaNama(input *string, n int){
// 	for i:= 0; i < n; i++{
// 		fmt.Scan(&input)
// 	}
// }
func cariProyekByNama(data Proyek) {
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
	for i := 0; i < jumlah; i++ {
		if insert[i].kategori == x {
			return i
		}
	}
	return -1
}

func investor(data *Proyek) {
	var pilih int

	fmt.Println("===========================")
	fmt.Println("          Investor")
	fmt.Println("---------------------------")
	fmt.Println()
	cariProyek(data)
	fmt.Println()
	fmt.Println("---------------------------")
	fmt.Println("0. Kembali ke menu Utama")
	fmt.Println("1. Mengubah Dana yang di investasikan")
	fmt.Println("2. Menambah Dana yang ingin di investasikan")
	fmt.Println("Pilih (0/1/2)?")
	fmt.Println("===========================")
	fmt.Scan(&pilih)

	for pilih != 0 {
		switch pilih {
		case 0:
			Dashboard()
		case 1:
			ubahDanaInvestor(data)
		case 2:
			tambahDanaInvestor(data, pilih)
		default:
			fmt.Println("Pilihan tidak valid, Silahkan pilih sesuai petunjuk")

		}
		fmt.Println("===========================")
		fmt.Println("         Investor")
		fmt.Println("---------------------------")
		fmt.Println("0. Kembali ke menu Utama")
		fmt.Println("1. Mengubah Dana yang di investasikan")
		fmt.Println("2. Menambah Dana yang ingin di investasikan")
		fmt.Println("Pilih (0/1/2)?")
		fmt.Println("===========================")
		fmt.Scan(&pilih)
	}
}

func ubahDanaInvestor(insert *Proyek) {
	var namaCari string
	var danaBaru int
	var ditemukan bool
	var idx int

	selectionSortByJumDana(insert)
	DisplayProjek(insert)

	fmt.Println("Masukkan nama proyek yang ingin Anda ubah dananya:")
	fmt.Scan(&namaCari)

	// fmt.Print("Sblm sort")
	// DisplayProjekTercapai(insert)
	// fmt.Println()
	// insertionSortByNama(insert)
	// fmt.Print("Stlh sort")
	// DisplayProjekTercapai(insert)

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

func tambahDanaInvestor(insert *Proyek, pilih int) {
	var danaInvest int

	selectionSortByJumDana(insert)
	DisplayProjek(insert)

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

	insert[pilih-1].jumlahDana += danaInvest
	insert[pilih-1].jumlahInvestor++

}

func selectionSortByJumDana(insert *Proyek) {
	//descending
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
	var i int
	fmt.Println("Berikut adalah proyek yang tersedia :")

	for i = 0; i < jumlah; i++ {
		fmt.Println("Nama Proyek :", insert[i].namaProyek)
		fmt.Println("Target Dana :", insert[i].targetDana)
		fmt.Println("Dana Terkumpul :", insert[i].jumlahDana)
	}

}
func DisplayProjekTercapai(insert *Proyek) {
	var i int
	fmt.Println("Berikut adalah proyek yang telah mencapai target ppendanaan :")

	insertionSortByJumInvestor(insert)
	for i = 0; i < jumlah; i++ {
		if insert[i].jumlahDana >= insert[i].targetDana {
			fmt.Println("Nama Proyek :", insert[i].namaProyek)
			fmt.Println("Target Dana :", insert[i].targetDana)
			fmt.Println("Dana Terkumpul :", insert[i].jumlahDana)

		} else {
			fmt.Println("Maaf untuk saat ini belum ada proyek yang telah mencapai target")
		}
	}
}
