package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tryy3/itg-lan/models"
)

// SignupJSON - json data expected for signin up
type SignupJSON struct {
	Name      string `json:"name"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Class     string `json:"class"`
	ClassRoom string `json:"classroom"`
}

// CreateSignup - Handler for when someone signs up
func (api *API) CreateSignup(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := SignupJSON{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	if jsondata.Name == "" {
		outputJSON(w, false, "Du har inte skrivit in ett förnamn", nil)
		return
	}

	if jsondata.Lastname == "" {
		outputJSON(w, false, "Du har inte skrivit in ett efternamn", nil)
		return
	}

	if jsondata.Email == "" {
		outputJSON(w, false, "Du har inte skrivit in en email address", nil)
		return
	}

	if jsondata.Class == "" {
		outputJSON(w, false, "Du har inte valt en klass", nil)
		return
	}

	if jsondata.ClassRoom == "" {
		outputJSON(w, false, "Du har inte valt ett klassrum", nil)
		return
	}

	if !emailCheck.MatchString(jsondata.Email) {
		outputJSON(w, false, "Du har inte skrivit in en giltig email address", nil)
		return
	}

	s := &models.Signup{
		Name:      jsondata.Name,
		Lastname:  jsondata.Lastname,
		Email:     jsondata.Email,
		Class:     jsondata.Class,
		ClassRoom: jsondata.ClassRoom,
	}

	if err := api.signup.Create(s); err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	outputJSON(w, true, "Du har nu anmält dig", s)
}

// GetSignup - Retrieves all people who signed up
func (api *API) GetSignup(w http.ResponseWriter, res *http.Request) {
	s, err := api.signup.FindAll()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	js, _ := json.Marshal(s)
	if _, err := w.Write(js); err != nil {
		log.Fatal(err)
	}
}
