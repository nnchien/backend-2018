package server

import (
	"time"
	"net/http"
	"strings"
	"fmt"
	"log"

	"github.com/nnchien/backend-2018/golang/practices/server/handlers"
)

var Session *http.Server
var r  *Router

func Run(port string) {
	Session = &http.Server{
		Addr:           port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(Session.ListenAndServe())
}

type Handle func(http.ResponseWriter, *http.Request)

type Router struct {
	mux map[string]Handle
}

func newRouter() *Router {
	return &Router{
		mux: make(map[string]Handle),
	}
}

func (r *Router) Add(path string, handle Handle) {
	r.mux[path] = handle
}

func GetHeader(url string) string {
	sl := strings.Split(url, "/")
	return fmt.Sprintf("/%s", sl[1])
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	head := GetHeader(r.URL.Path)
	h, ok := rt.mux[head]
	if ok {
		h(w, r)
		return
	}
	http.NotFound(w, r)
}

func Start() {
	handler := handlers.NewHandler()
	r = newRouter()
	r.Add("/push", handler.Push)
	Run(":8080")
}
