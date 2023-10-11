package router

import (
	"github.com/ShreyasSahoo/mongoapi/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/api/movies/{id}", controller.MarkAsViewed).Methods("PUT")
	router.HandleFunc("/api/movies", controller.DeleteAllMovies).Methods("DELETE")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	return router
}
