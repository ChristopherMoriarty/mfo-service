package domain

type ColdUsers struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Phone     int   `json:"phone"`
	Credits   int   `json:"credits"`
}