package main

import (
	"fmt"
)

// Function untuk menghitung rata-rata skor, dengan menerima parameter berupa array
func hitungRataRata(array [3]int) int {

	//deklrasi variable sum untuk menampung nilai rata-rata
	sum := 0

	//looping array yang diterima
	for _, nilai := range array {
		//menambahkan value hasil looping kedalam vaariable sum
		sum += nilai
	}

	//melakukan return berupa int
	return sum / len(array)
}

// function untuk verifikasi apakah skor nya kurang dari 100
func cekSkor(array [3]int) bool {
	//data yang diterima di lempar kembali ke function hitungRataRata
	skorRata := hitungRataRata(array)

	//setelah mendapatkan nilai return, maka tinggal lakukan if-else
	if skorRata >= 100 {
		return true
	} else {
		return false
	}
}

func Data1() {
	//inisialisasi variable dengan nilai array yang telah ditetapkan
	lumbaLumba := [3]int{96, 108, 89}
	koala := [3]int{88, 91, 110}

	//panggil function hitungRataRata
	rataRataLumba := hitungRataRata(lumbaLumba)
	rataRataKoala := hitungRataRata(koala)

	fmt.Println("====Data 1====")

	//Lakukan perbandingan nilai
	if rataRataLumba > rataRataKoala {
		fmt.Println("Skor Rata-Rata dimenangkan oleh Lumba-Lumba, dengan skor", rataRataLumba)
	} else if rataRataLumba < rataRataKoala {
		fmt.Println("Skor Rata-Rata dimenangkan oleh Koala, dengan skor", rataRataKoala)
	} else {
		fmt.Println("Skor Rata-Rata Seri, dengan skor", rataRataLumba)
	}

}

func dataBonus1() {
	//Sama aja kaya data1
	lumbaLumba := [3]int{97, 112, 101}
	koala := [3]int{109, 95, 123}

	fmt.Println("====Data Bonus 1====")

	//bedanya disini kita lakukan verifikasi dlu apakah nilai skor nya lebih dari 100 atau tidak, jika hasil nya true, maka proses akan dilanjutkan
	if cekSkor(lumbaLumba) && cekSkor(koala) {
		rataRataLumba := hitungRataRata(lumbaLumba)
		rataRataKoala := hitungRataRata(koala)

		if rataRataLumba > rataRataKoala {
			fmt.Println("Skor Rata-Rata dimenangkan oleh Lumba-Lumba, dengan skor", rataRataLumba)
		} else if rataRataLumba < rataRataKoala {
			fmt.Println("Skor Rata-Rata dimenangkan oleh Koala, dengan skor", rataRataKoala)
		} else {
			fmt.Println("Skor Rata-Rata Seri, dengan skor", rataRataLumba)
		}

	} else {
		fmt.Println("Skor kurang dari 100")
	}

}

func dataBonus2() {
	//sama seperti function diatas
	lumbaLumba := [3]int{97, 112, 101}
	koala := [3]int{109, 95, 106}

	fmt.Println("====Data Bonus 2====")

	if cekSkor(lumbaLumba) && cekSkor(koala) {
		rataRataLumba := hitungRataRata(lumbaLumba)
		rataRataKoala := hitungRataRata(koala)

		if rataRataLumba > rataRataKoala {
			fmt.Println("Skor Rata-Rata dimenangkan oleh Lumba-Lumba, dengan skor", rataRataLumba)
		} else if rataRataLumba < rataRataKoala {
			fmt.Println("Skor Rata-Rata dimenangkan oleh Koala, dengan skor", rataRataKoala)
		} else {
			fmt.Println("Skor Rata-Rata Seri, dengan skor", rataRataLumba)
		}
	} else {
		fmt.Println("Skor kurang dari 100")
	}

}

func main() {
	//memanggil function nya
	Data1()
	dataBonus1()
	dataBonus2()
}
