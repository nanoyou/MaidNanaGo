package password

type PlainPassword struct {
	password string
}

func NewPlainPassword(password string) *PlainPassword {
	return &PlainPassword{password}
}

func (pp *PlainPassword) Validate(password string) bool {
	return password == pp.password
}

func (pp *PlainPassword) String() string {
	return "PLAIN::" + pp.password
}
