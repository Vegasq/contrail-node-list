package main

import (
	"encoding/json"
	"fmt"
	"os"
)


// UNIFIED NODES OWNER
type ContrailNodes struct {
	Nodes []contrailNode
}

func (cn *ContrailNodes) GetAll() []contrailNode {
	return cn.Nodes
}
func (cn *ContrailNodes) GetKnownTypes() []string {
	return []string{"config-node", "analytics-node", "database-node",
		"virtual-router"}
}
func (cn *ContrailNodes) GetSplitByType() [][]contrailNode {
	knownKeys := map[string]bool{}
	for _, n := range cn.Nodes{
		knownKeys[n.Type] = true
	}

	nodes := [][]contrailNode{}

	for key := range knownKeys {
		nodesBlock := []contrailNode{}
		for _, node := range cn.Nodes {
			if node.Type == key {
				nodesBlock = append(nodesBlock, node)
			}
		}
		nodes = append(nodes, nodesBlock)
	}

	return nodes
}
func (cn *ContrailNodes) GetByType(Type string) []contrailNode {
	nodesBlock := []contrailNode{}
	for _, node := range cn.Nodes {
		if node.Type == Type {
			nodesBlock = append(nodesBlock, node)
		}
	}
	return nodesBlock
}


// UNIFIED NODE TYPE
type contrailNode struct {
	Type string
	UUID string
	Name string
	IP string
}

func (c contrailNode) AsList() []string {
	return []string{c.Type, c.UUID, c.Name, c.IP}
}


// PER NODE API RESPONSE

type jsonConfigNode struct {
	Node struct {
		NodeIPAddress string `json:"config_node_ip_address"`
		FQName []string `json:"fq_name"`
		UUID string `json:"uuid"`
		Name string `json:"name"`
		DisplayName string `json:"display_name"`
		ParentType string `json:"parent_type`
	} `json:"config-node"`
}

type jsonAnalyticsNode struct {
	Node struct {
		NodeIPAddress string `json:"analytics_node_ip_address"`
		FQName []string `json:"fq_name"`
		UUID string `json:"uuid"`
		Name string `json:"name"`
		DisplayName string `json:"display_name"`
		ParentType string `json:"parent_type`
	} `json:"analytics-node"`
}

type jsonDatabaseNode struct {
	Node struct {
		NodeIPAddress string `json:"database_node_ip_address"`
		FQName []string `json:"fq_name"`
		UUID string `json:"uuid"`
		Name string `json:"name"`
		DisplayName string `json:"display_name"`
		ParentType string `json:"parent_type`
	} `json:"database-node"`
}

type jsonVRouterNode struct {
	Node struct {
		NodeIPAddress string `json:"virtual_router_ip_address"`
		FQName []string `json:"fq_name"`
		UUID string `json:"uuid"`
		Name string `json:"name"`
		DisplayName string `json:"display_name"`
		ParentType string `json:"parent_type`
	} `json:"virtual_router"`
}


// PER BUCKET API RESPONSE
type jsonNodesStruct struct {
	Href string `json:"href"`
	FqName []string `json:"fq_name"`
	Uuid string `json:"uuid"`
}

type jsonConfigStruct struct {
	Nodes []jsonNodesStruct `json:"config-nodes"`
}
type jsonAnalyticsStruct struct {
	Nodes []jsonNodesStruct `json:"analytics-nodes"`
}
type jsonDatabaseStruct struct {
	Nodes []jsonNodesStruct `json:"database-nodes"`
}
type jsonVRouterStruct struct {
	Nodes []jsonNodesStruct `json:"virtual-routers"`
}


type configBucket struct {
	ContrailAPI ContrailAPIInterface
	BucketPath string
	Nodes []jsonConfigNode
}
func (c *configBucket) loadBucket(){
	resp := c.ContrailAPI.ApiCall(c.BucketPath)

	nodes := jsonConfigStruct{}
	if err := json.Unmarshal(resp, &nodes); err != nil {
		fmt.Println("Can't parse JSON response", string(resp))
		os.Exit(1)
	}

	for _, n := range nodes.Nodes{
		resp = c.ContrailAPI.ApiCall(n.Href)
		node := jsonConfigNode{}
		if err := json.Unmarshal(resp, &node); err != nil {
			fmt.Println("Can't parse JSON response", string(resp))
			os.Exit(1)
		}
		c.Nodes = append(c.Nodes, node)
	}
}
func (b *configBucket) getNodes() []contrailNode {
	nodes := []contrailNode{}
	for _, n := range b.Nodes{
		nodes = append(nodes, contrailNode{
			"config-node",
			n.Node.UUID,
			n.Node.Name,
			n.Node.NodeIPAddress,
		})
	}
	return nodes
}


