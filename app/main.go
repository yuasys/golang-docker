package main

import(
	"fmt"
	"html"
	"log"
	"net/http"
)

func main(){
	http.HndleFunc("/", func(w http.ResponseWriter, *http.
	Request){
		fmt.Fprintf(w, "Hi Hello, %q", html:EscapeString(r.Path)
		)
	})

	log.Fatal(http.ListenAndServe(":8080",nil))
}