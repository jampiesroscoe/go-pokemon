package main

import (
	"go-pokemon/internal/core/services"
	"go-pokemon/internal/handlers"
	"go-pokemon/internal/repositories"
	"go-pokemon/internal/server"
)

func main() {
	pokemonRepository := repositories.NewPokemonRepository()
	//services
	pokemonService := services.NewPokemonService(pokemonRepository)
	//handlers
	pokemonHandlers := handlers.NewPokemonHandlers(pokemonService)
	//server
	httpServer := server.NewServer(
		pokemonHandlers,
	)
	httpServer.Initialize()
}
