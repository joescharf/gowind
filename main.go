package main

import (
	"bufio"
	"bytes"
	"html/template"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func main() {
	render()
	http.HandleFunc("/", indexHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("dist/css"))))
	http.ListenAndServe(":3000", nil)
}

// https://betterprogramming.pub/how-to-generate-html-with-golang-templates-5fad0d91252
func render() {
	allFiles := []string{"index.html"}

	var allPaths []string
	for _, tmpl := range allFiles {
		allPaths = append(allPaths, "./templates/"+tmpl)
	}

	templates := template.Must(template.ParseFiles(allPaths...))

	var processed bytes.Buffer
	templates.ExecuteTemplate(&processed, "page", nil)

	outputPath := "./dist/index.html"
	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(string(processed.Bytes()))
	w.Flush()
}
