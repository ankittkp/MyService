package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"

	"MyService/models"
)

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}

func CreateToken(w http.ResponseWriter, r *http.Request) {
	var u models.User
	_ = json.NewDecoder(r.Body).Decode(&u)
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username": u.Username,
		"password": u.Password,
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
	}
	err = json.NewEncoder(w).Encode(JwtToken{Token: tokenString})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
