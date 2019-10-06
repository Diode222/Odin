package client

import (
	"github.com/Diode222/Odin/app/conf"
	pb "github.com/Diode222/Odin/proto_gen"
	"github.com/Diode222/etcd_service_discovery/etcdservice"
	"sync"
)

var wordSplitClient pb.WordSplitServiceClient
var wordSplitClientOnce sync.Once

func GetWordSplitClient() pb.WordSplitServiceClient {
	wordSplitClientOnce.Do(func() {
		wordSplitClient = etcdservice.NewServiceManager(conf.ETCD_ADDRESS).GetClient(conf.ALGORITHM_SERVICE_NAME, pb.NewWordSpliteWrapper).(pb.WordSplitServiceClient)
	})
	return wordSplitClient
}
