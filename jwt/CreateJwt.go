package jwt

import (
	"backend/api/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// Define los datos del usuario

func TokenCreate(user models.User) (string, error) {
	fmt.Println(user)
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.NameUser,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	// Crea un token con los datos del usuario y la firma especificada
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firma el token con una clave secreta
	signedToken, err := token.SignedString([]byte("mi_clave_secreta"))

	return signedToken, err
}
