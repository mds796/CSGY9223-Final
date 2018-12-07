package web

import (
	"context"
	"github.com/mds796/CSGY9223-Final/auth"
	"github.com/mds796/CSGY9223-Final/auth/authpb"
	"github.com/mds796/CSGY9223-Final/feed"
	"github.com/mds796/CSGY9223-Final/feed/feedpb"
	"github.com/mds796/CSGY9223-Final/follow"
	"github.com/mds796/CSGY9223-Final/follow/followpb"
	"github.com/mds796/CSGY9223-Final/post"
	"github.com/mds796/CSGY9223-Final/post/postpb"
	"github.com/mds796/CSGY9223-Final/storage"
	"github.com/mds796/CSGY9223-Final/user"
	"github.com/mds796/CSGY9223-Final/user/userpb"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
	"strconv"
)

type HttpService struct {
	StaticPath    string
	StaticUrl     *url.URL
	Multiplexer   *http.ServeMux
	Server        *http.Server
	UserService   userpb.UserClient
	AuthService   authpb.AuthClient
	PostService   postpb.PostClient
	FollowService followpb.FollowClient
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

	if srv.StaticUrl == nil {
		srv.Multiplexer.HandleFunc("/", srv.ServeStatic)
	} else {
		proxy := httputil.NewSingleHostReverseProxy(srv.StaticUrl)
		oldDirector := proxy.Director

		proxy.Director = func(req *http.Request) {
			oldDirector(req)
			req.Method = "GET"
		}

		srv.Multiplexer.Handle("/", proxy)
	}
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
	service := newService(config.Target(), config.StaticPath, config.StaticUrl)

	userService, err := user.NewClient(config.UserTarget())
	if err != nil {
		log.Println(err)
		return nil
	}

	authService, err := auth.NewClient(config.AuthTarget())
	if err != nil {
		log.Println(err)
		return nil
	}

	postService, err := post.NewClient(config.PostTarget())
	if err != nil {
		log.Println(err)
		return nil
	}

	followService, err := follow.NewClient(config.FollowTarget())
	if err != nil {
		log.Println(err)
		return nil
	}

	feedService, err := feed.NewClient(config.FeedTarget())
	if err != nil {
		log.Println(err)
		return nil
	}

	service.UserService = userService
	service.AuthService = authService
	service.PostService = postService
	service.FollowService = followService
	service.FeedService = feedService

	return service
}

func newService(target string, staticPath string, staticUrl string) *HttpService {
	mux := http.NewServeMux()
	server := &http.Server{Addr: target, Handler: mux}

	staticServer, err := url.Parse(staticUrl)
	if staticUrl == "" || err != nil {
		log.Printf("Did not receive a valid static server URL. Using static server directory instead. Error: %v", err)
		staticServer = nil
	}

	return &HttpService{StaticPath: staticPath, StaticUrl: staticServer, Multiplexer: mux, Server: server}
}

func newStubService(host string, port uint16, staticPath string) *HttpService {
	target := host + ":" + strconv.Itoa(int(port))
	service := newService(target, staticPath, "")

	userService := user.NewStubClient(user.CreateStub(storage.STUB))
	authService := auth.NewStubClient(auth.CreateStub(storage.STUB, userService))
	postService := post.NewStubClient()
	followService := follow.NewStubClient(follow.CreateStub(userService))
	feedService := feed.NewStubClient(feed.NewStubServer(postService, userService, followService))

	service.UserService = userService
	service.AuthService = authService
	service.PostService = postService
	service.FollowService = followService
	service.FeedService = feedService

	return service
}
