package routes

import (
	"github.com/TheoRev/OdontoSoft_Backend/controller_api"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetCrudWorkRouter estaablece la ruta para crear un trabajo
func SetCrudWorkRouter(router *mux.Router) {
	prefix := "/api/crud/work"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controller_api.CreateWork).Methods("POST")
	subRouter.HandleFunc("/", controller_api.UpdateWork).Methods("PUT")
	subRouter.HandleFunc("/", controller_api.DeleteWork).Methods("DELETE")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}

// SetFindAllWorksRouter estaablece la ruta para obtener todos los trabajos
func SetFindAllWorksRouter(router *mux.Router) {
	prefix := "/api/works"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controller_api.FindAllWork).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
