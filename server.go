package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("view").Parse(uploadPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("view").Parse(awsurl)
	file, header, err := r.FormFile("aws_file")
	if err != nil {
		log.Fatal(err)
	}
	data := struct {
		URL string
	}{
		uploadToAws(file, header),
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func webServer(host string, port int) {
	addr := fmt.Sprintf("%s:%d", host, port)

	http.HandleFunc("/aws/url", viewHandler)
	http.HandleFunc("/", rootHandler)

	log.Println("Starting web interface at", addr)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		log.Fatal("ERROR:", err)
	}
}
