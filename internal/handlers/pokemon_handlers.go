package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-pokemon/internal/core/ports"
)

type PokemonHandlers struct {
	pokemonService ports.IPokemonService
}

var _ ports.IPokemonHandlers = (*PokemonHandlers)(nil)

func NewPokemonHandlers(pokemonService ports.IPokemonService) *PokemonHandlers {
	return &PokemonHandlers{
		pokemonService: pokemonService,
	}
}

func (h *PokemonHandlers) IChooseYou(c *fiber.Ctx) error {
	var name = c.Params("name")
	response := h.pokemonService.IChooseYou(name)
	return c.Send(response)
}
