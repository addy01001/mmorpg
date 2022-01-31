package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"mmorpg-bot/src/helpers"
)

type content struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type chest struct {
	Status bool `json:"status"`
	Data   struct {
		Name     string    `json:"name"`
		XP       int       `json:"xp"`
		Coins    int       `json:"coins"`
		Contents []content `json:"contentItems"`
	} `json:"data"`
}

func OpenChest(userId string, chestId string, token string) (chest, error) {
	values := map[string]string{"userId": userId, "chest": chestId}
	json_data, err := json.Marshal(values)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(json_data)

	resp, err := helpers.Call("http://localhost:3000/chest", "POST",
		bytes.NewBuffer(json_data), token)

	var res chest

	if err != nil {
		return res, err
	}
	json.NewDecoder(resp).Decode(&res)

	return res, nil
}
