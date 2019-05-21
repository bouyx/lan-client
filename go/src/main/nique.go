package handler

import (
	"fmt"
	"net/http"
)

func Nique(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "prout")
}
