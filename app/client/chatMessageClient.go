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
		chatMessageClient = etcdservice.NewServiceManager(conf.ETCD_ADDRESS).GetClient(conf.DATABASE_SERVICE_NAME, pb.NewChatMessageWrapper).(pb.ChatMessageListServiceClient)
	})
	return chatMessageClient
}
