package main

import (
	"fmt"
	"log"
	"net/http"
)


func hissFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hiss" {
		http.Error(w, "404 not found!", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Sorry, this method is not supported.", http.StatusNotFound)
		return
	}
	
	fmt.Fprintf(w, "HISSSSSS!!!")
}

func inpHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request success")
	meowName := r.FormValue("name")
	meowEmail := r.FormValue("email")
	meowMessage := r.FormValue("message")
	fmt.Fprintf(w, "Name = %s\n", meowName)
	fmt.Fprintf(w, "Email = %s\n", meowEmail)
	fmt.Fprintf(w, "Message = %s\n", meowMessage)
}

func main() {
	fileServer := http.FileServer(http.Dir("./src"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hiss", hissFunc)
	http.HandleFunc("/form", inpHandler)

	fmt.Printf("Server started at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}