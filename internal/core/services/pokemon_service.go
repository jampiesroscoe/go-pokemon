package services

import (
	"fmt"
	"go-pokemon/internal/core/ports"
)

type PokemonService struct {
	pokemonRepository ports.IPokemonRepository
}

// This line is for get feedback in case we are not implementing the interface correctly
var _ ports.IPokemonService = (*PokemonService)(nil)

func NewPokemonService(repository ports.IPokemonRepository) *PokemonService {
	return &PokemonService{
		pokemonRepository: repository,
	}
}

func (s *PokemonService) IChooseYou(name string) []byte {
	response, err := s.pokemonRepository.GetPokemon(name)
	if err != nil {
		response, err := s.pokemonRepository.FindInWild(name)
		if err != nil {
			fmt.Println(err)
		}
		erro := s.pokemonRepository.StorePokemon(response, name)
		if erro != nil {
			fmt.Println(erro)
		}
		return response
	}
	return response
}
