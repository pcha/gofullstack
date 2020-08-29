package templatesUtil

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}
	_ = t.Execute(w, templateData)
}
