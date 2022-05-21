package main

import (
	"html/template"
	"log"
	"net/http"
)

type CountData struct {
	PageTitle string
	Counter   int
}

var c int = 0

func main() {
	fs := http.FileServer(http.Dir("css"))
	sp := http.StripPrefix("/css/", fs)
	http.Handle("/css/", sp)

	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := CountData{
		PageTitle: "My Counter",
		Counter:   c,
	}

	index, err := template.New("").ParseFiles(
		"index.html",
	)
	if err != nil {
		log.Println("Error creating new template.", err)
	}

	switch r.URL.Path {
	case "/counter/":
		log.Println("Handling: /counter/")
		err = index.ExecuteTemplate(w, "T", data)
	case "/counter/increment":
		log.Println("Handling: /counter/increment")
		c++
		data.Counter = c
		err = index.ExecuteTemplate(w, "V", data)
	}
	if err != nil {
		log.Printf("Template execution error: %v", err)
	}

}

/*==============================================================*/
