package web

import (
	"io/ioutil"
	"log"
	"net/http"
)

func (srv *HttpService) FetchFeed(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(200)
		w.Write([]byte(`
				{
					"feed":[
	                	{"name": "fake123", "text": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque ultrices leo sollicitudin nisl facilisis imperdiet. Nam a pellentesque enim. Donec sollicitudin placerat semper. Nam non neque quam. Suspendisse nec mauris rutrum dolor accumsan pellentesque nec vel tortor. Interdum et malesuada fames ac ante ipsum primis in faucibus. Cras et quam viverra nunc vulputate euismod nec in nisi. In vehicula faucibus erat, id ullamcorper sapien. Maecenas eu tristique ligula, a tempus ipsum. Nam vel pretium sed."},
	                	{"name": "fake234", "text": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque ultrices leo sollicitudin nisl facilisis imperdiet. Nam a pellentesque enim. Donec sollicitudin placerat semper. Nam non neque quam. Suspendisse nec mauris rutrum dolor accumsan pellentesque nec vel tortor. Interdum et malesuada fames ac ante ipsum primis in faucibus. Cras et quam viverra nunc vulputate euismod nec in nisi. In vehicula faucibus erat, id ullamcorper sapien. Maecenas eu tristique ligula, a tempus ipsum. Nam vel pretium sed."}
	            	]
				}
				`))
	}
}

func (srv *HttpService) MakePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if bytes, err := ioutil.ReadAll(r.Body); err == nil {
			log.Println(string(bytes))
		}
	}
}
