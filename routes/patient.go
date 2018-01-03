package routes

import (
	"github.com/TheoRev/OdontoSoftGo/controller_api"
	"github.com/TheoRev/OdontoSoftGo/controller"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// SetCrudPatientRouter estaablece la ruta para crear pacientes
func SetCrudPatientRouter(router *mux.Router) {
	prefix := "/api/crud/patients"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.Headers("Access-Control-Allow-Origin", "*")
	subRouter.HandleFunc("/", controller_api.CreatePatient).Methods("POST")
	subRouter.HandleFunc("/", controller_api.CreatePatient).Methods("OPTIONS")
	subRouter.HandleFunc("/", controller_api.UpdatePatient).Methods("PUT")
	subRouter.HandleFunc("/", controller_api.DeletePatient).Methods("DELETE")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}

// SetFindAllPatientRouter estaablece la ruta para obtener todos los pacientes
func SetFindAllPatientRouter(router *mux.Router) {
	prefix := "/api/patients"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controller_api.FindAllPatients).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}

// SetFindPatientWhitoutTreatmentRouter estaablece la ruta para obtener los pacientes que no han iniciado un tratamiento
func SetFindPatientWhitoutTreatmentRouter(router *mux.Router) {
	prefix := "/api/patients-whitout-treatment"
	subRouter := mux.NewRouter().PathPrefix(prefix).Subrouter().StrictSlash(true)
	subRouter.HandleFunc("/", controller_api.FindPatientsWithoutTreatment).Methods("GET")

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}

func ShowAllPatients(router *mux.Router){	
	prefix:="/patients"
	subRouter:=mux.NewRouter().PathPrefix(prefix)
	.Subrouter().StrictSlash(true)
	subRouter.HandlerFunc("/", controller.FindAllPatients)

	router.PathPrefix(prefix).Handler(
		negroni.New(
			negroni.Wrap(subRouter),
		),
	)
}
