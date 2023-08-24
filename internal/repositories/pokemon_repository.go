package repositories

import (
	"context"
	"fmt"
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

func (r *PokemonRepository) GetPokemon(name string) ([]byte, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancelFunc()

	var result []byte
	err := r.collection.FindOne(ctx, bson.M{"id": name}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No pokemon found in your collection with the name %s\n", name)
		fmt.Printf("Lets catch %s in the wild!", name)
	}
	return result, err
}

func (r *PokemonRepository) StorePokemon(pokemon []byte, name string) error {
	ctx, cancelFunc := context.WithTimeout(context.Background(), MongoClientTimeout*time.Second)
	defer cancelFunc()
	_, err := r.collection.InsertOne(ctx, bson.M{"id": name, "pokemonData": pokemon})
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func (r *PokemonRepository) FindInWild(name string) ([]byte, error) {
	response, err := http.Get("http://pokeapi.co/api/v2/pokemon/" + name)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body, err
}
