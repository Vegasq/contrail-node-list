package main

import (
	"net/http"
	"strings"
	"fmt"
	"os"
	"io/ioutil"
)

type ContrailAPI struct {
	Token string
	ContrailAPI string
}


func (c *ContrailAPI) ApiCall(Path string) []byte {
	client := &http.Client{}

	if strings.Contains(Path, c.ContrailAPI) {
		Path = strings.Replace(Path, c.ContrailAPI, "", 1)
	}

	req, err := http.NewRequest("GET", c.ContrailAPI + Path, nil)
	if err != nil {
		fmt.Printf("Can't create Request object.")
		os.Exit(1)
	}
	req.Header.Add("X-Auth-Token", c.Token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("No response from Contrail API.", c.ContrailAPI,
			Path, err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}
