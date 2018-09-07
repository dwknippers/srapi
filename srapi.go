package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/group/{groupCode}/schedule/week/{weekNumber}", retrieveSchedule).Methods("GET")
	//router.HandleFunc("/teacher/{teacherCode}/schedule/week/{weekNumber}", retrieveSchedule).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func httpError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
}

func retrieveSchedule(w http.ResponseWriter, r *http.Request) {
	jar, _ := cookiejar.New(&cookiejar.Options{})

	u, _ := url.Parse("https://roosters.saxion.nl/")

	if r.Header["Authorization"] == nil {
		httpError(w, http.StatusBadRequest, "Authorization cookie was not supplied.")
		return
	}

	jar.SetCookies(u, []*http.Cookie{
		&http.Cookie{
			Name:  "saxion_roosters[access_token]",
			Value: (r.Header["Authorization"][0])[7:],
		},
	})

	client := &http.Client{Jar: jar}

	resp, err := client.Get(fmt.Sprintf(
		"https://roosters.saxion.nl/schedule/week:%s/group:%s",
		mux.Vars(r)["weekNumber"], mux.Vars(r)["groupCode"]))
	if err != nil {
		log.Println(err)
		httpError(w, http.StatusInternalServerError,
			"Could not retrieve schedule.")
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)

	days := parse(string(body))

	resp.Body.Close()

	json.NewEncoder(w).Encode(days)

	// TODO: Connect to O365.

	// TODO: Create appointments according to timetable format.
	// TODO:	Check if appointments do not already exist.
}
