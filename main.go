package main

import (
	"github.com/docker/go-plugins-helpers/network"

	"github.com/omega/omega-netplugins/drivers"
	"github.com/opencontainers/runc/libcontainer/user"
	"github.com/Sirupsen/logrus"
	"github.com/docker/libkv"
	"github.com/docker/libkv/store"
	"github.com/samalba/dockerclient"
	"github.com/vishvananda/netlink"
)

func main(){

	store , err := libkv.NewStore(store.CONSUL, []string{"localhost:8500"}  , nil)
	if err!=nil {
		logrus.Fatal("init store client error:%s" , err)
	}


	dockerclient , err := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	if err!=nil {
		logrus.Fatal("init docker client error:%s", err)
	}


	mainDev , err := netlink.LinkByName("eth0")
	if err!=nil {
		logrus.Fatal("no such net device %s" , err.Error())
	}

	vlanDev , ok := link.(*netlink.Vlan)
	if !ok {
		logrus.Fatalf("%s is a vlan dev" , "eth0")
	}

	opts := make(map[string]interface{})

	h := network.NewHandler(drivers.NewDriver("vlan" , store ,  dockerclient, mainDev, opts))
	g, err := user.CurrentGroup()
	if err!=nil {
		logrus.Fatalf("can't get current group , err:%s" ,err)
	}
	h.ServeUnix("vlan",g.Gid)
}