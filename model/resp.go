package model

// TestResponse
type TestResponse struct {
	City        string      `json:"city"`
	Country     string      `json:"country"`
	CountryCode string      `json:"country_code"`
	Distinct    interface{} `json:"distinct"`
	IP          string      `json:"ip"`
	Isp         string      `json:"isp"`
	Lat         string      `json:"lat"`
	Lon         string      `json:"lon"`
	Operator    string      `json:"operator"`
	Province    string      `json:"province"`
}
