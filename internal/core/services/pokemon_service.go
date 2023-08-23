package services

import (
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
	response := s.pokemonRepository.IChooseYou(name)
	return response
}
