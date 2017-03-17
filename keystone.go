package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"os"
	"crypto/tls"
)


type keystoneResponse struct {
	Access struct {
		Token struct {
			Id string
		}
	}
}


// GetKeystoneToken extract new token from KeystoneV2
func GetKeystoneToken(endpoint string, username string, password string, tenant string) string {
	url := endpoint + "/v2.0/tokens"
	authBody := `{
		"auth": {
			"passwordCredentials": {
				"username": "` + username + `",
				"password": "` + password + `"
			},
			"tenantName": "` + tenant + `"
		}
	}`

	// For https we will drop cert verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    	}
	client := &http.Client{Transport: tr}

	res, err := client.Post(
		url,
		"application/json; charset=utf-8",
		strings.NewReader(authBody))

	if err != nil {
		fmt.Println("Can't get responce from Keystone.")
		os.Exit(1)
	}
	body2, _ := ioutil.ReadAll(res.Body)

	kr := keystoneResponse{}
	err = json.Unmarshal(body2, &kr)
	if err != nil {
		fmt.Println(err.Error())
	}

	return kr.Access.Token.Id
}
