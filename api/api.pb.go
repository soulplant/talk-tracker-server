// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	ServerState
	User
	Talk
	FetchAllRequest
	FetchAllResponse
	AddUserRequest
	AddUserResponse
	AddTalkRequest
	AddTalkResponse
	ReorderRequest
	ReorderResponse
	GetUsersRequest
	GetUsersResponse
	UpdateUserRequest
	UpdateUserResponse
	RemoveUserRequest
	RemoveUserResponse
	CompleteTalkRequest
	CompleteTalkResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

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

type ServerState struct {
	User   []*User `protobuf:"bytes,1,rep,name=user" json:"user,omitempty"`
	Talk   []*Talk `protobuf:"bytes,2,rep,name=talk" json:"talk,omitempty"`
	NextId int64   `protobuf:"varint,3,opt,name=nextId" json:"nextId,omitempty"`
}

func (m *ServerState) Reset()                    { *m = ServerState{} }
func (m *ServerState) String() string            { return proto.CompactTextString(m) }
func (*ServerState) ProtoMessage()               {}
func (*ServerState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServerState) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *ServerState) GetTalk() []*Talk {
	if m != nil {
		return m.Talk
	}
	return nil
}

func (m *ServerState) GetNextId() int64 {
	if m != nil {
		return m.NextId
	}
	return 0
}

type User struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	NextTalk string `protobuf:"bytes,3,opt,name=nextTalk" json:"nextTalk,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetNextTalk() string {
	if m != nil {
		return m.NextTalk
	}
	return ""
}

type Talk struct {
	Id        int64                       `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name      string                      `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	SpeakerId int64                       `protobuf:"varint,3,opt,name=speakerId" json:"speakerId,omitempty"`
	Done      bool                        `protobuf:"varint,4,opt,name=done" json:"done,omitempty"`
	Url       []string                    `protobuf:"bytes,5,rep,name=url" json:"url,omitempty"`
	Completed *google_protobuf1.Timestamp `protobuf:"bytes,6,opt,name=completed" json:"completed,omitempty"`
}

func (m *Talk) Reset()                    { *m = Talk{} }
func (m *Talk) String() string            { return proto.CompactTextString(m) }
func (*Talk) ProtoMessage()               {}
func (*Talk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Talk) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Talk) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Talk) GetSpeakerId() int64 {
	if m != nil {
		return m.SpeakerId
	}
	return 0
}

func (m *Talk) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Talk) GetUrl() []string {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *Talk) GetCompleted() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Completed
	}
	return nil
}

type FetchAllRequest struct {
}

func (m *FetchAllRequest) Reset()                    { *m = FetchAllRequest{} }
func (m *FetchAllRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchAllRequest) ProtoMessage()               {}
func (*FetchAllRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type FetchAllResponse struct {
	Version string  `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	User    []*User `protobuf:"bytes,2,rep,name=user" json:"user,omitempty"`
	Talk    []*Talk `protobuf:"bytes,3,rep,name=talk" json:"talk,omitempty"`
}

func (m *FetchAllResponse) Reset()                    { *m = FetchAllResponse{} }
func (m *FetchAllResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchAllResponse) ProtoMessage()               {}
func (*FetchAllResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FetchAllResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *FetchAllResponse) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *FetchAllResponse) GetTalk() []*Talk {
	if m != nil {
		return m.Talk
	}
	return nil
}

type AddUserRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AddUserRequest) Reset()                    { *m = AddUserRequest{} }
func (m *AddUserRequest) String() string            { return proto.CompactTextString(m) }
func (*AddUserRequest) ProtoMessage()               {}
func (*AddUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AddUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AddUserResponse struct {
	// The newly added user.
	UserId int64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
}

func (m *AddUserResponse) Reset()                    { *m = AddUserResponse{} }
func (m *AddUserResponse) String() string            { return proto.CompactTextString(m) }
func (*AddUserResponse) ProtoMessage()               {}
func (*AddUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AddUserResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type AddTalkRequest struct {
	UserId int64  `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *AddTalkRequest) Reset()                    { *m = AddTalkRequest{} }
func (m *AddTalkRequest) String() string            { return proto.CompactTextString(m) }
func (*AddTalkRequest) ProtoMessage()               {}
func (*AddTalkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *AddTalkRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AddTalkRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AddTalkResponse struct {
	TalkId int64 `protobuf:"varint,1,opt,name=talkId" json:"talkId,omitempty"`
}

func (m *AddTalkResponse) Reset()                    { *m = AddTalkResponse{} }
func (m *AddTalkResponse) String() string            { return proto.CompactTextString(m) }
func (*AddTalkResponse) ProtoMessage()               {}
func (*AddTalkResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AddTalkResponse) GetTalkId() int64 {
	if m != nil {
		return m.TalkId
	}
	return 0
}

type ReorderRequest struct {
	Version      int64 `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	MoveUserId   int64 `protobuf:"varint,2,opt,name=moveUserId" json:"moveUserId,omitempty"`
	AnchorUserId int64 `protobuf:"varint,3,opt,name=anchorUserId" json:"anchorUserId,omitempty"`
	Before       bool  `protobuf:"varint,4,opt,name=before" json:"before,omitempty"`
}

func (m *ReorderRequest) Reset()                    { *m = ReorderRequest{} }
func (m *ReorderRequest) String() string            { return proto.CompactTextString(m) }
func (*ReorderRequest) ProtoMessage()               {}
func (*ReorderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReorderRequest) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ReorderRequest) GetMoveUserId() int64 {
	if m != nil {
		return m.MoveUserId
	}
	return 0
}

func (m *ReorderRequest) GetAnchorUserId() int64 {
	if m != nil {
		return m.AnchorUserId
	}
	return 0
}

func (m *ReorderRequest) GetBefore() bool {
	if m != nil {
		return m.Before
	}
	return false
}

type ReorderResponse struct {
	Accepted bool   `protobuf:"varint,1,opt,name=accepted" json:"accepted,omitempty"`
	Version  string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *ReorderResponse) Reset()                    { *m = ReorderResponse{} }
func (m *ReorderResponse) String() string            { return proto.CompactTextString(m) }
func (*ReorderResponse) ProtoMessage()               {}
func (*ReorderResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ReorderResponse) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *ReorderResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type GetUsersRequest struct {
}

func (m *GetUsersRequest) Reset()                    { *m = GetUsersRequest{} }
func (m *GetUsersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUsersRequest) ProtoMessage()               {}
func (*GetUsersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type GetUsersResponse struct {
	Version int64   `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	User    []*User `protobuf:"bytes,2,rep,name=user" json:"user,omitempty"`
}

func (m *GetUsersResponse) Reset()                    { *m = GetUsersResponse{} }
func (m *GetUsersResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUsersResponse) ProtoMessage()               {}
func (*GetUsersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *GetUsersResponse) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *GetUsersResponse) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateUserRequest struct {
	UserId      int64  `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	HasName     bool   `protobuf:"varint,3,opt,name=hasName" json:"hasName,omitempty"`
	NextTalk    string `protobuf:"bytes,4,opt,name=nextTalk" json:"nextTalk,omitempty"`
	HasNextTalk bool   `protobuf:"varint,5,opt,name=hasNextTalk" json:"hasNextTalk,omitempty"`
}

func (m *UpdateUserRequest) Reset()                    { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()               {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *UpdateUserRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UpdateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateUserRequest) GetHasName() bool {
	if m != nil {
		return m.HasName
	}
	return false
}

func (m *UpdateUserRequest) GetNextTalk() string {
	if m != nil {
		return m.NextTalk
	}
	return ""
}

func (m *UpdateUserRequest) GetHasNextTalk() bool {
	if m != nil {
		return m.HasNextTalk
	}
	return false
}

type UpdateUserResponse struct {
}

func (m *UpdateUserResponse) Reset()                    { *m = UpdateUserResponse{} }
func (m *UpdateUserResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserResponse) ProtoMessage()               {}
func (*UpdateUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

type RemoveUserRequest struct {
	UserId int64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
}

func (m *RemoveUserRequest) Reset()                    { *m = RemoveUserRequest{} }
func (m *RemoveUserRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveUserRequest) ProtoMessage()               {}
func (*RemoveUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *RemoveUserRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type RemoveUserResponse struct {
}

func (m *RemoveUserResponse) Reset()                    { *m = RemoveUserResponse{} }
func (m *RemoveUserResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveUserResponse) ProtoMessage()               {}
func (*RemoveUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

type CompleteTalkRequest struct {
	UserId int64 `protobuf:"varint,1,opt,name=userId" json:"userId,omitempty"`
}

func (m *CompleteTalkRequest) Reset()                    { *m = CompleteTalkRequest{} }
func (m *CompleteTalkRequest) String() string            { return proto.CompactTextString(m) }
func (*CompleteTalkRequest) ProtoMessage()               {}
func (*CompleteTalkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *CompleteTalkRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CompleteTalkResponse struct {
}

func (m *CompleteTalkResponse) Reset()                    { *m = CompleteTalkResponse{} }
func (m *CompleteTalkResponse) String() string            { return proto.CompactTextString(m) }
func (*CompleteTalkResponse) ProtoMessage()               {}
func (*CompleteTalkResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func init() {
	proto.RegisterType((*ServerState)(nil), "api.ServerState")
	proto.RegisterType((*User)(nil), "api.User")
	proto.RegisterType((*Talk)(nil), "api.Talk")
	proto.RegisterType((*FetchAllRequest)(nil), "api.FetchAllRequest")
	proto.RegisterType((*FetchAllResponse)(nil), "api.FetchAllResponse")
	proto.RegisterType((*AddUserRequest)(nil), "api.AddUserRequest")
	proto.RegisterType((*AddUserResponse)(nil), "api.AddUserResponse")
	proto.RegisterType((*AddTalkRequest)(nil), "api.AddTalkRequest")
	proto.RegisterType((*AddTalkResponse)(nil), "api.AddTalkResponse")
	proto.RegisterType((*ReorderRequest)(nil), "api.ReorderRequest")
	proto.RegisterType((*ReorderResponse)(nil), "api.ReorderResponse")
	proto.RegisterType((*GetUsersRequest)(nil), "api.GetUsersRequest")
	proto.RegisterType((*GetUsersResponse)(nil), "api.GetUsersResponse")
	proto.RegisterType((*UpdateUserRequest)(nil), "api.UpdateUserRequest")
	proto.RegisterType((*UpdateUserResponse)(nil), "api.UpdateUserResponse")
	proto.RegisterType((*RemoveUserRequest)(nil), "api.RemoveUserRequest")
	proto.RegisterType((*RemoveUserResponse)(nil), "api.RemoveUserResponse")
	proto.RegisterType((*CompleteTalkRequest)(nil), "api.CompleteTalkRequest")
	proto.RegisterType((*CompleteTalkResponse)(nil), "api.CompleteTalkResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ApiService service

type ApiServiceClient interface {
	FetchAll(ctx context.Context, in *FetchAllRequest, opts ...grpc.CallOption) (*FetchAllResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error)
	// Change the position of one user in the list of upcoming talks.
	Reorder(ctx context.Context, in *ReorderRequest, opts ...grpc.CallOption) (*ReorderResponse, error)
	AddTalk(ctx context.Context, in *AddTalkRequest, opts ...grpc.CallOption) (*AddTalkResponse, error)
	CompleteTalk(ctx context.Context, in *CompleteTalkRequest, opts ...grpc.CallOption) (*CompleteTalkResponse, error)
}

type apiServiceClient struct {
	cc *grpc.ClientConn
}

func NewApiServiceClient(cc *grpc.ClientConn) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) FetchAll(ctx context.Context, in *FetchAllRequest, opts ...grpc.CallOption) (*FetchAllResponse, error) {
	out := new(FetchAllResponse)
	err := grpc.Invoke(ctx, "/api.ApiService/FetchAll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := grpc.Invoke(ctx, "/api.ApiService/GetUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	out := new(AddUserResponse)
	err := grpc.Invoke(ctx, "/api.ApiService/AddUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := grpc.Invoke(ctx, "/api.ApiService/UpdateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error) {
	out := new(RemoveUserResponse)
	err := grpc.Invoke(ctx, "/api.ApiService/RemoveUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) Reorder(ctx context.Context, in *ReorderRequest, opts ...grpc.CallOption) (*ReorderResponse, error) {
	out := new(ReorderResponse)
	err := grpc.Invoke(ctx, "/api.ApiService/Reorder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) AddTalk(ctx context.Context, in *AddTalkRequest, opts ...grpc.CallOption) (*AddTalkResponse, error) {
	out := new(AddTalkResponse)
	err := grpc.Invoke(ctx, "/api.ApiService/AddTalk", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) CompleteTalk(ctx context.Context, in *CompleteTalkRequest, opts ...grpc.CallOption) (*CompleteTalkResponse, error) {
	out := new(CompleteTalkResponse)
	err := grpc.Invoke(ctx, "/api.ApiService/CompleteTalk", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ApiService service

type ApiServiceServer interface {
	FetchAll(context.Context, *FetchAllRequest) (*FetchAllResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error)
	// Change the position of one user in the list of upcoming talks.
	Reorder(context.Context, *ReorderRequest) (*ReorderResponse, error)
	AddTalk(context.Context, *AddTalkRequest) (*AddTalkResponse, error)
	CompleteTalk(context.Context, *CompleteTalkRequest) (*CompleteTalkResponse, error)
}

func RegisterApiServiceServer(s *grpc.Server, srv ApiServiceServer) {
	s.RegisterService(&_ApiService_serviceDesc, srv)
}

func _ApiService_FetchAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).FetchAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/FetchAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).FetchAll(ctx, req.(*FetchAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetUsers(ctx, req.(*GetUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).RemoveUser(ctx, req.(*RemoveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_Reorder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReorderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).Reorder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/Reorder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).Reorder(ctx, req.(*ReorderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_AddTalk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTalkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).AddTalk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/AddTalk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).AddTalk(ctx, req.(*AddTalkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_CompleteTalk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteTalkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CompleteTalk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ApiService/CompleteTalk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CompleteTalk(ctx, req.(*CompleteTalkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApiService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchAll",
			Handler:    _ApiService_FetchAll_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _ApiService_GetUsers_Handler,
		},
		{
			MethodName: "AddUser",
			Handler:    _ApiService_AddUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _ApiService_UpdateUser_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _ApiService_RemoveUser_Handler,
		},
		{
			MethodName: "Reorder",
			Handler:    _ApiService_Reorder_Handler,
		},
		{
			MethodName: "AddTalk",
			Handler:    _ApiService_AddTalk_Handler,
		},
		{
			MethodName: "CompleteTalk",
			Handler:    _ApiService_CompleteTalk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 796 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x96, 0xed, 0x34, 0x8d, 0x27, 0x6d, 0x93, 0x6c, 0xd2, 0xe0, 0x5a, 0x05, 0xa2, 0x15, 0x87,
	0x50, 0xd4, 0x44, 0x84, 0x0b, 0xaa, 0xb8, 0x54, 0x48, 0xad, 0x22, 0x50, 0x0f, 0x6e, 0x2b, 0x21,
	0x71, 0xda, 0xc6, 0xdb, 0xd6, 0xd4, 0xf1, 0x1a, 0xdb, 0x89, 0x90, 0x10, 0x17, 0x0e, 0xbc, 0x00,
	0x07, 0x9e, 0x00, 0x89, 0xf7, 0xe1, 0x15, 0x78, 0x10, 0xb4, 0x3f, 0xfe, 0x6d, 0x80, 0x72, 0xf3,
	0xfc, 0x7d, 0x33, 0xfe, 0xe6, 0x1b, 0x1b, 0x4c, 0x12, 0x7a, 0xa3, 0x30, 0x62, 0x09, 0x43, 0x06,
	0x09, 0x3d, 0x7b, 0xf7, 0x8a, 0xb1, 0x2b, 0x9f, 0x8e, 0x49, 0xe8, 0x8d, 0x49, 0x10, 0xb0, 0x84,
	0x24, 0x1e, 0x0b, 0x62, 0x99, 0x62, 0x3f, 0x54, 0x51, 0x61, 0x5d, 0x2c, 0x2e, 0xc7, 0x67, 0xde,
	0x9c, 0xc6, 0x09, 0x99, 0x87, 0x32, 0x01, 0xcf, 0xa0, 0x79, 0x4a, 0xa3, 0x25, 0x8d, 0x4e, 0x13,
	0x92, 0x50, 0x74, 0x1f, 0x6a, 0x8b, 0x98, 0x46, 0x96, 0x36, 0x30, 0x86, 0xcd, 0x89, 0x39, 0xe2,
	0xcd, 0xce, 0x63, 0x1a, 0x39, 0xc2, 0xcd, 0xc3, 0x09, 0xf1, 0x6f, 0x2c, 0xbd, 0x10, 0x3e, 0x23,
	0xfe, 0x8d, 0x23, 0xdc, 0xa8, 0x0f, 0xf5, 0x80, 0x7e, 0x48, 0xa6, 0xae, 0x65, 0x0c, 0xb4, 0xa1,
	0xe1, 0x28, 0x0b, 0x1f, 0x41, 0x8d, 0x83, 0xa0, 0x2d, 0xd0, 0x3d, 0xd7, 0xd2, 0x44, 0x4c, 0xf7,
	0x5c, 0x84, 0xa0, 0x16, 0x90, 0x39, 0xb5, 0xf4, 0x81, 0x36, 0x34, 0x1d, 0xf1, 0x8c, 0x6c, 0x68,
	0xf0, 0x2a, 0x8e, 0x2a, 0x50, 0x4c, 0x27, 0xb3, 0xf1, 0x0f, 0x0d, 0x6a, 0xfc, 0xe1, 0x4e, 0x40,
	0xbb, 0x60, 0xc6, 0x21, 0x25, 0x37, 0x34, 0xca, 0xe6, 0xc9, 0x1d, 0xbc, 0xc2, 0x65, 0x01, 0xb5,
	0x6a, 0x03, 0x6d, 0xd8, 0x70, 0xc4, 0x33, 0x6a, 0x83, 0xb1, 0x88, 0x7c, 0x6b, 0x6d, 0x60, 0x0c,
	0x4d, 0x87, 0x3f, 0xa2, 0xe7, 0x60, 0xce, 0xd8, 0x3c, 0xf4, 0x69, 0x42, 0x5d, 0xab, 0x3e, 0xd0,
	0x86, 0xcd, 0x89, 0x3d, 0x92, 0x94, 0x8e, 0x52, 0x4a, 0x47, 0x19, 0xa5, 0x4e, 0x9e, 0x8c, 0x3b,
	0xd0, 0x3a, 0xa2, 0xc9, 0xec, 0xfa, 0xd0, 0xf7, 0x1d, 0xfa, 0x7e, 0x41, 0xe3, 0x04, 0xbf, 0x83,
	0x76, 0xee, 0x8a, 0x43, 0x16, 0xc4, 0x14, 0x59, 0xb0, 0xbe, 0xa4, 0x51, 0xec, 0xb1, 0x40, 0xbc,
	0x8d, 0xe9, 0xa4, 0x66, 0xb6, 0x09, 0xfd, 0xef, 0x9b, 0x30, 0x56, 0x6e, 0x02, 0x3f, 0x82, 0xad,
	0x43, 0xd7, 0x15, 0xf9, 0xb2, 0x7b, 0x46, 0x91, 0x96, 0x53, 0x84, 0x1f, 0x43, 0x2b, 0xcb, 0x52,
	0x03, 0xf5, 0xa1, 0xce, 0xf1, 0xa7, 0x29, 0xbb, 0xca, 0xc2, 0x2f, 0x04, 0xa0, 0xe8, 0xa0, 0x00,
	0xff, 0x90, 0xb9, 0x6a, 0x17, 0xaa, 0x91, 0xac, 0xce, 0x1b, 0xf1, 0x49, 0xf3, 0x72, 0x69, 0xe1,
	0x2f, 0x1a, 0x6c, 0x39, 0x94, 0x45, 0x6e, 0x3e, 0x7a, 0x85, 0x24, 0x23, 0x27, 0xe9, 0x01, 0xc0,
	0x9c, 0x2d, 0xe9, 0xb9, 0x9c, 0x43, 0x17, 0xc1, 0x82, 0x07, 0x61, 0xd8, 0x20, 0xc1, 0xec, 0x9a,
	0x45, 0x2a, 0x43, 0xca, 0xa0, 0xe4, 0xe3, 0x83, 0x5c, 0xd0, 0x4b, 0x16, 0xa5, 0x5a, 0x50, 0x16,
	0x3e, 0x86, 0x56, 0x36, 0x87, 0x9a, 0xd9, 0x86, 0x06, 0x99, 0xcd, 0x68, 0xc8, 0xd5, 0xa0, 0x89,
	0xe4, 0xcc, 0x2e, 0x0e, 0xa9, 0x97, 0x36, 0xc9, 0xa5, 0x70, 0x4c, 0x13, 0xde, 0x2d, 0x4e, 0xa5,
	0xf0, 0x0a, 0xda, 0xb9, 0x6b, 0xb5, 0x14, 0x8c, 0xbb, 0x4a, 0x01, 0x7f, 0xd3, 0xa0, 0x73, 0x1e,
	0xba, 0x24, 0xa1, 0xc5, 0x7d, 0xff, 0xc7, 0x7a, 0x78, 0xeb, 0x6b, 0x12, 0x9f, 0x70, 0xb7, 0x21,
	0x5e, 0x2b, 0x35, 0x4b, 0xd7, 0x58, 0x2b, 0x5f, 0x23, 0x1a, 0x40, 0x93, 0xa7, 0xa5, 0xe1, 0x35,
	0x51, 0x59, 0x74, 0xe1, 0x1e, 0xa0, 0xe2, 0x60, 0xf2, 0x45, 0xf1, 0x13, 0xe8, 0x38, 0x34, 0x5d,
	0xd2, 0x3f, 0xc6, 0xe5, 0x10, 0xc5, 0x64, 0x05, 0xb1, 0x0f, 0xdd, 0x97, 0xea, 0xd4, 0xee, 0x20,
	0x49, 0xdc, 0x87, 0x5e, 0x39, 0x5d, 0xc2, 0x4c, 0xbe, 0xaf, 0x01, 0x1c, 0x86, 0x1e, 0xff, 0x00,
	0x7a, 0x33, 0x8a, 0x4e, 0xa0, 0x91, 0x1e, 0x28, 0xea, 0x09, 0x96, 0x2b, 0x27, 0x6c, 0x6f, 0x57,
	0xbc, 0x6a, 0x9c, 0xed, 0xcf, 0x3f, 0x7f, 0x7d, 0xd5, 0x5b, 0x68, 0x73, 0xbc, 0x7c, 0x3a, 0xbe,
	0xe4, 0xd1, 0x7d, 0xe2, 0xfb, 0x68, 0x0a, 0x8d, 0x74, 0xcb, 0x0a, 0xaf, 0xa2, 0x03, 0x85, 0x57,
	0x95, 0x02, 0x6e, 0x0b, 0x3c, 0x40, 0x0d, 0x8e, 0x27, 0xce, 0x7d, 0x0a, 0xeb, 0xea, 0x52, 0x51,
	0x57, 0xd4, 0x94, 0xaf, 0xdb, 0xee, 0x95, 0x9d, 0x0a, 0xa7, 0x2b, 0x70, 0x36, 0x71, 0x86, 0x73,
	0xa0, 0xed, 0xa1, 0xb7, 0x00, 0xf9, 0x52, 0x50, 0x5f, 0xaa, 0xa9, 0x2a, 0x1f, 0xfb, 0xde, 0x2d,
	0xbf, 0xc2, 0xdc, 0x15, 0x98, 0xfd, 0x49, 0x27, 0xc5, 0x1c, 0x7f, 0x94, 0x2c, 0x7f, 0xe2, 0xe0,
	0x6f, 0x00, 0xf2, 0x75, 0x29, 0xf0, 0x5b, 0xcb, 0x56, 0xe0, 0x2b, 0xf6, 0xba, 0x23, 0xc0, 0xbb,
	0x7b, 0xb7, 0xc1, 0xd1, 0x6b, 0x58, 0x57, 0xe7, 0xa8, 0x18, 0x28, 0x7f, 0x24, 0x14, 0x03, 0x95,
	0x8b, 0xc5, 0x7d, 0x01, 0xd8, 0xc6, 0x4d, 0x0e, 0x18, 0xc9, 0x20, 0x9f, 0x53, 0xf2, 0x29, 0x64,
	0x9c, 0xf1, 0x59, 0x50, 0x52, 0xce, 0x67, 0x51, 0x2f, 0x65, 0x3e, 0xf9, 0xf7, 0x8a, 0x43, 0x11,
	0xd8, 0x28, 0x8a, 0x0b, 0x59, 0xa2, 0x74, 0x85, 0x3c, 0xed, 0x9d, 0x15, 0x91, 0x32, 0xab, 0xb8,
	0x93, 0x22, 0x8f, 0xd3, 0x5f, 0xc9, 0x81, 0xb6, 0x77, 0x51, 0x17, 0xff, 0x9a, 0x67, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xcf, 0xfc, 0x1a, 0xfb, 0xfc, 0x07, 0x00, 0x00,
}
