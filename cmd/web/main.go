package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/girdhar1982/go-learning/config"
	"github.com/girdhar1982/go-learning/pkg/handlers"
	"github.com/girdhar1982/go-learning/pkg/render"
)

const portNumber = ":8080"
// each file must have main
func main() {
 var app config.AppConfig;
 //get the TemplateCache for the app config
  tc, err := render.CreateTemplateCache();
	app.TemplateCache=tc;
	app.UseCache=false;
	repo:=handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	if err != nil {log.Fatal(err)}
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println("Starting Application on Port: ",portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
