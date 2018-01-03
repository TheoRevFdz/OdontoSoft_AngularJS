package templates

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/TheoRev/OdontoSoftGo/config"
	"github.com/TheoRev/OdontoSoftGo/models"
)

// FindAllPatients obtiene todos los pacientes de la db
func FindllPatients(w http.ResponseWriter, r *http.Request) {
	patients := models.Patients{}
	msg := models.Message{}

	db := config.GetConnection()
	defer db.Close()

	// err := db.Where("state = TRUE").Order("id desc").Find(&patients).Error
	err := db.Order("id desc").Find(&patients).Error
	if err != nil {
		msg.Message = fmt.Sprintf("Error al obtener los datos: %s", err)
	}

	t, _ := template.ParseFiles("public/patient/patients.html")
	t.Execute(w, patients)
}
