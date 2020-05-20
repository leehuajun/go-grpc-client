// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.0
// source: proto/Student.proto

package service

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Class int32

const (
	Class_C1901 Class = 0
	Class_C1902 Class = 1
	Class_C1903 Class = 2
	Class_C1910 Class = 9
)

// Enum value maps for Class.
var (
	Class_name = map[int32]string{
		0: "C1901",
		1: "C1902",
		2: "C1903",
		9: "C1910",
	}
	Class_value = map[string]int32{
		"C1901": 0,
		"C1902": 1,
		"C1903": 2,
		"C1910": 9,
	}
)

func (x Class) Enum() *Class {
	p := new(Class)
	*p = x
	return p
}

func (x Class) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Class) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_Student_proto_enumTypes[0].Descriptor()
}

func (Class) Type() protoreflect.EnumType {
	return &file_proto_Student_proto_enumTypes[0]
}

func (x Class) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Class.Descriptor instead.
func (Class) EnumDescriptor() ([]byte, []int) {
	return file_proto_Student_proto_rawDescGZIP(), []int{0}
}

type StudentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Class Class `protobuf:"varint,1,opt,name=class,proto3,enum=service.Class" json:"class,omitempty"`
}

func (x *StudentRequest) Reset() {
	*x = StudentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Student_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentRequest) ProtoMessage() {}

func (x *StudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Student_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentRequest.ProtoReflect.Descriptor instead.
func (*StudentRequest) Descriptor() ([]byte, []int) {
	return file_proto_Student_proto_rawDescGZIP(), []int{0}
}

func (x *StudentRequest) GetClass() Class {
	if x != nil {
		return x.Class
	}
	return Class_C1901
}

type StudentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Students int32 `protobuf:"varint,1,opt,name=students,proto3" json:"students,omitempty"`
}

func (x *StudentResponse) Reset() {
	*x = StudentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_Student_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentResponse) ProtoMessage() {}

func (x *StudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_Student_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentResponse.ProtoReflect.Descriptor instead.
func (*StudentResponse) Descriptor() ([]byte, []int) {
	return file_proto_Student_proto_rawDescGZIP(), []int{1}
}

func (x *StudentResponse) GetStudents() int32 {
	if x != nil {
		return x.Students
	}
	return 0
}

var File_proto_Student_proto protoreflect.FileDescriptor

var file_proto_Student_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x36,
	0x0a, 0x0e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x24, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52,
	0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x22, 0x2d, 0x0a, 0x0f, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x73, 0x2a, 0x33, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x09,
	0x0a, 0x05, 0x43, 0x31, 0x39, 0x30, 0x31, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x31, 0x39,
	0x30, 0x32, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x31, 0x39, 0x30, 0x33, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x43, 0x31, 0x39, 0x31, 0x30, 0x10, 0x09, 0x32, 0x59, 0x0a, 0x0e, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x12, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x75,
	0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x21, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x75, 0x6e,
	0x6a, 0x65, 0x74, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x69, 0x5a, 0x0b, 0x61, 0x70, 0x69,
	0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_Student_proto_rawDescOnce sync.Once
	file_proto_Student_proto_rawDescData = file_proto_Student_proto_rawDesc
)

func file_proto_Student_proto_rawDescGZIP() []byte {
	file_proto_Student_proto_rawDescOnce.Do(func() {
		file_proto_Student_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_Student_proto_rawDescData)
	})
	return file_proto_Student_proto_rawDescData
}

var file_proto_Student_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_Student_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_Student_proto_goTypes = []interface{}{
	(Class)(0),              // 0: service.Class
	(*StudentRequest)(nil),  // 1: service.StudentRequest
	(*StudentResponse)(nil), // 2: service.StudentResponse
}
var file_proto_Student_proto_depIdxs = []int32{
	0, // 0: service.StudentRequest.class:type_name -> service.Class
	1, // 1: service.StudentService.GetStudentsByClass:input_type -> service.StudentRequest
	2, // 2: service.StudentService.GetStudentsByClass:output_type -> service.StudentResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_Student_proto_init() }
func file_proto_Student_proto_init() {
	if File_proto_Student_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_Student_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentRequest); i {
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
		file_proto_Student_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StudentResponse); i {
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
			RawDescriptor: file_proto_Student_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_Student_proto_goTypes,
		DependencyIndexes: file_proto_Student_proto_depIdxs,
		EnumInfos:         file_proto_Student_proto_enumTypes,
		MessageInfos:      file_proto_Student_proto_msgTypes,
	}.Build()
	File_proto_Student_proto = out.File
	file_proto_Student_proto_rawDesc = nil
	file_proto_Student_proto_goTypes = nil
	file_proto_Student_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StudentServiceClient interface {
	GetStudentsByClass(ctx context.Context, in *StudentRequest, opts ...grpc.CallOption) (*StudentResponse, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) GetStudentsByClass(ctx context.Context, in *StudentRequest, opts ...grpc.CallOption) (*StudentResponse, error) {
	out := new(StudentResponse)
	err := c.cc.Invoke(ctx, "/service.StudentService/GetStudentsByClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
type StudentServiceServer interface {
	GetStudentsByClass(context.Context, *StudentRequest) (*StudentResponse, error)
}

// UnimplementedStudentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (*UnimplementedStudentServiceServer) GetStudentsByClass(context.Context, *StudentRequest) (*StudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudentsByClass not implemented")
}

func RegisterStudentServiceServer(s *grpc.Server, srv StudentServiceServer) {
	s.RegisterService(&_StudentService_serviceDesc, srv)
}

func _StudentService_GetStudentsByClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).GetStudentsByClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.StudentService/GetStudentsByClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).GetStudentsByClass(ctx, req.(*StudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StudentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStudentsByClass",
			Handler:    _StudentService_GetStudentsByClass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/Student.proto",
}