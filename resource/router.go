package resource

import (
	"github.com/gorilla/mux"
	"sync"
)

var (
	router *mux.Router
	once sync.Once
)

func GetRouter() *mux.Router {
	once.Do(func() {
		router = mux.NewRouter() // match incoming request to their respective register routes (handlers)
	})
	return router
}