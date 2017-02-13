package drivers

import (
	"github.com/omega/omega-netplugins/drivers/vlan"
	"github.com/docker/go-plugins-helpers/network"
	"github.com/docker/libkv/store"
	"github.com/samalba/dockerclient"
	"github.com/vishvananda/netlink"
)



func NewDriver( _name string ,  store store.Store , dClient *dockerclient.DockerClient , mainDev netlink.Link ,  opts map[string]interface{}) network.Driver {
	return  vlan.NewDriver(store , dClient, mainDev , opts)
}
