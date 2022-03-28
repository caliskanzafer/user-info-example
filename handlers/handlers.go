package handlers

import (
	d "example.com/dataloaders"
	b "example.com/models"

	"encoding/json"
	"net/http"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
func Handler(w http.ResponseWriter, r *http.Request) {
	//page nesnesi

	page := b.Page{ID: 7, Name: "Kullanici", Description: "Kullanici bilgisi", URI: "/users"}

	// data loaders

	users := d.LoadUsers()
	interests := d.LoadInterests()
	userInterestsMapping := d.LoadInterestMapping()

	// // islem

	var newUser []b.User

	for _, user := range users {
		for _, userInterest := range userInterestsMapping {
			if user.ID == userInterest.UserID {
				for _, interest := range interests {
					if interest.ID == userInterest.InterestID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUser = append(newUser, user)

	}

	viewModel := b.UserViewModel{Page: page, User: newUser}

	data, _ := json.Marshal(viewModel)

	w.Write(data)
}
