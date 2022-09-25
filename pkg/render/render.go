package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//render template
func RenderTemplates_Old(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl,"./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}

var tc=make(map[string]*template.Template);
func RenderTemplates(w http.ResponseWriter, t string) {
var tmpl *template.Template;
var err error;
//check if template is already there in cache
_, inMap := tc[t]
if !inMap {
	log.Println("creating template and adding to cache")
	//need to create the template
	err=createTemplateCache(t)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}
}else{
	//return template
  log.Println("using cached template")
}
tmpl=tc[t]
err= tmpl.Execute(w,nil)
if err != nil {
	fmt.Println("error parsing template: ", err)
}
}
func createTemplateCache(t string) error{
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
//parse the template
tmpl,err:=template.ParseFiles(templates...)
if err != nil {
	fmt.Println("error parsing template: ", err)
	return err
}
//add template to cache (map)
tc[t]=tmpl
return nil
}