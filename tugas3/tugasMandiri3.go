package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var daftarTeman = []Teman{
	{Nama: "Adi", Alamat: "Semarang", Pekerjaan: "Mahasiswa", Alasan: "Ingin belajar backend"},
	{Nama: "Budi", Alamat: "Grobogan", Pekerjaan: "Freelance", Alasan: "Memperdalam ilmu"},
	{Nama: "Cipto", Alamat: "Surabaya", Pekerjaan: "System Analyst", Alasan: "Menambah pengetahuan bahasa pemrograman baru"},
	{Nama: "Darma", Alamat: "Tegal", Pekerjaan: "DevOps", Alasan: "Belajar backend"},
}

func tampilData(id int) {
	if id < 1 || id > len(daftarTeman) {
		fmt.Println("Data dengan id tersebut tidak ditemukan.")
		return
	}
	teman := daftarTeman[id-1]
	fmt.Printf("Nama: %s\n", teman.Nama)
	fmt.Printf("Alamat: %s\n", teman.Alamat)
	fmt.Printf("Pekerjaan: %s\n", teman.Pekerjaan)
	fmt.Printf("Alasan memilih kelas Golang: %s\n", teman.Alasan)
}

func main() {
	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("ID harus berupa angka.")
		return
	}

	tampilData(id)
}
