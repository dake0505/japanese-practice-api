package firebase

import (
	"context"
	"encoding/json"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func InitFirebase() *firebase.App {
	configJSON := os.Getenv("FIREBASE_CONFIG")
	if configJSON == "" {
		log.Fatalf("FIREBASE_CONFIG environment variable is not set")
	}

	var config map[string]interface{}
	err := json.Unmarshal([]byte(configJSON), &config)
	if err != nil {
		log.Fatalf("Error parsing FIREBASE_CONFIG: %v", err)
	}

	opt := option.WithCredentialsJSON([]byte(configJSON))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	log.Println("Firebase app initialized successfully")
	return app
}
