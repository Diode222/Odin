package proto_gen

import "google.golang.org/grpc"

func NewWordFreqListClientWrapper(cc *grpc.ClientConn) interface{} {
	return NewWordFreqListServiceClient(cc)
}

func NewChatMessageListWrapper(cc *grpc.ClientConn) interface{} {
	return NewChatMessageListServiceClient(cc)
}
