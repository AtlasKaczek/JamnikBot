package aplikacja

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
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
			url = res[i+1:]
			break
		}
	}
	return res1, url, nil
}

//Saves to "dat" file
func Stf(res1, res2 string) error {
	_, err := os.Stat("dat")
	if os.IsNotExist(err) {
		f, err := os.Create("dat")
		if err != nil {
			fmt.Printf("STF | An error occured: %e\n", err)
		}
		defer f.Close()

		res, rerr := f.Write([]byte("!" + res1 + " " + res2 + "\n"))
		if rerr != nil {
			fmt.Printf("STF | An error occured: %e\n", rerr)
		} else {
			fmt.Printf("STF | wrote %d bytes\n", res)
		}
	} else {
		f, err := os.OpenFile("dat", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("STF | An error occured: %e\n", err)
		}
		defer f.Close()

		res, rerr := f.Write([]byte("!" + res1 + " " + res2 + "\n"))
		if rerr != nil {
			fmt.Printf("STF | An error occured: %e\n", rerr)
		} else {
			fmt.Printf("STF | wrote %d bytes\n", res)
		}

	}

	return nil
}

// Checks if URL is OK
func CheckURL(url string) (bool, error) {
	if url[:25] != "https://www.reddit.com/r/" {
		err := fmt.Sprintf(" Bad URL %s (does not match pattern) ", url)
		return false, errors.New(err)
	}
	text := ReadDat()
	for _, each_ln := range text {
		if strings.Contains(each_ln, url) {
			err := fmt.Sprintf(" This %s URL already in used by command ", url)
			return false, errors.New(err)
		}
	}
	return true, nil
}

// Checks if command is OK
func CheckCMD(cmd string) (bool, error) {
	text := ReadDat()
	for _, each_ln := range text {
		if strings.Contains(each_ln, cmd) {
			err := fmt.Sprintf(" %s command already exists ", cmd)
			return false, errors.New(err)
		}
	}
	return true, nil
}

// Reads data line by line from "dat" file
func ReadDat() []string {
	file, err := os.Open("dat")
	if err != nil {
		fmt.Printf("ReadDat | An error occured: %e\n", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	return text
}
