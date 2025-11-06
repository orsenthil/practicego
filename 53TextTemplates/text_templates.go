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

	// Create t1 := template.New("t1")
	t1 := template.New("t1")
	// Parse t1 with "Value is {{.}} " and handle error
	t1, err := t1.Parse("Value is {{.}} ")
	if err != nil {
		panic(err)
	}

	// Alternatively, we can use the `template.Must` function to
	// panic in case `Parse` returns an error. This is especially
	// useful for templates initialized in the global scope.

	// Use template.Must to parse t1 with "Value: {{.}} " and handle error
	t1 = template.Must(template.New("t1").Parse("Value: {{.}} "))

	// By "executing" the template we generate its text with
	// specific values for its actions. The `{{.}}` action is
	// replaced by the value passed as a parameter to `Execute`.

	// Execute t1 with "some text"
	t1.Execute(os.Stdout, "some text")
	// TODO: Execute t1 with 5
	t1.Execute(os.Stdout, 5)
	// Execute t1 with []string{"Go", "Rust", "C++", "C#"}
	t1.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})
	

	// Helper function we'll use below.

	// Create Create function with name and t string that returns template.Must(template.New(name).Parse(t))
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}
	

	// If the data is a struct we can use the `{{.FieldName}}` action to access
	// its fields. The fields should be exported to be accessible when a
	// template is executing.

	// Create t2 := Create("t2", "Name: {{.Name}} ")
	t2 := Create("t2", "Name: {{.Name}} ")

	// Execute t2 with struct {Name string}{"Jane Doe"}
	t2.Execute(os.Stdout, struct {Name string}{"Jane Doe"})


	// The same applies to maps; with maps there is no restriction on the
	// case of key names.

	// Execute t2 with map[string]string{"Name": "Mickey Mouse"}
	t2.Execute(os.Stdout, map[string]string{"Name": "Mickey Mouse"})

	// if/else provide conditional execution for templates. A value is considered
	// false if it's the default value of a type, such as 0, an empty string,
	// nil pointer, etc.
	// This sample demonstrates another
	// feature of templates: using `-` in actions to trim whitespace.


	// Create t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}} ")
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}} ")
	// Execute t3 with "not empty"
	t3.Execute(os.Stdout, "not empty")
	// Execute t3 with ""
	t3.Execute(os.Stdout, "")

	// range blocks let us loop through slices, arrays, maps or channels. Inside
	// the range block `{{.}}` is set to the current item of the iteration.

	// Create t4 := Create("t4", "Range: {{range .}}{{.}} {{end}} ")
	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}} ")
	// Execute t4 with []string{"Go", "Rust", "C++", "C#"}
	t4.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})
	
}