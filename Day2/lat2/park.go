package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Log struct {
	time1  int64
	time2  int64
	idpark int
}

func main() {
	var idpark int
	var input int
	var status bool = true
	datein := time.Now()
	dateout := time.Now()
	dateinUnix := datein.Unix()
	dateoutUnix := dateout.Unix()
	for input != 99 {
		fmt.Println("====Menu====") //Pilihan Menu
		fmt.Println("1. Masuk Parkir")
		fmt.Println("2. Keluar Parkir")
		fmt.Println("99. EXIT")
		fmt.Scanln(&input) //Scan menu
		switch input {
		case 1:
			//======================KENDARAAN MASUK=========================
			if status {
				datein = time.Now()
				idpark = rand.Intn(100)
				fmt.Println("Anda Masuk pada :", datein.Format("2006-01-02 15:04:05 Monday"))
				dateinUnix = datein.Unix()
				fmt.Printf("datein unix timestamp is :%v\n", dateinUnix)
				fmt.Println("ID Parkir anda:", idpark)
				status = false
			} else {
				fmt.Println("Error : Anda Sudah Masuk")
			}
		case 2:
			//======================KENDARAAN KELUAR=========================
			if status == false {
				fmt.Println("====Menu====") //Pilihan Menu
				fmt.Println("1. Keluar Parkir Motor")
				fmt.Println("2. Keluar Parkir Mobil")
				fmt.Println("Pilih Jenis Kendaraan")
				var opsi int
				fmt.Scanln(&opsi)
				var noplat int
				var idparktemp int
				if opsi == 1 {
					//======================MOTOR MASUK=========================
					fmt.Println("Masukan Plat Nomor Kendaraan Anda")
					fmt.Scanln(&noplat)
					fmt.Println("Masukan Nomor Parkir Anda")
					fmt.Scanln(&idparktemp)
					// fmt.Println(idparktemp, idpark)
					if idparktemp == idpark {
						dateout = time.Now()
						dateoutUnix = dateout.Unix() //<<<<<<<<<<<<<<<<<<<<<<< CONVERT KE UNIX
						count := Log{dateinUnix, dateoutUnix, idpark}
						fmt.Println("======================================")
						fmt.Println("Anda Masuk pada :", datein.Format("2006-01-02 15:04:05 Monday"))
						fmt.Println("Anda Keluar pada :", dateout.Format("2006-01-02 15:04:05 Monday"))
						fmt.Println("======================================")
						count.TarifMotor()
						status = true
					}

					//======================MOTOR MASUK=========================
				} else if opsi == 2 {
					fmt.Println("Masukan Plat Nomor Kendaraan Anda")
					fmt.Scanln(&noplat)
					fmt.Println("Masukan Nomor Parkir Anda")
					fmt.Scanln(&idparktemp)
					// fmt.Println(idparktemp, idpark)
					if idparktemp == idpark {
						dateout = time.Now()
						dateoutUnix = dateout.Unix() //<<<<<<<<<<<<<<<<<<<<<<< CONVERT KE UNIX
						count := Log{dateinUnix, dateoutUnix, idpark}
						fmt.Println("======================================")
						fmt.Println("Anda Masuk pada :", datein.Format("2006-01-02 15:04:05 Monday"))
						fmt.Println("Anda Keluar pada :", dateout.Format("2006-01-02 15:04:05 Monday"))
						fmt.Println("======================================")
						count.TarifMobil()
						status = true
					} else {
						fmt.Println("Nomor Parkir Anda Salah!")
					}
				} else {
					fmt.Println("Menu Pilihan Anda Salah!")
				}
			} else {
				fmt.Println("Error : Anda Belum Masuk")
			}
		}
	}
}

func (t Log) TarifMotor() {
	var tarif int64
	counttime := t.time2 - t.time1
	if counttime >= 1 {
		tarif = 3000
	}
	tarif = tarif + ((counttime - 1) * 2000)
	fmt.Println("======================================")
	fmt.Println("Lama anda parkir: ", counttime, "detik")
	fmt.Println("Nomor Parkir Motor Anda:", t.idpark)
	fmt.Println("Tarif Parkir Motor anda: Rp.", tarif)
	fmt.Println("======================================")
}

func (t Log) TarifMobil() {
	var tarif int64
	counttime := t.time2 - t.time1
	if counttime >= 1 {
		tarif = 5000
	}
	tarif = tarif + ((counttime - 1) * 3000)
	fmt.Println("======================================")
	fmt.Println("Lama anda parkir: ", counttime, "detik")
	fmt.Println("Nomor Parkir Mobil Anda:", t.idpark)
	fmt.Println("Tarif Parkir Mobil anda: Rp.", tarif)
	fmt.Println("======================================")
}
