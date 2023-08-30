package beneficiary

import (
	"bernardolsp/invoice-generator/helpers"
	"database/sql"
	"log"
	"net/http"
)

type BeneficiaryStruct struct {
	DB     *sql.DB
	Logger *log.Logger
}

func (b *BeneficiaryStruct) Handle(w http.ResponseWriter, r *http.Request) {
	b.Logger.Println("Handling request...")

	if r.Method == http.MethodGet {
		b.handleGet(w, r)
	} else if r.Method == http.MethodPost {
		b.handlePost(w, r)
	}
}

func (b *BeneficiaryStruct) handleGet(w http.ResponseWriter, r *http.Request) {
	beneficiaries, err := b.get_beneficiaries()
	if err != nil {
		b.Logger.Println("Error getting beneficiaries", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b.Logger.Println("Successfully got beneficiaries", beneficiaries)

	w.WriteHeader(http.StatusOK)
}

func (b *BeneficiaryStruct) handlePost(w http.ResponseWriter, r *http.Request) {
	b.Logger.Println("Handling POST request...")
	// Decode the request body into an Beneficiary struct
	// The beneficiary has a customer ID, so we need to check if the customer exists
	// If the customer does not exist, return an error
	// If the customer exists, create the beneficiary
	// Return the beneficiary

	var beneficiary Beneficiary
	err := helpers.DecodeJSONBody(r, &beneficiary)
	if err != nil {
		b.Logger.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	b.Logger.Printf("Successfully decoded JSON: %+v\n", beneficiary)

}
