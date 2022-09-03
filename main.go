package main

import (
	"net/http"
)

func main() {
	repo := NewRepo()
	defer repo.Close()

	control := NewControl(repo)

	s := http.NewServeMux()

	s.HandleFunc("/newentry", control.newEntry)

	http.ListenAndServe(":8080", s)
}
