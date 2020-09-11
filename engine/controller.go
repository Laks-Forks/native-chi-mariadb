package engine

import (
	"native-chi-mariadb/model"
	"os"
	"text/template"
)

func getTemplateController() string {
	return `
package controller

func ReadAll{{.Name}}s(w http.ResponseWriter,r *http.Request){
	var
}

func Read{{.Name}}(w http.ResponseWriter,r *http.Request){
	var
}

func Create{{.Name}}(w http.ResponseWriter,r *http.Request){
	
}

func Update{{.Name}}(w http.ResponseWriter,r *http.Request){
	
}

func Delete{{.Name}}(w http.ResponseWriter,r *http.Request){
	
}
`
}

func MakeControllerForm(b model.Pair) error {
	t := template.Must(template.New("oo").Parse(getTemplateController()))
	return t.Execute(os.Stdout, b)
}
