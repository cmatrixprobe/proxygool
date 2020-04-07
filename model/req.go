package model

import (
	"strings"
)

// Request
type Request struct {
	WebName   string
	WebURL    string
	TrRegexp  string
	Pages     int
	HostIndex int
	PortIndex int
	ProtIndex int
	Trim      bool
	Protocol  func(string) string
}

// NewRequest
func NewRequest() *Request {
	return &Request{
		Pages: 1,
		Protocol: func(s string) string {
			return strings.ToLower(s)
		},
	}
}
