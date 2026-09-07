package slice

// MENYISIPKAN ANGKA DIANTARA DATA DALAM SLICE , MINITASK 3 ---------------
func SisipAngka(originalSlice []int) []int {
	toInsert := []int{88}

	mid := len(originalSlice) / 2
	return append(originalSlice[:mid], append(toInsert, originalSlice[mid:]...)...)
}
