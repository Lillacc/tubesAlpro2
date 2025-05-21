package main

import (
	"fmt"
	"strings"
)

const NMAX int = 99

// Struct untuk menyimpan data investasi
type structInvestasi struct {
	id           int
	jenisAset    string
	jumlahDana   float64
	nilaiAwal    float64
	nilaiAkhir   float64
	persenUntung float64
}

// Array untuk menyimpan data investasi
type arrInvestasi [NMAX]structInvestasi

// Function untuk menampilkan header aplikasi
func header() {
	fmt.Println("***============================================***")
	fmt.Println("*** Aplikasi Pengelola Investasi         ***")
	fmt.Println("***============================================***")
}

// Function untuk menampilkan footer aplikasi
func footer() {
	fmt.Println("\n==================================================")
	fmt.Println("==================================================\n")
}

// Function untuk menampilkan menu utama
func menuUtama() int {
	header()
	var inputUser int

	fmt.Println("\n=== Menu Utama ===")
	fmt.Println("1. Laporan Investasi")
	fmt.Println("2. Tambah Data Investasi")
	fmt.Println("3. Edit Data Investasi")
	fmt.Println("4. Hapus Data Investasi")
	fmt.Println("5. Cari Data Investasi")
	fmt.Println("6. Urutkan Data Investasi")
	fmt.Println("0. Keluar")

	fmt.Print("\nMasukkan Menu : ")
	fmt.Scanln(&inputUser)
	footer()
	return inputUser
}

// Function untuk menambahkan data investasi
func tambahInvestasi(dataInvestasi *arrInvestasi, nInvestasi *int) {
	header()
	fmt.Println("\n=== Tambah Data Investasi ===")

	if *nInvestasi == 0 {
		dataInvestasi[*nInvestasi].id = 1
	} else {
		dataInvestasi[*nInvestasi].id = dataInvestasi[*nInvestasi-1].id + 1
	}

	fmt.Print("Jenis Aset (Saham/Obligasi/ReksaDana): ")
	fmt.Scanln(&dataInvestasi[*nInvestasi].jenisAset)

	fmt.Print("Jumlah Dana yang Diinvestasikan: ")
	fmt.Scanln(&dataInvestasi[*nInvestasi].jumlahDana)
	dataInvestasi[*nInvestasi].nilaiAwal = dataInvestasi[*nInvestasi].jumlahDana

	fmt.Print("Nilai Aset Terkini: ")
	fmt.Scanln(&dataInvestasi[*nInvestasi].nilaiAkhir)

	if dataInvestasi[*nInvestasi].nilaiAwal != 0 {
		dataInvestasi[*nInvestasi].persenUntung = (dataInvestasi[*nInvestasi].nilaiAkhir - dataInvestasi[*nInvestasi].nilaiAwal) / dataInvestasi[*nInvestasi].nilaiAwal * 100
	} else {
		dataInvestasi[*nInvestasi].persenUntung = 0.0
	}

	*nInvestasi++

	fmt.Println("\nData investasi berhasil ditambahkan.")
	footer()
}

