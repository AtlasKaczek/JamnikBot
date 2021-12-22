package aplikacja

import (
	"encoding/json"
	"fmt"

	"github.com/go-rod/rod"
)

func GetJamnikObj() (Images, error) {
	browser := rod.New().MustConnect()

	defer browser.MustClose()

	page := browser.MustPage("https://www.reddit.com/r/Dachshund.json")
	resp, err := page.GetResource("https://www.reddit.com/r/Dachshund.json")
	if err != nil {
		fmt.Printf("GetRandomJamnik 1: An error occured: %v\n", err)
		return Images{}, err
	}

	var jamnik Images

	jsonErr := json.Unmarshal(resp, &jamnik)
	if jsonErr != nil {
		fmt.Printf("GetRandomJamnik 2: An error occured: %v\n", jsonErr)
		return Images{}, jsonErr
	}

	return jamnik, nil
}
