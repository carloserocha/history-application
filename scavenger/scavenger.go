package scavenger

import (
    "encoding/json"
    "net/http"
)

type Scavenger struct {
	Metadata struct {
		ReceiveAt  string        `json:"receive_at"`
		IPAddress  string        `json:"ip_address"`
		Topic      string        `json:"topic"`
		Components []interface{} `json:"components"`
	} `json:"metadata"`
	User struct {
		Name    string `json:"name"`
		Address struct {
			City     string `json:"city"`
			District string `json:"district"`
			Street   string `json:"street"`
			Number   string `json:"number"`
			ZipCode  string `json:"zip_code"`
		} `json:"address"`
		Phones []struct {
			Phone string `json:"phone"`
		} `json:"phones"`
		BirthDate     string `json:"birth_date"`
		DoumentNumber string `json:"doument_number"`
		Email         string `json:"email"`
	} `json:"user"`
}

func CreateScavenger(w http.ResponseWriter, r *http.Request) {
	fmt.Println(json.NewEncoder(w).Encode(Scavenger))
}