package main

import (
	"fmt"
	luaskelilingpersegipanjang "koda-b9-go/internal/luasKelilingPersegiPanjang"
)

func main() {
	// manifest
	var hello string = "Hellow"
	fmt.Println(hello)

	// inference
	world := "World"
	fmt.Println(world)

	greets("rama")

	//panggil LUAS & KELILING
	Luas, Keliling := luaskelilingpersegipanjang.LuasKeliling(5, 10)
	fmt.Println("Luas", Luas)
	fmt.Println("Keliling", Keliling)

	//PANGGIL WINDOW
	fmt.Print(window(5))

	//PANGGIL SISIPANGKA ------- MINITASK 3
	numbers := []int{50, 75, 66, 20, 32, 90}
	hasil := sisipAngka(numbers)
	fmt.Println(hasil)

	for _, result := range numbers {
		fmt.Println(result)
	}

	// STRUCT VARIABEL DATA
	user := UserData{
		Name:      "Rama Lana Komara",
		Photo:     "rama.png",
		Email:     "rlanakomara7@gmail.com",
		Age:       27,
		Phone:     "089614238447",
		IsMarried: true,
		Education: []Education{{
			Univ:  "Universitas PGRI,",
			Study: "Computer Science",
		}},
	}
	fmt.Println(user)
}

// GREET / SAPA
func greets(name string) {
	fmt.Printf("hello %s", name)
}

//WINDOWS PATTERN DENGAN * MINITASK 2 --------------------

func window(n int) error {
	if n <= 0 {
		return fmt.Errorf(" ")
	}

	for x := 1; x <= n; x++ {
		for y := 1; y <= n; y++ {
			if x == 1 || x == n || y == 1 || y == n {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	return nil
}

// JADI DISINI TUGASNYA MENYISIPKAN NILAI , MENGGUNAKAN SIFAT SLICE YANG REFRENCE
// BELAH DATA SILCE MENGGUNAKAN SLICE PARTIONING
// SISIPKAN ANGKA TERSEBUT MENGGUNAKAN CONCAT / APPEND
// GABUNGKAN LAGI MENJADI DATA AWAL YANG SUDAH DISISIPKAN NILAI TERSEBUT
// DAN TINGGAL SEBUT LAGI 1 PER 1 GUNAKAN FOR LOOP

// func sisipAngka(numbers []int) []int {
// 	toInsert := []int{88}

// 	fmt.Println("Original Slice:", numbers)

// 	numbers = append(numbers, num)

// 	fmt.Println("after appending : " , numbers)
// }

// MENYISIPKAN ANGKA DIANTARA DATA DALAM SLICE , MINITASK 3 ---------------
func sisipAngka(originalSlice []int) []int {
	toInsert := []int{88}

	mid := len(originalSlice) / 2
	return append(originalSlice[:mid], append(toInsert, originalSlice[mid:]...)...)
}

// STRUCT DAT RIWAYAT MINITASK 4 ---------------

type UserData struct {
	Name      string
	Photo     string
	Email     string
	Age       uint8
	Phone     string
	IsMarried bool
	Education []Education
}

type Education struct {
	Univ  string
	Study string
}
