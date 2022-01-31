package services

import (
	"bytes"
	"encoding/json"
	"log"
	"mmorpg-bot/src/helpers"
	"net/http"
)

type team struct {
	UserId string `json:"userId"`
	Hp     int    `json:"hp"`
	MaxHp  int    `json:"maxHp"`
	Sp     int    `json:"sp"`
	MaxSp  int    `json:"maxSp"`
	Name   string `json:"name"`
}

type Match struct {
	Status 		bool 	`json:"status"`
	Data   		struct {
		TeamA  	[]team 	`json:"teamA"`
		TeamB  	[]team 	`json:"teamB"`
		TurnNo	int    	`json:"turnNo"`
		TurnOf	string 	`json:"turnOf"`
	} 					`json:"data"`
}

func RandomMatch(userId string) (Match, error) {
	values := map[string]string{"userId": userId}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:3000/match", "application/json",
		bytes.NewBuffer(json_data))

	var res Match
	if err != nil {
		return res, err
	}

	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}

func HandleAttack(matchId string, charId string, Type string, attackType string, abilityId string) error {
	json_data, _ := json.Marshal(map[string]string{
		"match":   matchId,
		"id":      charId,
		"type":    Type,
		"attack":  attackType,
		"ability": abilityId,
	})
	
	_, err := helpers.Call("http://localhost:3000/match/attack", "POST", bytes.NewBuffer(json_data), "")
	if err != nil {
		return err
	}
	return nil
}
