package service

import (
	"context"
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
func (s *FriendService) SendFriendRequest(ctx context.Context, req *pb.SendFriendRequestRequest) (*pb.SendFriendRequestResponse, error) {
	// 根据请求中的 UniqueId 获取用户 ID（调用用户服务）
	sid, err := s.UserClient.GetIdByUnique(ctx, &v1.GetIdByUniqueRequest{
		UniqueId: req.SelfUniqueId,
	})
	if err != nil {
		return nil, RequestSendError(err)
	}
	tid, err := s.UserClient.GetIdByUnique(ctx, &v1.GetIdByUniqueRequest{
		UniqueId: req.SelfUniqueId,
	})
	if err != nil {
		return nil, RequestSendError(err)
	}
	sendFriend, err := s.UC.SendFriend(ctx, uint(sid.GetUserId()), uint(tid.GetUserId()))
	if err != nil {
		return nil, RequestSendError(err)
	}

	return &pb.SendFriendRequestResponse{Msg: sendFriend}, nil
}

// AcceptFriendRequest 根据 SelfUniqueId 和对方 UniqueId，同意好友请求，更新数据库中的好友关系
func (s *FriendService) AcceptFriendRequest(ctx context.Context, req *pb.AcceptFriendRequestRequest) (*pb.AcceptFriendRequestResponse, error) {

	return &pb.AcceptFriendRequestResponse{}, nil
}

// RejectFriendRequest 根据 SelfUniqueId 和对方 UniqueId，拒绝好友请求，删除或标记对应的请求记录
func (s *FriendService) RejectFriendRequest(ctx context.Context, req *pb.RejectFriendRequestRequest) (*pb.RejectFriendRequestResponse, error) {

	return &pb.RejectFriendRequestResponse{}, nil
}

// GetFriendList 根据用户的 UniqueId 获取其好友列表
func (s *FriendService) GetFriendList(ctx context.Context, req *pb.GetFriendListRequest) (*pb.GetFriendListResponse, error) {

	return &pb.GetFriendListResponse{}, nil
}

// DeleteFriend 根据 SelfUniqueId 和要删除的好友 UniqueId，从数据库中删除好友关系
func (s *FriendService) DeleteFriend(ctx context.Context, req *pb.DeleteFriendRequest) (*pb.DeleteFriendResponse, error) {

	return &pb.DeleteFriendResponse{}, nil
}

// GetFriendProfile 根据好友的 UniqueId 获取其好友资料信息
func (s *FriendService) GetFriendProfile(ctx context.Context, req *pb.GetFriendProfileRequest) (*pb.GetFriendProfileReply, error) {

	return &pb.GetFriendProfileReply{}, nil
}
