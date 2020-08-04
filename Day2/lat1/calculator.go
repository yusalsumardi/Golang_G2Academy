package main

import (
	"fmt"
	"math"
	"strconv"
)

type Hitung struct {
	ang1, ang2 int
}

type BangunRuang struct {
	p, l, t int
}

type BangunDatar struct {
	p, l int
}

type Akar struct {
	s float64
}

func main() {
	fmt.Println("====Menu====") //Pilihan Menu
	fmt.Println("1. kali")
	fmt.Println("2. bagi")
	fmt.Println("3. tambah")
	fmt.Println("4. kurang")
	fmt.Println("5. akar")
	fmt.Println("6. pangkat")
	fmt.Println("7. luas persegi")
	fmt.Println("8. luas lingkaran")
	fmt.Println("9. volume tabung")
	fmt.Println("10. volume balok")
	fmt.Println("11. volume prisma")
	fmt.Print("Enter menu: ")
	var menu string
	fmt.Scanln(&menu) //Scan menu
	menuswitch, _ := strconv.Atoi(menu)
	switch menuswitch {
	case 1:
		fmt.Print("Enter Angka Pertama: ")
		var Ang1 int
		fmt.Scanln(&Ang1)
		fmt.Print("Enter Angka Kedua: ")
		var Ang2 int
		fmt.Scanln(&Ang2)
		perhitungan := Hitung{Ang1, Ang2}
		hasil := perhitungan.Perkalian()
		fmt.Println("Hasil Perkalian: ", hasil)
	case 2:
		fmt.Print("Enter Angka Pertama: ")
		var Ang1 int
		fmt.Scanln(&Ang1)
		fmt.Print("Enter Angka Kedua: ")
		var Ang2 int
		fmt.Scanln(&Ang2)
		perhitungan := Hitung{Ang1, Ang2}
		hasil := perhitungan.Pembagian()
		fmt.Println("Hasil Pembagian: ", hasil)
	case 3:
		fmt.Print("Enter Angka Pertama: ")
		var Ang1 int
		fmt.Scanln(&Ang1)
		fmt.Print("Enter Angka Kedua: ")
		var Ang2 int
		fmt.Scanln(&Ang2)
		perhitungan := Hitung{Ang1, Ang2}
		hasil := perhitungan.Penambahan()
		fmt.Println("Hasil Penambahan: ", hasil)
	case 4:
		fmt.Print("Enter Angka Pertama: ")
		var Ang1 int
		fmt.Scanln(&Ang1)
		fmt.Print("Enter Angka Kedua: ")
		var Ang2 int
		fmt.Scanln(&Ang2)
		perhitungan := Hitung{Ang1, Ang2}
		hasil := perhitungan.Pengurangan()
		fmt.Println("Hasil Pengurangan: ", hasil)
	case 5:
		fmt.Print("Enter Angka Pertama: ")
		var Ang1 float64
		fmt.Scanln(&Ang1)
		perhitungan := Akar{Ang1}
		hasil := perhitungan.Akar()
		fmt.Println("Hasil Akar: ", hasil)
	case 6:
		fmt.Print("Enter Angka Pertama: ")
		var Ang1 int
		fmt.Scanln(&Ang1)
		fmt.Print("Enter Jumlah Pangkat: ")
		var Ang2 int
		fmt.Scanln(&Ang2)
		perhitungan := Hitung{Ang1, Ang2}
		hasil := perhitungan.Pangkat()
		fmt.Println("Hasil Pangkat: ", hasil)
	case 7:
		fmt.Print("Enter Panjang Sisi: ")
		var Panjang int
		fmt.Scanln(&Panjang)
		fmt.Print("Enter Lebar Sisi: ")
		var Lebar int
		fmt.Scanln(&Lebar)
		bangundatar := BangunDatar{Panjang, Lebar}
		hasil := bangundatar.LuasPersegi()
		fmt.Println("Luas Persegi: ", hasil)
	case 8:
		fmt.Print("Enter Jari-Jari: ")
		var Jari int
		fmt.Scanln(&Jari)
		bangundatar := BangunDatar{Jari, 0}
		hasil := bangundatar.LuasLingkaran()
		fmt.Println("Luas Lingkaran: ", hasil)
	case 9:
		fmt.Print("Enter Jari-Jari: ")
		var Jari int
		fmt.Scanln(&Jari)
		fmt.Print("Enter Tinggi Tabung: ")
		var Tinggi int
		fmt.Scanln(&Tinggi)
		bangunruang := BangunRuang{Jari, 0, Tinggi}
		hasil := bangunruang.VolumTabung()
		fmt.Println("Volume Tabung: ", hasil)
	case 10:
		fmt.Print("Enter Panjang: ")
		var Panjang int
		fmt.Scanln(&Panjang)
		fmt.Print("Enter Lebar: ")
		var Lebar int
		fmt.Scanln(&Lebar)
		fmt.Print("Enter Tinggi: ")
		var Tinggi int
		fmt.Scanln(&Tinggi)
		bangunruang := BangunRuang{Panjang, Lebar, Tinggi}
		hasil := bangunruang.VolumBalok()
		fmt.Println("Volume Balok: ", hasil)
	case 11:
		fmt.Print("Enter Panjang Alas: ")
		var Panjang int
		fmt.Scanln(&Panjang)
		fmt.Print("Enter Tinggi Alas: ")
		var Lebar int
		fmt.Scanln(&Lebar)
		fmt.Print("Enter Tinggi Prisma: ")
		var Tinggi int
		fmt.Scanln(&Tinggi)
		bangunruang := BangunRuang{Panjang, Lebar, Tinggi}
		hasil := bangunruang.VolumPrisma()
		fmt.Println("Volume Prisma: ", hasil)
	}
}

func (h Hitung) Perkalian() int {
	return h.ang1 * h.ang2
}

func (h Hitung) Pembagian() int {
	return h.ang1 / h.ang2
}

func (h Hitung) Penambahan() int {
	return h.ang1 + h.ang2
}

func (h Hitung) Pengurangan() int {
	return h.ang1 - h.ang2
}

func (h Akar) Akar() float64 {
	return math.Sqrt(h.s)
}

func (h Hitung) Pangkat() int {
	var temp int = h.ang1
	for i := 1; i < h.ang2; i++ {
		temp = temp * h.ang1
	}
	return temp
}

func (l BangunDatar) LuasPersegi() int {
	return l.p * l.l
}

func (l BangunDatar) LuasLingkaran() int {
	var luas int
	luas = l.p * l.p * (22 / 7)
	return luas
}

func (v BangunRuang) VolumTabung() int {
	var vol int
	vol = (v.p * v.p * (22 / 7)) * v.t
	return vol
}

func (v BangunRuang) VolumBalok() int {
	return v.p * v.l * v.t
}

func (v BangunRuang) VolumPrisma() int {
	var vol int
	vol = ((1 / 2) * v.p * v.l) * v.t
	return vol
}
