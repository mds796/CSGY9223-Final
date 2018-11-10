package web

import (
	"github.com/mds796/CSGY9223-Final/post"
	"io/ioutil"
	"net/http"
)

func (srv *HttpService) FetchFeed(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
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
	defer r.Body.Close()

	err := srv.createPost(r)

	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
func (srv *HttpService) createPost(r *http.Request) error {
	response, err := srv.verifyToken(r)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	request := post.CreatePostRequest{UserID: response.UserID, Text: string(bytes)}

	_, err = srv.PostService.Create(request)

	return err
}
