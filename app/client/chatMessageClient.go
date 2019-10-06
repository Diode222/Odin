package client

import (
	"github.com/Diode222/Odin/app/conf"
	pb "github.com/Diode222/Odin/proto_gen"
	"github.com/Diode222/etcd_service_discovery/etcdservice"
	"sync"
)

var chatMessageClient pb.ChatMessageListServiceClient
var chatMessageClientOnce sync.Once

func GetChatMessageClient() pb.ChatMessageListServiceClient {
	chatMessageClientOnce.Do(func() {
		// TODO 修复服务发现库，目前在同一个serviceManager对象下面GetClient会报错，暂时每次都new一个新的serviceManager
		chatMessageClient = etcdservice.NewServiceManager(conf.ETCD_ADDRESS).GetClient(conf.DATABASE_SERVICE_NAME, pb.NewChatMessageWrapper).(pb.ChatMessageListServiceClient)
	})
	return chatMessageClient
}
