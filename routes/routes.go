package routes

import (
	"github.com/gorilla/mux"
	"github.com/tryy3/itg-lan/api"
	"github.com/tryy3/itg-lan/auth"
	"github.com/urfave/negroni"
	"net/http"
)

// NewRoutes builds the routes for the api
func NewRoutes(api *api.API) *mux.Router {
	mux := mux.NewRouter()

	// client static files
	mux.Handle("/", http.FileServer(http.Dir("./client/dist"))).Methods("GET")
	mux.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./client/dist/static/"))))

	// api
	a := mux.PathPrefix("/api/v1").Subrouter()

	// users
	u := a.PathPrefix("/user").Subrouter()
	// u.HandleFunc("/signup", api.UserSignup).Methods("POST")
	u.HandleFunc("/login", api.UserLogin).Methods("POST")
	u.Handle("/info", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.UserInfo)),
	))

	// signup
	s := a.PathPrefix("/signup").Subrouter()
	s.HandleFunc("/create", api.CreateSignup).Methods("POST")
	s.Handle("/find", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.GetSignup)),
	)).Methods("POST") // Need to be logged in/admin

	// faq
	faq := a.PathPrefix("/faq").Subrouter()
	faq.HandleFunc("/create", api.CreateFAQ).Methods("POST")
	faq.HandleFunc("/find/answered", api.GetAnsweredFaq).Methods("POST")
	faq.Handle("/find/all", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.GetFaq)),
	)).Methods("POST") // Need to be logged in/admin
	faq.Handle("/answer", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.AnswerQuestion)),
	)).Methods("POST") // Need to be logged in/admin

	// tournament
	tour := a.PathPrefix("/tournament").Subrouter()
	tour.Handle("/create", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.CreateTournament)),
	)).Methods("POST") // Need to be logged in/admin
	tour.Handle("/edit", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.EditTournament)),
	)).Methods("POST") // Need to be logged in/admin
	tour.Handle("/delete", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.DeleteTournament)),
	)).Methods("POST") // Need to be logged in/admin
	tour.Handle("/find", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.FindAllTournaments)),
	)).Methods("POST") // Need to be logged in/admin

	// tournamentSignups
	tourSigns := a.PathPrefix("/tournament-signups").Subrouter()
	tourSigns.HandleFunc("/create", api.CreateTeam).Methods("POST")
	tourSigns.HandleFunc("/find/active", api.FindAllActiveTournaments).Methods("POST")
	tourSigns.Handle("/find/all", negroni.New(
		negroni.HandlerFunc(auth.JwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(api.FindAllTournamentsSignups)),
	)).Methods("POST") // Need to be logged in/admin

	// Redirect everything else to /
	mux.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "./client/dist/index.html")
	})

	return mux
}
