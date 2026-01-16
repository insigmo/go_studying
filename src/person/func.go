package person

func (p *NewPerson) UpdateAge(age int) {
	p.Age = age
}

func (p *NewPerson) GetSecret() string {
	return p.secret
}
