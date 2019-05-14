package main

import (
	"fmt"
	"net/http"
	"time"
)

func Mime(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, time.Now().Format(time.RFC1123Z))
}

func Socisse(w http.ResponseWriter, r *http.Request){
	fmt.Printf("nique")
}