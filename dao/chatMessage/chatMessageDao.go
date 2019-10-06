package chatMessage

import (
	"context"
	"github.com/Diode222/Odin/app/client"
	pb "github.com/Diode222/Odin/proto_gen"
)

type ChatMessageDao struct {
	client pb.ChatMessageListServiceClient
}

func NewChatMessageDao() *ChatMessageDao {
	return &ChatMessageDao{
		client: client.GetChatMessageClient(),
	}
}

func (d *ChatMessageDao) PutChatMessageList(context context.Context, chatMessageList *pb.ChatMessageList) (*pb.ChatMessageListServiceStatus, error) {
	return d.client.PutChatMessageList(context, chatMessageList)
}
