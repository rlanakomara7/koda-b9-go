package main

import (
	"fmt"
	"koda-b9-go/internal/looping"
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
	fmt.Print(looping.Window(5))

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
