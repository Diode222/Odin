package wordSplit

import (
	"context"
	pb "github.com/Diode222/Odin/proto_gen"
)

type WordSplitDaoInterface interface {
	GetWordSplittedMessageList(context context.Context, chatMessageList *pb.ChatMessageList) (*pb.ChatMessageList, error)
}
