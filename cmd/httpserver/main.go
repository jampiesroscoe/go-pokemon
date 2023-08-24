package main

import (
	"go-pokemon/internal/core/services"
	"go-pokemon/internal/handlers"
	"go-pokemon/internal/repositories"
	"go-pokemon/internal/server"
)

func main() {
	pokemonRepository := repositories.NewPokemonRepository("mongodb+srv://pokemon:charizard@cluster0.014fu.mongodb.net/?retryWrites=true&w=majority")
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
