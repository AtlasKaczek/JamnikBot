package aplikacja

import (
	"encoding/json"
	"fmt"

	"stankryj/JamnikBot/rest"
)

func GetRandomJamnik(url string) (Images, error) {
	resp, err := rest.SendGET(url)
	if err != nil {
		fmt.Printf("GetRandomJamnik 1: An error occured: %v\n", err)
		return Images{}, err
	}

	var jamnik Images

	jsonErr := json.Unmarshal(*resp, &jamnik)
	if jsonErr != nil {
		fmt.Printf("GetRandomJamnik 2: An error occured: %v\n", jsonErr)
		return Images{}, jsonErr
	}

	return jamnik, nil
}
