package authentication

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var JWTKEY = []byte("scavenger_secrets")
var SCAVENGERS = map[string]string{
	"carlinhos": "password_super_segura",
}

func AuthenticateLogin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the expected password from our in memory map
	expectedPassword, ok := SCAVENGERS[creds.Username]

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(JWTKEY)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func AuthenticateAuthorize(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("token")
	if err != nil {
		return nil, errors.New("unauthorized access")

		/*if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return JWTKEY, errors.New("unauthorized access")
		}
		// For any other type of error, return a bad request status
		return JWTKEY, errors.New("bad request")*/
	}

	// Get the JWT string from the cookie
	tknStr := c.Value

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTKEY, nil
	})
	if err != nil {
		return nil, errors.New("unauthorized access")
		/*if err == jwt.ErrSignatureInvalid {
			return JWTKEY, errors.New("unauthorized access")
		}
		return JWTKEY, errors.New("bad request")*/
	}
	if !tkn.Valid {
		return nil, errors.New("unauthorized access")
	}

	// Finally, return the welcome message to the user, along with their
	// username given in the token
	return JWTKEY, nil
}
