package web

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

func Start(host string, port uint16, staticPath string) {
	multiplexer := http.NewServeMux()

	multiplexer.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "username", Value: "mds796", Expires: time.Now().Add(24 * time.Hour)})
		http.Redirect(w, r, "/", 307)
	})
	multiplexer.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "username", Value: "", Expires: time.Unix(0, 0)})
		http.Redirect(w, r, "/", 307)
	})

	multiplexer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(staticPath, r.URL.Path))
	})

	log.Printf("Now listening on port %v.\n", port)
	log.Fatal(http.ListenAndServe(host+":"+strconv.Itoa(int(port)), multiplexer))
}
