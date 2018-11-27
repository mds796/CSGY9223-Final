package web

import (
	"context"
	"github.com/mds796/CSGY9223-Final/auth"
	"github.com/mds796/CSGY9223-Final/feed"
	"github.com/mds796/CSGY9223-Final/feed/feedpb"
	"github.com/mds796/CSGY9223-Final/follow"
	"github.com/mds796/CSGY9223-Final/post"
	"github.com/mds796/CSGY9223-Final/post/postpb"
	"github.com/mds796/CSGY9223-Final/user"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

type HttpService struct {
	StaticPath    string
	Multiplexer   *http.ServeMux
	Server        *http.Server
	UserService   user.Service
	AuthService   auth.Service
	PostService   postpb.PostClient
	FollowService follow.Service
	FeedService   feedpb.FeedClient
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
	srv.Multiplexer.HandleFunc("/follow", srv.ToggleFollow)
	srv.Multiplexer.HandleFunc("/follows", srv.ListFollows)
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
	log.Printf("Web now listening on %v.\n", srv.Address())
	err := srv.Server.ListenAndServe()
	log.Println(err)
}

// New creates a new Web service with nil dependencies.
func New(config *Config) Service {
	service := newService(config.Target(), config.StaticPath)

	service.UserService = user.CreateStub()
	service.AuthService = auth.CreateStub(service.UserService)
	service.PostService = post.NewStubClient(post.CreateStub())
	service.FollowService = follow.CreateStub(service.UserService)

	// Use this value once the feed service is updated to use thew gRPC user, post, and follow clients
	_, err := feed.NewClient(config.FeedTarget())
	if err != nil {
		log.Println(err)
		return nil
	}

	service.FeedService = feed.NewStubClient(feed.NewStubServer(service.PostService, service.UserService, service.FollowService))

	return service
}

func newService(target string, staticPath string) *HttpService {
	mux := http.NewServeMux()
	server := &http.Server{Addr: target, Handler: mux}

	return &HttpService{StaticPath: staticPath, Multiplexer: mux, Server: server}
}

func newStubService(host string, port uint16, staticPath string) *HttpService {
	service := newService(host+":"+strconv.Itoa(int(port)), staticPath)

	userService := user.CreateStub()
	authService := auth.CreateStub(userService)
	postService := post.NewStubClient(post.CreateStub())
	followService := follow.CreateStub(userService)
	feedService := feed.NewStubClient(feed.NewStubServer(postService, userService, followService))

	service.UserService = userService
	service.AuthService = authService
	service.PostService = postService
	service.FollowService = followService
	service.FeedService = feedService

	return service
}
