package client

import (
	"github.com/Diode222/Odin/app/conf"
	"github.com/Diode222/Odin/app/service"
	pb "github.com/Diode222/Odin/proto_gen"
	"sync"
)

var wordFreqListClient pb.WordFreqListServiceClient
var wordFreqListClientOnce sync.Once

func GetWordFreqListClient() pb.WordFreqListServiceClient {
	wordFreqListClientOnce.Do(func() {
		wordFreqListClient = service.GetServiceManager(conf.ETCD_ADDRESS).GetClient(conf.DATABASE_SERVICE_NAME, pb.NewWordFreqListClientWrapper).(pb.WordFreqListServiceClient)
	})
	return wordFreqListClient
}
