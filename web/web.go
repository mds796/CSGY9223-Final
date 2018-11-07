package web

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func Start(host string, port uint16, staticPath string) {
	const redirect = 307

	multiplexer := http.NewServeMux()

	multiplexer.Handle("/feed", http.RedirectHandler("/", redirect))
	multiplexer.Handle("/follow", http.RedirectHandler("/", redirect))
	multiplexer.Handle("/about", http.RedirectHandler("/", redirect))
	multiplexer.Handle("/login", http.RedirectHandler("/", redirect))
	multiplexer.Handle("/register", http.RedirectHandler("/", redirect))

	multiplexer.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		if r.Body != nil {
			if _, err := io.Copy(w, r.Body); err != nil {
				log.Printf("Encountered an error while echoing the body: %v\n.", err)
			}
		}
	})

	multiplexer.Handle("/", http.FileServer(http.Dir(staticPath)))

	log.Printf("Now listening on port %v.\n", port)
	log.Fatal(http.ListenAndServe(host+":"+strconv.Itoa(int(port)), multiplexer))
}
