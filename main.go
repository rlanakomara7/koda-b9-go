package main

import (
	"fmt"

	checkout "koda-b9-go/internal/interfacee"
	"koda-b9-go/internal/looping"
	luaskelilingpersegipanjang "koda-b9-go/internal/luasKelilingPersegiPanjang"
	"koda-b9-go/internal/method"
	panicrecover "koda-b9-go/internal/panicRecover"
	"koda-b9-go/internal/slice"
	"koda-b9-go/internal/userdata"
)

func main() {
	var pilihan int

	for {
		fmt.Println("\n==============================")
		fmt.Println("       MENU MINITASK GO")
		fmt.Println("==============================")
		fmt.Println("1. Hitung luas dan keliling")
		fmt.Println("2. Tampilkan pola kotak")
		fmt.Println("3. Sisipkan angka pada slice")
		fmt.Println("4. Tampilkan data pengguna")
		fmt.Println("5. Procces File Path")
		fmt.Println("6. Method dan constructor Person")
		fmt.Println("7. Sistem Checkout Interface")
		fmt.Println("0. Keluar")
		fmt.Println("==============================")
		fmt.Print("Masukkan pilihan: ")

		fmt.Scan(&pilihan)
		fmt.Println()

		switch pilihan {
		case 1:
			// Minitask luas dan keliling
			luas, keliling :=
				luaskelilingpersegipanjang.LuasKeliling(5, 10)

			fmt.Println("Luas:", luas)
			fmt.Println("Keliling:", keliling)

		case 2:
			// Minitask pola jendela
			err := looping.Window(5)

			if err != nil {
				fmt.Println("Terjadi kesalahan:", err)
			}

		case 3:
			// Minitask slice
			numbers := []int{50, 75, 66, 20, 32, 90}
			hasil := slice.SisipAngka(numbers)

			fmt.Println("Data awal:", numbers)
			fmt.Println("Hasil setelah disisipkan:", hasil)

		case 4:
			// Minitask struct data user
			user := userdata.UserData{
				Name:      "Rama Lana Komara",
				Photo:     "rama.png",
				Email:     "rlanakomara7@gmail.com",
				Age:       27,
				Phone:     "089614238447",
				IsMarried: true,
				Education: []userdata.Education{
					{
						Univ:  "Universitas PGRI",
						Study: "Computer Science",
					},
				},
			}

			fmt.Println(user)

		case 5:
			var filepath string

			fmt.Print("Masukan file path:")
			fmt.Print(&filepath)

			panicrecover.ProcessFile(filepath)

		case 6:
			person := method.NewPerson(
				"Rama",
				"Bogor",
				"0897654213",
			)

			fmt.Printf(person.Print())
			fmt.Printf(person.Greet())

			person.SetName("Sonia Cahya")
			fmt.Println("\nSetelah nama diubah:")
			fmt.Println(person.Greet())

		case 7:
			bank := []int{10000, 20000}
			online := []int{15000, 25000}
			fiktif := &checkout.Fiktif{}

			if err := checkout.Checkout(checkout.Bank{}, bank); err != nil {
				fmt.Println("Error:", err)
			}

			if err := checkout.Checkout(checkout.Online{}, online); err != nil {
				fmt.Println("Error:", err)
			}

			if err := checkout.Checkout(fiktif, []int{5000, 10000}); err != nil {
				fmt.Println("Error:", err)
			}

			fmt.Println("Total pembayaran fiktif:", fiktif.Total)
		case 0:
			fmt.Println("Program selesai. Terima kasih!")
			return

		default:
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}
