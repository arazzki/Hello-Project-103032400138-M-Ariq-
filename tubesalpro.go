// ini komentar untuk tugas latihan Git
package main 
import "fmt"
func menu() {
	fmt.Println("\n=== MENU SEA GAMES MANAGER ===")
	fmt.Println("1. Tambah Negara")
	fmt.Println("2. Edit Nama Negara")
	fmt.Println("3. Hapus Negara")
	fmt.Println("4. Edit Data Medali")
	fmt.Println("5. Hapus Medali (reset ke 0)")
	fmt.Println("6. Tampilkan Peringkat")
	fmt.Println("7. Keluar")
	fmt.Print("Pilih menu: ")
}	

func editMedali(dataN DataNegara, dataM *DataMedali, n int) {
	var nama string
	var e, p, pr int
	fmt.Print("Nama negara: ")
	fmt.Scan(&nama)
	idx := cariNegara(dataN, n, nama)
	if idx != -1 {
		fmt.Print("Emas: ")
		fmt.Scan(&e)
		fmt.Print("Perak: ")
		fmt.Scan(&p)
		fmt.Print("Perunggu: ")
		fmt.Scan(&pr)
		dataM[idx] = Medali{e, p, pr}
		fmt.Println("Medali berhasil diubah.")
	} else {
		fmt.Println("Negara tidak ditemukan.")
	}
}

func hapusMedali(dataN DataNegara, dataM *DataMedali, n int) {
	var nama string
	fmt.Print("Nama negara: ")
	fmt.Scan(&nama)
	idx := cariNegara(dataN, n, nama)
	if idx != -1 {
		dataM[idx] = Medali{0, 0, 0}
		fmt.Println("Data medali berhasil direset.")
	} else {
		fmt.Println("Negara tidak ditemukan.")
	}
}

const MAX = 100
type Negara struct {
	Nama string
}
type Medali struct {
	Emas, Perak, Perunggu int
}
type DataNegara [MAX]Negara
type DataMedali [MAX]Medali

func cariNegara(data DataNegara, n int, nama string) int {
	for i := 0; i < n; i++ {
		if data[i].Nama == nama {
			return i
		}
	}
	return -1
}

func tambahNegara(dataN *DataNegara, dataM *DataMedali, n *int) {
	var nama string
	fmt.Print("Nama Negara: ")
	fmt.Scan(&nama)
	if cariNegara(*dataN, *n, nama) == -1 {
		dataN[*n].Nama = nama
		dataM[*n] = Medali{0, 0, 0}
		*n++
		fmt.Println("Negara berhasil ditambahkan.")
	} else {
		fmt.Println("Negara sudah terdaftar.")
	}
}

func editNegara(dataN *DataNegara, n int) {
	var lama, baru string
	fmt.Print("Nama negara yang ingin diedit: ")
	fmt.Scan(&lama)
	idx := cariNegara(*dataN, n, lama)
	if idx != -1 {
		fmt.Print("Nama baru: ")
		fmt.Scan(&baru)
		dataN[idx].Nama = baru
		fmt.Println("Nama negara berhasil diubah.")
	} else {
		fmt.Println("Negara tidak ditemukan.")
	}
}

func hapusNegara(dataN *DataNegara, dataM *DataMedali, n *int) {
	var nama string
	fmt.Print("Nama negara yang ingin dihapus: ")
	fmt.Scan(&nama)
	idx := cariNegara(*dataN, *n, nama)
	if idx != -1 {
		for i := idx; i < *n-1; i++ {
			dataN[i] = dataN[i+1]
			dataM[i] = dataM[i+1]
		}
		*n--
		fmt.Println("Negara berhasil dihapus.")
	} else {
		fmt.Println("Negara tidak ditemukan.")
	}
}

func insertionSort(dataN *DataNegara, dataM *DataMedali, n int) {
	for pass := 1; pass < n; pass++ {
		tempN := dataN[pass]
		tempM := dataM[pass]
		i := pass

		for i > 0 && (
			tempM.Emas > dataM[i-1].Emas ||
				(tempM.Emas == dataM[i-1].Emas && tempM.Perak > dataM[i-1].Perak) ||
				(tempM.Emas == dataM[i-1].Emas && tempM.Perak == dataM[i-1].Perak && tempM.Perunggu > dataM[i-1].Perunggu)) {

			dataN[i] = dataN[i-1]
			dataM[i] = dataM[i-1]
			i--
		}

		dataN[i] = tempN
		dataM[i] = tempM
	}
}

func tampilkan(dataN DataNegara, dataM DataMedali, n int) {
	insertionSort(&dataN, &dataM, n)
	fmt.Println("\nPeringkat Negara:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s | Emas: %d, Perak: %d, Perunggu: %d\n",
			i+1, dataN[i].Nama, dataM[i].Emas, dataM[i].Perak, dataM[i].Perunggu)
	}
}

func main() {
	var dataNegara DataNegara
	var dataMedali DataMedali
	var nData, pilihan int

	for pilihan != 7 {
		menu()
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			tambahNegara(&dataNegara, &dataMedali, &nData)
		case 2:
			editNegara(&dataNegara, nData)
		case 3:
			hapusNegara(&dataNegara, &dataMedali, &nData)
		case 4:
			editMedali(dataNegara, &dataMedali, nData)
		case 5:
			hapusMedali(dataNegara, &dataMedali, nData)
		case 6:
			tampilkan(dataNegara, dataMedali, nData)
		case 7:
			fmt.Println("Terima kasih!")
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}