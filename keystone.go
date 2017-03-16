package main

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type KeystoneResponce struct {
	Access struct {
		Token struct {
			Id string
		      }
	       }
}


func GetKeystoneToken(
	endpoint string,
	username string, password string,
	tenant string) string {

	url := endpoint+"/v2.0/tokens"
	authBody := `{
		"auth": {
			"passwordCredentials": {
				"username": "`+username+`",
				"password": "`+password+`"
			},
			"tenantName": "`+tenant+`"
		}
	}`

	res, err := http.Post(url, "application/json; charset=utf-8",
		strings.NewReader(authBody))
	if err != nil {
		fmt.Println("Can't get responce from Keystone.")
		fmt.Errorf(err.Error())
	}

	body2, _ := ioutil.ReadAll(res.Body)

	kr := KeystoneResponce{}
	err = json.Unmarshal(body2, &kr)
	if err != nil {
		fmt.Println(err.Error())
	}
	return kr.Access.Token.Id
}