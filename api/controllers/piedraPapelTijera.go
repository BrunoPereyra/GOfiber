package controllers

import (
	"fmt"
	"math/rand"

	"github.com/gofiber/fiber/v2"
)

type OpcionUser struct {
	OpcionUser int `json:"OpcionUsers"`
}

func PiedraPapelTijera(c *fiber.Ctx) error {
	var userOption OpcionUser

	if err := c.BodyParser(&userOption); err != nil {
		return c.JSON(fiber.Map{
			"res": "errror body parcer",
		})
	}
	valueOpcionUser := userOption.OpcionUser
	fmt.Println(valueOpcionUser)

	opcion := make([]string, 3)
	copy(opcion, []string{"piedra", "papel", "tijera"})

	if !(valueOpcionUser > 0) || !(valueOpcionUser <= 3) {
		return c.JSON(fiber.Map{
			"res": "eleji entre 1,2,3 ",
		})
	}
	computerOption := rand.Intn(len(opcion))
	fmt.Println(computerOption)
	fmt.Println(len(opcion), "LEN")

	if valueOpcionUser == computerOption {
		return c.JSON(fiber.Map{
			"res": "empate",
			"A":   userOption,
		})
	} else if (valueOpcionUser == 0 && computerOption == 2) || (valueOpcionUser == 1 && computerOption == 0) || (valueOpcionUser == 2 && computerOption == 1) {
		return c.JSON(fiber.Map{
			"res": "perdiste",
			"A":   userOption,
		})
	} else {
		return c.JSON(fiber.Map{
			"res": "ganaste",
			"A":   userOption,
		})
	}
}
