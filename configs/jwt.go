package configs

import (
	"net/http"
	"time"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("mysupersecretphrase")

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("tehere was an error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				fmt.Fprintf(rw, err.Error())
			}
			if token.Valid {
				endpoint(rw, r)
			}

		} else {
			fmt.Fprintf(rw, "not authorized")
		}

	})
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = "Elliot Forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("ada yang error")
		return "", err
	}

	return tokenString, nil
}