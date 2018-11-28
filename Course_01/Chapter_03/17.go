package main

import (
	"html/template"
	"log"
	"os"
)

const tpl = `{{.Name}}
------------------------
Age:	{{.Age}}
Title:	{{.Title}}
`

func main() {

	/* ---------------------------------
	John
	------------------------
	Age:	17
	Title:	student
	--------------------------------- */
	var t = template.Must(template.New("user").Parse(tpl))
	err := t.Execute(os.Stdout, map[string]string{"Name": "John", "Age": "17", "Title": "student"})
	if err != nil {
		log.Fatal(err)
	}
}
