package engine

import (
	"native-chi-mariadb/model"
	"os"
	"text/template"
)

func getTemplateHTML() string {
	return `
	<DOCTYPE! html>
	<html lang=en>

	<body>
		<!-- Create Table -->
		<form action="/{{.Name}}" method="post" enctype="multipart/form-data">
		{{range $nama,$tipe := .Elements}}
			<div>
				<label>{{$tipe.Label}}</label>
				<input name="{{$tipe.Label}}" type="{{$tipe.DataType}}">
			</div>
		{{end}}</form>
		<!-- Update Text -->
		<form action="/{{.Name}}" method="put">
		{{range $nama,$tipe := .Elements}} 
			{{if (ne $tipe.DataType "file")}}
				<div>
					<label>{{$tipe.Label}}</label>
					<input name="{{$tipe.Label}}" type="{{$tipe.DataType}}">
				</div>
			{{end}}
		{{end}}
		</form>
		<!-- Update Image -->
		<form action="/{{.Name}}" method="put">
		{{range $nama,$tipe := .Elements}}
			{{if (eq $tipe.DataType "file")}}
				<div>
					<label>{{$tipe.Label}}</label>
					<input name="{{$tipe.Label}}" type="file">
				</div>
			{{end}}
		{{end}}
		</form>
		<!-- Delete Table -->
		<form action="/{{.Name}}" method="delete">{{range $nama,$tipe := .Elements}}
			<div>
				<label>{{$tipe.Label}}</label>
				<input name="{{$tipe.Label}}" type="{{$tipe.DataType}}">
			</div>{{end}}</form>
	</body>

	</html>
`
}

func MakeHTMLForm(b model.Pair) error {
	t := template.Must(template.New("oo").Parse(getTemplateHTML()))
	return t.Execute(os.Stdout, b)
}
