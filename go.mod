module backend

go 1.16

require (
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/gofiber/fiber/v2 v2.43.0
	github.com/leodido/go-urn v1.2.2 // indirect
	go.mongodb.org/mongo-driver v1.11.3
	golang.org/x/crypto v0.7.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)

replace backend => ./ // Agregamos esta línea para que Go sepa que el módulo está en el directorio actual
