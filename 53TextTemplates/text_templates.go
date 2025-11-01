// Go offers built-in support for creating dynamic content or showing customized
// output to the user with the `text/template` package. A sibling package
// named `html/template` provides the same API but has additional security
// features and should be used for generating HTML.

package main

import (
	"os"
	"text/template"
)

func main() {

	// We can create a new template and parse its body from
	// a string.
	// Templates are a mix of static text and "actions" enclosed in
	// `{{...}}` that are used to dynamically insert content.

	// TODO: Create t1 := template.New("t1")
	// TODO: Parse t1 with "Value is {{.}} " and handle error
	t1, err := template.New("t1").Parse("Value is {{.}}")
	if err != nil {
		panic(err)
	}

	// Alternatively, we can use the `template.Must` function to
	// panic in case `Parse` returns an error. This is especially
	// useful for templates initialized in the global scope.

	// TODO: Use template.Must to parse t1 with "Value: {{.}} " and handle error
	t1 = template.Must(template.New("t1").Parse("Value: {{.}}"))
	// By "executing" the template we generate its text with
	// specific values for its actions. The `{{.}}` action is
	// replaced by the value passed as a parameter to `Execute`.

	// TODO: Execute t1 with "some text"
	// TODO: Execute t1 with 5
	// TODO: Execute t1 with []string{"Go", "Rust", "C++", "C#"}
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})
	

	// Helper function we'll use below.

	// TODO: Create Create function with name and t string that returns template.Must(template.New(name).Parse(t))
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// If the data is a struct we can use the `{{.FieldName}}` action to access
	// its fields. The fields should be exported to be accessible when a
	// template is executing.

	// TODO: Create t2 := Create("t2", "Name: {{.Name}} ")

	t2 := Create("t2", "Name: {{.Name}}")

	// TODO: Execute t2 with struct {Name string}{"Jane Doe"}
	t2.Execute(os.Stdout, struct {Name string}{"Jane Doe"})

	// The same applies to maps; with maps there is no restriction on the
	// case of key names.

	// TODO: Execute t2 with map[string]string{"Name": "Mickey Mouse"}
	t2.Execute(os.Stdout, map[string]string{"Name": "Mickey Mouse"})

	// if/else provide conditional execution for templates. A value is considered
	// false if it's the default value of a type, such as 0, an empty string,
	// nil pointer, etc.
	// This sample demonstrates another
	// feature of templates: using `-` in actions to trim whitespace.


	// TODO: Create t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}")
	// TODO: Execute t3 with "not empty"
	// TODO: Execute t3 with ""
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")
	// range blocks let us loop through slices, arrays, maps or channels. Inside
	// the range block `{{.}}` is set to the current item of the iteration.

	// TODO: Create t4 := Create("t4", "Range: {{range .}}{{.}} {{end}} ")
	// TODO: Execute t4 with []string{"Go", "Rust", "C++", "C#"}
	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}")	
	t4.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})
}