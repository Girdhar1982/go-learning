package handlers

import (
	"net/http"

	"github.com/girdhar1982/go-learning/config"
	"github.com/girdhar1982/go-learning/pkg/render"
)
//TemplateData holds data sent from template to  templates
type TemplateData struct {
StringMap map[string]string
IntMap map[string]int
FloatMap map[string]float32
Data map[string] interface{}
CSRFToken string
Flash string
Warning string
Error string
}

//Repo the repository used by the handlers
var Repo *Respository;

//Repository is the repository type
type Respository struct{
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Respository{
	return &Respository{ App: a	}
}


//New Handlers sets the repository for the handlers
func NewHandlers(r *Respository){
	Repo = r
}

func (m *Respository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "home.page.tmpl",&TemplateData{})
}

func (m *Respository)  About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplates(w, "about.page.tmpl",&TemplateData{})
}
