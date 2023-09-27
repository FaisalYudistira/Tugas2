package main

import "fmt"

func main(){
	//mengurutkan bilangan dari yang terbesar hingga yang terkecil
	var angka [5]int64
	angka = [5]int64{
		84,14,73,32,58,
	}

	fmt.Printf("Angka sebelum diurutkan : %d\n",angka)

	//metode bubblesort
	for i := 0; i < len(angka) - 1; i++ {
		for j := 0; j < len(angka) - 1; j++ {
			if angka[j] < angka[j + 1] {
				var temp int64 = angka[j]
				angka[j] = angka[j + 1]
				angka[j + 1] = temp
			}
		}
	}
	
	fmt.Printf("Angka setelah diurutkan : %d\n",angka)

	/*
	.
	.
	.
	.
	.
	.
	.
	*/

	//membuat slice dengan jumlah data 5 dan menambahkan 3 data ke dalam slice tersebut
	var BarangAwal []string
	BarangAwal = []string{
		"buku","pulpen","pensil","penghapus","binder",
	}

	fmt.Printf("barang sebelum ditambahkan : %s\n",BarangAwal)

	//menambahkan data ke dalam slice BarangAwal
	var BarangAkhir = append(BarangAwal,"laptop","hp","charger")

	fmt.Printf("barang setelah ditambahkan : %s",BarangAkhir)
}