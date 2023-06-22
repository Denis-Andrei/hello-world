package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Denis-Andrei/goapp/pkg/config"
	"github.com/Denis-Andrei/goapp/pkg/handlers"
	"github.com/Denis-Andrei/goapp/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.Appconfig
var session *scs.SessionManager
func main() {

	//change this to true when in prod
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true //this will persist the cookie even if the user exits the browser
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // we set it to false because on local http is not an encrypted conexion

	app.Session = session

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
