// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             v6.30.0--rc1
// source: friend/v1/friend.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationFriendAcceptFriendRequest = "/api.friend.v1.Friend/AcceptFriendRequest"
const OperationFriendDeleteFriend = "/api.friend.v1.Friend/DeleteFriend"
const OperationFriendGetFriendList = "/api.friend.v1.Friend/GetFriendList"
const OperationFriendGetFriendProfile = "/api.friend.v1.Friend/GetFriendProfile"
const OperationFriendRejectFriendRequest = "/api.friend.v1.Friend/RejectFriendRequest"
const OperationFriendSendFriendRequest = "/api.friend.v1.Friend/SendFriendRequest"

type FriendHTTPServer interface {
	AcceptFriendRequest(context.Context, *AcceptFriendRequestRequest) (*AcceptFriendRequestResponse, error)
	DeleteFriend(context.Context, *DeleteFriendRequest) (*DeleteFriendResponse, error)
	GetFriendList(context.Context, *GetFriendListRequest) (*GetFriendListResponse, error)
	GetFriendProfile(context.Context, *GetFriendProfileRequest) (*GetFriendProfileReply, error)
	RejectFriendRequest(context.Context, *RejectFriendRequestRequest) (*RejectFriendRequestResponse, error)
	SendFriendRequest(context.Context, *SendFriendRequestRequest) (*SendFriendRequestResponse, error)
}

func RegisterFriendHTTPServer(s *http.Server, srv FriendHTTPServer) {
	r := s.Route("/")
	r.POST("/friend/send", _Friend_SendFriendRequest0_HTTP_Handler(srv))
	r.POST("/friend/accept", _Friend_AcceptFriendRequest0_HTTP_Handler(srv))
	r.POST("/friend/reject", _Friend_RejectFriendRequest0_HTTP_Handler(srv))
	r.GET("/friend/list/{unique_id}", _Friend_GetFriendList0_HTTP_Handler(srv))
	r.POST("/friend/delete", _Friend_DeleteFriend0_HTTP_Handler(srv))
	r.GET("/friend/profile/{unique_id}", _Friend_GetFriendProfile0_HTTP_Handler(srv))
}

func _Friend_SendFriendRequest0_HTTP_Handler(srv FriendHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendFriendRequestRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFriendSendFriendRequest)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendFriendRequest(ctx, req.(*SendFriendRequestRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendFriendRequestResponse)
		return ctx.Result(200, reply)
	}
}

func _Friend_AcceptFriendRequest0_HTTP_Handler(srv FriendHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AcceptFriendRequestRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFriendAcceptFriendRequest)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AcceptFriendRequest(ctx, req.(*AcceptFriendRequestRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AcceptFriendRequestResponse)
		return ctx.Result(200, reply)
	}
}

func _Friend_RejectFriendRequest0_HTTP_Handler(srv FriendHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RejectFriendRequestRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFriendRejectFriendRequest)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RejectFriendRequest(ctx, req.(*RejectFriendRequestRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RejectFriendRequestResponse)
		return ctx.Result(200, reply)
	}
}

func _Friend_GetFriendList0_HTTP_Handler(srv FriendHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetFriendListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFriendGetFriendList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFriendList(ctx, req.(*GetFriendListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetFriendListResponse)
		return ctx.Result(200, reply)
	}
}

func _Friend_DeleteFriend0_HTTP_Handler(srv FriendHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteFriendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFriendDeleteFriend)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteFriend(ctx, req.(*DeleteFriendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteFriendResponse)
		return ctx.Result(200, reply)
	}
}

func _Friend_GetFriendProfile0_HTTP_Handler(srv FriendHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetFriendProfileRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationFriendGetFriendProfile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetFriendProfile(ctx, req.(*GetFriendProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetFriendProfileReply)
		return ctx.Result(200, reply)
	}
}

type FriendHTTPClient interface {
	AcceptFriendRequest(ctx context.Context, req *AcceptFriendRequestRequest, opts ...http.CallOption) (rsp *AcceptFriendRequestResponse, err error)
	DeleteFriend(ctx context.Context, req *DeleteFriendRequest, opts ...http.CallOption) (rsp *DeleteFriendResponse, err error)
	GetFriendList(ctx context.Context, req *GetFriendListRequest, opts ...http.CallOption) (rsp *GetFriendListResponse, err error)
	GetFriendProfile(ctx context.Context, req *GetFriendProfileRequest, opts ...http.CallOption) (rsp *GetFriendProfileReply, err error)
	RejectFriendRequest(ctx context.Context, req *RejectFriendRequestRequest, opts ...http.CallOption) (rsp *RejectFriendRequestResponse, err error)
	SendFriendRequest(ctx context.Context, req *SendFriendRequestRequest, opts ...http.CallOption) (rsp *SendFriendRequestResponse, err error)
}

type FriendHTTPClientImpl struct {
	cc *http.Client
}

func NewFriendHTTPClient(client *http.Client) FriendHTTPClient {
	return &FriendHTTPClientImpl{client}
}

func (c *FriendHTTPClientImpl) AcceptFriendRequest(ctx context.Context, in *AcceptFriendRequestRequest, opts ...http.CallOption) (*AcceptFriendRequestResponse, error) {
	var out AcceptFriendRequestResponse
	pattern := "/friend/accept"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFriendAcceptFriendRequest))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FriendHTTPClientImpl) DeleteFriend(ctx context.Context, in *DeleteFriendRequest, opts ...http.CallOption) (*DeleteFriendResponse, error) {
	var out DeleteFriendResponse
	pattern := "/friend/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFriendDeleteFriend))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FriendHTTPClientImpl) GetFriendList(ctx context.Context, in *GetFriendListRequest, opts ...http.CallOption) (*GetFriendListResponse, error) {
	var out GetFriendListResponse
	pattern := "/friend/list/{unique_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationFriendGetFriendList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FriendHTTPClientImpl) GetFriendProfile(ctx context.Context, in *GetFriendProfileRequest, opts ...http.CallOption) (*GetFriendProfileReply, error) {
	var out GetFriendProfileReply
	pattern := "/friend/profile/{unique_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationFriendGetFriendProfile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FriendHTTPClientImpl) RejectFriendRequest(ctx context.Context, in *RejectFriendRequestRequest, opts ...http.CallOption) (*RejectFriendRequestResponse, error) {
	var out RejectFriendRequestResponse
	pattern := "/friend/reject"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFriendRejectFriendRequest))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *FriendHTTPClientImpl) SendFriendRequest(ctx context.Context, in *SendFriendRequestRequest, opts ...http.CallOption) (*SendFriendRequestResponse, error) {
	var out SendFriendRequestResponse
	pattern := "/friend/send"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationFriendSendFriendRequest))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
