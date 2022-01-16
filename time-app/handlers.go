package timeApp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Time struct {
	CurrentTime string `json:"current_time" xml:"name"`
}

func loadDefaultsHeaders(w http.ResponseWriter) {
	w.Header().Add("Content-type", "application/json")
}

func getTime(rw http.ResponseWriter, r *http.Request) {
	currentTime := Time{time.Now().UTC().String()}
	fmt.Print(mux.Vars(r)["tz"])

	loadDefaultsHeaders(rw)
	json.NewEncoder(rw).Encode(currentTime)
}

func getTimeWithTimezone(rw http.ResponseWriter, r *http.Request) {
	fmt.Print("query")
	tz := mux.Vars(r)["tz"]
	if !isValidTimezone(tz) {
		http.Error(rw, "Bad Request", 400)
	}
	loadDefaultsHeaders(rw)

	response, err := makeTimeZoneResponse(tz)
	if err != nil {
		http.Error(rw, "Internal server error", 500)
	}
	json.NewEncoder(rw).Encode(response)
}

func makeTimeZoneResponse(tzs string) (map[string]string, error) {
	response := make(map[string]string)

	tzsList := strings.Split(tzs, ",")

	for _, v := range tzsList {
		loc, err := time.LoadLocation(v)
		fmt.Println(loc)
		if err != nil {
			return nil, err
		}
		response[v] = time.Now().In(loc).String()
	}
	return response, nil
}

func isValidTimezone(tz string) bool {
	return tz != ""
}
