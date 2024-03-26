package main

import (
	"fmt"
)

// hitung BMI nya dengan menerima parameter berat dan tinggi
func hitungBMI(berat, tinggi float64) float64 {
	//return value nya berupa float64
	return berat / (tinggi * tinggi)
}

// function untuk pengecekan bmi antara jhon dan mark dengan hasil boolean
func markHigherBMI(beratMark, tinggiMark, beratJohn, tinggiJohn float64) bool {

	//memanggil fungsi hitungBMI dengan mengirim parameter
	bmiMark := hitungBMI(beratMark, tinggiMark)
	bmiJohn := hitungBMI(beratJohn, tinggiJohn)

	//melakukan perbandingan apakah nilai bmiMark lebih besar dari bmiJhon, jika iya, return true
	return bmiMark > bmiJohn
}

func main() {
	// Data 1
	beratMark1 := 78.0  // kg
	tinggiMark1 := 1.69 // m
	beratJohn1 := 92.0  // kg
	tinggiJohn1 := 1.95 // m
	markHigherBMI1 := markHigherBMI(beratMark1, tinggiMark1, beratJohn1, tinggiJohn1)
	fmt.Println("Data 1: Mark memiliki BMI lebih tinggi dari John: ", markHigherBMI1)

	// Data 2
	beratMark2 := 95.0  // kg
	tinggiMark2 := 1.88 // m
	beratJohn2 := 85.0  // kg
	tinggiJohn2 := 1.76 // m
	markHigherBMI2 := markHigherBMI(beratMark2, tinggiMark2, beratJohn2, tinggiJohn2)
	fmt.Println("Data 2: Mark memiliki BMI lebih tinggi dari John: ", markHigherBMI2)
}
