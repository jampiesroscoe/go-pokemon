package repositories

import (
	"go-pokemon/internal/core/ports"
	"io"
	"log"
	"net/http"
)

type PokemonRepository struct {
}

var _ ports.IPokemonRepository = (*PokemonRepository)(nil)

func NewPokemonRepository() *PokemonRepository {
	return &PokemonRepository{}
}

func (r *PokemonRepository) IChooseYou(name string) []byte {
	response, err := http.Get("http://pokeapi.co/api/v2/pokemon/" + name)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(`here`)
	}
	return body
}
