package wordSplit

import (
	"context"
	"github.com/Diode222/Odin/app/client"
	pb "github.com/Diode222/Odin/proto_gen"
)

type WordSplitDao struct {
	client pb.WordSplitServiceClient
}

func NewWordSplitDao() *WordSplitDao {
	return &WordSplitDao{
		client: client.GetWordSplitClient(),
	}
}

func (d *WordSplitDao) GetWordSplittedMessageList(context context.Context, chatMessageList *pb.ChatMessageList) (*pb.ChatMessageList, error) {
	return d.client.GetWordSplittedMessageList(context, chatMessageList)
}
