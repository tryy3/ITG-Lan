package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tryy3/itg-lan/models"
)

// CreateTournament - Handler for when someone creates a tournament
func (api *API) CreateTournament(w http.ResponseWriter, res *http.Request) {
	t := &models.Tournament{}
	if err := api.tournament.Create(t); err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}
	outputJSON(w, true, "Du har nu skapat en ny turnering", t)
}

// FindAllTournaments - Handler for when someone wants to see all the tournaments
func (api *API) FindAllTournaments(w http.ResponseWriter, res *http.Request) {
	t, err := api.tournament.FindAll()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	js, _ := json.Marshal(t)
	if _, err := w.Write(js); err != nil {
		log.Fatal(err)
	}
}

// TournamentJSON - json data expected for editing a tournament
type TournamentJSON struct {
	Image  string `json:"image"`
	Amount uint   `json:"amount"`
	Name   string `json:"name"`
	ID     uint   `json:"id"`
	Active bool   `json:"active"`
}

// EditTournament - Handler for when someone wants to edit a tournament
func (api *API) EditTournament(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := TournamentJSON{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	if jsondata.Image == "" {
		outputJSON(w, false, "Du har inte skrivit in en bild URL", nil)
		return
	}

	if jsondata.Amount == 0 {
		outputJSON(w, false, "Antal måste vara högre än 0", nil)
		return
	}

	if jsondata.Name == "" {
		outputJSON(w, false, "Du har inte skrivit in ett turnerings namn", nil)
		return
	}

	t, err := api.tournament.Find(jsondata.ID)
	if err != nil {
		outputJSON(w, false, "Turneringens ID finns inte", nil)
		return
	}

	// Replace exisitng tournament data
	t.Image = jsondata.Image
	t.Amount = jsondata.Amount
	t.Name = jsondata.Name
	t.Active = jsondata.Active

	// Save tournament
	if err := api.tournament.Save(t); err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	outputJSON(w, true, "Du har nu uppdaterat turneringen", t)
}

// DeleteTournament - Deletes an existing tournament
func (api *API) DeleteTournament(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := DeleteJSON{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	if err := api.tournament.DeleteWithID(jsondata.ID); err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	if err := api.tournamentSignups.DeleteWithTournamentID(jsondata.ID); err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	outputJSON(w, true, "Du har nu tagit bort turneringen", nil)
}
