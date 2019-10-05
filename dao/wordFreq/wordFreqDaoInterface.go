package wordFreq

import (
	"context"
	pb "github.com/Diode222/Odin/proto_gen"
)

type WordFreqDaoInterface interface {
	GetFreqList(Context context.Context, pos pb.PartOfSpeech) (*pb.WordFreqList, error)
}
