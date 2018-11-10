package web

import (
	"github.com/mds796/CSGY9223-Final/auth"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"time"
)

func (srv *HttpService) RegisterUser(w http.ResponseWriter, r *http.Request) {
	username, password, err := getUserAndPassword(r.Body, true)
	if err == nil {
		response, err := srv.AuthService.Register(auth.RegisterAuthRequest{Username: username, Password: password})

		if err == nil {
			http.SetCookie(w, &http.Cookie{Name: "error", Value: "", Expires: time.Unix(0, 0)})
			http.SetCookie(w, &http.Cookie{Name: "username", Value: username, Expires: time.Now().Add(1 * time.Minute)})
			http.SetCookie(w, &response.Cookie)

			http.Redirect(w, r, "/", 307)
		} else {
			setErrorCookie(w, "Invalid register request.")
			http.Redirect(w, r, "/#/register", 307)
		}
	} else {
		setErrorCookie(w, "Invalid register request.")
		http.Redirect(w, r, "/#/register", 307)
	}
}

func getUserAndPassword(reader io.ReadCloser, repeatPassword bool) (username string, password string, err error) {
	values, err := getParameters(reader)

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
	if repeatPassword && err != nil {
		return "", "", err
	}

	if repeatPassword && password != password2 {
		return "", "", errors.New("The provided passwords do not match.")
	}

	return username, password, nil
}

func setErrorCookie(w http.ResponseWriter, message string) {
	http.SetCookie(w, &http.Cookie{Name: "error", Value: message, Expires: time.Now().Add(5 * time.Minute)})
}

func (srv *HttpService) LogInUser(w http.ResponseWriter, r *http.Request) {
	username, password, err := getUserAndPassword(r.Body, false)
	if err == nil {
		response, err := srv.AuthService.Login(auth.LoginAuthRequest{Username: username, Password: password})

		if err == nil {
			http.SetCookie(w, &http.Cookie{Name: "error", Value: "", Expires: time.Unix(0, 0)})
			http.SetCookie(w, &http.Cookie{Name: "username", Value: username, Expires: time.Now().Add(1 * time.Minute)})
			http.SetCookie(w, &response.Cookie)

			http.Redirect(w, r, "/", 307)
		} else {
			setErrorCookie(w, "Invalid login request.")
			http.Redirect(w, r, "/#/login", 307)
		}
	} else {
		setErrorCookie(w, "Invalid login request.")
		http.Redirect(w, r, "/#/login", 307)
	}
}

func (srv *HttpService) LogOutUser(w http.ResponseWriter, r *http.Request) {
	username, err := srv.logout(r)
	if err == nil {
		expireCookie(w, username)
		expireCookie(w, "username")
		expireCookie(w, "error")
	} else {
		setErrorCookie(w, "Invalid logout request.")
	}

	http.Redirect(w, r, "/", 307)
}

func (srv *HttpService) logout(r *http.Request) (string, error) {
	username, token, err := getUsernameAndToken(r)
	if err != nil {
		return "", err
	}

	_, err = srv.AuthService.Logout(auth.LogoutAuthRequest{Username: username, Cookie: *token})
	if err != nil {
		return "", err
	}

	return username, nil
}

func expireCookie(w http.ResponseWriter, name string) {
	http.SetCookie(w, &http.Cookie{Name: name, Value: "", Expires: time.Unix(0, 0)})
}

func getUsernameAndToken(r *http.Request) (username string, token *http.Cookie, err error) {
	usernameCookie, err := r.Cookie("username")
	if err != nil {
		return "", nil, err
	}

	tokenCookie, err := r.Cookie(usernameCookie.Value)
	if err != nil {
		return usernameCookie.Value, tokenCookie, err
	}

	return usernameCookie.Value, tokenCookie, nil
}

func (srv *HttpService) verifyToken(r *http.Request) (string, error) {
	_, token, err := getUsernameAndToken(r)
	if err != nil {
		return "", err
	}

	response, err := srv.AuthService.Verify(auth.VerifyAuthRequest{Cookie: *token})
	if err != nil {
		return "", err
	}

	return response.Username, nil
}
