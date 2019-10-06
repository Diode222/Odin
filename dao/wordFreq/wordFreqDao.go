package wordFreq

import (
	"context"
	"github.com/Diode222/Odin/app/client"
	pb "github.com/Diode222/Odin/proto_gen"
)

type WordFreqDao struct {
	client pb.WordFreqListServiceClient
}

func NewWordFreqDao() *WordFreqDao {
	return &WordFreqDao{
		client: client.GetWordFreqClient(),
	}
}

func (d *WordFreqDao) GetWordFreqList(context context.Context, pos pb.PartOfSpeech) (*pb.WordFreqList, error) {
	return d.client.GetWordFreqList(context, &pos)
}
