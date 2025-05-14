// The text/template package allows you to define templates with placeholders
// that can be replaced with actual values.
package basics

import (
	"fmt"
	"os"
	"text/template"
)

func TemplateString() {
	name := "Bob"
	age := 72
	message := fmt.Sprintf("Hello, %s! You are %d years old.", name, age)
	fmt.Println(message)
}

type Person struct {
	Name string
	Age  int
}

var person = Person{Name: "Bob", Age: 72}

func TemplateWithTextPackage() {

	tmpl := `Hello, {{.Name}}! You are {{.Age}} years old.`

	t, err := template.New("personTemplate").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, person)
	if err != nil {
		panic(err)
	}
}

func TemplateWithHTMLPackage() {

	tmpl := `<h1>Hello, {{.Name}}!</h1><p>You are {{.Age}} years old.</p>`

	t, err := template.New("personTemplate").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, person)
	if err != nil {
		panic(err)
	}
}
