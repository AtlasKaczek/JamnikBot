package aplikacja

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"stankryj/JamnikBot/rest"
)

func GetRandomJamnik() string {
	resp, err := rest.SendGET("https://www.reddit.com/r/Dachshund.json")
	if err != nil {
		fmt.Println(err)
	}

	var jamnik Images

	jsonErr := json.Unmarshal(*resp, &jamnik)
	if jsonErr != nil {
		fmt.Printf("ParseJSON 1: An error occured: %v", jsonErr)
	}

	return jamnik.GetImageURL(rand.Intn(jamnik.GetChildrenLen()))
}
