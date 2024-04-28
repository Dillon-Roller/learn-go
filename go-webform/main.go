package main

import (
    "fmt"
    "html/template"
    "net/http"
)

// Define a template.
const formTemplateHTML = `
<!DOCTYPE html>
<html>
<head>
    <title>Simple Form</title>
</head>
<body>
    <h1>Input Form</h1>
    <form method="POST" action="/submit">
        <label for="input">Enter something:</label>
        <input type="text" id="input" name="input">
        <input type="submit" value="Submit">
    </form>
    {{if .}}
    <h2>Received:</h2>
    <p>{{.}}</p>
    {{end}}
</body>
</html>
`

func formHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.New("webpage").Parse(formTemplateHTML)
    if err != nil {
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        return
    }
    t.Execute(w, nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    r.ParseForm()
    input := r.FormValue("input")

    t, err := template.New("webpage").Parse(formTemplateHTML)
    if err != nil {
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        return
    }
    t.Execute(w, input)
}

func main() {
    http.HandleFunc("/", formHandler)
    http.HandleFunc("/submit", submitHandler)
    fmt.Println("Server starting at port 8080...")
    http.ListenAndServe(":8080", nil)
}