// Function untuk mengedit data investasi
func editInvestasi(dataInvestasi *arrInvestasi, nInvestasi int) {
	header()
	var inputId int
	fmt.Println("\n=== Edit Data Investasi ===")

	fmt.Print("Masukkan ID Investasi yang akan Diedit: ")
	fmt.Scanln(&inputId)

	indeks := cariInvestasiId(*dataInvestasi, nInvestasi, inputId)

	if indeks != -1 {
		fmt.Println("\nData Investasi Sebelumnya:")
		fmt.Println("ID:", dataInvestasi[indeks].id)
		fmt.Println("Jenis Aset:", dataInvestasi[indeks].jenisAset)
		fmt.Println("Jumlah Dana:", dataInvestasi[indeks].jumlahDana)
		fmt.Println("Nilai Awal:", dataInvestasi[indeks].nilaiAwal)
		fmt.Println("Nilai Akhir:", dataInvestasi[indeks].nilaiAkhir)

		fmt.Print("\nJenis Aset Baru: ")
		fmt.Scanln(&dataInvestasi[indeks].jenisAset)

		fmt.Print("Jumlah Dana Baru: ")
		fmt.Scanln(&dataInvestasi[indeks].jenisAset)
		dataInvestasi[indeks].nilaiAwal = dataInvestasi[indeks].jumlahDana

		fmt.Print("Nilai Aset Terkini Baru: ")
		fmt.Scanln(&dataInvestasi[indeks].nilaiAkhir)

		if dataInvestasi[indeks].nilaiAwal != 0 {
			dataInvestasi[indeks].persenUntung = (dataInvestasi[indeks].nilaiAkhir - dataInvestasi[indeks].nilaiAwal) / dataInvestasi[indeks].nilaiAwal * 100
		} else {
			dataInvestasi[indeks].persenUntung = 0.0
		}

		fmt.Println("\nData investasi berhasil diubah.")

	} else {
		fmt.Println("\nData investasi tidak ditemukan.")
	}
	footer()
}

// Function untuk menghapus data investasi
func hapusInvestasi(dataInvestasi *arrInvestasi, nInvestasi *int) {
	header()
	var inputId int
	fmt.Println("\n=== Hapus Data Investasi ===")

	fmt.Print("Masukkan ID Investasi yang akan Dihapus: ")
	fmt.Scanln(&inputId)

	indeks := cariInvestasiId(*dataInvestasi, *nInvestasi, inputId)

	if indeks != -1 {
		for i := indeks; i < *nInvestasi-1; i++ {
			dataInvestasi[i] = dataInvestasi[i+1]
		}
		*nInvestasi--
		fmt.Println("\nData investasi berhasil dihapus.")
	} else {
		fmt.Println("\nData investasi tidak ditemukan.")
	}
	footer()
}

// Function untuk mencari data investasi (Sequential Search)
func cariInvestasiSequential(dataInvestasi arrInvestasi, nInvestasi int, keyword string, jenisPencarian int) int {
	for i := 0; i < nInvestasi; i++ {
		if jenisPencarian == 1 && strings.Contains(strings.ToLower(dataInvestasi[i].jenisAset), strings.ToLower(keyword)) {
			return i
		} else if jenisPencarian == 2 && strings.ToLower(dataInvestasi[i].jenisAset) == strings.ToLower(keyword) {
			return i
		}
	}
	return -1
}

