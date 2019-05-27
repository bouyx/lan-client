package path

import (
	"main/handler"
	"encoding/json"
	"net/http"
)
type pong struct {
	ping string
}

func Login(w http.ResponseWriter, r *http.Request){
	if(r.Method == "POST"){
		w.Header().Set("Content-Type", "application/json")
		if (handler.Login(r.URL.Query()["password"][0]) != nil){
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(map[string]string{"message": "not logged"})
		}else{
					
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"message": "logged"})
		}
	}
}