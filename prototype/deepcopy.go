package prototype

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	return &Person{
		Name:    p.Name,
		Address: p.Address.DeepCopy(),
		Friends: append([]string{}, p.Friends...),
	}
}
