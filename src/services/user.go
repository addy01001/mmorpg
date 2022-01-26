package services

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type userRes struct{
	Status 	    bool		`json:"status"`
	Data	    struct  {
		UserId 	string		`json:"userId"`
		Name 	string  	`json:"name"`
        Level   int         `json:"lvl"`
        XP      int         `json:"xp"`
	}		                `json:"data"`
}

func StartUser(userId string, name string) (userRes,error) {
    values := map[string]string{"userId": userId, "name": name}
    json_data, err := json.Marshal(values)

    if err != nil {
        log.Fatal(err)
    }

    resp, err := http.Post("http://localhost:3000/user", "application/json",
        bytes.NewBuffer(json_data))
	var res userRes

    if err != nil {
        return res, err
    }
    json.NewDecoder(resp.Body).Decode(&res)

	return res, nil
}