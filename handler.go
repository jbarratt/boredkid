package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"time"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	alexa "github.com/mikeflynn/go-alexa/skillserver"
)

func randString(s []string) string {
	n := rand.Intn(len(s))
	return s[n]
}

func randomBoredString() (string, error) {
	result, err := parseBoredSeed("{{ seed }}")
	return result, err
}

func parseBoredSeed(seed string) (string, error) {

	bored := boredMap()
	boredFunc := template.FuncMap{}
	for k, _ := range bored {
		item := k // keep it around for closure
		boredFunc[item] = func() string {
			return randString(bored[item])
		}
	}

	finalString := seed
	for {
		parser := template.New("").Funcs(boredFunc)
		parsed, err := parser.Parse(finalString)
		if err != nil {
			return "", err
		}
		var content bytes.Buffer
		if err := parsed.Execute(&content, ""); err != nil {
			return "", err
		}
		// when the strings match, there is nothing left to expand.
		if content.String() == finalString {
			break
		}
		finalString = content.String()
	}
	return finalString, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	boredSolution, err := randomBoredString()
	if err != nil {
		log.Fatalf("error coming up with a solution: %v", err)
	}
	fmt.Println(boredSolution)
}

func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	rand.Seed(time.Now().UnixNano())

	echoResponse := alexa.NewEchoResponse()
	boredSolution, err := parseBoredSeed("{{ alexaseed }}")

	if err != nil {
		log.Printf("error coming up with a solution: %v", err)
		echoResponse.OutputSpeech("I had a problem coming up with an idea. Looks like this one is on you.")
	} else {
		echoResponse.OutputSpeech(boredSolution)
	}

	return echoResponse, nil
}
