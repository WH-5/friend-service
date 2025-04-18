package service

import (
	"context"
	"errors"
	pb "github.com/WH-5/friend-service/api/friend/v1"
	v1 "github.com/WH-5/friend-service/api/user/v1"
	"github.com/WH-5/friend-service/internal/biz"
	"github.com/WH-5/friend-service/internal/conf"
	"github.com/WH-5/friend-service/internal/pkg"
	"github.com/go-kratos/kratos/v2/log"
)

type FriendService struct {
	pb.UnimplementedFriendServer
	UC         *biz.FriendUsecase
	UserClient v1.UserClient
}

func NewFriendService(c *conf.Server, usecase *biz.FriendUsecase) *FriendService {

	uc := pkg.UserClient(c.Registry.GetConsul())
	if uc == nil {
		log.Fatal("user client is nil — check consul config")
	}
	return &FriendService{
		UC:         usecase,
		UserClient: uc,
	}
}

// SendFriendRequest 根据 SelfUniqueId 和对方 UniqueId，发送好友请求
// 1. 从 context 中获取当前用户 ID（即请求发送者）
// 2. 调用用户服务将目标方 UniqueId 转换为用户 ID
// 3. 调用 UseCase 层逻辑进行好友请求的业务处理
// 4. 返回处理结果消息
func (s *FriendService) SendFriendRequest(ctx context.Context, req *pb.SendFriendRequestRequest) (*pb.SendFriendRequestResponse, error) {
	// 获取 user_id
	uidValue := ctx.Value("user_id")
	sid, ok := uidValue.(float64)
	if !ok {
		return nil, RequestSendError(errors.New("invalid or missing user_id in context"))
	}

	// 调用用户服务获取目标 user_id
	tid, err := s.UserClient.GetIdByUnique(ctx, &v1.GetIdByUniqueRequest{
		UniqueId: req.GetTargetUniqueId(),
	})
	if err != nil {
		return nil, RequestSendError(err)
	}

	// 调用 UseCase 方法
	err = s.UC.SendFriend(ctx, uint(sid), uint(tid.GetUserId()))
	if err != nil {
		return nil, RequestSendError(err)
	}

	// 返回结果
	return &pb.SendFriendRequestResponse{Msg: "send success"}, nil
}

// AcceptFriendRequest 根据 SelfUniqueId 和对方 UniqueId，同意好友请求
// 1. 获取当前用户 ID（接受者）
// 2. 获取请求发起者 ID（通过 UniqueId 查询）
// 3. 调用 UseCase 执行业务逻辑（写入好友关系，更新请求状态）
// 4. 返回处理结果
func (s *FriendService) AcceptFriendRequest(ctx context.Context, req *pb.AcceptFriendRequestRequest) (*pb.AcceptFriendRequestResponse, error) {
	// 获取 user_id
	uidValue := ctx.Value("user_id")
	sid, ok := uidValue.(float64)
	if !ok {
		return nil, RequestAcceptError(errors.New("invalid or missing user_id in context"))
	}

	// 调用用户服务获取目标 user_id（如果有）
	tid, err := s.UserClient.GetIdByUnique(ctx, &v1.GetIdByUniqueRequest{
		UniqueId: req.GetOtherUniqueId(),
	})
	if err != nil {
		return nil, RequestAcceptError(err)
	}

	// 调用 UseCase 方法
	err = s.UC.AcceptFriend(ctx, uint(sid), uint(tid.GetUserId()))
	if err != nil {
		return nil, RequestAcceptError(err)
	}

	// 返回结果
	return &pb.AcceptFriendRequestResponse{Msg: ""}, nil
}

// RejectFriendRequest 根据 SelfUniqueId 和对方 UniqueId，拒绝好友请求
// 1. 获取当前用户 ID（拒绝者）
// 2. 获取请求发起者 ID（通过 UniqueId 查询）
// 3. 调用 UseCase 拒绝好友请求（删除请求记录或标记为 rejected）
// 4. 返回处理结果
func (s *FriendService) RejectFriendRequest(ctx context.Context, req *pb.RejectFriendRequestRequest) (*pb.RejectFriendRequestResponse, error) {
	// 获取 user_id
	uidValue := ctx.Value("user_id")
	sid, ok := uidValue.(float64)
	if !ok {
		return nil, RequestRejectError(errors.New("invalid or missing user_id in context"))
	}

	// 调用用户服务获取目标 user_id（如果有）
	tid, err := s.UserClient.GetIdByUnique(ctx, &v1.GetIdByUniqueRequest{
		UniqueId: req.GetOtherUniqueId(),
	})
	if err != nil {
		return nil, RequestRejectError(err)
	}

	// 调用 UseCase 方法
	err = s.UC.RejectFriend(ctx, uint(sid), uint(tid.GetUserId()))
	if err != nil {
		return nil, RequestRejectError(err)
	}

	// 返回结果
	return &pb.RejectFriendRequestResponse{Msg: ""}, nil
}

