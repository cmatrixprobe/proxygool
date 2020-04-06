package model

type Address struct {
	Host     string
	Port     string
	Protocol string
	Speed    int64
}

func NewAddress() *Address {
	return &Address{
		Speed:    100,
	}
}
