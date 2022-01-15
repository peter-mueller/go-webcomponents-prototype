package exampleapi

import (
	"encoding/json"
	"net/http"
)

type AnimalFact struct {
	Animal string `json:"type"`
	Text   string `json:"text"`
}

func GetFacts() chan []AnimalFact {
	c := make(chan []AnimalFact)
	go func() {
		defer close(c)
		res, err := http.Get("https://cat-fact.herokuapp.com/facts/")
		if err != nil {
			panic(err)
		}
		facts := make([]AnimalFact, 0)
		json.NewDecoder(res.Body).Decode(&facts)
		c <- facts
	}()

	return c
}
