package handlers

import (
	"net/http"

	"github.com/alg3n/bookings/pkg/config"
	"github.com/alg3n/bookings/pkg/models"
	"github.com/alg3n/bookings/pkg/render"
)

// Repo represents the repository used by the handlers.
var Repo *Repository

// Repository represents a data repository for the application.
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new instance of the Repository struct with the provided AppConfig.
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers initializes a new instance of Handlers with the given repository.
func NewHandlers(r *Repository) {
	Repo = r
}

// Home handles the HTTP request for the home page.
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About handles the HTTP request for the about page.
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
