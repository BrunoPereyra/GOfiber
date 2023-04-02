package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func parseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("mi_clave_secreta"), nil
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("B")
	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return token, nil
}

func ExtractDataFromToken(tokenString string) (string, error) {
	// Primero, parsea el token

	token, err := parseToken(tokenString)
	if err != nil {
		return "", err
	}
	// Luego, accede a los claims del token para obtener los datos que necesitas
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("Invalid claims")
	}
	nameUser, ok := claims["username"].(string)
	if !ok {
		return "", fmt.Errorf("Invalid nameUser")
	}
	return nameUser, nil
}
