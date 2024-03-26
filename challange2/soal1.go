package main

import (
	"fmt"
)

func hitungRataRata(array [3]int) int {
	sum := 0
	for _, nilai := range array {
		sum += nilai
	}
	return sum / len(array)
}

func cekSkor(array [3]int) bool {
	skorRata := hitungRataRata(array)

	if skorRata < 100 {
		return false
	} else {
		return true
	}
}

func Data1() {
	lumbaLumba := [3]int{96, 108, 89}
	koala := [3]int{88, 91, 110}

	rataRataLumba := hitungRataRata(lumbaLumba)
	rataRataKoala := hitungRataRata(koala)

	fmt.Println("====Data 1====")
	if rataRataLumba > rataRataKoala {
		fmt.Println("Skor Rata-Rata dimenangkan oleh Lumba-Lumba, dengan skor", rataRataLumba)
	} else if rataRataLumba < rataRataKoala {
		fmt.Println("Skor Rata-Rata dimenangkan oleh Koala, dengan skor", rataRataKoala)
	} else {
		fmt.Println("Skor Rata-Rata Seri, dengan skor", rataRataLumba)
	}

}

func dataBonus1() {
	lumbaLumba := [3]int{97, 112, 101}
	koala := [3]int{109, 95, 123}

	fmt.Println("====Data Bonus 1====")
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
	Data1()
	dataBonus1()
	dataBonus2()
}
