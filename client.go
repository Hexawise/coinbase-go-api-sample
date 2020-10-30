package main

import (
	"log"
	"os"

	coinbasepro "github.com/preichenberger/go-coinbasepro/v2"

	"github.com/joho/godotenv"
)

// NewClient creates a coinbasePro client and returns it for usage in API
func NewClient() *coinbasepro.Client {
	// Load .env file for secret values
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
	}

	// Create our coinbase pro client
	client := coinbasepro.NewClient()

	// And set the various options for it
	client.UpdateConfig(&coinbasepro.ClientConfig{
		BaseURL: "https://api-public.sandbox.pro.coinbase.com",
		Key: os.Getenv("COINBASE_KEY"),
		Passphrase: os.Getenv("COINBASE_PASSPHRASE"),
		Secret: os.Getenv("COINBASE_SECRET"),
	})

	return client
}