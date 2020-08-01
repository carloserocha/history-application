package scavenger

import (
	"encoding/json"
	"log"
	"net/http"
)

// estrutura de um gincaneiro
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

// Schema default to model scavenger
type Schema struct {
	Metadata struct {
		ReceiveAt  string        `db:"receive_at"`
		IPAddress  string        `db:"ip_address"`
		Topic      string        `db:"topic"`
		Components []interface{} `db:"components"`
	} `db:"metadata"`
	User struct {
		Name    string `db:"name"`
		Address struct {
			City     string `db:"city"`
			District string `db:"district"`
			Street   string `db:"street"`
			Number   string `db:"number"`
			ZipCode  string `db:"zip_code"`
		} `db:"address"`
		Phones []struct {
			Phone string `db:"phone"`
		} `db:"phones"`
		BirthDate     string `db:"birth_date"`
		DoumentNumber string `db:"doument_number"`
		Email         string `db:"email"`
	} `db:"user"`
}

// Cria um novo gincaneiro e retorna
func CreateScavenger(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s Scavenger

	json.NewDecoder(r.Body).Decode(&s)

	// do stuff

	encoding, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(encoding)
}
