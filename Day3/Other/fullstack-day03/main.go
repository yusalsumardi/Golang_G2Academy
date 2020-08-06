package main

import (
	"fmt"
	"math"
)

type hitung interface {
	luas() float64
	keliling() float64
	// jariJari() float64
}

type persegi struct {
	sisi float64
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}

func (p persegi) keliling() float64 {
	return 4 * p.sisi
}

func main() {
	var bangunDatar hitung
	var bangunPersegi hitung

	bangunDatar = lingkaran{14.0}
	bangunPersegi = persegi{7}

	fmt.Println("luas\t\t:", bangunPersegi.luas())
	fmt.Println("keliling\t:", bangunPersegi.keliling())

	fmt.Println("luas\t\t:", bangunDatar.luas())
	fmt.Println("keliling\t:", bangunDatar.keliling())
	// fmt.Println("jari-jari\t:", bangunDatar.jariJari())
	fmt.Println("jari-jari\t:", bangunDatar.(lingkaran).jariJari())
}
