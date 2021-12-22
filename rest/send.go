package rest

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendGET(url string) (*[]byte, error) {
	fmt.Printf("get\n")
	response, err := http.Get(url)
	//response.Header.Set("Cookie", "csv=2; edgebucket=IK3PSmXX4D0hiW6qiN; loid=0000000000hhozc2jz.2.1639324324224.Z0FBQUFBQmh0aHFrOGZ5NDNaZW01MU52alp1WUNCdjFPdW10WkRrYzJXM2xTRzgwdEhvWk5PSHJzS1dhYzNsa2V3VlRRVVN1RnVCLXVRdXRHY0l1cENXQUtUTC15RnV0LTJUZXE2Nnl4dlVZd2xudXlaYUFCcnBfRW9iRzNkZTJ2UHpuZmszTWx4Rzc")
	response.Header.Set("user-Agent", "my-user-agent")
	response.Header.Set("client-id", "my-client-id")
	response.Header.Set("client-secret", "my-client-secret")
	//response.Header.Set("Accept", "*/*")
	//response.Header.Set("Accept-Encoding", "gzip, deflate, br")
	//response.Header.Set("Connection", "keep-alive")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if response.Body != nil {
		defer response.Body.Close()
	}

	if response.StatusCode != 200 {
		eror := fmt.Sprintf("GET: Bad Request StatusCode = %d \n", response.StatusCode)
		return nil, errors.New(eror)
	}

	byteValue, _ := ioutil.ReadAll(response.Body)

	return &byteValue, nil
}
