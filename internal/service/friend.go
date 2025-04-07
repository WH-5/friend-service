package service

import (
	"context"
	"friend-service/internal/biz"

	pb "friend-service/api/friend/v1"
)

type FriendService struct {
	pb.UnimplementedFriendServer
	uc *biz.FriendUsecase
}

func NewFriendService() *FriendService {
	return &FriendService{}
}

func (s *FriendService) SendFriendRequest(ctx context.Context, req *pb.SendFriendRequestRequest) (*pb.SendFriendRequestResponse, error) {
	return &pb.SendFriendRequestResponse{}, nil
}
func (s *FriendService) AcceptFriendRequest(ctx context.Context, req *pb.AcceptFriendRequestRequest) (*pb.AcceptFriendRequestResponse, error) {
	return &pb.AcceptFriendRequestResponse{}, nil
}
func (s *FriendService) RejectFriendRequest(ctx context.Context, req *pb.RejectFriendRequestRequest) (*pb.RejectFriendRequestResponse, error) {
	return &pb.RejectFriendRequestResponse{}, nil
}
func (s *FriendService) GetFriendList(ctx context.Context, req *pb.GetFriendListRequest) (*pb.GetFriendListResponse, error) {
	return &pb.GetFriendListResponse{}, nil
}
func (s *FriendService) DeleteFriend(ctx context.Context, req *pb.DeleteFriendRequest) (*pb.DeleteFriendResponse, error) {
	return &pb.DeleteFriendResponse{}, nil
}
func (s *FriendService) GetFriendProfile(ctx context.Context, req *pb.GetFriendProfileRequest) (*pb.GetFriendProfileReply, error) {
	return &pb.GetFriendProfileReply{}, nil
}