type analyticsBucket struct {
	ContrailAPI ContrailAPIInterface
	BucketPath string
	Nodes []jsonAnalyticsNode
}
func (a *analyticsBucket) loadBucket(){
	resp := a.ContrailAPI.ApiCall(a.BucketPath)

	nodes := jsonAnalyticsStruct{}
	if err := json.Unmarshal(resp, &nodes); err != nil {
		fmt.Println("Can't parse JSON response from " +
			"Keystone. Actual response:", string(resp))
		os.Exit(1)
	}

	for _, n := range nodes.Nodes{
		resp = a.ContrailAPI.ApiCall(n.Href)
		node := jsonAnalyticsNode{}
		if err := json.Unmarshal(resp, &node); err != nil {
			fmt.Println("Can't parse JSON response", string(resp))
			os.Exit(1)
		}
		a.Nodes = append(a.Nodes, node)
	}
}
func (b *analyticsBucket) getNodes() []contrailNode {
	nodes := []contrailNode{}
	for _, n := range b.Nodes{
		nodes = append(nodes, contrailNode{
			"analytics-node",
			n.Node.UUID,
			n.Node.Name,
			n.Node.NodeIPAddress,
		})
	}
	return nodes
}


type databaseBucket struct {
	ContrailAPI ContrailAPIInterface
	BucketPath string
	Nodes []jsonDatabaseNode
}
func (d *databaseBucket) loadBucket(){
	resp := d.ContrailAPI.ApiCall(d.BucketPath)

	nodes := jsonDatabaseStruct{}
	if err := json.Unmarshal(resp, &nodes); err != nil {
		fmt.Println("Can't parse JSON response from " +
			"Keystone. Actual response:", string(resp))
		os.Exit(1)
	}

	for _, n := range nodes.Nodes{
		resp = d.ContrailAPI.ApiCall(n.Href)
		node := jsonDatabaseNode{}
		if err := json.Unmarshal(resp, &node); err != nil {
			fmt.Println("Can't parse JSON response", string(resp))
			os.Exit(1)
		}
		d.Nodes = append(d.Nodes, node)
	}
}
func (b *databaseBucket) getNodes() []contrailNode {
	nodes := []contrailNode{}
	for _, n := range b.Nodes{
		nodes = append(nodes, contrailNode{
			"database-node",
			n.Node.UUID,
			n.Node.Name,
			n.Node.NodeIPAddress,
		})
	}
	return nodes
}


type vRouterBucket struct {
	ContrailAPI ContrailAPIInterface
	BucketPath string
	Nodes []jsonVRouterNode
}
func (b *vRouterBucket) loadBucket(){
	resp := b.ContrailAPI.ApiCall(b.BucketPath)

	nodes := jsonVRouterStruct{}
	if err := json.Unmarshal(resp, &nodes); err != nil {
		fmt.Println("Can't parse JSON response from " +
			"Keystone. Actual response:", string(resp))
		os.Exit(1)
	}

	for _, n := range nodes.Nodes{
		resp = b.ContrailAPI.ApiCall(n.Href)
		node := jsonVRouterNode{}
		if err := json.Unmarshal(resp, &node); err != nil {
			fmt.Println("Can't parse JSON response", string(resp))
			os.Exit(1)
		}
		b.Nodes = append(b.Nodes, node)
	}
}
func (b *vRouterBucket) getNodes() []contrailNode {
	nodes := []contrailNode{}
	for _, n := range b.Nodes{
		nodes = append(nodes, contrailNode{
			"vrouter-node",
			n.Node.UUID,
			n.Node.Name,
			n.Node.NodeIPAddress,
		})
	}
	return nodes
}


func ContrailNodeManager(capi ContrailAPIInterface) ContrailNodes {
	cn := ContrailNodes{}

	cb := configBucket{capi, "/config-nodes", nil}
	cb.loadBucket()
	cn.Nodes = append(cn.Nodes, cb.getNodes()...)

	ab := analyticsBucket{capi, "/analytics-nodes", nil}
	ab.loadBucket()
	cn.Nodes = append(cn.Nodes, ab.getNodes()...)

	db := databaseBucket{capi, "/database-nodes", nil}
	db.loadBucket()
	cn.Nodes = append(cn.Nodes, db.getNodes()...)

	vb := vRouterBucket{capi, "/virtual-routers", nil}
	vb.loadBucket()
	cn.Nodes = append(cn.Nodes, vb.getNodes()...)

	return cn
}
