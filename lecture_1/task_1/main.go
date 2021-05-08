package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type serverData struct {
	Name string
	Age int
	Passed bool
	Skills []string
}

const applicationsURL = "https://run.mocky.io/v3/f7ceece5-47ee-4955-b974-438982267dc8"

func linearBackoff(retry int) time.Duration {
	return time.Duration(retry) * time.Second
}

func containsJavaOrGo(skills []string) bool {
	for _, val := range skills {
		if val == "Java" || val == "Go" {
			return true
		}
	}
	return false
}

func writePeopleInFile(f *os.File, data []serverData) {
	for _, val := range data {
		if val.Passed && containsJavaOrGo(val.Skills){
			var allSkills string
			for idx, value := range val.Skills {
				if idx != 0 {
					allSkills += ", " + value
				} else {
					allSkills += " - " + value
				}
			}
			defer f.WriteString(fmt.Sprint(val.Name + allSkills) + "\n")
		}
	}
}

func main() {
	httpClient := pester.New()
	httpClient.Backoff = linearBackoff

	httpResponse, err := httpClient.Get(applicationsURL)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "HTTP get towards yesno API"),
		)
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "reading body of yesno API response"),
		)
	}

	var decodedContent []serverData
	err = json.Unmarshal(bodyContent, &decodedContent)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}

	defer f.Close()
	writePeopleInFile(f, decodedContent)
	f.Sync()
}
