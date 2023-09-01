package ports

import (
	"github.com/gofiber/fiber/v2"
	"go-pokemon/internal/core/domain"
)

type IPokemonService interface {
	IChooseYou(name string) domain.Pokemon
}

type IPokemonRepository interface {
	StorePokemon(pokemon domain.Pokemon, name string) error
	GetPokemon(name string) (domain.Pokemon, error)
	FindInWild(name string) (domain.Pokemon, error)
}

type IPokemonHandlers interface {
	IChooseYou(c *fiber.Ctx) error
}
