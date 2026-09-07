package userdata

// STRUCT DAT RIWAYAT MINITASK 4 ---------------

type UserData struct {
	Name      string
	Photo     string
	Email     string
	Age       uint8
	Phone     string
	IsMarried bool
	Education []Education
}

type Education struct {
	Univ  string
	Study string
}
