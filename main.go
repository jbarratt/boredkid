package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"time"
)

func randString(s []string) string {
	n := rand.Intn(len(s))
	return s[n]
}

func boredString() (string, error) {

	bored := boredMap()
	boredFunc := template.FuncMap{}
	for k, _ := range bored {
		item := k // keep it around for closure
		boredFunc[item] = func() string {
			return randString(bored[item])
		}
	}

	finalString := "{{ seed }}"
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

	boredSolution, err := boredString()
	if err != nil {
		log.Fatalf("error coming up with a solution: %v", err)
	}
	fmt.Println(boredSolution)
}
