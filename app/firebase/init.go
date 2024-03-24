package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func InitFirebase() *auth.Client {
	opt := option.WithCredentialsFile("./config.json")
	print(opt, "opt----------------------------------")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
	authClient, err := app.Auth(context.Background())
	return authClient
}
