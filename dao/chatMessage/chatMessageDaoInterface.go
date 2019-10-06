package chatMessage

import (
	"context"
	pb "github.com/Diode222/Odin/proto_gen"
)

type ChatMessageDaoInterface interface {
	PutChatMessageList(context context.Context, chatMessageList *pb.ChatMessageList) (*pb.ChatMessageListServiceStatus, error)
}
