package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tryy3/itg-lan/models"
)

// CreateTeamJSON - Data expected when someone creates a tournmanet team
type CreateTeamJSON struct {
	Team         string `json:"team"`
	Name         string `json:"name"`
	IGN          string `json:"ign"`
	Leader       bool   `json:"leader"`
	Email        string `json:"email"`
	IDTournament uint   `json:"id_tournament"`
}

// CreateTeam - Handler for when someone creates a tournament team
func (api *API) CreateTeam(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := []CreateTeamJSON{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	// TODO: Check for multiple leaders
	for _, team := range jsondata {
		if team.Team == "" {
			outputJSON(w, false, "Du har inte skrivit in ett team namn", nil)
			return
		}

		if team.Name == "" {
			outputJSON(w, false, "Du har inte skrivit in ditt namn", nil)
			return
		}

		if team.IGN == "" {
			outputJSON(w, false, "Du har inte skrivit in ditt IGN", nil)
			return
		}

		if team.Email == "" {
			outputJSON(w, false, "Du har inte skrivit in en email address", nil)
			return
		}

		if !emailCheck.MatchString(team.Email) {
			outputJSON(w, false, "Du har inte skrivit in en giltig email address", nil)
			return
		}
	}

	if api.tournamentSignups.Exists(jsondata[0].Team) {
		outputJSON(w, false, "Någon har redan anmält sig med detta team namn", nil)
		return
	}

	if _, err := api.tournament.Find(jsondata[0].IDTournament); err != nil {
		outputJSON(w, false, "Detta turnerings ID finns inte", nil)
		return
	}

	t := []*models.TournamentSignup{}
	for _, team := range jsondata {
		tt := &models.TournamentSignup{
			Team:         team.Team,
			Name:         team.Name,
			IGN:          team.IGN,
			Leader:       team.Leader,
			Email:        team.Email,
			IDTournament: team.IDTournament,
		}

		if err := api.tournamentSignups.Create(tt); err != nil {
			log.Print(err)
			outputJSON(w, false, "Ett internt error har hänt", nil)
			return
		}
		t = append(t, tt)
	}
	outputJSON(w, true, "Du har nu anmält ditt team till turneringen", t)
}

// ActiveTournaments - is used when outputting all active tournaments and who signed up
type ActiveTournaments struct {
	Tournament models.Tournament `json:"tournament"`
	Teams      []string          `json:"teams"`
}

// FindAllActiveTournaments - Will get all active tournaments and also get all teams who signed up
func (api *API) FindAllActiveTournaments(w http.ResponseWriter, res *http.Request) {
	tournaments, err := api.tournament.FindActive()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	a := []ActiveTournaments{}

	for _, tournament := range tournaments {
		teams, err := api.tournamentSignups.FindTeams(tournament.ID)
		if err != nil {
			log.Print(err)
			outputJSON(w, false, "Ett internt error har hänt", nil)
			return
		}

		noDups := removeDuplicates(teams)
		a = append(a, ActiveTournaments{Tournament: tournament, Teams: noDups})
	}

	js, _ := json.Marshal(a)
	if _, err := w.Write(js); err != nil {
		log.Fatal(err)
	}
}

// ActiveTournamentSignups - is used when outputting all active tournaments and who signed up
type ActiveTournamentSignups struct {
	Tournament models.Tournament         `json:"tournament"`
	Teams      []models.TournamentSignup `json:"teams"`
}

// FindAllTournamentsSignups - Will get all active tournaments and also all people who signed up
func (api *API) FindAllTournamentsSignups(w http.ResponseWriter, res *http.Request) {
	tournaments, err := api.tournament.FindActive()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	a := []ActiveTournamentSignups{}
	for _, tournament := range tournaments {
		teams, err := api.tournamentSignups.FindAllTeams(tournament.ID)
		if err != nil {
			log.Print(err)
			outputJSON(w, false, "Ett internt error har hänt", nil)
			return
		}

		a = append(a, ActiveTournamentSignups{Tournament: tournament, Teams: teams})
	}

	js, _ := json.Marshal(a)
	if _, err := w.Write(js); err != nil {
		log.Fatal(err)
	}
}

func removeDuplicates(dups []string) []string {
	set := make(map[string]struct{})
	for _, value := range dups {
		set[value] = struct{}{}
	}

	noDups := make([]string, len(set))
	i := 0
	for key := range set {
		noDups[i] = key
		i++
	}
	return noDups
}
