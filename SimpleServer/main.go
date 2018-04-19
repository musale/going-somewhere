package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// User information of a github user. This is the
// structure of the JSON data you are rendering so you
// customise or make other structs that are inline with
// the API responses for the data you are displaying
type User struct {
	Name     string `json:"name"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Email    string `json:"email"`
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))

	// The endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		user, err := getGithubUser("musale")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		if err := templates.ExecuteTemplate(w, "index.html", user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Start the server on 8080
	fmt.Println(http.ListenAndServe(":8080", nil))
}

// One of your API endpoints
func getGithubUser(username string) (User, error) {
	var resp *http.Response
	var err error
	var user User
	// The endpoint
	const githubUserAPI = "https://api.github.com/users/"

	// Get the required data in json
	if resp, err = http.Get(githubUserAPI + username); err != nil {
		return user, err
	}

	defer resp.Body.Close()

	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return user, err
	}

	// Unmarshal the response into the struct
	if err = json.Unmarshal(body, &user); err != nil {
		return user, err
	}

	return user, nil
}
