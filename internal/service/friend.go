package service

import (
	"context"
	pb "github.com/WH-5/friend-service/api/friend/v1"
	v1 "github.com/WH-5/friend-service/api/user/v1"
	"github.com/WH-5/friend-service/internal/biz"
	"github.com/WH-5/friend-service/internal/conf"
	"github.com/WH-5/friend-service/internal/pkg"
	"github.com/go-kratos/kratos/v2/log"
	"strconv"
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

// SendFriendRequest
func (s *FriendService) SendFriendRequest(ctx context.Context, req *pb.SendFriendRequestRequest) (*pb.SendFriendRequestResponse, error) {

	id, err := s.UserClient.GetIdByUnique(ctx, &v1.GetIdByUniqueRequest{
		UniqueId: req.SelfUniqueId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.SendFriendRequestResponse{Msg: strconv.Itoa(int(id.GetUserId()))}, nil
}

// AcceptFriendRequest
func (s *FriendService) AcceptFriendRequest(ctx context.Context, req *pb.AcceptFriendRequestRequest) (*pb.AcceptFriendRequestResponse, error) {
	return &pb.AcceptFriendRequestResponse{}, nil
}

// RejectFriendRequest
func (s *FriendService) RejectFriendRequest(ctx context.Context, req *pb.RejectFriendRequestRequest) (*pb.RejectFriendRequestResponse, error) {
	return &pb.RejectFriendRequestResponse{}, nil
}

// GetFriendList
func (s *FriendService) GetFriendList(ctx context.Context, req *pb.GetFriendListRequest) (*pb.GetFriendListResponse, error) {
	return &pb.GetFriendListResponse{}, nil
}

// DeleteFriend
func (s *FriendService) DeleteFriend(ctx context.Context, req *pb.DeleteFriendRequest) (*pb.DeleteFriendResponse, error) {
	return &pb.DeleteFriendResponse{}, nil
}

// GetFriendProfile
func (s *FriendService) GetFriendProfile(ctx context.Context, req *pb.GetFriendProfileRequest) (*pb.GetFriendProfileReply, error) {
	return &pb.GetFriendProfileReply{}, nil
}
