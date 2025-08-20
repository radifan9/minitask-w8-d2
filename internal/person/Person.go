package person

import "fmt"

type Person struct {
	Name    string
	Address string
	Phone   string
}

// Constructor function
func NewPerson(name, address, phone string) *Person {
	return &Person{
		Name:    name,
		Address: address,
		Phone:   phone,
	}
}

// p *Person --> pointer receiver
// Dengan pointer receiver, method bisa mengakses dan mengubah data pada struct aslinya
func (p *Person) PrintData() {
	fmt.Printf("Name: %s\nAddress: %s\nPhone: %s\n", p.Name, p.Address, p.Phone)
}

func (p *Person) Greet(target string) {
	fmt.Printf("Halo %s! Dari %s.\n", target, p.Name)
}

func (p *Person) UpdateName(newName string) {
	p.Name = newName
}