// Function untuk mencari data investasi (Binary Search)
func cariInvestasiBinary(dataInvestasi arrInvestasi, nInvestasi int, keyword string) int {
	left, right := 0, nInvestasi-1
	for left <= right {
		mid := (left + right) / 2
		if strings.ToLower(dataInvestasi[mid].jenisAset) == strings.ToLower(keyword) {
			return mid
		} else if strings.ToLower(dataInvestasi[mid].jenisAset) < strings.ToLower(keyword) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// Function untuk mencari investasi berdasarkan ID
func cariInvestasiId(dataInvestasi arrInvestasi, nInvestasi int, id int) int {
	for i := 0; i < nInvestasi; i++ {
		if dataInvestasi[i].id == id {
			return i
		}
	}
	return -1
}

// Function untuk menampilkan hasil pencarian
func tampilkanHasilPencarian(dataInvestasi arrInvestasi, indeks int) {
	if indeks != -1 {
		fmt.Println("\nData Investasi Ditemukan:")
		fmt.Println("ID:", dataInvestasi[indeks].id)
		fmt.Println("Jenis Aset:", dataInvestasi[indeks].jenisAset)
		fmt.Println("Jumlah Dana:", dataInvestasi[indeks].jumlahDana)
		fmt.Println("Nilai Awal:", dataInvestasi[indeks].nilaiAwal)
		fmt.Println("Nilai Akhir:", dataInvestasi[indeks].nilaiAkhir)
		fmt.Printf("Persentase Keuntungan: %.2f%%\n", dataInvestasi[indeks].persenUntung)
	} else {
		fmt.Println("\nData investasi tidak ditemukan.")
	}
}

// Function untuk menu pencarian investasi
func menuCariInvestasi(dataInvestasi arrInvestasi, nInvestasi int) {
	header()
	var inputUser int
	var keyword string

	fmt.Println("\n=== Cari Data Investasi ===")
	fmt.Println("1. Cari Berdasarkan Jenis Aset (Sequential)")
	fmt.Println("2. Cari Berdasarkan Jenis Aset (Binary) - Data harus terurut")
	fmt.Println("0. Kembali")

	fmt.Print("Masukkan Menu : ")
	fmt.Scanln(&inputUser)

	if inputUser == 1 {
		fmt.Print("Masukkan Jenis Aset: ")
		fmt.Scanln(&keyword)
		indeks := cariInvestasiSequential(dataInvestasi, nInvestasi, keyword, 1)
		tampilkanHasilPencarian(dataInvestasi, indeks)
	} else if inputUser == 2 {
		fmt.Print("Masukkan Jenis Aset: ")
		fmt.Scanln(&keyword)
		indeks := cariInvestasiBinary(dataInvestasi, nInvestasi, keyword)
		tampilkanHasilPencarian(dataInvestasi, indeks)
	} else if inputUser == 0 {
	} else {
		fmt.Println("Input Invalid")
	}
	footer()
}

// Function untuk mengurutkan data investasi (Selection Sort)
func urutkanInvestasiSelection(dataInvestasi *arrInvestasi, nInvestasi int, jenisUrut int) {
	for i := 0; i < nInvestasi-1; i++ {
		indeksMax := i
		for j := i + 1; j < nInvestasi; j++ {
			if jenisUrut == 1 && dataInvestasi[j].nilaiAkhir > dataInvestasi[indeksMax].nilaiAkhir {
				indeksMax = j
			} else if jenisUrut == 2 && dataInvestasi[j].persenUntung > dataInvestasi[indeksMax].persenUntung {
				indeksMax = j
			}
		}
		dataInvestasi[i], dataInvestasi[indeksMax] = dataInvestasi[indeksMax], dataInvestasi[i]
	}
}

// Function untuk mengurutkan data investasi (Insertion Sort)
func urutkanInvestasiInsertion(dataInvestasi *arrInvestasi, nInvestasi int, jenisUrut int) {
	for i := 1; i < nInvestasi; i++ {
		temp := dataInvestasi[i]
		j := i - 1
		for j >= 0 {
			if jenisUrut == 1 && temp.nilaiAkhir > dataInvestasi[j].nilaiAkhir {
				dataInvestasi[j+1] = dataInvestasi[j]
			} else if jenisUrut == 2 && temp.persenUntung > dataInvestasi[j].persenUntung {
				dataInvestasi[j+1] = dataInvestasi[j]
			} else {
				break
			}
			j--
		}
		dataInvestasi[j+1] = temp
	}
}

// Function untuk menu pengurutan investasi
func menuUrutInvestasi(dataInvestasi *arrInvestasi, nInvestasi int) {
	header()
	var inputUser, jenisUrut int

	fmt.Println("\n=== Urutkan Data Investasi ===")
	fmt.Println("1. Urutkan Berdasarkan Nilai Investasi (Selection Sort)")
	fmt.Println("2. Urutkan Berdasarkan Persentase Keuntungan (Selection Sort)")
	fmt.Println("3. Urutkan Berdasarkan Nilai Investasi (Insertion Sort)")
	fmt.Println("4. Urutkan Berdasarkan Persentase Keuntungan (Insertion Sort)")
	fmt.Println("0. Kembali")

	fmt.Print("Masukkan Menu : ")
	fmt.Scanln(&inputUser)

	if inputUser >= 1 && inputUser <= 4 {
		if inputUser <= 2 {
			jenisUrut = inputUser
			urutkanInvestasiSelection(dataInvestasi, nInvestasi, jenisUrut)
		} else {
			jenisUrut = inputUser - 2
			urutkanInvestasiInsertion(dataInvestasi, nInvestasi, jenisUrut)
		}
		fmt.Println("\nData berhasil diurutkan.")
		tampilkanLaporanInvestasi(*dataInvestasi, nInvestasi)
	} else if inputUser == 0 {
	} else {
		fmt.Println("Input Invalid")
	}
	footer()
}

// Function untuk menampilkan laporan investasi
func tampilkanLaporanInvestasi(dataInvestasi arrInvestasi, nInvestasi int) {
	header()
	fmt.Println("\n=== Laporan Investasi ===")
	if nInvestasi == 0 {
		fmt.Println("Belum ada data investasi yang tersedia.")
		fmt.Println("--------------------------------------------------------------------------------------------\n")
		footer()
		return
	}
	fmt.Println("--------------------------------------------------------------------------------------------")
	fmt.Printf("| %-4s | %-15s | %-12s | %-12s | %-12s | %-17s |\n", "ID", "Jenis Aset", "Dana Awal", "Nilai Awal", "Nilai Akhir", "Persentase Untung")
	fmt.Println("--------------------------------------------------------------------------------------------")
	for i := 0; i < nInvestasi; i++ {
		fmt.Printf("| %-4d | %-15s | %-12.2f | %-12.2f | %-12.2f | %-16.2f%% |\n", dataInvestasi[i].id, dataInvestasi[i].jenisAset, dataInvestasi[i].jumlahDana, dataInvestasi[i].nilaiAwal, dataInvestasi[i].nilaiAkhir, dataInvestasi[i].persenUntung)
	}
	fmt.Println("--------------------------------------------------------------------------------------------\n")
	footer()
}

// Fungsi menuLogin yang sudah dilanjutkan
func menuLogin() int {
	header()
	const max int = 5
	var (
		username string
		password int
		attempts int
	)

	fmt.Println("\n=== Halaman Login ===")

	for attempts < max {
		fmt.Printf("Percobaan %d dari %d\n", attempts+1, max)
		fmt.Print("Username : ")
		fmt.Scanln(&username)
		fmt.Print("Password : ")
		fmt.Scanln(&password)

		processedUsername := strings.TrimSpace(strings.ToLower(username))

		if processedUsername == "admin" && password == 123123 {
			fmt.Println("\nLogin Berhasil!")
			footer()
			return 1
		} else {
			fmt.Println("Username atau Password Salah.")
			attempts++
		}
	}

	fmt.Println("\nAnda telah melebihi batas percobaan login.")
	footer()
	return -1
}

func main() {
	var (
		dataInvestasi arrInvestasi
		nInvestasi    int
		inputUser     int
		statusLogin   int
	)

	nInvestasi = 0
	statusLogin = -1

	for statusLogin != 0 {
		if statusLogin == -1 {
			statusLogin = menuLogin()
			if statusLogin == -1 {
				fmt.Println("Aplikasi akan ditutup.")
				return
			}
		} else {
			inputUser = menuUtama()

			if inputUser == 1 {
				tampilkanLaporanInvestasi(dataInvestasi, nInvestasi)
			} else if inputUser == 2 {
				tambahInvestasi(&dataInvestasi, &nInvestasi)
			} else if inputUser == 3 {
				editInvestasi(&dataInvestasi, nInvestasi)
			} else if inputUser == 4 {
				hapusInvestasi(&dataInvestasi, &nInvestasi)
			} else if inputUser == 5 {
				menuCariInvestasi(dataInvestasi, nInvestasi)
			} else if inputUser == 6 {
				menuUrutInvestasi(&dataInvestasi, nInvestasi)
			} else if inputUser == 0 {
				fmt.Println("\nTerima kasih telah menggunakan aplikasi ini.")
				statusLogin = 0
			} else {
				fmt.Println("\nInput Invalid.\n")
			}
		}
	}
	fmt.Println("Terima kasih telah menggunakan Aplikasi Pengelola Investasi.")
}
