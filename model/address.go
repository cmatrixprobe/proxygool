package model

type Address struct {
	Host     string
	Port     string
	Protocol string
	Speed    int64
	Origin   string
}

func NewAddress() *Address {
	return new(Address)
}
