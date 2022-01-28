package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type content struct{
	Name 		string  	`json:"name"`
	Icon		string		`json:"icon"`
}

type chest struct{
	Status 	    	bool		`json:"status"`
	Data	    	struct  {
		Name 		string  	`json:"name"`
        XP      	int         `json:"xp"`
		Coins		int			`json:"coins"`
		Contents	[]content	`json:"contentItems"`
	}		                	`json:"data"`
}

func OpenChest(userId string, chestId string) (chest,error) {
    values := map[string]string{"userId": userId, "chest": chestId}
    json_data, err := json.Marshal(values)

    if err != nil {
        log.Fatal(err)
    }

    resp, err := http.Post("http://localhost:3000/chest", "application/json",
        bytes.NewBuffer(json_data))
	var res chest

    if err != nil {
        return res, err
    }
    json.NewDecoder(resp.Body).Decode(&res)

	return res, nil
}