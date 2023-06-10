package main

import (
	"fmt"
	"os"
	"text/template"
)

type Profile struct {
	FirstName string
	LastName  string
}

// this is a template creation
const tpl = `FirstName := {{.FirstName}} LastName := {{.LastName}}`
const tpl2 = `
	{{range .}}
		FirstName : {{.FirstName}} LastName : {{.LastName}}
	{{end}}
`

func main() {
	// passing a single data structure instance
	user := Profile{
		FirstName: "Abhinand",
		LastName:  "K R",
	}
	t := template.New("profile")

	te, err := t.Parse(tpl)
	if err != nil {
		fmt.Println(err)
	}

	err = te.Execute(os.Stdout, user)
	if err != nil {
		fmt.Println(err)
	}

	// passing a slice of struct instance
	users := []Profile{
		{
			FirstName: "abhinand",
			LastName:  "kr",
		},
		{
			FirstName: "athira",
			LastName:  "k radhakrishnan",
		},
	}

	// parsing the template in order for some data to be mapped to the template
	// we can only pass data to parsed template
	t2, err := t.Parse(tpl2)
	if err != nil {
		fmt.Println(err)
	}

	err = t2.Execute(os.Stdout, users)
	if err != nil {
		fmt.Println(err)
	}

}
