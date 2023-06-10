package main

import (
	"fmt"
	"os"
	"text/template"
)

// here we are defining a template using define keyword and naming in "T", it's kind of like an html only
func main() {
	t, err := template.New("test").Parse(`{{define "T"}} Hello {{.}} {{end}}`)
	if err != nil {
		fmt.Println(err)
	}

	err = t.ExecuteTemplate(os.Stdout, "T", "World")
	if err != nil {
		fmt.Println(err)
	}
	// os.Stdout means command line

}
