package aplikacja

import (
	"encoding/json"
	"fmt"
	"strings"

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

// Gets command and url from !add <command> <url>
func GetCMDvariables(messegeContents string) (string, string, error) {
	res := strings.TrimPrefix(messegeContents, "!add ")
	var res1, url string
	for i := 0; i < len(res); i++ {
		if res[i] == ' ' {
			res1 = res[:i]
			url = res[i:]
			break
		}
	}
	return res1, url, nil
}
