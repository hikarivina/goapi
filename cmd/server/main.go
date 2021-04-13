package main

import (
	"fmt"
	"net/http"

	transportHTTP "go-rest/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {
	fmt.Printf("Setting Up Our App")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {

	fmt.Println("Go Rest API")

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our RESR API")
		fmt.Println(err)
	}
}
