package web

import (
	"context"
	"github.com/mds796/CSGY9223-Final/auth"

	"github.com/mds796/CSGY9223-Final/auth/authpb"
	"github.com/pkg/errors"
	"io"
	"log"
	"net/http"
	"time"
)

func (srv *HttpService) RegisterUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	username, password, err := getUserAndPassword(r.Body, true)

	if err == nil {
		response, err := srv.AuthService.Register(context.Background(), &authpb.RegisterAuthRequest{Username: username, Password: password})

		if err == nil {
			http.SetCookie(w, &http.Cookie{Name: "error", Value: "", Expires: time.Unix(0, 0)})
			http.SetCookie(w, &http.Cookie{Name: "username", Value: username, Expires: time.Now().Add(24 * time.Hour)})
			responseCookie := auth.DecodeCookie(response.Cookie)
			http.SetCookie(w, responseCookie)

			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		} else {
			log.Println(err)
			setErrorCookie(w, "Invalid register request.")
			http.Redirect(w, r, "/#/register", http.StatusTemporaryRedirect)
		}
	} else {
		log.Println(err)
		setErrorCookie(w, "Invalid register request.")
		http.Redirect(w, r, "/#/register", http.StatusTemporaryRedirect)
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
	http.SetCookie(w, &http.Cookie{Name: "error", Value: message, Expires: time.Now().Add(10 * time.Second)})
}

func (srv *HttpService) LogInUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	username, password, err := getUserAndPassword(r.Body, false)

	if err == nil {
		response, err := srv.AuthService.Login(context.Background(), &authpb.LoginAuthRequest{Username: username, Password: password})

		if err == nil {
			http.SetCookie(w, &http.Cookie{Name: "error", Value: "", Expires: time.Unix(0, 0)})
			http.SetCookie(w, &http.Cookie{Name: "username", Value: username, Expires: time.Now().Add(24 * time.Hour)})
			responseCookie := auth.DecodeCookie(response.Cookie)
			http.SetCookie(w, responseCookie)

			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		} else {
			setErrorCookie(w, "Invalid login request.")
			http.Redirect(w, r, "/#/login", http.StatusTemporaryRedirect)
		}
	} else {
		log.Println(err)
		setErrorCookie(w, "Invalid login request.")
		http.Redirect(w, r, "/#/login", http.StatusTemporaryRedirect)
	}
}

func (srv *HttpService) LogOutUser(w http.ResponseWriter, r *http.Request) {
	username, err := srv.logout(r)
	if err == nil {
		expireCookie(w, username)
		expireCookie(w, "username")
		expireCookie(w, "error")
	} else {
		log.Println(err)
		setErrorCookie(w, "Invalid logout request.")
	}

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (srv *HttpService) logout(r *http.Request) (string, error) {
	username, token, err := getUsernameAndToken(r)
	if err != nil {
		return "", err
	}

	_, err = srv.AuthService.Logout(context.Background(), &authpb.LogoutAuthRequest{Username: username, Cookie: token.String()})
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

func (srv *HttpService) verifyToken(r *http.Request) (*authpb.VerifyAuthResponse, error) {
	_, token, err := getUsernameAndToken(r)
	if err != nil {
		return nil, err
	}

	response, err := srv.AuthService.Verify(context.Background(), &authpb.VerifyAuthRequest{Cookie: token.String()})
	if err != nil {
		return nil, err
	}

	return response, nil
}
