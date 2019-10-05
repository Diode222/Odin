package service

import (
	"github.com/Diode222/etcd_service_discovery/etcdservice"
	"sync"
)

var serviceManager *etcdservice.ServiceManager
var serviceManagerOnce sync.Once

func GetServiceManager(etcdAddr string) *etcdservice.ServiceManager {
	serviceManagerOnce.Do(func() {
		serviceManager = etcdservice.NewServiceManager(etcdAddr)
	})
	return serviceManager
}
