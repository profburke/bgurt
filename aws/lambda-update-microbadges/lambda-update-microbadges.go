package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/profburke/bgurt/aws/utilities"
	"github.com/profburke/bgurt/bggclient"
	"github.com/profburke/bgurt/microbadge"
)

var badges []uint

// TODO: migrate away from having a list of IDs in an environment variable
// and instead pull from a file in an S3 bucket (or ??).
// File to be maintained by a separate lambda running microbadge.GetAll
func init() {
	badgesString := utilities.GetEnvOrDie("BGGBADGEIDS")
	err := json.Unmarshal([]byte(badgesString), &badges)
	if err != nil {
		log.Fatalf("Could not parse badge IDs: %v", err)
	}
}

func HandleRequest() {
	user := utilities.GetEnvOrDie("BGGUSERNAME")
	passhash := utilities.GetEnvOrDie("BGGPASSHASH")

	bggclient.SetCredentials(bggclient.Credentials{user, passhash})

	log.Println("Updating badges...")

	newbadges := utilities.Picksome(badges, microbadge.TotalSlots)
	log.Printf("New microbadges: %v.", newbadges)

	_, err := microbadge.SetAll(newbadges)
	if err != nil {
		log.Printf("Error updating microbadges: %v.", err)
	} else {
		log.Println("Microbadges updated.")
	}
}

func main() {
	lambda.Start(HandleRequest)
}

// Local Variables:
// compile-command: "GOOS=linux go build"
// End:
