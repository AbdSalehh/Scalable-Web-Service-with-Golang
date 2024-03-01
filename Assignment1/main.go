package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func getDataTeman(absen int, teman []Teman) Teman {
	if absen < 1 || absen > len(teman) {
		return Teman{}
	}

	return teman[absen-1]
}

func tampilkanSemuaTeman(teman []Teman) {
	for i, t := range teman {
		fmt.Printf("%d. %s\n", i+1, t.Nama)
		fmt.Println("   Alamat:", t.Alamat)
		fmt.Println("   Pekerjaan:", t.Pekerjaan)
		fmt.Println("   Alasan memilih kelas Golang:", t.Alasan)
		fmt.Println()
	}
}

func main() {
	teman := []Teman{
		{"Wahyu Ginanjar", "Jalan Merdeka No. 10, Jakarta Pusat", " Software Engineer", "Ingin memperdalam pemrograman Go untuk meningkatkan keterampilan pengembangan perangkat lunak."},
		{"I Made Putra", "Jalan Raya Ubud No. 25, Gianyar, Bali", "Data Scientist", "Melihat potensi Golang dalam pengolahan data dan analisis yang cepat dan efisien"},
		{"Ratna Sari", "Perumahan Permata Hijau Blok C2 No. 15, Surabaya, Jawa Timur", "Web Developer", "Mendengar tentang kemampuan konkuren dan efisiensi Golang dalam pengembangan aplikasi web"},
		{"Ratna Sari", "Komplek Puri Indah Blok D3 No. 8, Bandung, Jawa Barat", "UI/UX Designer", "Tertarik dengan kesederhanaan sintaksis Golang dan potensi untuk pengembangan antarmuka pengguna yang responsif"},
		{"Sri Wahyuni", "Jalan Mawar 17, Kelurahan Kebon Jeruk, Malang, Jawa Timur", "Systems Analyst", "Menginginkan kehandalan dan performa yang tinggi yang ditawarkan oleh Golang dalam pengembangan sistem informasi"},
	}

	if len(os.Args) < 2 {
		tampilkanSemuaTeman(teman)
		return
	}

	absen := os.Args[1]

	var absenInt int
	_, err := fmt.Sscanf(absen, "%d", &absenInt)
	if err != nil {
		fmt.Println("Nomor absen harus berupa angka")
		return
	}

	temanTerpilih := getDataTeman(absenInt, teman)

	if temanTerpilih.Nama == "" {
		fmt.Println("Nomor absen tidak valid. Menampilkan semua daftar teman:")
		tampilkanSemuaTeman(teman)
		return
	}

	fmt.Println("Nama:", temanTerpilih.Nama)
	fmt.Println("Alamat:", temanTerpilih.Alamat)
	fmt.Println("Pekerjaan:", temanTerpilih.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", temanTerpilih.Alasan)
}
