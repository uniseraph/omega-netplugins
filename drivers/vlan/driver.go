package vlan

import (
	"github.com/docker/go-plugins-helpers/network"
	"github.com/docker/libkv/store"
	"github.com/samalba/dockerclient"
	"github.com/vishvananda/netlink"
)

func NewDriver(store store.Store , dockerClient *dockerclient.DockerClient , mainDev netlink.Link,opts  map[string]interface{}) network.Driver {

	return &Driver{
		store:store ,
		dockerClient:dockerClient,
		mainDev: mainDev,

	}
}

type Driver struct {
	store			store.Store
	dockerClient    	*dockerclient.DockerClient
	mainDev                 netlink.Link

}

func (d *Driver) GetCapabilities() (*network.CapabilitiesResponse, error) {
	return nil, nil
}
func (d *Driver) CreateNetwork(*network.CreateNetworkRequest) error {
	return nil
}
func (d *Driver) AllocateNetwork(*network.AllocateNetworkRequest) (*network.AllocateNetworkResponse, error) {
	return nil, nil
}
func (d *Driver) DeleteNetwork(*network.DeleteNetworkRequest) error {
	return nil
}
func (d *Driver) FreeNetwork(*network.FreeNetworkRequest) error {
	return nil
}
func (d *Driver) CreateEndpoint(*network.CreateEndpointRequest) (*network.CreateEndpointResponse, error) {
	return nil, nil
}
func (d *Driver) DeleteEndpoint(*network.DeleteEndpointRequest) error {
	return nil
}
func (d *Driver) EndpointInfo(*network.InfoRequest) (*network.InfoResponse, error) {
	return nil, nil
}
func (d *Driver) Join(*network.JoinRequest) (*network.JoinResponse, error) {
	return nil, nil
}
func (d *Driver) Leave(*network.LeaveRequest) error {
	return nil
}
func (d *Driver) DiscoverNew(*network.DiscoveryNotification) error {
	return nil
}
func (d *Driver) DiscoverDelete(*network.DiscoveryNotification) error {
	return nil
}
func (d *Driver) ProgramExternalConnectivity(*network.ProgramExternalConnectivityRequest) error {
	return nil
}
func (d *Driver) RevokeExternalConnectivity(*network.RevokeExternalConnectivityRequest) error {
	return nil
}
