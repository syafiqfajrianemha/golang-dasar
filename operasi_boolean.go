package main

import "fmt"

func main() {
	nilaiAkhir := 90
	presensi := 80

	lulusNilaiAkhir := nilaiAkhir > 80
	lulusPresensi := presensi > 80

	lulus := lulusNilaiAkhir && lulusPresensi

	fmt.Println(lulus)
}