// GetFriendList 获取当前用户的好友列表
// 1. 从 context 获取当前用户 ID
// 2. 调用 UseCase 层获取好友列表信息
// 3. 返回好友列表结构体
func (s *FriendService) GetFriendList(ctx context.Context, req *pb.GetFriendListRequest) (*pb.GetFriendListResponse, error) {
	// 获取 user_id
	uidValue := ctx.Value("user_id")
	sid, ok := uidValue.(float64)
	if !ok {
		return nil, ListFetchError(errors.New("invalid or missing user_id in context"))
	}

	// 调用 UseCase 方法
	res, count, err := s.UC.ListFriends(ctx, uint(sid))
	if err != nil {
		return nil, ListFetchError(err)
	}
	friends := make([]*pb.FriendInfo, 0)
	//这里会产生大量请求，后续必须要优化
	for i := 0; i < count; i++ {
		many, err := s.UserClient.GetUniqueByIdMany(ctx, &v1.GetUniqueByIdManyRequest{
			UserId: uint64(res[i].FriendId),
		})
		if err != nil {
			return nil, err
		}
		friends = append(friends, &pb.FriendInfo{
			UniqueId: many.GetUniqueId(),
			Nickname: res[i].Nickname,
		})
	}

	// 返回结果
	return &pb.GetFriendListResponse{
		Friends: friends,
		Count:   int32(count),
	}, nil
}

//
//// DeleteFriend 删除指定好友关系
//// 1. 获取当前用户 ID
//// 2. 获取目标好友 ID（通过 UniqueId 查询）
//// 3. 调用 UseCase 删除好友关系（双向删除）
//// 4. 返回处理结果
//func (s *FriendService) DeleteFriend(ctx context.Context, req *pb.DeleteFriendRequest) (*pb.DeleteFriendResponse, error) {
//	// 获取 user_id
//	uidValue := ctx.Value("user_id")
//	sid, ok := uidValue.(float64)
//	if !ok {
//		return nil, DeleteError(errors.New("invalid or missing user_id in context"))
//	}
//
//	// 调用用户服务获取目标 user_id（如果有）
//	tid, err := s.UserClient.GetIdByUnique(ctx, &v1.GetIdByUniqueRequest{
//		UniqueId: req.GetTargetUniqueId(),
//	})
//	if err != nil {
//		return nil, DeleteError(err)
//	}
//
//	// 调用 UseCase 方法
//	msg, err := s.UC.DeleteFriend(ctx, uint(sid), uint(tid.GetUserId()))
//	if err != nil {
//		return nil, DeleteError(err)
//	}
//
//	// 返回结果
//	return &pb.DeleteFriendResponse{Msg: msg}, nil
//}
//
//// GetFriendProfile 获取好友详细信息
//// 1. 获取当前用户 ID
//// 2. 获取目标好友 ID（通过 UniqueId 查询）
//// 3. 调用 UseCase 获取双方好友关系对应的 Profile 信息
//// 4. 返回好友 Profile
//func (s *FriendService) GetFriendProfile(ctx context.Context, req *pb.GetFriendProfileRequest) (*pb.GetFriendProfileReply, error) {
//	// 获取 user_id
//	uidValue := ctx.Value("user_id")
//	sid, ok := uidValue.(float64)
//	if !ok {
//		return nil, RequestSendError(errors.New("invalid or missing user_id in context"))
//	}
//
//	// 调用用户服务获取目标 user_id（如果有）
//	tid, err := s.UserClient.GetIdByUnique(ctx, &v1.GetIdByUniqueRequest{
//		UniqueId: req.GetFriendId(),
//	})
//	if err != nil {
//		return nil, RequestSendError(err)
//	}
//
//	// 调用 UseCase 方法
//	profile, err := s.UC.GetProfile(ctx, uint(sid), uint(tid.GetUserId()))
//	if err != nil {
//		return nil, RequestSendError(err)
//	}
//
//	// 返回结果
//	return profile, nil
//}
