package main

import (
	"fmt"
	"net/http"
	"main/path"
)

func main() {
	http.HandleFunc("/api/login", path.Login)

	fmt.Print("listen 8080")
	http.ListenAndServe(":8080",nil)
}