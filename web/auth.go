package web

import (
	"github.com/mds796/CSGY9223-Final/auth"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func (srv *HttpService) RegisterUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, err := getUserAndPassword(r.Body)
		if err == nil {
			response, err := srv.AuthService.Register(auth.RegisterAuthRequest{Username: username, Password: password})

			if err == nil {
				http.SetCookie(w, &http.Cookie{Name: "error", Value: "", Expires: time.Unix(0, 0)})
				http.SetCookie(w, &http.Cookie{Name: "username", Value: username, Expires: time.Now().Add(1 * time.Minute)})
				http.SetCookie(w, &response.Cookie)

				http.Redirect(w, r, "/#/login", 307)
			}
		}

		if err != nil {
			log.Println(err)
			setErrorCookie(w, "Invalid register request.")
			http.Redirect(w, r, "/#/register", 307)
		}
	}
}

func getUserAndPassword(reader io.ReadCloser) (username string, password string, err error) {
	bytes, err := ioutil.ReadAll(reader)

	if err != nil {
		return "", "", err
	}

	parameters := string(bytes)
	values, err := url.ParseQuery(parameters)

	if err != nil {
		return "", "", err
	}

	username, err = getKey(values, "username")
	if err != nil {
		return "", "", err
	}

	password, err = getKey(values, "password")
	if err != nil {
		return "", "", err
	}

	password2, err := getKey(values, "password2")
	if err != nil {
		return "", "", err
	}

	if password != password2 {
		return "", "", errors.New("The provided passwords do not match.")
	}

	return username, password, nil
}

func getKey(parameters map[string][]string, key string) (string, error) {
	values, ok := parameters[key]
	if !ok || len(values) == 0 || len(values[0]) == 0 {
		return "", errors.New("Please provide a " + key + ".")
	}

	return values[0], nil
}

func setErrorCookie(w http.ResponseWriter, message string) {
	http.SetCookie(w, &http.Cookie{Name: "error", Value: message, Expires: time.Now().Add(5 * time.Minute)})
}

func (srv *HttpService) LogInUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "username", Value: "mds796", Expires: time.Now().Add(24 * time.Hour)})
		http.Redirect(w, r, "/", 307)
	}
}

func (srv *HttpService) LogOutUser() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{Name: "username", Value: "", Expires: time.Unix(0, 0)})
		http.Redirect(w, r, "/", 307)
	}
}
