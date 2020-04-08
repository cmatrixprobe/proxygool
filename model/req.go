package model

import (
	"strings"
)

// Request
type Request struct {
	WebName   string              // website name
	WebURL    string              // website url
	TrRegular string              // css selector of table row
	Pages     int                 // pages to be crawled
	HostIndex int                 // host column number
	PortIndex int                 // port column number
	ProtIndex int                 // protocol column number
	Trim      bool                // whether to remove "\t\n"
	Protocol  func(string) string // protocol
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
