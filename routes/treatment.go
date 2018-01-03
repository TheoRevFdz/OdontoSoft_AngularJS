package routes

import (
	"github.com/TheoRev/OdontoSoft_Backend/controller_api"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetCrudTreatmentRouter estaablece la ruta para crear un tratamiento
func SetCrudTreatmentRouter(router *mux.Router) {
	prefix := "/api/crud/treatment"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controller_api.CreateTreatment).Methods("POST")
	subRouter.HandleFunc("/", controller_api.UpdateTreatment).Methods("PUT")
	subRouter.HandleFunc("/", controller_api.DeleteTreatment).Methods("DELETE")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}

// SetFindAllTreatmentsRouter establece la ruta para obtener todos los tratamientos
func SetFindAllTreatmentsRouter(router *mux.Router) {
	prefix := "/api/treatments"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controller_api.FindAllTreatments).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}

// SetLastTreatmentRouter establece la ruta para obtener el ultimo tratamiento registrado
func SetLastTreatmentRouter(router *mux.Router) {
	prefix := "/api/last-treatment"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controller_api.FindLastTreatment).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
