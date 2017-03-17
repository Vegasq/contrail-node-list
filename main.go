package main

import (
	"flag"
)


func main() {
	var contrail = flag.String(
		"endpoint", "http://10.0.0.1:8182", "Contrail endpoint")

	var keystone = flag.String(
		"kyestone", "https://10.0.0.2:5000", "Keystone endpoint")

	var username = flag.String("username", "admin", "")
	var password = flag.String("password", "admin", "")
	var tenant = flag.String("tenant", "admin", "")

	flag.Parse()

	token := GetKeystoneToken(*keystone, *username, *password, *tenant)

	c := Contrail{token, *contrail, nil}
	c.LoadNodes()
	c.DisplayNodes()
}
