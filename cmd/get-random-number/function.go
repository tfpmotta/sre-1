package getrandomnumber

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"fmt"
)

const (
	RandomNumberGetSource = "http://www.randomnumberapi.com/api/v1.0/random?min=1&max=52&count=52"
)

var (
	RandomNumberRegexMatch = regexp.MustCompile("\n")
)

type RandomNumberResponse struct {
	RandomNumber string `json:"random_number"`
}

// GetRandomNumber is a function to return a random number
func GetRandomNumber(response http.ResponseWriter, request *http.Request) {

	sourceResponse, err := http.Get(RandomNumberGetSource)
	if err != nil {
		log.Println(err)
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	defer sourceResponse.Body.Close()

	sourceBody, err := ioutil.ReadAll(sourceResponse.Body)
	if err != nil {
		log.Println(err)
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	sourceBody1 := strings.Replace(string(sourceBody), "[", "", -1)
	sourceBody2 := strings.Replace(string(sourceBody1), "]", "", -1)
	sourceBody3 := strings.Replace(string(sourceBody2), ",", "", -1)
	sourceBody4 := strings.Replace(string(sourceBody3), " ", "", -1)
	fmt.Println(sourceBody4)

	out := &RandomNumberResponse{
		RandomNumber: RandomNumberRegexMatch.ReplaceAllString(string(sourceBody4), ""),
	}

	response.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(response).Encode(out)
	if err != nil {
		log.Printf("failed to encode json to HTTP response: %v", err)
	}

}
