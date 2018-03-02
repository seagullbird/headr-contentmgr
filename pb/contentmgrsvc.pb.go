// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contentmgrsvc.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	contentmgrsvc.proto

It has these top-level messages:
	CreateNewPostRequest
	CreateNewPostReply
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The new post request contains zero parameters.
type CreateNewPostRequest struct {
	Title    string   `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Summary  string   `protobuf:"bytes,2,opt,name=summary" json:"summary,omitempty"`
	Content  string   `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
	Tags     []string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
	Author   string   `protobuf:"bytes,5,opt,name=author" json:"author,omitempty"`
	Sitename string   `protobuf:"bytes,6,opt,name=sitename" json:"sitename,omitempty"`
	Date     string   `protobuf:"bytes,7,opt,name=date" json:"date,omitempty"`
}

func (m *CreateNewPostRequest) Reset()                    { *m = CreateNewPostRequest{} }
func (m *CreateNewPostRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateNewPostRequest) ProtoMessage()               {}
func (*CreateNewPostRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateNewPostRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreateNewPostRequest) GetSummary() string {
	if m != nil {
		return m.Summary
	}
	return ""
}

func (m *CreateNewPostRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CreateNewPostRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *CreateNewPostRequest) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *CreateNewPostRequest) GetSitename() string {
	if m != nil {
		return m.Sitename
	}
	return ""
}

func (m *CreateNewPostRequest) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

// The new post response contains the result of the creation.
type CreateNewPostReply struct {
	Id  string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *CreateNewPostReply) Reset()                    { *m = CreateNewPostReply{} }
func (m *CreateNewPostReply) String() string            { return proto.CompactTextString(m) }
func (*CreateNewPostReply) ProtoMessage()               {}
func (*CreateNewPostReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateNewPostReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateNewPostReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateNewPostRequest)(nil), "pb.CreateNewPostRequest")
	proto.RegisterType((*CreateNewPostReply)(nil), "pb.CreateNewPostReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Contentmgr service

type ContentmgrClient interface {
	// Create a new post.
	NewPost(ctx context.Context, in *CreateNewPostRequest, opts ...grpc.CallOption) (*CreateNewPostReply, error)
}

type contentmgrClient struct {
	cc *grpc.ClientConn
}

func NewContentmgrClient(cc *grpc.ClientConn) ContentmgrClient {
	return &contentmgrClient{cc}
}

func (c *contentmgrClient) NewPost(ctx context.Context, in *CreateNewPostRequest, opts ...grpc.CallOption) (*CreateNewPostReply, error) {
	out := new(CreateNewPostReply)
	err := grpc.Invoke(ctx, "/pb.Contentmgr/NewPost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Contentmgr service

type ContentmgrServer interface {
	// Create a new post.
	NewPost(context.Context, *CreateNewPostRequest) (*CreateNewPostReply, error)
}

func RegisterContentmgrServer(s *grpc.Server, srv ContentmgrServer) {
	s.RegisterService(&_Contentmgr_serviceDesc, srv)
}

func _Contentmgr_NewPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentmgrServer).NewPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Contentmgr/NewPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentmgrServer).NewPost(ctx, req.(*CreateNewPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Contentmgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Contentmgr",
	HandlerType: (*ContentmgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewPost",
			Handler:    _Contentmgr_NewPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contentmgrsvc.proto",
}

func init() { proto.RegisterFile("contentmgrsvc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xbd, 0x4e, 0xc3, 0x40,
	0x10, 0x84, 0xb1, 0x9d, 0xd8, 0x64, 0x0b, 0x84, 0x96, 0x28, 0x5a, 0xa5, 0x8a, 0x5c, 0xa5, 0x72,
	0x01, 0x12, 0x1d, 0x55, 0x4a, 0x24, 0x84, 0xfc, 0x06, 0xe7, 0x78, 0x15, 0x2c, 0xf9, 0x7c, 0xc7,
	0xdd, 0x1a, 0xe4, 0x77, 0xe3, 0xe1, 0x90, 0x2f, 0x47, 0x0a, 0x94, 0x6e, 0xbe, 0x99, 0xfb, 0x19,
	0x0d, 0x3c, 0x1c, 0xcd, 0x20, 0x3c, 0x88, 0x3e, 0x39, 0xff, 0x75, 0xac, 0xac, 0x33, 0x62, 0x30,
	0xb5, 0x4d, 0xf9, 0x93, 0xc0, 0xfa, 0xe0, 0x58, 0x09, 0xbf, 0xf1, 0xf7, 0xbb, 0xf1, 0x52, 0xf3,
	0xe7, 0xc8, 0x5e, 0x70, 0x0d, 0x4b, 0xe9, 0xa4, 0x67, 0x4a, 0x76, 0xc9, 0x7e, 0x55, 0x9f, 0x01,
	0x09, 0x0a, 0x3f, 0x6a, 0xad, 0xdc, 0x44, 0x69, 0xf0, 0xff, 0x70, 0x4e, 0xe2, 0x1f, 0x94, 0x9d,
	0x93, 0x88, 0x88, 0xb0, 0x10, 0x75, 0xf2, 0xb4, 0xd8, 0x65, 0xfb, 0x55, 0x1d, 0x34, 0x6e, 0x20,
	0x57, 0xa3, 0x7c, 0x18, 0x47, 0xcb, 0x70, 0x38, 0x12, 0x6e, 0xe1, 0xd6, 0x77, 0xc2, 0x83, 0xd2,
	0x4c, 0x79, 0x48, 0x2e, 0x3c, 0xbf, 0xd3, 0x2a, 0x61, 0x2a, 0x82, 0x1f, 0x74, 0xf9, 0x0c, 0xf8,
	0xaf, 0xbd, 0xed, 0x27, 0xbc, 0x83, 0xb4, 0x6b, 0x63, 0xf1, 0xb4, 0x6b, 0xf1, 0x1e, 0x32, 0x76,
	0x2e, 0x36, 0x9e, 0xe5, 0xe3, 0x2b, 0xc0, 0xe1, 0xb2, 0x08, 0xbe, 0x40, 0x11, 0xef, 0x23, 0x55,
	0xb6, 0xa9, 0xae, 0x0d, 0xb2, 0xdd, 0x5c, 0x49, 0x6c, 0x3f, 0x95, 0x37, 0x4d, 0x1e, 0xe6, 0x7c,
	0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x00, 0xe2, 0xfe, 0x33, 0x65, 0x01, 0x00, 0x00,
}
