package web

import (
	"net/http"
	"time"
)

func (srv *HttpService) LogOutUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "username", Value: "", Expires: time.Unix(0, 0)})
		http.Redirect(w, r, "/", 307)
	}
}

func (srv *HttpService) LogInUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "username", Value: "mds796", Expires: time.Now().Add(24 * time.Hour)})
		http.Redirect(w, r, "/", 307)
	}
}

func (srv *HttpService) RegisterUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/#/login", 307)
	}
}
