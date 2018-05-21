package api

import (
	"encoding/json"
	"github.com/tryy3/itg-lan/models"
	"io"
	"log"
	"regexp"
)

var emailCheck *regexp.Regexp

func init() {
	emailCheck = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
}

// API struct
type API struct {
	users             *models.UserManager
	signup            *models.SignupManager
	faq               *models.FAQManager
	tournament        *models.TournamentManager
	tournamentSignups *models.TournamentSignupManager
}

// NewAPI - Creates a new API
func NewAPI(db *models.DB) *API {
	usermgr, err := models.NewUserManager(db)
	if err != nil {
		log.Fatal(err)
	}

	signupmgr, err := models.NewSignupManager(db)
	if err != nil {
		log.Fatal(err)
	}

	faqmgr, err := models.NewFAQManager(db)
	if err != nil {
		log.Fatal(err)
	}

	tourmgr, err := models.NewTournamentManager(db)
	if err != nil {
		log.Fatal(err)
	}

	toursignmgr, err := models.NewTournamentSignupManager(db)
	if err != nil {
		log.Fatal(err)
	}

	return &API{
		users:             usermgr,
		signup:            signupmgr,
		faq:               faqmgr,
		tournament:        tourmgr,
		tournamentSignups: toursignmgr,
	}
}

// MessageJSON - json data for outputting
type MessageJSON struct {
	Success bool        `json:"success"` // Wether an error occured or not
	Message string      `json:"message"` // The message
	Data    interface{} `json:"data"`    // Extra data, generally it will contain a struct
}

// DeleteJSON - json data expected when deleting something
type DeleteJSON struct {
	ID uint `json:"id"`
}

func outputJSON(w io.Writer, success bool, message string, data interface{}) {
	out, _ := json.Marshal(MessageJSON{Success: success, Message: message, Data: data})
	if _, err := w.Write(out); err != nil {
		log.Fatal(err)
	}
}
