package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Denis-Andrei/goapp/pkg/config"
	"github.com/Denis-Andrei/goapp/pkg/handlers"
	"github.com/Denis-Andrei/goapp/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app config.Appconfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))

	srv := &http.Server{
		Addr: portNumber, 
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
