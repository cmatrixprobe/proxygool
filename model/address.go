package model

// Address .
type Address struct {
	Host     string
	Port     string
	Protocol string
	Speed    int64
	Origin   string
}

// NewAddress returns a new pointer to address.
func NewAddress() *Address {
	return new(Address)
}
