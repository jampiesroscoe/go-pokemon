package ports

import (
	"github.com/gofiber/fiber/v2"
)

type IPokemonService interface {
	IChooseYou(name string) []byte
}

type IPokemonRepository interface {
	IChooseYou(name string) []byte
}

type IPokemonHandlers interface {
	IChooseYou(c *fiber.Ctx) error
}
