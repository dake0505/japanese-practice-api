package firebase

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	// "firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func InitFirebase() *firebase.App {
	dir, err := os.Getwd() // 获取当前工作目录和可能的错误
	if err != nil {
		log.Fatalf("Error getting current directory: %v", err)
	}

	log.Println("Current working directory:", dir)
	opt := option.WithCredentialsFile("./config.json")
	log.Println(opt, "-----------------------------------opt")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}
	// println("InitFirebase", &app)
	// authClient, err := app.Auth(context.Background())
	return app
}
