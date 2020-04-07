package model

// Address
type Address struct {
	Host     string
	Port     string
	Protocol string
	Speed    int64
	Origin   string
}

// NewAddress
func NewAddress() *Address {
	return new(Address)
}
