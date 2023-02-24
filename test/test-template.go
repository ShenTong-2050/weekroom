package main

import (
	"bytes"
	"fmt"
	"html/template"
)

func main() {

	var buf bytes.Buffer

	temp := template.Must(template.New("test").Parse("Hello {{.}} !"))

	temp.Execute(&buf,"world")

	fmt.Println(buf.String())
}
