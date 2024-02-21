package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rajath002/bookings/internal/config"
	"github.com/rajath002/bookings/internal/forms"
	"github.com/rajath002/bookings/internal/models"
	"github.com/rajath002/bookings/internal/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// New Handlers : sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("HOME")
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplateDynamicCache(w, r, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	log.Println("ABOUT")
	render.RenderTemplateDynamicCache(w, r, "about.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	log.Println("Contact")
	render.RenderTemplateDynamicCache(w, r, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	log.Println("Majors")
	render.RenderTemplateDynamicCache(w, r, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) GeneralsQuarters(w http.ResponseWriter, r *http.Request) {
	log.Println("GeneralsQuarters")
	render.RenderTemplateDynamicCache(w, r, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) SearchAvailability(w http.ResponseWriter, r *http.Request) {
	log.Println("SearchAvailability")
	render.RenderTemplateDynamicCache(w, r, "search-availability.page.tmpl", &models.TemplateData{})
}

// PostAvailability handles post
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is %s and end is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// PostAvailability handles request for availability and send the JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "     ")

	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	log.Println("Reservation")
	render.RenderTemplateDynamicCache(w, r, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// Handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	log.Println("Post Reservation")
	// render.RenderTemplateDynamicCache(w, r, "make-reservation.page.tmpl", &models.TemplateData{
	// 	Form: forms.New(nil),
	// })
}
