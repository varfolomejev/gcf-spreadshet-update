package googletablesfunction

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Contact is main structure which represents a row in the google spreadsheet
type Contact struct {
	Email     string `json:"email"`
	Telephone string `json:"tel"`
	Address   string `json:"address"`
}

func responseError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func contactExists(cs []*Contact, c *Contact) bool {
	for _, cc := range cs {
		if cc.Email == c.Email && cc.Telephone == c.Telephone && cc.Address == c.Address {
			return true
		}
	}
	return false
}

// UpdateSheetHandler is an HTTP Cloud Function.
func UpdateSheetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	var newContacts []*Contact
	err := json.NewDecoder(r.Body).Decode(&newContacts)
	if err != nil {
		responseError(w, err)
		return
	}
	curContacts, err := getSheet(r.Context())
	if err != nil {
		responseError(w, err)
		return
	}
	for _, c := range newContacts {
		if !contactExists(curContacts, c) {
			curContacts = append(curContacts, c)
		}
	}
	err = writeSheet(r.Context(), curContacts)
	if err != nil {
		responseError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Writed %d rows", len(curContacts))
}
