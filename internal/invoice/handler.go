package invoice

import (
	"database/sql"
	"log"
	"net/http"
	// ... other imports
)

type Application struct {
	Logger *log.Logger
	DB     *sql.DB
}

func (h *Application) Handle(w http.ResponseWriter, r *http.Request) {
	h.Logger.Println("Handling request...")

	if r.Method == http.MethodGet {
		h.handleGet(w, r)
	} else if r.Method == http.MethodPost {
		h.handlePost(w, r)
	}
}

func (h *Application) handleGet(w http.ResponseWriter, r *http.Request) {
	invoices, err := h.GetInvoices()
	if err != nil {
		h.Logger.Println("Error getting invoices", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	h.Logger.Println("Successfully got invoices", invoices)

	w.WriteHeader(http.StatusOK)
}

func (h *Application) handlePost(w http.ResponseWriter, r *http.Request) {
	h.Logger.Println("Handling POST request...")
}
