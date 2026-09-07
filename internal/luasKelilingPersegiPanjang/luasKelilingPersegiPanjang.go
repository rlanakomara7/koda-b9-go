package luaskelilingpersegipanjang

// HITUNG LUAS
func HitungLuas(p int16, l int16) int16 {
	hasilLuas := p * l
	return hasilLuas
}

// HITUNG KELILING
func HitungKeliling(p int16, l int16) int16 {
	hasilKeliling := 2 * (p + l)
	return hasilKeliling
}

// LUAS DAN KELILING
func LuasKeliling(p int16, l int16) (Luas int16, Keliling int16) {
	return HitungLuas(p, l), HitungKeliling(p, l)
}
