package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifica que el algoritmo de firma sea HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid signing method")
		}
		// Debe devolver la clave secreta que se usó para firmar el token
		return []byte("secreto"), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}
	return token, nil
}

func ExtractDataFromToken(tokenString string) (string, error) {
	// Primero, parsea el token
	token, err := ParseToken(tokenString)
	if err != nil {
		return "", err
	}

	// Luego, accede a los claims del token para obtener los datos que necesitas
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("Invalid claims")
	}
	nameUser, ok := claims["nameUser"].(string)
	if !ok {
		return "", fmt.Errorf("Invalid nameUser")
	}
	return nameUser, nil
}
