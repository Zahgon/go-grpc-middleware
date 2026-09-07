package testvalidatev1

import (
	sync "sync"

	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SendRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *SendRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*SendRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SendRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *SendRequest) GetMessage() string { _ = "STUB: not implemented"; return "" }

type SendResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *SendResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*SendResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SendResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SendResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

type SendStreamRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendStreamRequest) Reset() { _ = "STUB: not implemented"; return }

func (x *SendStreamRequest) String() string { _ = "STUB: not implemented"; return "" }

func (*SendStreamRequest) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SendStreamRequest) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SendStreamRequest) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

func (x *SendStreamRequest) GetMessage() string { _ = "STUB: not implemented"; return "" }

type SendStreamResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendStreamResponse) Reset() { _ = "STUB: not implemented"; return }

func (x *SendStreamResponse) String() string { _ = "STUB: not implemented"; return "" }

func (*SendStreamResponse) ProtoMessage() { _ = "STUB: not implemented"; return }

func (x *SendStreamResponse) ProtoReflect() protoreflect.Message {
	_ = "STUB: not implemented"
	return *new(protoreflect.Message)
}

func (*SendStreamResponse) Descriptor() ([]byte, []int) { _ = "STUB: not implemented"; return nil, nil }

var File_testing_testvalidate_v1_test_validate_proto protoreflect.FileDescriptor

const file_testing_testvalidate_v1_test_validate_proto_rawDesc = "" +
	"\n" +
	"+testing/testvalidate/v1/test_validate.proto\x12\x17testing.testvalidate.v1\x1a\x1bbuf/validate/validate.proto\"0\n" +
	"\vSendRequest\x12!\n" +
	"\amessage\x18\x01 \x01(\tB\a\xbaH\x04r\x02`\x01R\amessage\"\x0e\n" +
	"\fSendResponse\"6\n" +
	"\x11SendStreamRequest\x12!\n" +
	"\amessage\x18\x01 \x01(\tB\a\xbaH\x04r\x02`\x01R\amessage\"\x14\n" +
	"\x12SendStreamResponse2\xd7\x01\n" +
	"\x13TestValidateService\x12U\n" +
	"\x04Send\x12$.testing.testvalidate.v1.SendRequest\x1a%.testing.testvalidate.v1.SendResponse\"\x00\x12i\n" +
	"\n" +
	"SendStream\x12*.testing.testvalidate.v1.SendStreamRequest\x1a+.testing.testvalidate.v1.SendStreamResponse\"\x000\x01BUZSgithub.com/grpc-ecosystem/go-grpc-middleware/testing/testvalidate/v1;testvalidatev1b\x06proto3"

var (
	file_testing_testvalidate_v1_test_validate_proto_rawDescOnce sync.Once
	file_testing_testvalidate_v1_test_validate_proto_rawDescData []byte
)

func file_testing_testvalidate_v1_test_validate_proto_rawDescGZIP() []byte {
	_ = "STUB: not implemented"
	return nil
}

var file_testing_testvalidate_v1_test_validate_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_testing_testvalidate_v1_test_validate_proto_goTypes = []any{
	(*SendRequest)(nil),
	(*SendResponse)(nil),
	(*SendStreamRequest)(nil),
	(*SendStreamResponse)(nil),
}
var file_testing_testvalidate_v1_test_validate_proto_depIdxs = []int32{
	0,
	2,
	1,
	3,
	2,
	0,
	0,
	0,
	0,
}

func init()                                                  { file_testing_testvalidate_v1_test_validate_proto_init() }
func file_testing_testvalidate_v1_test_validate_proto_init() { _ = "STUB: not implemented"; return }
