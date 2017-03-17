package main

import (
	"flag"
	"gopkg.in/vegasq/GoPrintTable.v1"
	"fmt"
	"strings"
)


func main() {
	var contrail = flag.String(
		"endpoint", "http://172.16.18.3:8182", "Contrail endpoint")

	var keystone = flag.String(
		"kyestone", "http://172.16.18.3:5000", "Keystone endpoint")

	var username = flag.String("username", "admin", "")
	var password = flag.String("password", "admin", "")
	var tenant = flag.String("tenant", "admin", "")

	flag.Parse()

	token := GetKeystoneToken(*keystone, *username, *password, *tenant)

	capi := ContrailAPI{token, *contrail}
	cnm := ContrailNodeManager(capi)

	for _, Type := range cnm.GetKnownTypes(){

		fmt.Println(strings.ToUpper(Type))

		table := [][]string{}
		table = append(table, []string{"Node type", "UUID", "Node name", "IP"})

		for _, node := range cnm.GetByType(Type) {
			table = append(table, node.AsList())
		}

		GoPrintTable.PrintTableWithHeader(table)
	}

}
