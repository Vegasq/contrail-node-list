package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/vegasq/GoPrintTable"

)

type IGenericNodes interface {
	NodesList() []GenericNode
	GetType() string
}

type GenericNode struct {
	Href string `json:"href"`
	FqName []string `json:"fq_name"`
	Uuid string `json:"uuid"`
}

type ConfigNodes struct {
	NodesBucket []GenericNode `json:"config-nodes"`
}

func (nodes *ConfigNodes) NodesList() []GenericNode {
	return nodes.NodesBucket
}
func (nodes *ConfigNodes) GetType() string {
	return "Config Nodes"
}

type DatabaseNodes struct {
	NodesBucket []GenericNode `json:"database-nodes"`
}

func (nodes *DatabaseNodes) NodesList() []GenericNode {
	return nodes.NodesBucket
}
func (nodes *DatabaseNodes) GetType() string {
	return "Database Nodes"
}


type AnalyticsNodes struct {
	NodesBucket []GenericNode `json:"analytics-nodes"`
}
func (nodes *AnalyticsNodes) NodesList() []GenericNode {
	return nodes.NodesBucket
}
func (nodes *AnalyticsNodes) GetType() string {
	return "Analytics Nodes"
}


type VRouterNodes struct {
	NodesBucket []GenericNode `json:"vrouters-nodes"`
}
func (nodes *VRouterNodes) NodesList() []GenericNode {
	return nodes.NodesBucket
}
func (nodes *VRouterNodes) GetType() string {
	return "VRouter Nodes"
}


type Contrail struct {
	Token string
	ContrailAPI string
	Nodes []IGenericNodes
}


func (c *Contrail) ApiCall(Path string) []byte {
	client := &http.Client{}

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


func (c *Contrail) LoadNodes() {
	nodeTypes := []string{
		"/config-nodes",
		"/database-nodes",
		"/analytics-nodes",
		"/virtual-routers"}
	c.Nodes = append(c.Nodes, &ConfigNodes{})
	c.Nodes = append(c.Nodes, &DatabaseNodes{})
	c.Nodes = append(c.Nodes, &AnalyticsNodes{})
	c.Nodes = append(c.Nodes, &VRouterNodes{})

	for i, nodeType := range nodeTypes {
		resp := c.ApiCall(nodeType)

		if err := json.Unmarshal(resp, &c.Nodes[i]); err != nil {
			fmt.Println("Can't parse JSON response from " +
				"Keystone. Actual response:", string(resp))
			os.Exit(1)
		}
	}
}

func (c *Contrail) DisplayNodes() {
	for _, node := range c.Nodes {
		buildTable := [][]string{}
		fmt.Println("\n", node.GetType())

		buildTable = append(buildTable, []string{"Domain", "UUID"})

		exactNodes := node.NodesList()
		for _, enode := range exactNodes {
			buildTable = append(buildTable,
				[]string{enode.FqName[1], enode.Uuid})
		}
		GoPrintTable.PrintTableWithHeader(buildTable)
	}

}