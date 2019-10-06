package proto_gen

import "google.golang.org/grpc"

func NewWordFreqClientWrapper(cc *grpc.ClientConn) interface{} {
	return NewWordFreqListServiceClient(cc)
}

func NewChatMessageWrapper(cc *grpc.ClientConn) interface{} {
	return NewChatMessageListServiceClient(cc)
}

func NewWordSpliteWrapper(cc *grpc.ClientConn) interface{} {
	return NewWordSplitServiceClient(cc)
}
