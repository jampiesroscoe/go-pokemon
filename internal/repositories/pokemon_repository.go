package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"go-pokemon/internal/core/domain"
	"go-pokemon/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"io"
	"net/http"
	"time"
)

const (
	MongoClientTimeout = 5
)

type PokemonRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

var _ ports.IPokemonRepository = (*PokemonRepository)(nil)

func NewPokemonRepository(conn string) *PokemonRepository {
	ctx, cancelFunc := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancelFunc()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		conn,
	))
	if err != nil {
		fmt.Println(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connected to MongoDB")
	}
	return &PokemonRepository{
		client:     client,
		database:   client.Database("go-pokemon"),
		collection: client.Database("go-pokemon").Collection("pokemon"),
	}
}

func (r *PokemonRepository) GetPokemon(name string) (domain.Pokemon, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancelFunc()

	var result domain.Pokemon
	err := r.collection.FindOne(ctx, bson.M{"id": name}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No pokemon found in your mongodb collection with the name %s\n", name)
		fmt.Printf("Finding %s from pokemon api", name)
	}
	return result, err
}

func (r *PokemonRepository) StorePokemon(pokemon domain.Pokemon, name string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancelFunc()
	_, err := r.collection.InsertOne(ctx, bson.M{"id": name, "pokemonData": pokemon})
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (r *PokemonRepository) FindInWild(name string) (domain.Pokemon, error) {
	var pokemon domain.Pokemon

	response, responseErr := http.Get("http://pokeapi.co/api/v2/pokemon/" + name)
	if responseErr != nil {
		fmt.Println(responseErr)
	}

	body, parseErr := io.ReadAll(response.Body)
	if parseErr != nil {
		fmt.Println(parseErr)
	}

	jsonErr := json.Unmarshal(body, &pokemon)
	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	return pokemon, jsonErr
}
