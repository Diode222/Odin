package client

import (
	"github.com/Diode222/Odin/app/conf"
	pb "github.com/Diode222/Odin/proto_gen"
	"github.com/Diode222/etcd_service_discovery/etcdservice"
	"sync"
)

var chatMessageListClient pb.ChatMessageListServiceClient
var chatMessageListClientOnce sync.Once

func GetChatMessageListClient() pb.ChatMessageListServiceClient {
	chatMessageListClientOnce.Do(func() {
		chatMessageListClient = etcdservice.NewServiceManager(conf.ETCD_ADDRESS).GetClient(conf.DATABASE_SERVICE_NAME, pb.NewChatMessageListWrapper).(pb.ChatMessageListServiceClient)
	})
	return chatMessageListClient
}
