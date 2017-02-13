package vlan

import (

	"fmt"
	"github.com/docker/libnetwork/netlabel"
	"github.com/omega/network-plugins/option"
	"github.com/docker/go-plugins-helpers/network"
	"github.com/coreos/etcd/store"
)

type Network struct {
	NetworkId string
	VlanId    int
	IPv4Data  []*network.IPAMData
}


func save(store store.Store , *Network) {
	
}

func parseNetworkOption(ipv4Data []*network.IPAMData, options map[string]interface{}) ( Network,error ) {
	if len(ipv4Data) == 0 {
		return Network{}, fmt.Errorf("At least one IPv4 Data is needed")
	}

	for _, ipv4 := range ipv4Data {
		if ipv4.Gateway == "" {
			return Network{}, fmt.Errorf("IPv4 Gateway cannot be nil, data %v", ipv4)
		} else if ipv4.Pool == "" {
			return Network{}, fmt.Errorf("IPv4 Subnet cannot be nil, data %v", ipv4)
		}
	}

	gopt := option.GenericOption(options[netlabel.GenericData].(map[string]interface{}))

	n := Network{
		VlanId:   gopt.Int("VlanId"),
		IPv4Data: ipv4Data,
	}
	if n.VlanId <= 0 || n.VlanId >= 4096 {
		return Network{}, fmt.Errorf(`opt "VlanId" invalid, must 0 < VlanId < 4096`)
	}
	return n, nil
}