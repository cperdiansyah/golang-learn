package slice

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3] //ngambil value index ke 0 sampai 3
	var bFruits = fruits[1:4] //ngambil value index ke 1 sampai 4

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	// Buah "grape" diubah menjadi "pinnaple"
	baFruits[0] = "pinnaple"
	// ini terjdi karena sifat slice itu reference, jadi satu berubah semuanya berubah

	fmt.Println(fruits)   // [apple pinnaple banana melon]
	fmt.Println(aFruits)  // [apple pinnaple banana]
	fmt.Println(bFruits)  // [pinnaple banana melon]
	fmt.Println(aaFruits) // [pinnaple]
	fmt.Println(baFruits) // [pinnaple]
}

func learnLen() {
	//Fungsi len() digunakan untuk menghitung jumlah elemen slice yang ada. Sebagai contoh jika sebuah variabel adalah slice dengan data 4 buah, maka fungsi ini pada variabel tersebut akan mengembalikan angka 4.
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits)) // 4
}

func learCap() {
	// Fungsi cap() digunakan untuk menghitung lebar atau kapasitas maksimum slice. Nilai kembalian fungsi ini untuk slice yang baru dibuat pasti sama dengan len, tapi bisa berubah seiring operasi slice yang dilakukan. Agar lebih jelas, silakan pelajari kode berikut.

	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	var aFruits = fruits[0:3]
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	var bFruits = fruits[1:4]
	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3
}

func learAppend() {
	//Fungsi append() digunakan untuk menambahkan elemen pada slice. Elemen baru tersebut diposisikan setelah indeks paling akhir. Nilai balik fungsi ini adalah slice yang sudah ditambahkan nilai barunya. Contoh penggunaannya bisa dilihat di kode berikut.
	//apped tidak bisa digunakan di fixed aray

	var fruits = []string{"apple", "grape", "banana"}
	var cFruits = append(fruits, "papaya")

	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(cFruits) // ["apple", "grape", "banana", "papaya"]
}
