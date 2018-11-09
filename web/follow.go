package web

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (srv *HttpService) ToggleFollow(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(200)
		w.Write([]byte(`
								{ 
									"follows":[
					                	{"name": "fake123", "followed": true},
					                	{"name": "fake234", "followed": false}
					            	]
								}
								`))
	} else if r.Method == http.MethodPost {

	} else if r.Method == http.MethodDelete {

	}
}

func (srv *HttpService) ListFollows(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if bytes, err := ioutil.ReadAll(r.Body); err == nil {
			follow := new(Follow)
			json.Unmarshal(bytes, follow)
			log.Printf("%v", follow)
		}
	}
}
