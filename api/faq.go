package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tryy3/itg-lan/models"
)

// FAQJSON - json data expected for signin up
type FAQJSON struct {
	Question string `json:"question"`
}

// CreateFAQ - Handler for when creates a question
func (api *API) CreateFAQ(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := FAQJSON{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	if jsondata.Question == "" {
		outputJSON(w, false, "Du saknar en fråga", nil)
		return
	}

	faq := &models.FAQ{
		Question: jsondata.Question,
	}

	if err := api.faq.Create(faq); err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	outputJSON(w, true, "Frågan har skickats in", faq)
}

// GetFaq - Retrieves all questions
func (api *API) GetFaq(w http.ResponseWriter, res *http.Request) {
	faq, err := api.faq.FindAll()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	js, _ := json.Marshal(faq)
	if _, err := w.Write(js); err != nil {
		log.Fatal(err)
	}
}

// GetAnsweredFaq - Retrieves all answered questions
func (api *API) GetAnsweredFaq(w http.ResponseWriter, res *http.Request) {
	faq, err := api.faq.FindAnswered()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	js, _ := json.Marshal(faq)
	if _, err := w.Write(js); err != nil {
		log.Fatal(err)
	}
}

// AnswerJSON - json data expected for answering a question
type AnswerJSON struct {
	ID     uint   `json:"id"`
	Answer string `json:"answer"`
}

// AnswerQuestion - Adds an answer to a question
func (api *API) AnswerQuestion(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := AnswerJSON{}

	w.Header().Set("Content-Type", "application/json")
	if err := decoder.Decode(&jsondata); err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}

	question, err := api.faq.Find(jsondata.ID)
	if err != nil {
		log.Print(err)
		outputJSON(w, false, "Kan inte hitta en fråga med detta ID.", nil)
		return
	}

	question.Answer = jsondata.Answer
	err = api.faq.Save(question)
	if err != nil {
		log.Print(err)
		outputJSON(w, false, "Ett internt error har hänt", nil)
		return
	}
	outputJSON(w, true, "Du har nu svarat på frågan", question)
}
