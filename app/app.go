package app

import (
	"KVStore/bigcache"
	c "KVStore/config"
	"KVStore/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type App struct {
	Config c.Configurations
	Router *mux.Router
}

func (a *App) Initialize() {
	env := os.Getenv("ENV")

	if env == "DEV" {
		viper.SetConfigName("config.development")
	} else {
		log.Fatalf("Unknown environment: %s", env)
	}
	log.Printf("Enviroment: %s", env)

	viper.SetConfigName("config.development")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	var configuration c.Configurations
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to read config into struct, %v", err)
	}

	config := bigcache.Config{
		Ttl: time.Duration(configuration.Cache.TTLMinutes) * time.Minute,
	}

	cacheClient, err := bigcache.New(config)

	if err != nil {
		log.Println("Failed to initialize cache...")
		log.Fatal(err)
	}

	log.Printf("Cache initialized with TTL %d minutes", configuration.Cache.TTLMinutes)

	handler := handlers.StoreHandler{
		Store: cacheClient,
	}
	r := mux.NewRouter()
	r.HandleFunc("/{key}", handler.Get).Methods(http.MethodGet)
	r.HandleFunc("/{key}", handler.Post).Methods(http.MethodPost).HeadersRegexp("Content-Type", "text/plain;charset=utf-8")
	r.HandleFunc("/", handler.NotImplemented)

	a.Config = configuration
	a.Router = r
}

func (a *App) Run() {
	log.Printf("Starting server on port :%d", a.Config.Server.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(a.Config.Server.Port), a.Router))
}
