package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/context"

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

func Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorisationHeader := r.Header.Get("authorization")
		if authorisationHeader != "" {
			token, err := jwt.Parse(authorisationHeader, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return []byte("secret"), nil
			})
			if err != nil {
				err := json.NewEncoder(w).Encode(Exception{Message: error.Error(err)})
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}
			if token.Valid {
				context.Set(r, "decoded", token.Claims)
				next(w, r)
			} else {
				err := json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
		} else {
			err := json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
