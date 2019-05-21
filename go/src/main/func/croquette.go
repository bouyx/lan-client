package handler

import (
	"fmt"
	"net/http"
)


func Croc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"quette")
}
