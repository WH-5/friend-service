// Package service  friend_err.go
// Author: 王辉
// Created: 2025-04-11 01:54
// 使用时直接传入业务层返回的error即可
// 直接传服务器报错原因会有安全风险
package service

import (
	"github.com/go-kratos/kratos/v2/errors"
)

type ErrStruct struct {
	code int
	name string
}

var (
	// 参数错误
	FriendSelfIdEmptyError   = errors.New(451, "FRIEND_SELF_ID_EMPTY", "self unique id is empty")
	FriendTargetIdEmptyError = errors.New(451, "FRIEND_TARGET_ID_EMPTY", "target unique id is empty")
	FriendGreetingEmptyError = errors.New(451, "FRIEND_GREETING_EMPTY", "greeting content is empty")

	// 业务逻辑错误
	FriendRequestSendError   = ErrStruct{code: 452, name: "FRIEND_REQUEST_SEND_ERROR"}
	FriendRequestAcceptError = ErrStruct{code: 453, name: "FRIEND_REQUEST_ACCEPT_ERROR"}
	FriendRequestRejectError = ErrStruct{code: 454, name: "FRIEND_REQUEST_REJECT_ERROR"}
	FriendListFetchError     = ErrStruct{code: 455, name: "FRIEND_LIST_FETCH_ERROR"}
	FriendDeleteError        = ErrStruct{code: 456, name: "FRIEND_DELETE_ERROR"}
	FriendInternalError      = ErrStruct{code: 500, name: "FRIEND_INTERNAL_ERROR"}
)

func FriendError(e ErrStruct, err error) *errors.Error {
	return errors.New(e.code, e.name, err.Error())
}
func RequestSendError(err error) *errors.Error {
	e := FriendRequestSendError
	return errors.New(e.code, e.name, err.Error())
}
func RequestAcceptError(err error) *errors.Error {
	e := FriendRequestAcceptError
	return errors.New(e.code, e.name, err.Error())
}
func RequestRejectError(err error) *errors.Error {
	e := FriendRequestRejectError
	return errors.New(e.code, e.name, err.Error())
}
func ListFetchError(err error) *errors.Error {
	e := FriendListFetchError
	return errors.New(e.code, e.name, err.Error())
}
func DeleteError(err error) *errors.Error {
	e := FriendDeleteError
	return errors.New(e.code, e.name, err.Error())
}
func InternalError(err error) *errors.Error {
	e := FriendInternalError
	return errors.New(e.code, e.name, err.Error())
}
