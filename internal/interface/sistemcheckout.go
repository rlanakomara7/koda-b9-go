package interface

type Payment interface {
	Pay() string
}
type Bank struct{}

func (b Bank) Pay(list []int) string{
	return fmt.Sprintf("List Harga: %v, Metode Pembayaran Bank",list)
}

type Online struct{}

func (o Online) Pay(list []int) string{
	return fmt.Sprintf("List Harga: %v, Metode Pembayaran Online", list)
}