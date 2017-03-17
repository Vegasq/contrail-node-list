package main

import (
	"testing"
	"strings"
)


type FaceContrailAPI struct {}
func (c FaceContrailAPI) ApiCall(Path string) []byte {
	if strings.Contains(Path, "http://172.16.18.3:8182") {
		Path = strings.Replace(Path, "http://172.16.18.3:8182", "", 1)
	}

	if  Path == `/config-nodes` {
	 return []byte(`{"config-nodes": [{"href": "http://172.16.18.3:8182/config-node/0281ba31-3c5b-4cdf-998c-2849ab139598", "fq_name": ["default-global-system-config", "node-59.test.domain.local"], "uuid": "0281ba31-3c5b-4cdf-998c-2849ab139598"}, {"href": "http://172.16.18.3:8182/config-node/377a5385-6516-4092-b582-dd261a17fb5a", "fq_name": ["default-global-system-config", "node-61.test.domain.local"], "uuid": "377a5385-6516-4092-b582-dd261a17fb5a"}, {"href": "http://172.16.18.3:8182/config-node/721ed194-301a-4c1a-bc82-f78964674fab", "fq_name": ["default-global-system-config", "node-60.test.domain.local"], "uuid": "721ed194-301a-4c1a-bc82-f78964674fab"}]}`)
	}
	if  Path == `/config-node/0281ba31-3c5b-4cdf-998c-2849ab139598` {
	 return []byte(`{"config-node": {"fq_name": ["default-global-system-config", "node-59.test.domain.local"], "uuid": "0281ba31-3c5b-4cdf-998c-2849ab139598", "parent_uuid": "e55fefa1-5e0e-4d83-b342-583fcaf95997", "parent_href": "http://172.16.18.3:8182/global-system-config/e55fefa1-5e0e-4d83-b342-583fcaf95997", "parent_type": "global-system-config", "perms2": {"owner": "55251e647b954127bf528a73ba096ca5", "owner_access": 7, "global_access": 0, "share": []}, "href": "http://172.16.18.3:8182/config-node/0281ba31-3c5b-4cdf-998c-2849ab139598", "config_node_ip_address": "172.16.19.28", "id_perms": {"enable": true, "uuid": {"uuid_mslong": 180630180700310751, "uuid_lslong": 11064262681410377112}, "created": "2017-03-07T17:55:32.524967", "description": null, "creator": null, "user_visible": true, "last_modified": "2017-03-07T17:55:32.524967", "permissions": {"owner": "neutron", "owner_access": 7, "other_access": 7, "group": "admin", "group_access": 7}}, "display_name": "node-59.test.domain.local", "name": "node-59.test.domain.local"}}`)
	}
	if  Path == `/config-node/377a5385-6516-4092-b582-dd261a17fb5a` {
	 return []byte(`{"config-node": {"fq_name": ["default-global-system-config", "node-61.test.domain.local"], "uuid": "377a5385-6516-4092-b582-dd261a17fb5a", "parent_uuid": "e55fefa1-5e0e-4d83-b342-583fcaf95997", "parent_href": "http://172.16.18.3:8182/global-system-config/e55fefa1-5e0e-4d83-b342-583fcaf95997", "parent_type": "global-system-config", "perms2": {"owner": "55251e647b954127bf528a73ba096ca5", "owner_access": 7, "global_access": 0, "share": []}, "href": "http://172.16.18.3:8182/config-node/377a5385-6516-4092-b582-dd261a17fb5a", "config_node_ip_address": "172.16.19.3", "id_perms": {"enable": true, "uuid": {"uuid_mslong": 3997599451636449426, "uuid_lslong": 13079259423553616730}, "created": "2017-03-07T17:55:28.641200", "description": null, "creator": null, "user_visible": true, "last_modified": "2017-03-07T17:55:28.641200", "permissions": {"owner": "neutron", "owner_access": 7, "other_access": 7, "group": "admin", "group_access": 7}}, "display_name": "node-61.test.domain.local", "name": "node-61.test.domain.local"}}`)
	}
	if  Path == `/config-node/721ed194-301a-4c1a-bc82-f78964674fab` {
	 return []byte(`{"config-node": {"fq_name": ["default-global-system-config", "node-60.test.domain.local"], "uuid": "721ed194-301a-4c1a-bc82-f78964674fab", "parent_uuid": "e55fefa1-5e0e-4d83-b342-583fcaf95997", "parent_href": "http://172.16.18.3:8182/global-system-config/e55fefa1-5e0e-4d83-b342-583fcaf95997", "parent_type": "global-system-config", "perms2": {"owner": "55251e647b954127bf528a73ba096ca5", "owner_access": 7, "global_access": 0, "share": []}, "href": "http://172.16.18.3:8182/config-node/721ed194-301a-4c1a-bc82-f78964674fab", "config_node_ip_address": "172.16.19.29", "id_perms": {"enable": true, "uuid": {"uuid_mslong": 8223240404017499162, "uuid_lslong": 13583691595569909675}, "created": "2017-03-07T17:55:34.799875", "description": null, "creator": null, "user_visible": true, "last_modified": "2017-03-07T17:55:34.799875", "permissions": {"owner": "neutron", "owner_access": 7, "other_access": 7, "group": "admin", "group_access": 7}}, "display_name": "node-60.test.domain.local", "name": "node-60.test.domain.local"}}`)
	}
	if  Path == `/analytics-nodes` {
	 return []byte(`{"analytics-nodes": []}`)
	}
	if  Path == `/database-nodes` {
	 return []byte(`{"database-nodes": [{"href": "http://172.16.18.3:8182/database-node/b9a2036e-3318-448d-b243-2061fdfa1166", "fq_name": ["default-global-system-config", "node-59.test.domain.local"], "uuid": "b9a2036e-3318-448d-b243-2061fdfa1166"}, {"href": "http://172.16.18.3:8182/database-node/02a26c34-48ba-4d98-9720-359b8ec349dc", "fq_name": ["default-global-system-config", "node-61.test.domain.local"], "uuid": "02a26c34-48ba-4d98-9720-359b8ec349dc"}, {"href": "http://172.16.18.3:8182/database-node/e8a46bf3-d385-4bab-9353-ce404f6daddb", "fq_name": ["default-global-system-config", "node-60.test.domain.local"], "uuid": "e8a46bf3-d385-4bab-9353-ce404f6daddb"}]}`)
	}
	if  Path == `/database-node/b9a2036e-3318-448d-b243-2061fdfa1166` {
	 return []byte(`{"database-node": {"database_node_ip_address": "172.16.19.28", "fq_name": ["default-global-system-config", "node-59.test.domain.local"], "uuid": "b9a2036e-3318-448d-b243-2061fdfa1166", "parent_uuid": "e55fefa1-5e0e-4d83-b342-583fcaf95997", "parent_href": "http://172.16.18.3:8182/global-system-config/e55fefa1-5e0e-4d83-b342-583fcaf95997", "parent_type": "global-system-config", "perms2": {"owner": "55251e647b954127bf528a73ba096ca5", "owner_access": 7, "global_access": 0, "share": []}, "href": "http://172.16.18.3:8182/database-node/b9a2036e-3318-448d-b243-2061fdfa1166", "id_perms": {"enable": true, "uuid": {"uuid_mslong": 13376257615082308749, "uuid_lslong": 12845146167435727206}, "created": "2017-03-07T17:55:40.103035", "description": null, "creator": null, "user_visible": true, "last_modified": "2017-03-07T17:55:40.103035", "permissions": {"owner": "neutron", "owner_access": 7, "other_access": 7, "group": "admin", "group_access": 7}}, "display_name": "node-59.test.domain.local", "name": "node-59.test.domain.local"}}`)
	}
	if  Path == `/database-node/02a26c34-48ba-4d98-9720-359b8ec349dc` {
	 return []byte(`{"database-node": {"database_node_ip_address": "172.16.19.3", "fq_name": ["default-global-system-config", "node-61.test.domain.local"], "uuid": "02a26c34-48ba-4d98-9720-359b8ec349dc", "parent_uuid": "e55fefa1-5e0e-4d83-b342-583fcaf95997", "parent_href": "http://172.16.18.3:8182/global-system-config/e55fefa1-5e0e-4d83-b342-583fcaf95997", "parent_type": "global-system-config", "perms2": {"owner": "55251e647b954127bf528a73ba096ca5", "owner_access": 7, "global_access": 0, "share": []}, "href": "http://172.16.18.3:8182/database-node/02a26c34-48ba-4d98-9720-359b8ec349dc", "id_perms": {"enable": true, "uuid": {"uuid_mslong": 189833106117250456, "uuid_lslong": 10889762841213225436}, "created": "2017-03-07T17:55:37.899902", "description": null, "creator": null, "user_visible": true, "last_modified": "2017-03-07T17:55:37.899902", "permissions": {"owner": "neutron", "owner_access": 7, "other_access": 7, "group": "admin", "group_access": 7}}, "display_name": "node-61.test.domain.local", "name": "node-61.test.domain.local"}}`)
	}
	if  Path == `/database-node/e8a46bf3-d385-4bab-9353-ce404f6daddb` {
	 return []byte(`{"database-node": {"database_node_ip_address": "172.16.19.29", "fq_name": ["default-global-system-config", "node-60.test.domain.local"], "uuid": "e8a46bf3-d385-4bab-9353-ce404f6daddb", "parent_uuid": "e55fefa1-5e0e-4d83-b342-583fcaf95997", "parent_href": "http://172.16.18.3:8182/global-system-config/e55fefa1-5e0e-4d83-b342-583fcaf95997", "parent_type": "global-system-config", "perms2": {"owner": "55251e647b954127bf528a73ba096ca5", "owner_access": 7, "global_access": 0, "share": []}, "href": "http://172.16.18.3:8182/database-node/e8a46bf3-d385-4bab-9353-ce404f6daddb", "id_perms": {"enable": true, "uuid": {"uuid_mslong": 16763642407949781931, "uuid_lslong": 10616055522248207835}, "created": "2017-03-07T17:55:41.993064", "description": null, "creator": null, "user_visible": true, "last_modified": "2017-03-07T17:55:41.993064", "permissions": {"owner": "neutron", "owner_access": 7, "other_access": 7, "group": "admin", "group_access": 7}}, "display_name": "node-60.test.domain.local", "name": "node-60.test.domain.local"}}`)
	}
	if  Path == `/virtual-routers` {
	 return []byte(`{"virtual-routers": []}`)
	}
	return []byte("{}")
}


