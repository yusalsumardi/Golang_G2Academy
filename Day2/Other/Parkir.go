package main

import (
	"fmt"
	"time"
	"strconv"
)

func main() {
	menu := "\n=== Menu Parkir ===\n" +					//String menu
		"1.  Masuk\n" +
		"2.  Keluar\n" +
		"99. Keluar Program"

	kendaraan := "\n=== Jenis Kendaraan ===\n" +		//String kendaraan
		"1. Mobil\n" +
		"2. Motor"
	
	customer := make(map[string]time.Time)				//Map customer
	id := 210200										//Inisialisasi id
	var pil int32
	for pil != 99 {
		fmt.Println(menu)
		fmt.Print("Pilihan\t: ")
		fmt.Scan(&pil)									//Pilih menu
		
		switch pil {
		case 1:
			idKendaraan := "RSS" + strconv.Itoa(id)		//Id kendaraan masuk
			customer[idKendaraan] = time.Now()
			fmt.Println("\nId Customer\t:", idKendaraan)
			id++

		case 2:
			var (
				jamMasuk time.Time
				jenisKen int32
				idKen string							//Id kendaraan keluar
				)
						
			for (jenisKen != 1) && (jenisKen != 2) {
				fmt.Println(kendaraan)
				fmt.Print("Pilihan (1-2)\t: ")
				fmt.Scan(&jenisKen)
				if !(jenisKen == 1 || jenisKen == 2) {
					fmt.Println("Pilihan Salah!")
				}
			}

			inputIdKen:
			fmt.Print("Id Kendaraan\t: ")
			fmt.Scan(&idKen)

			value, ok := customer[idKen];				//Cek Id Kendaraan
			if ok {
				jamMasuk = value
			} else {
				fmt.Println("Id Kendaraan tidak ditemukan")
				goto inputIdKen
			}

			jamKeluar := time.Now()						//Jam Keluar
			Keluar(jamMasuk, jamKeluar, jenisKen)
		case 99:
			fmt.Println("Berhasil Keluar")
		default:
			fmt.Println("Pilihan salah")
		}
	}
}

func Keluar(jamMasuk time.Time, jamKeluar time.Time, jenisKen int32){
	var tarif, tarifAwal, tarifMaksimal int64

	switch jenisKen {									//Cek tarif sesuai jenis
		case 1:											//kendaraan
			tarif = 3000
			tarifAwal = 5000
			tarifMaksimal = 100000
		case 2:
			tarif = 2000
			tarifAwal = 3000
			tarifMaksimal = 50000
	}

	durasi := jamKeluar.Sub(jamMasuk)
	totalTarif := tarifAwal + ((int64 (durasi.Seconds())) - 1) * tarif

	if totalTarif > tarifMaksimal{						//Cek tarif maksimal
		totalTarif = tarifMaksimal
	}

	fmt.Println("\nJam Masuk\t:", jamMasuk.Format("2006-01-02 15:04:05"))
	fmt.Println("Jam Keluar\t:", jamKeluar.Format("2006-01-02 15:04:05"))
	fmt.Println("Durasi Parkir\t:", int64 (durasi.Seconds()))
	fmt.Println("Total Tarif\t:", totalTarif)
}