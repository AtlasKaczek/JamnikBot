package rest

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendGET(url string) (*[]byte, error) {

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}

	if response.StatusCode != 200 {
		eror := fmt.Sprintf("GET: Bad Request StatusCode = %d", response.StatusCode)
		return nil, errors.New(eror)
	}

	byteValue, _ := ioutil.ReadAll(response.Body)

	return &byteValue, nil
}
