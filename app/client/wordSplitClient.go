package client

import (
	"github.com/Diode222/Odin/app/conf"
	"github.com/Diode222/Odin/app/service"
	pb "github.com/Diode222/Odin/proto_gen"
	"sync"
)

var wordSplitClient pb.WordSplitServiceClient
var wordSplitClientOnce sync.Once

func GetWordSplitClient() pb.WordSplitServiceClient {
	wordSplitClientOnce.Do(func() {
		wordSplitClient = service.GetServiceManager(conf.ETCD_ADDRESS).GetClient(conf.ALGORITHM_SERVICE_NAME, pb.NewWordSpliteWrapper).(pb.WordSplitServiceClient)
	})
	return wordSplitClient
}
