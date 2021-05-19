package main

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/profburke/bgurt/aws/utilities"
	"github.com/profburke/bgurt/bggclient"
	"github.com/profburke/bgurt/geekbadge"
	"golang.org/x/image/colornames"
)

var borderColors []string
var boxColors []string
var leftWords []string
var rightWords []string

func init() {
	borderColors = strings.Split(utilities.GetEnvOrDie("BGGBORDERCOLORS"), ",")
	boxColors = strings.Split(utilities.GetEnvOrDie("BGGBOXCOLORS"), ",")
	leftWords = strings.Split(utilities.GetEnvOrDie("BGGLEFTWORDS"), ",")
	rightWords = strings.Split(utilities.GetEnvOrDie("BGGRIGHTWORDS"), ",")

	rand.Seed(time.Now().UnixNano())
}

func pickColors() (firstColor, secondColor string) {
	firstColor = boxColors[rand.Intn(len(boxColors))]
	for {
		secondColor = boxColors[rand.Intn(len(boxColors))]
		if firstColor != secondColor {
			break
		}
	}

	return
}

func HandleRequest() {
	user := utilities.GetEnvOrDie("BGGUSERNAME")
	passhash := utilities.GetEnvOrDie("BGGPASSHASH")

	bggclient.SetCredentials(bggclient.Credentials{user, passhash})

	log.Println("Updating geekbadge...")

	firstColor, secondColor := pickColors()
	log.Printf("Box colors: %s, %s\n", firstColor, secondColor)

	leftWord := leftWords[rand.Intn(len(leftWords))]

	lb := geekbadge.Box{
		Text:       leftWord,
		Background: colornames.Map[firstColor],
		TextColor:  colornames.Map[secondColor],
		TextStart:  4,
	}

	rightWord := rightWords[rand.Intn(len(rightWords))]

	log.Printf("Left word: %s, Right word: %s", leftWord, rightWord)

	rb := geekbadge.Box{
		Text:       rightWord,
		Background: colornames.Map[secondColor],
		TextColor:  colornames.Map[firstColor],
		TextStart:  44,
	}

	outerColor := borderColors[rand.Intn(len(borderColors))]
	innerColor := borderColors[rand.Intn(len(borderColors))]

	log.Printf("Border colors: %s %s\n", outerColor, innerColor)

	gb := geekbadge.Geekbadge{
		OuterBorder: colornames.Map[outerColor],
		InnerBorder: colornames.Map[innerColor],
		BarPosition: 40,
		LeftBox:     lb,
		RightBox:    rb,
	}

	_, err := geekbadge.Set(gb)
	if err != nil {
		log.Printf("Error updating geekbadge: %v", err)
	} else {
		log.Println("Geekbadge updated.")
	}
}

func main() {
	lambda.Start(HandleRequest)
}

// Local Variables:
// compile-command: "GOOS=linux go build"
// End:
