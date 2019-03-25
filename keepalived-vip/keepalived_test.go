package main

import "testing"

func TestNodeCapturedCorrectVIP(t *testing.T) {
	currentNodeIP := "10.0.1.1"
	expectedMasterVirtualIP := "172.16.4.1"
	tables := []struct {
		nodeIP    string
		neighbors []string
		vips      []string
	}{
		{currentNodeIP, []string{"10.0.1.2", "10.0.1.4", "10.0.1.5"}, []string{"172.16.4.4", expectedMasterVirtualIP}},
		{currentNodeIP, []string{"10.0.1.2", "10.0.1.4", "10.0.1.5"}, []string{expectedMasterVirtualIP, "172.16.4.4"}},
		{currentNodeIP, []string{"10.0.1.4", "10.0.1.2", "10.0.1.5"}, []string{"172.16.4.4", expectedMasterVirtualIP}},
	}

	for _, table := range tables {
		vipsInfo := GetVIPsForNode(table.nodeIP, table.neighbors, table.vips)
		if vipsInfo.Master != expectedMasterVirtualIP {
			t.Errorf("Node %v captured wrong master virtual IP (%v).", currentNodeIP, vipsInfo.Master)
		}
	}

}

// Node have to not capture vip if all vips already captured by other nodes
func TestNodeDontCapturedMaster(t *testing.T) {
	currentNodeIP := "10.0.1.6"
	tables := []struct {
		nodeIP    string
		neighbors []string
		vips      []string
	}{
		{currentNodeIP, []string{"10.0.1.2", "10.0.1.4", "10.0.1.5"}, []string{"172.16.4.4", "172.16.4.1"}},
		{currentNodeIP, []string{"10.0.1.2", "10.0.1.4", "10.0.1.5"}, []string{"172.16.4.1", "172.16.4.4"}},
		{currentNodeIP, []string{"10.0.1.4", "10.0.1.2", "10.0.1.5"}, []string{"172.16.4.4", "172.16.4.1"}},
	}

	for _, table := range tables {
		vipsInfo := GetVIPsForNode(table.nodeIP, table.neighbors, table.vips)
		if vipsInfo.Master != "" {
			t.Errorf("Node %v captured wrong master virtual IP (%v).", currentNodeIP, vipsInfo.Master)
		}
	}

}
