package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	banner "./src"
)

type ViewData struct {
	Data    string
	OldData string
	Font    string
}

func main() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	fs = http.FileServer(http.Dir("images"))
	http.Handle("/images/", http.StripPrefix("/images/", fs))
	data := ViewData{}
	data.Font = "Standard"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			tmpl, _ := template.ParseFiles("./html/ERROR.html")
			tmpl.Execute(w, nil)
			return
		}

		tmpl, _ := template.ParseFiles("./html/index.html")

		if r.Method == "POST" {

			data = ViewData{Data: banner.Get(r.FormValue("Text"), r.FormValue("fs")), OldData: r.FormValue("Text"), Font: r.FormValue("fs")}
			// data = ViewData{Data: "Wow"}
			file, _ := os.Create("./files/output.txt")
			file.WriteString(data.Data)
			tmpl.Execute(w, data)
		}
		if r.Method == "GET" {
			tmpl.Execute(w, nil)
		}
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}

// banner.Get(w, text, "standard")
