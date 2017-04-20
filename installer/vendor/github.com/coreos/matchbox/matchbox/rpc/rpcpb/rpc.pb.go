// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

/*
Package rpcpb is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
*/
package rpcpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import serverpb "github.com/coreos/matchbox/matchbox/server/serverpb"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Groups service

type GroupsClient interface {
	// Create a Group.
	GroupPut(ctx context.Context, in *serverpb.GroupPutRequest, opts ...grpc.CallOption) (*serverpb.GroupPutResponse, error)
	// Get a machine Group by id.
	GroupGet(ctx context.Context, in *serverpb.GroupGetRequest, opts ...grpc.CallOption) (*serverpb.GroupGetResponse, error)
	// Delete a machine Group by id.
	GroupDelete(ctx context.Context, in *serverpb.GroupDeleteRequest, opts ...grpc.CallOption) (*serverpb.GroupDeleteResponse, error)
	// List all machine Groups.
	GroupList(ctx context.Context, in *serverpb.GroupListRequest, opts ...grpc.CallOption) (*serverpb.GroupListResponse, error)
}

type groupsClient struct {
	cc *grpc.ClientConn
}

func NewGroupsClient(cc *grpc.ClientConn) GroupsClient {
	return &groupsClient{cc}
}

func (c *groupsClient) GroupPut(ctx context.Context, in *serverpb.GroupPutRequest, opts ...grpc.CallOption) (*serverpb.GroupPutResponse, error) {
	out := new(serverpb.GroupPutResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Groups/GroupPut", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsClient) GroupGet(ctx context.Context, in *serverpb.GroupGetRequest, opts ...grpc.CallOption) (*serverpb.GroupGetResponse, error) {
	out := new(serverpb.GroupGetResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Groups/GroupGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsClient) GroupDelete(ctx context.Context, in *serverpb.GroupDeleteRequest, opts ...grpc.CallOption) (*serverpb.GroupDeleteResponse, error) {
	out := new(serverpb.GroupDeleteResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Groups/GroupDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsClient) GroupList(ctx context.Context, in *serverpb.GroupListRequest, opts ...grpc.CallOption) (*serverpb.GroupListResponse, error) {
	out := new(serverpb.GroupListResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Groups/GroupList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Groups service

type GroupsServer interface {
	// Create a Group.
	GroupPut(context.Context, *serverpb.GroupPutRequest) (*serverpb.GroupPutResponse, error)
	// Get a machine Group by id.
	GroupGet(context.Context, *serverpb.GroupGetRequest) (*serverpb.GroupGetResponse, error)
	// Delete a machine Group by id.
	GroupDelete(context.Context, *serverpb.GroupDeleteRequest) (*serverpb.GroupDeleteResponse, error)
	// List all machine Groups.
	GroupList(context.Context, *serverpb.GroupListRequest) (*serverpb.GroupListResponse, error)
}

func RegisterGroupsServer(s *grpc.Server, srv GroupsServer) {
	s.RegisterService(&_Groups_serviceDesc, srv)
}

func _Groups_GroupPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(serverpb.GroupPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).GroupPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Groups/GroupPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).GroupPut(ctx, req.(*serverpb.GroupPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Groups_GroupGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(serverpb.GroupGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).GroupGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Groups/GroupGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).GroupGet(ctx, req.(*serverpb.GroupGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Groups_GroupDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(serverpb.GroupDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).GroupDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Groups/GroupDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).GroupDelete(ctx, req.(*serverpb.GroupDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Groups_GroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(serverpb.GroupListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).GroupList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Groups/GroupList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).GroupList(ctx, req.(*serverpb.GroupListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Groups_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.Groups",
	HandlerType: (*GroupsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GroupPut",
			Handler:    _Groups_GroupPut_Handler,
		},
		{
			MethodName: "GroupGet",
			Handler:    _Groups_GroupGet_Handler,
		},
		{
			MethodName: "GroupDelete",
			Handler:    _Groups_GroupDelete_Handler,
		},
		{
			MethodName: "GroupList",
			Handler:    _Groups_GroupList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// Client API for Profiles service

type ProfilesClient interface {
	// Create a Profile.
	ProfilePut(ctx context.Context, in *serverpb.ProfilePutRequest, opts ...grpc.CallOption) (*serverpb.ProfilePutResponse, error)
	// Get a Profile by id.
	ProfileGet(ctx context.Context, in *serverpb.ProfileGetRequest, opts ...grpc.CallOption) (*serverpb.ProfileGetResponse, error)
	// Delete a Profile by id.
	ProfileDelete(ctx context.Context, in *serverpb.ProfileDeleteRequest, opts ...grpc.CallOption) (*serverpb.ProfileDeleteResponse, error)
	// List all Profiles.
	ProfileList(ctx context.Context, in *serverpb.ProfileListRequest, opts ...grpc.CallOption) (*serverpb.ProfileListResponse, error)
}

type profilesClient struct {
	cc *grpc.ClientConn
}

func NewProfilesClient(cc *grpc.ClientConn) ProfilesClient {
	return &profilesClient{cc}
}

func (c *profilesClient) ProfilePut(ctx context.Context, in *serverpb.ProfilePutRequest, opts ...grpc.CallOption) (*serverpb.ProfilePutResponse, error) {
	out := new(serverpb.ProfilePutResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Profiles/ProfilePut", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesClient) ProfileGet(ctx context.Context, in *serverpb.ProfileGetRequest, opts ...grpc.CallOption) (*serverpb.ProfileGetResponse, error) {
	out := new(serverpb.ProfileGetResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Profiles/ProfileGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesClient) ProfileDelete(ctx context.Context, in *serverpb.ProfileDeleteRequest, opts ...grpc.CallOption) (*serverpb.ProfileDeleteResponse, error) {
	out := new(serverpb.ProfileDeleteResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Profiles/ProfileDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profilesClient) ProfileList(ctx context.Context, in *serverpb.ProfileListRequest, opts ...grpc.CallOption) (*serverpb.ProfileListResponse, error) {
	out := new(serverpb.ProfileListResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Profiles/ProfileList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Profiles service

type ProfilesServer interface {
	// Create a Profile.
	ProfilePut(context.Context, *serverpb.ProfilePutRequest) (*serverpb.ProfilePutResponse, error)
	// Get a Profile by id.
	ProfileGet(context.Context, *serverpb.ProfileGetRequest) (*serverpb.ProfileGetResponse, error)
	// Delete a Profile by id.
	ProfileDelete(context.Context, *serverpb.ProfileDeleteRequest) (*serverpb.ProfileDeleteResponse, error)
	// List all Profiles.
	ProfileList(context.Context, *serverpb.ProfileListRequest) (*serverpb.ProfileListResponse, error)
}

func RegisterProfilesServer(s *grpc.Server, srv ProfilesServer) {
	s.RegisterService(&_Profiles_serviceDesc, srv)
}

func _Profiles_ProfilePut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(serverpb.ProfilePutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).ProfilePut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Profiles/ProfilePut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).ProfilePut(ctx, req.(*serverpb.ProfilePutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profiles_ProfileGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(serverpb.ProfileGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).ProfileGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Profiles/ProfileGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).ProfileGet(ctx, req.(*serverpb.ProfileGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profiles_ProfileDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(serverpb.ProfileDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).ProfileDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Profiles/ProfileDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).ProfileDelete(ctx, req.(*serverpb.ProfileDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Profiles_ProfileList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(serverpb.ProfileListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfilesServer).ProfileList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Profiles/ProfileList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfilesServer).ProfileList(ctx, req.(*serverpb.ProfileListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Profiles_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.Profiles",
	HandlerType: (*ProfilesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProfilePut",
			Handler:    _Profiles_ProfilePut_Handler,
		},
		{
			MethodName: "ProfileGet",
			Handler:    _Profiles_ProfileGet_Handler,
		},
		{
			MethodName: "ProfileDelete",
			Handler:    _Profiles_ProfileDelete_Handler,
		},
		{
			MethodName: "ProfileList",
			Handler:    _Profiles_ProfileList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// Client API for Ignition service

type IgnitionClient interface {
	// Create or update an Container Linux Config template.
	IgnitionPut(ctx context.Context, in *serverpb.IgnitionPutRequest, opts ...grpc.CallOption) (*serverpb.IgnitionPutResponse, error)
	// Get a Container Linux Config template by name.
	IgnitionGet(ctx context.Context, in *serverpb.IgnitionGetRequest, opts ...grpc.CallOption) (*serverpb.IgnitionGetResponse, error)
	// Delete a Container Linux Config template by name.
	IgnitionDelete(ctx context.Context, in *serverpb.IgnitionDeleteRequest, opts ...grpc.CallOption) (*serverpb.IgnitionDeleteResponse, error)
}

type ignitionClient struct {
	cc *grpc.ClientConn
}

func NewIgnitionClient(cc *grpc.ClientConn) IgnitionClient {
	return &ignitionClient{cc}
}

func (c *ignitionClient) IgnitionPut(ctx context.Context, in *serverpb.IgnitionPutRequest, opts ...grpc.CallOption) (*serverpb.IgnitionPutResponse, error) {
	out := new(serverpb.IgnitionPutResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Ignition/IgnitionPut", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ignitionClient) IgnitionGet(ctx context.Context, in *serverpb.IgnitionGetRequest, opts ...grpc.CallOption) (*serverpb.IgnitionGetResponse, error) {
	out := new(serverpb.IgnitionGetResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Ignition/IgnitionGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ignitionClient) IgnitionDelete(ctx context.Context, in *serverpb.IgnitionDeleteRequest, opts ...grpc.CallOption) (*serverpb.IgnitionDeleteResponse, error) {
	out := new(serverpb.IgnitionDeleteResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Ignition/IgnitionDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ignition service

type IgnitionServer interface {
	// Create or update an Container Linux Config template.
	IgnitionPut(context.Context, *serverpb.IgnitionPutRequest) (*serverpb.IgnitionPutResponse, error)
	// Get a Container Linux Config template by name.
	IgnitionGet(context.Context, *serverpb.IgnitionGetRequest) (*serverpb.IgnitionGetResponse, error)
	// Delete a Container Linux Config template by name.
	IgnitionDelete(context.Context, *serverpb.IgnitionDeleteRequest) (*serverpb.IgnitionDeleteResponse, error)
}

func RegisterIgnitionServer(s *grpc.Server, srv IgnitionServer) {
	s.RegisterService(&_Ignition_serviceDesc, srv)
}

func _Ignition_IgnitionPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(serverpb.IgnitionPutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IgnitionServer).IgnitionPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Ignition/IgnitionPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IgnitionServer).IgnitionPut(ctx, req.(*serverpb.IgnitionPutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ignition_IgnitionGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(serverpb.IgnitionGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IgnitionServer).IgnitionGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Ignition/IgnitionGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IgnitionServer).IgnitionGet(ctx, req.(*serverpb.IgnitionGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ignition_IgnitionDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(serverpb.IgnitionDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IgnitionServer).IgnitionDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Ignition/IgnitionDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IgnitionServer).IgnitionDelete(ctx, req.(*serverpb.IgnitionDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ignition_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.Ignition",
	HandlerType: (*IgnitionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IgnitionPut",
			Handler:    _Ignition_IgnitionPut_Handler,
		},
		{
			MethodName: "IgnitionGet",
			Handler:    _Ignition_IgnitionGet_Handler,
		},
		{
			MethodName: "IgnitionDelete",
			Handler:    _Ignition_IgnitionDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

// Client API for Select service

type SelectClient interface {
	// SelectGroup returns the Group matching the given labels.
	SelectGroup(ctx context.Context, in *serverpb.SelectGroupRequest, opts ...grpc.CallOption) (*serverpb.SelectGroupResponse, error)
	// SelectProfile returns the Profile matching the given labels.
	SelectProfile(ctx context.Context, in *serverpb.SelectProfileRequest, opts ...grpc.CallOption) (*serverpb.SelectProfileResponse, error)
}

type selectClient struct {
	cc *grpc.ClientConn
}

func NewSelectClient(cc *grpc.ClientConn) SelectClient {
	return &selectClient{cc}
}

func (c *selectClient) SelectGroup(ctx context.Context, in *serverpb.SelectGroupRequest, opts ...grpc.CallOption) (*serverpb.SelectGroupResponse, error) {
	out := new(serverpb.SelectGroupResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Select/SelectGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *selectClient) SelectProfile(ctx context.Context, in *serverpb.SelectProfileRequest, opts ...grpc.CallOption) (*serverpb.SelectProfileResponse, error) {
	out := new(serverpb.SelectProfileResponse)
	err := grpc.Invoke(ctx, "/rpcpb.Select/SelectProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Select service

type SelectServer interface {
	// SelectGroup returns the Group matching the given labels.
	SelectGroup(context.Context, *serverpb.SelectGroupRequest) (*serverpb.SelectGroupResponse, error)
	// SelectProfile returns the Profile matching the given labels.
	SelectProfile(context.Context, *serverpb.SelectProfileRequest) (*serverpb.SelectProfileResponse, error)
}

func RegisterSelectServer(s *grpc.Server, srv SelectServer) {
	s.RegisterService(&_Select_serviceDesc, srv)
}

func _Select_SelectGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(serverpb.SelectGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SelectServer).SelectGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Select/SelectGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SelectServer).SelectGroup(ctx, req.(*serverpb.SelectGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Select_SelectProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(serverpb.SelectProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SelectServer).SelectProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpcpb.Select/SelectProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SelectServer).SelectProfile(ctx, req.(*serverpb.SelectProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Select_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.Select",
	HandlerType: (*SelectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SelectGroup",
			Handler:    _Select_SelectGroup_Handler,
		},
		{
			MethodName: "SelectProfile",
			Handler:    _Select_SelectProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x94, 0xd1, 0x4e, 0xc2, 0x30,
	0x14, 0x86, 0x85, 0x44, 0x02, 0xc7, 0xe8, 0xc5, 0xee, 0x44, 0x10, 0xe3, 0x03, 0x8c, 0x04, 0xdf,
	0x40, 0x8d, 0x84, 0x84, 0x0b, 0x82, 0xf1, 0x01, 0xd8, 0x72, 0x84, 0x25, 0x40, 0x6b, 0xdb, 0x19,
	0x1f, 0xc9, 0xf8, 0x18, 0xbe, 0x97, 0x89, 0x69, 0xd7, 0x96, 0xb3, 0xad, 0xbd, 0xe2, 0xe4, 0xff,
	0xd6, 0x8f, 0xc3, 0x5f, 0x32, 0x18, 0x08, 0x9e, 0xa7, 0x5c, 0x30, 0xc5, 0x92, 0x73, 0xc1, 0x73,
	0x9e, 0x0d, 0x1f, 0xb7, 0x85, 0xda, 0x95, 0x59, 0x9a, 0xb3, 0xc3, 0x34, 0x67, 0x02, 0x99, 0x9c,
	0x1e, 0x36, 0x2a, 0xdf, 0x65, 0xec, 0xeb, 0x34, 0x48, 0x14, 0x9f, 0x28, 0xec, 0x07, 0xcf, 0xa6,
	0x07, 0x94, 0x72, 0xb3, 0x45, 0x59, 0xa9, 0x66, 0xdf, 0x5d, 0xe8, 0xcd, 0x05, 0x2b, 0xb9, 0x4c,
	0x9e, 0xa0, 0x6f, 0xa6, 0x55, 0xa9, 0x92, 0xeb, 0xd4, 0x1d, 0x48, 0x5d, 0xb6, 0xc6, 0x8f, 0x12,
	0xa5, 0x1a, 0x0e, 0x43, 0x48, 0x72, 0x76, 0x94, 0x78, 0x7f, 0xe6, 0x25, 0x73, 0x6c, 0x4b, 0xe6,
	0x18, 0x95, 0x18, 0xe4, 0x25, 0x4b, 0xb8, 0x30, 0xe9, 0x33, 0xee, 0x51, 0x61, 0x32, 0x6a, 0x3c,
	0x5c, 0xc5, 0x4e, 0x35, 0x8e, 0x50, 0x6f, 0x7b, 0x81, 0x81, 0x01, 0xcb, 0x42, 0xaa, 0xa4, 0xf9,
	0xc5, 0x3a, 0x74, 0xa6, 0x9b, 0x20, 0x73, 0x9e, 0xd9, 0x6f, 0x17, 0xfa, 0x2b, 0xc1, 0xde, 0x8b,
	0x3d, 0xca, 0x64, 0x01, 0x60, 0x67, 0x5d, 0x17, 0x39, 0x79, 0x4a, 0x9d, 0x76, 0x14, 0x86, 0x7e,
	0xbf, 0x93, 0x4a, 0x97, 0xd6, 0x56, 0x91, 0xda, 0x46, 0x61, 0xe8, 0x55, 0x6b, 0xb8, 0xb4, 0xb9,
	0xad, 0xee, 0xb6, 0x75, 0xa0, 0x5e, 0xde, 0x24, 0xca, 0xe9, 0x65, 0x58, 0x64, 0x0a, 0x6c, 0xaf,
	0x40, 0x2b, 0x1c, 0x47, 0xa8, 0x2f, 0xf1, 0xaf, 0x03, 0xfd, 0xc5, 0xf6, 0x58, 0xa8, 0x82, 0x1d,
	0xb5, 0xda, 0xcd, 0xba, 0x45, 0xa2, 0x26, 0x71, 0x40, 0x5d, 0xa3, 0x74, 0x51, 0x07, 0x74, 0x91,
	0x01, 0x1b, 0x69, 0x72, 0x1c, 0xa1, 0xde, 0xf6, 0x06, 0x57, 0x0e, 0xd8, 0x2e, 0x27, 0xed, 0x23,
	0xf5, 0x32, 0xef, 0xe2, 0x0f, 0xf8, 0xdf, 0xff, 0xd3, 0x81, 0xde, 0x2b, 0xee, 0x31, 0x57, 0x7a,
	0xdf, 0x6a, 0x32, 0x7f, 0x36, 0xba, 0x2f, 0x89, 0x03, 0xfb, 0xd6, 0x28, 0xbd, 0xfa, 0x0a, 0xd8,
	0xde, 0xe9, 0xd5, 0xd7, 0x40, 0xe0, 0xea, 0x1b, 0xdc, 0x39, 0xb3, 0x9e, 0x79, 0x47, 0x3c, 0xfc,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x05, 0x60, 0x8e, 0x1b, 0x7b, 0x04, 0x00, 0x00,
}