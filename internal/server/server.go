package server

import (
	"github.com/gofiber/fiber/v2"
	"go-pokemon/internal/core/ports"
	"log"
)

type Server struct {
	//We will add every new Handler here
	pokemonHandlers ports.IPokemonHandlers
	//middlewares ports.IMiddlewares
	//paymentHandlers ports.IPaymentHandlers
}

func NewServer(pHandlers ports.IPokemonHandlers) *Server {
	return &Server{
		pokemonHandlers: pHandlers,
		//paymentHandlers: pHandlers
	}
}

func (s *Server) Initialize() {
	app := fiber.New()
	app.Get(
		"pokemon/iChooseYou/:name",
		s.pokemonHandlers.IChooseYou,
	)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
