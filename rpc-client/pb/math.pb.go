// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: math.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumA float32 `protobuf:"fixed32,1,opt,name=numA,proto3" json:"numA,omitempty"`
	NumB float32 `protobuf:"fixed32,2,opt,name=numB,proto3" json:"numB,omitempty"`
}

func (x *MathRequest) Reset() {
	*x = MathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MathRequest) ProtoMessage() {}

func (x *MathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_math_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MathRequest.ProtoReflect.Descriptor instead.
func (*MathRequest) Descriptor() ([]byte, []int) {
	return file_math_proto_rawDescGZIP(), []int{0}
}

func (x *MathRequest) GetNumA() float32 {
	if x != nil {
		return x.NumA
	}
	return 0
}

func (x *MathRequest) GetNumB() float32 {
	if x != nil {
		return x.NumB
	}
	return 0
}

type MathResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MathResponse) Reset() {
	*x = MathResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_math_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MathResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MathResponse) ProtoMessage() {}

func (x *MathResponse) ProtoReflect() protoreflect.Message {
	mi := &file_math_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MathResponse.ProtoReflect.Descriptor instead.
func (*MathResponse) Descriptor() ([]byte, []int) {
	return file_math_proto_rawDescGZIP(), []int{1}
}

func (x *MathResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_math_proto protoreflect.FileDescriptor

var file_math_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x0b,
	0x4d, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x75, 0x6d, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x41, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x42, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6e,
	0x75, 0x6d, 0x42, 0x22, 0x26, 0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x5a, 0x0a, 0x0b, 0x4d,
	0x61, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x03, 0x41, 0x64,
	0x64, 0x12, 0x0c, 0x2e, 0x4d, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x4d, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x25, 0x0a, 0x04, 0x41, 0x64, 0x64, 0x32, 0x12, 0x0c, 0x2e, 0x4d, 0x61, 0x74, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x4d, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x65, 0x6c, 0x6d, 0x2e, 0x73,
	0x61, 0x2f, 0x61, 0x68, 0x6d, 0x61, 0x73, 0x69, 0x72, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2d, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_math_proto_rawDescOnce sync.Once
	file_math_proto_rawDescData = file_math_proto_rawDesc
)

func file_math_proto_rawDescGZIP() []byte {
	file_math_proto_rawDescOnce.Do(func() {
		file_math_proto_rawDescData = protoimpl.X.CompressGZIP(file_math_proto_rawDescData)
	})
	return file_math_proto_rawDescData
}

var file_math_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_math_proto_goTypes = []interface{}{
	(*MathRequest)(nil),  // 0: MathRequest
	(*MathResponse)(nil), // 1: MathResponse
}
var file_math_proto_depIdxs = []int32{
	0, // 0: MathService.Add:input_type -> MathRequest
	0, // 1: MathService.Add2:input_type -> MathRequest
	1, // 2: MathService.Add:output_type -> MathResponse
	1, // 3: MathService.Add2:output_type -> MathResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_math_proto_init() }
func file_math_proto_init() {
	if File_math_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_math_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MathRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_math_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MathResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_math_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_math_proto_goTypes,
		DependencyIndexes: file_math_proto_depIdxs,
		MessageInfos:      file_math_proto_msgTypes,
	}.Build()
	File_math_proto = out.File
	file_math_proto_rawDesc = nil
	file_math_proto_goTypes = nil
	file_math_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MathServiceClient is the client API for MathService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MathServiceClient interface {
	Add(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
	Add2(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error)
}

type mathServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMathServiceClient(cc grpc.ClientConnInterface) MathServiceClient {
	return &mathServiceClient{cc}
}

func (c *mathServiceClient) Add(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	out := new(MathResponse)
	err := c.cc.Invoke(ctx, "/MathService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mathServiceClient) Add2(ctx context.Context, in *MathRequest, opts ...grpc.CallOption) (*MathResponse, error) {
	out := new(MathResponse)
	err := c.cc.Invoke(ctx, "/MathService/Add2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MathServiceServer is the server API for MathService service.
type MathServiceServer interface {
	Add(context.Context, *MathRequest) (*MathResponse, error)
	Add2(context.Context, *MathRequest) (*MathResponse, error)
}

// UnimplementedMathServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMathServiceServer struct {
}

func (*UnimplementedMathServiceServer) Add(context.Context, *MathRequest) (*MathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (*UnimplementedMathServiceServer) Add2(context.Context, *MathRequest) (*MathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add2 not implemented")
}

func RegisterMathServiceServer(s *grpc.Server, srv MathServiceServer) {
	s.RegisterService(&_MathService_serviceDesc, srv)
}

func _MathService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MathService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Add(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MathService_Add2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MathServiceServer).Add2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MathService/Add2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MathServiceServer).Add2(ctx, req.(*MathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MathService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MathService",
	HandlerType: (*MathServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _MathService_Add_Handler,
		},
		{
			MethodName: "Add2",
			Handler:    _MathService_Add2_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "math.proto",
}
