package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	database "rest-api-go/database"
	repository "rest-api-go/repository"

	"github.com/gorilla/mux"
)

type Config struct {
	Port        string // Puerto donde se va a ejecutar
	JWTSecret   string // Llave secreta para generar tokens
	DatabaseUrl string // Conexión a la Base de datos
}

// Para que sea considerado servidor debe tener un método config - Las interfaces son contratos de implementación.
type Server interface {
	Config() *Config
}

// Encargado de manejar los servidores
type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

// Constructor
// context para entender errores en las go rutines
// Cuando se trabajoa con GoRutins, estas son bastantes independientes, aveces no sabemos porque estan fallando
// el ctx contexto sera clave para identificar este tipo de problemas
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("Port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("Secret is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("Database url is required")
	}

	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}

	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	repo, err := database.NewPostgresRepository(b.config.DatabaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	repository.SetRepository(repo)
	log.Println("Starting server on port", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAdnServer: ", err)
	}
}
