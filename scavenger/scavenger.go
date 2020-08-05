package scavenger

import (
	"encoding/json"
	"net/http"
	"time"

	boom "github.com/darahayes/go-boom"
	auth "github.com/carloserocha/history-application/authentication"
)

const TOPIC = "scavenger"

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
			Phone     string `json:"phone"`
			PhoneType string `json:"phone_type"`
		} `json:"phones"`
		BirthDate     string `json:"birth_date"`
		DoumentNumber string `json:"doument_number"`
		Email         string `json:"email"`
	} `json:"user"`
}

type ScavengerError struct {
	Error   string    `json:"error"`
	Payload Scavenger `json:"payload"`
}

const schemaScavenger = `
CREATE TABLE IF NOT EXISTS User (
	user_id INT AUTO_INCREMENT PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    birth_date DATE,
    document_number VARCHAR(11) NOT NULL,
    email VARCHAR(100) NOT NULL,
    city VARCHAR(30) NOT NULL,
    district VARCHAR (50) NOT NULL,
    street VARCHAR(50) NOT NULL,
    number_residential VARCHAR(10) NOT NULL,
    zip_code VARCHAR(8) NOT NULL,
    mobile_phone VARCHAR(15) NOT NULL,
    home_phone VARCHAR(15) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    payload JSON
)
`

func handleScanveger(s *Scavenger) {
	// s.User.Name = "Bruninho da manga"
	// fmt.Println(s)
}

// Cria um novo gincaneiro e retorna
func CreateScavenger(w http.ResponseWriter, r *http.Request) {

	//var sError ScavengerError
	_, err := auth.AuthenticateAuthorize(w, r)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		boom.Unathorized(w, err)
		return
	}

	var s Scavenger

	json.NewDecoder(r.Body).Decode(&s)

	s.Metadata.IPAddress = r.RemoteAddr
	s.Metadata.ReceiveAt = time.Now().UTC().Local().Format(time.RFC3339Nano)
	s.Metadata.Topic = TOPIC

	handleScanveger(&s)

	encoding, _ := json.Marshal(s)
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(encoding)
}
