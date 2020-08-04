package main

import (
	"fmt"
	"time"
)

func main() {

	day := 1
	// tahun := 2020
	// // space := "-"
	// jumlahhari := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	// firstday := [12]int{}

	// if (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0 {
	// 	jumlahhari[1] = 29
	// }

	var arrbul [13]time.Month
	arrbul[1] = time.January
	arrbul[2] = time.February
	arrbul[3] = time.March
	arrbul[4] = time.April
	arrbul[5] = time.May
	arrbul[6] = time.July
	arrbul[7] = time.June
	arrbul[8] = time.August
	arrbul[9] = time.September
	arrbul[10] = time.October
	arrbul[11] = time.November
	arrbul[12] = time.December

	for i := 1; i <= len(arrbul); i++ {
		fmt.Println("=======", arrbul[i], "=======")
		fmt.Println("---------------------------------------------------")
		fmt.Print("M\tS\tS\tR\tK\tJ\tS\n")
		fmt.Println("---------------------------------------------------")
		var cal = [5][7]int{}
		for a := 0; a < 5; a++ {
			for b := 0; b < 7; b++ {
				cal[a][b] = day
				day++
				if arrbul[i] == time.February {
					if day > 29 {
						day = 1
						if cal[a][b] == 29 {
							break
						}
					}
				} else {
					if day > 30 {
						day = 1
						if cal[a][b] == 30 {
							break
						}
					}
				}
			}
		}

		for a := 0; a < 5; a++ {
			for b := 0; b < 7; b++ {
				fmt.Print(cal[a][b], "\t")
				if cal[a][b] == 30 {
					break
				}
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	}
}