func TestContrailNodeManager (t *testing.T) {
	capi := FaceContrailAPI{}
	cnm := ContrailNodeManager(capi)

	allNodes := cnm.GetAll()
	if len(allNodes) != 6 {
		t.Errorf("Incorrect node count: %d != %d .", len(allNodes), 6)
	}

	byTypes := cnm.GetSplitByType()
	if len(byTypes) != 2 {
		t.Errorf("Incorrect type count: %d != %d .", len(byTypes), 2)
	}

	if len(byTypes[0]) != 3 {
		t.Errorf("Incorrect node count: %d != %d .", len(byTypes[0]), 3)
	}
	if len(byTypes[1]) != 3 {
		t.Errorf("Incorrect node count: %d != %d .", len(byTypes[1]), 3)
	}

	if allNodes[0].IP != "172.16.19.28" {
		t.Errorf("Incorrect node IP: %s != %s .", allNodes[0].IP, "172.16.19.28")
	}
	if allNodes[2].IP != "172.16.19.29" {
		t.Errorf("Incorrect node IP: %s != %s .", allNodes[2].IP, "172.16.19.29")
	}
	if allNodes[1].IP != "172.16.19.3" {
		t.Errorf("Incorrect node IP: %s != %s .", allNodes[1].IP, "172.16.19.3")
	}
}
