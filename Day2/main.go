package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Hitung struct {
	No1, No2 int
}

type Count struct {
	No1, No2 int
}

func main() {

	var m map[string]int

	fmt.Println("m:", m)

	if m == nil {
		fmt.Println("m is nil")
	}

	m = make(map[string]int) // byte

	m["januari"] = 10000
	m["februari"] = 11000

	// delete(m, "januari")
	fmt.Println("m:", m)

	_, check := m["januari"]
	// fmt.Println("value:", value)
	fmt.Println("check:", check)

	// value, check = m["februari"]
	// fmt.Println("value:", value)
	// fmt.Println("check:", check)

	fmt.Println("januari:", m["januari"])

	total := 0
	for _, val := range m {
		// fmt.Println("key:", key)

		// if key == "februari" {
		// 	fmt.Println("februari:", m[key])
		// }
		fmt.Println("val:", val)
		total += val
	}
	fmt.Println("total:", total)

	ret, age := printName("admins", 10)
	fmt.Println("ret:", ret)
	fmt.Println("age:", age)

	// showArray(1, 2, 3, 4, 5)
	// showArray(6, 7)
	// showArray([]byte("a"))

	func(msg string) {
		fmt.Println(msg)
	}("Hi,,")

	number := 9

	change(&number, 10) // 0x00001
	// number = change(number, 10)
	fmt.Println("number:", number)

	b := biodata()
	c := Person{"Admins", "Trator", 22}
	d := Person{"Adminis", "Trator", 23}

	c.Age = 20

	fmt.Println("Biodata:", b)
	fmt.Println("Biodata:", c)
	fmt.Println("Biodata:", d)
	fmt.Println("firstname:", b.FirstName)

	perhitungan := Hitung{5, 10}
	couting := Count{5, 10}
	fmt.Println("Hasil:", perhitungan.Pertambahan(13))
	// fmt.Println("Hasil:", perhitungan.Perkalian(13))
	fmt.Println("Hasil:", couting.Pertambahan(13))
	fmt.Println("Hasil:", couting.Perkalian(13))
	// adduser = "admin", "user", "client"
	// fmt.Println("tambah:", Pertambahan(3))

}

func printName(name string, no int) (nama string, total int) {
	nama = "admin"
	fmt.Println("Name:", name)
	fmt.Println("Age:", no)
	total = no + 2
	return rename(name), total
}

func rename(name string) string {
	return name + "istrator"
}

func showArray(arrayNo ...[]byte) {
	fmt.Println("arrayNo:", arrayNo)
	// fmt.Println("arrayNo 0:", arrayNo[0])
	// total := 0
	// for _, val := range arrayNo {
	// 	total += int(val)
	// }
	// fmt.Println("Total:", total)
}

func change(original *int, value int) {
	*original = 2 * value // 0x00001 = value
}

func biodata() Person {
	// return Person{"Admin", "Istrator", 20}
	return Person{LastName: "Istrator", FirstName: "Admin"}
	// return Person{
	// 	LastName:  "Istrator",
	// 	FirstName: "Admin",
	// 	Age:       21,
	// }
}

func (h Hitung) Pertambahan(no int) int {
	return h.No1 + h.No2 + no
}

func (h Count) Pertambahan(no int) int {
	return h.No1 - h.No2 - no
}

func (h Count) Perkalian(no int) int {
	return h.No1 * h.No2 * no
}

/*
 Buat kalkulator dengan fungsi:
	- kali
	- bagi
	- tambah
	- kurang
	- akar
	- pangkat
	- luas persegi
	- luas lingkaran
	- volume tabung
	- volume balok
	- volume prisma
menggunkan:
	- struct & method
	- pointer
	- function
*/
