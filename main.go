package main

import (
	"flag"
)


func main() {
	//header := []string{"Node Name", "IP", "Status"}
	//node1 := []string{"controller", "192.168.0.100", "Online"}
	//node2 := []string{"minion1", "192.168.0.101", "Offline"}
	//node3 := []string{"minion2", "192.168.0.102"}
	//t := [][]string{header, node1, node2, node3}
	//GoPrintTable.PrintTableWithHeader(t)

	var contrail = flag.String(
		"endpoint", "http://172.16.18.3:8182", "Contrail endpoint")

	var keystone = flag.String(
		"kyestone", "http://172.16.18.3:5000", "Keystone endpoint")

	var username = flag.String("username", "admin", "")
	var password = flag.String("password", "admin", "")
	var tenant = flag.String("tenant", "admin", "")

	flag.Parse()

	token := GetKeystoneToken(*keystone, *username, *password, *tenant)

	c := Contrail{token, *contrail, nil}
	c.LoadNodes()
	c.DisplayNodes()
}
