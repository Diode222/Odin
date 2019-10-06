package client

import (
	"github.com/Diode222/Odin/app/conf"
	pb "github.com/Diode222/Odin/proto_gen"
	"github.com/Diode222/etcd_service_discovery/etcdservice"
	"sync"
)

var wordFreqClient pb.WordFreqListServiceClient
var wordFreqClientOnce sync.Once

func GetWordFreqClient() pb.WordFreqListServiceClient {
	wordFreqClientOnce.Do(func() {
		wordFreqClient = etcdservice.NewServiceManager(conf.ETCD_ADDRESS).GetClient(conf.DATABASE_SERVICE_NAME, pb.NewWordFreqWrapper).(pb.WordFreqListServiceClient)
	})
	return wordFreqClient
}
