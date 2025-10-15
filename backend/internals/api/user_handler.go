package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Rituraj-13/userReg/backend/internals/store"
)

type UserHandler struct {
	user store.UserStore
}

func NewUserHandler(userStore store.UserStore)(*UserHandler){
	return &UserHandler{
		user: userStore,
	}
}

func(uh *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request){
	var user store.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		fmt.Println(err)
		http.Error(w, "Couldn't parse the body", http.StatusInternalServerError)
	}

	createdUser, err := uh.user.CreateUser(&user)
	if err != nil{
		fmt.Println(err)
		http.Error(w, "Couldn't create the user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(createdUser)
}

func(uh *UserHandler) HandleUserById(w http.ResponseWriter, r *http.Request){
	user, err := uh.user.GetUserById(10)
	if err != nil{
		http.Error(w,"Find by ID error", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(&user)
}