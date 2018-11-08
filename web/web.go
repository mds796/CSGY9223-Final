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

	multiplexer.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/#/login", 307)
	})
	multiplexer.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "username", Value: "mds796", Expires: time.Now().Add(24 * time.Hour)})
		http.Redirect(w, r, "/", 307)
	})
	multiplexer.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "username", Value: "", Expires: time.Unix(0, 0)})
		http.Redirect(w, r, "/", 307)
	})
	multiplexer.HandleFunc("/feed", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.WriteHeader(200)
			w.Write([]byte(`
			{ 
				"posts":[
                	{"name": "fake123", "text": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque ultrices leo sollicitudin nisl facilisis imperdiet. Nam a pellentesque enim. Donec sollicitudin placerat semper. Nam non neque quam. Suspendisse nec mauris rutrum dolor accumsan pellentesque nec vel tortor. Interdum et malesuada fames ac ante ipsum primis in faucibus. Cras et quam viverra nunc vulputate euismod nec in nisi. In vehicula faucibus erat, id ullamcorper sapien. Maecenas eu tristique ligula, a tempus ipsum. Nam vel pretium sed."},
                	{"name": "fake123", "text": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque ultrices leo sollicitudin nisl facilisis imperdiet. Nam a pellentesque enim. Donec sollicitudin placerat semper. Nam non neque quam. Suspendisse nec mauris rutrum dolor accumsan pellentesque nec vel tortor. Interdum et malesuada fames ac ante ipsum primis in faucibus. Cras et quam viverra nunc vulputate euismod nec in nisi. In vehicula faucibus erat, id ullamcorper sapien. Maecenas eu tristique ligula, a tempus ipsum. Nam vel pretium sed."}
            	]
			}
			`))
		}
	})
	multiplexer.HandleFunc("/follow", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {

		} else if r.Method == http.MethodPost {

		} else if r.Method == http.MethodDelete {

		}
	})
	multiplexer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(staticPath, r.URL.Path))
	})

	log.Printf("Now listening on %v port %v.\n", host, port)
	log.Fatal(http.ListenAndServe(host+":"+strconv.Itoa(int(port)), multiplexer))
}
