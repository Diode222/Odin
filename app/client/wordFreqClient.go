package client

import (
	"github.com/Diode222/Odin/app/conf"
	"github.com/Diode222/Odin/app/service"
	pb "github.com/Diode222/Odin/proto_gen"
	"sync"
)

var wordFreqClient pb.WordFreqListServiceClient
var wordFreqClientOnce sync.Once

func GetWordFreqClient() pb.WordFreqListServiceClient {
	wordFreqClientOnce.Do(func() {
		wordFreqClient = service.GetServiceManager(conf.ETCD_ADDRESS).GetClient(conf.DATABASE_SERVICE_NAME, pb.NewWordFreqClientWrapper).(pb.WordFreqListServiceClient)
	})
	return wordFreqClient
}
