package web

import (
	"context"
	"github.com/mds796/CSGY9223-Final/auth"
	"github.com/mds796/CSGY9223-Final/post"
	"github.com/mds796/CSGY9223-Final/user"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

type Follow struct {
	Name     string
	Followed bool
}

type HttpService struct {
	StaticPath  string
	Multiplexer *http.ServeMux
	Server      *http.Server
	UserService user.Service
	AuthService auth.Service
	PostService post.Service
}

func (srv *HttpService) Address() string {
	return "http://" + srv.Server.Addr
}

type Service interface {
	Start()
	Stop()
	Address() string
}

func (srv *HttpService) Start() {
	srv.configureRoutes()
	srv.listenAndServe()
}

func (srv *HttpService) configureRoutes() {
	srv.Multiplexer.HandleFunc("/register", srv.RegisterUser)
	srv.Multiplexer.HandleFunc("/login", srv.LogInUser)
	srv.Multiplexer.HandleFunc("/logout", srv.LogOutUser)
	srv.Multiplexer.HandleFunc("/feed", srv.FetchFeed)
	srv.Multiplexer.HandleFunc("/post", srv.MakePost)
	srv.Multiplexer.HandleFunc("/follow", srv.ListFollows)
	srv.Multiplexer.HandleFunc("/follows", srv.ToggleFollow)
	srv.Multiplexer.HandleFunc("/", srv.ServeStatic)
}

func (srv *HttpService) ServeStatic(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(srv.StaticPath, r.URL.Path)
	http.ServeFile(w, r, path)
}

func (srv *HttpService) Stop() {
	if err := srv.Server.Shutdown(context.Background()); err != nil {
		log.Printf("HTTP server Shutdown: %v\n", err)
	}
}

func (srv *HttpService) listenAndServe() {
	log.Printf("Now listening on %v.\n", srv.Address())
	srv.Server.ListenAndServe()
}

func New(host string, port uint16, staticPath string) Service {
	mux := http.NewServeMux()
	address := host + ":" + strconv.Itoa(int(port))
	server := &http.Server{Addr: address, Handler: mux}
	userService := user.CreateStub()
	authService := auth.CreateStub(userService)
	postService := post.CreateStub()

	service := &HttpService{
		StaticPath:  staticPath,
		Multiplexer: mux,
		Server:      server,
		UserService: userService,
		AuthService: authService,
		PostService: postService}

	return service
}
