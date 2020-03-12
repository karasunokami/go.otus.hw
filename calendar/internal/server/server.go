package server

import (
	"github.com/karasunokami/go.otus.hw/calendar/internal/config"
	"net/http"
)

type HttpServer interface {
	Run() error
}

type httpServerImpl struct {
	conf config.HttpListen

	server *http.Server
	mux    *http.ServeMux
}

func NewHttpServer(conf config.HttpListen) HttpServer {
	return &httpServerImpl{conf: conf}
}

func (s *httpServerImpl) Run() error {
	s.mux = http.NewServeMux()
	s.configureRoutes()

	s.server = &http.Server{
		Addr:    s.conf.Ip + ":" + s.conf.Port,
		Handler: s.mux,
	}

	return s.server.ListenAndServe()
}

func (s *httpServerImpl) configureRoutes() {
	s.mux.HandleFunc("/hello", s.hello)
}

func (s *httpServerImpl) hello(w http.ResponseWriter, r *http.Request) {
	_ = r
	_, _ = w.Write([]byte("hello"))
}
