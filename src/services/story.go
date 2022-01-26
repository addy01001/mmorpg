package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type story struct {
	Status 			bool 		`json:"status"`
	Data   			struct {
		Order       int      	`json:"order"`
		StoryType   string   	`json:"storyType"`
		Choices     []string 	`json:"choices"`
		Branch      string   	`json:"branch"`
		DisplayText string   	`json:"displayText"`
	} 							`json:"data"`
}

func GetStory(userId string) (story, error) {
	values := map[string]string{"userId": userId}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:3000/story", "application/json",
		bytes.NewBuffer(json_data))
	var res story

	if err != nil {
		return res, err
	}
	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}

func AdvanceStory(userId string) (story, error) {
	values := map[string]string{"userId": userId}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Post("http://localhost:3000/story/advance", "application/json",
		bytes.NewBuffer(json_data))
	var res story

	if err != nil {
		return res, err
	}
	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}
