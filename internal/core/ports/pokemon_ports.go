package ports

import (
	"github.com/gofiber/fiber/v2"
)

type IPokemonService interface {
	IChooseYou(name string) []byte
}

type IPokemonRepository interface {
	StorePokemon(pokemon []byte, name string) error
	GetPokemon(name string) ([]byte, error)
	FindInWild(name string) ([]byte, error)
}

type IPokemonHandlers interface {
	IChooseYou(c *fiber.Ctx) error
}
