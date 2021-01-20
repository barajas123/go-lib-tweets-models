package db

// struct for company object based on the iex cloud api

type Company struct {
	Ticker      string `json:"name"`
	Name        string `json:"name"`
	Industry    string `json:"industry"`
	Description string `json:"description"`
	Sector      string `json:"sector"`
	Country     string `json:"country"`
}
