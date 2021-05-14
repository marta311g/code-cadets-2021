package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type jobApplication struct {
	Name string
	Age int
	Passed bool
	Skills []string
}

const jobApplicationsURL = "https://run.mocky.io/v3/f7ceece5-47ee-4955-b974-438982267dc8"

func linearBackoff(retry int) time.Duration {
	return time.Duration(retry) * time.Second
}

func containsJavaOrGo(skills []string) bool {
	for _, skill := range skills {
		if skill == "Java" || skill == "Go" {
			return true
		}
	}
	return false
}

func writePeopleInFile(f *os.File, jobApplications []jobApplication) {
	for _, application := range jobApplications {
		if application.Passed && containsJavaOrGo(application.Skills){
			f.WriteString(fmt.Sprintf("%s - %s\n", application.Name, strings.Join(application.Skills, ", ")))
		}
	}
}

func main() {
	httpClient := pester.New()
	httpClient.Backoff = linearBackoff

	httpResponse, err := httpClient.Get(jobApplicationsURL)
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

	var jobApplications []jobApplication
	err = json.Unmarshal(bodyContent, &jobApplications)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}

	file, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}
	defer file.Close()

	writePeopleInFile(file, jobApplications)
	file.Sync()
}
