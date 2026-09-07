package interceptors

import (
	"google.golang.org/grpc"
)

func splitFullMethodName(fullMethod string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

type CallMeta struct {
	ReqOrNil any
	Typ      GRPCType
	Service  string
	Method   string
	IsClient bool
}

func NewClientCallMeta(fullMethod string, streamDesc *grpc.StreamDesc, reqOrNil any) CallMeta {
	_ = "STUB: not implemented"
	return *new(CallMeta)
}

func NewServerCallMeta(fullMethod string, streamInfo *grpc.StreamServerInfo, reqOrNil any) CallMeta {
	_ = "STUB: not implemented"
	return *new(CallMeta)
}

func (c CallMeta) FullMethod() string { _ = "STUB: not implemented"; return "" }

func clientStreamType(desc *grpc.StreamDesc) GRPCType {
	_ = "STUB: not implemented"
	return *new(GRPCType)
}

func serverStreamType(info *grpc.StreamServerInfo) GRPCType {
	_ = "STUB: not implemented"
	return *new(GRPCType)
}
