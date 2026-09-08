package method

import "fmt"

// CONSTRUCTOR MINITASK 7

// STRUCT DATA
type Person struct {
	Name    string
	Address string
	Phone   string
}

// CONSTRUCTOR
func NewPerson(name string, address string, phone string) *Person {
	return &Person{
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

// METHODE PRINT

func (p Person) Print() string {
	return fmt.Sprintf(
		"Name: %s , Address: %s , Phone: %s",
		p.Name,
		p.Address,
		p.Phone,
	)
}

// METHODE GREET

func (p Person) Greet() string {
	return fmt.Sprintf("Hello my name is %s", p.Name)
}

// METHODE SETTER
func (p *Person) setName(name string) {
	p.Name = name
}
