package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tobiaGasparoni/Go-Webapp/pkg/config"
	"github.com/tobiaGasparoni/Go-Webapp/pkg/handlers"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":5000"

var app config.AppConfig

var session *scs.SessionManager

// main is the main application function
func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	server := &http.Server{
		Addr:    portNumber,
		Handler: router(&app),
	}

	err := server.ListenAndServe()
	log.Fatal(err)
}
