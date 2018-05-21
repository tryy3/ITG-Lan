package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/tryy3/itg-lan/auth"
	"github.com/tryy3/itg-lan/models"
)

// UserJSON - json data expected for login/signup
type UserJSON struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserSignup - Handler for signing up a user
func (api *API) UserSignup(w http.ResponseWriter, res *http.Request) {
	decoder := json.NewDecoder(res.Body)
	jsondata := UserJSON{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.Username == "" || jsondata.Password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}

	if api.users.HasUser(jsondata.Username) {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	if api.users.HasUserByEmail(jsondata.Email) {
		http.Error(w, "A user with this email already exists", http.StatusBadRequest)
		return
	}

	user := api.users.AddUser(jsondata.Username, jsondata.Email, jsondata.Password)
	jsontoken := auth.GetJSONToken(user)

	w.Header().Set("Conent-Type", "application/json")
	if _, err := w.Write([]byte(jsontoken)); err != nil {
		log.Fatal(err)
	}
}

// UserLogin - Handler for login a user
func (api *API) UserLogin(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	jsondata := UserJSON{}
	err := decoder.Decode(&jsondata)

	if err != nil || jsondata.Username == "" || jsondata.Password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}

	user := api.users.FindUser(jsondata.Username)
	if user.Username == "" {
		user = api.users.FindUserByEmail(jsondata.Username)
		if user.Username == "" {
			http.Error(w, "Username or email not found", http.StatusBadRequest)
			return
		}
	}

	if !api.users.CheckPassword(user.Password, jsondata.Password) {
		http.Error(w, "Bad password", http.StatusBadRequest)
		return
	}

	jsontoken := auth.GetJSONToken(user)
	w.Header().Set("Conent-Type", "application/json")
	if _, err := w.Write([]byte(jsontoken)); err != nil {
		log.Fatal(err)
	}
}

// GetUserFromContext - return User reference from header token
func (api *API) GetUserFromContext(req *http.Request) *models.User {
	userclaims := auth.GetUserClaimsFromContext(req)
	user := api.users.FindUserByUUID(userclaims["uuid"].(string))
	return user
}

// UserInfo - example to get
func (api *API) UserInfo(w http.ResponseWriter, req *http.Request) {
	user := api.GetUserFromContext(req)
	js, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(js); err != nil {
		log.Fatal(err)
	}
}
