package main

import (
	"fmt"
	"log"
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"myapp/pkg/renderer"
	"net/http"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tmpl_cache, err := renderer.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tmpl_cache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	renderer.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	server := &http.Server{
		Addr:    portNumber,
		Handler: router(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}
