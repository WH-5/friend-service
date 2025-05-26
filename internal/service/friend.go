package service

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	pb "github.com/WH-5/friend-service/api/friend/v1"
	v2 "github.com/WH-5/friend-service/api/push/v1"
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
	PushClient v2.PushClient
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

	// 原始数据
	data := map[string]interface{}{
		"type": 1,
	}

	// 转换为 JSON 字节
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// base64 编码
	encoded := base64.StdEncoding.EncodeToString(jsonBytes)
	//通知好友，如果失败不返回错误，打印日志
	_, err = s.PushClient.PushMsg(ctx, &v2.PushMsgRequest{
		ToUnique:   req.GetTargetUniqueId(),
		SelfUserId: uint64(sid),
		MsgType:    2, //2代表好友消息
		Payload:    []byte(encoded),
	})
	if err != nil {
		log.Error(err)
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

	// 原始数据
	data := map[string]interface{}{
		"type": 2,
	}

	// 转换为 JSON 字节
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// base64 编码
	encoded := base64.StdEncoding.EncodeToString(jsonBytes)
	//通知好友，如果失败不返回错误，打印日志
	_, err = s.PushClient.PushMsg(ctx, &v2.PushMsgRequest{
		ToUnique:   req.GetOtherUniqueId(),
		SelfUserId: uint64(sid),
		MsgType:    2, //2代表好友消息
		Payload:    []byte(encoded),
	})
	if err != nil {
		log.Error(err)
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
	// 原始数据
	data := map[string]interface{}{
		"type": 3,
	}

	// 转换为 JSON 字节
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// base64 编码
	encoded := base64.StdEncoding.EncodeToString(jsonBytes)
	//通知好友，如果失败不返回错误，打印日志
	_, err = s.PushClient.PushMsg(ctx, &v2.PushMsgRequest{
		ToUnique:   req.GetOtherUniqueId(),
		SelfUserId: uint64(sid),
		MsgType:    2, //2代表好友消息
		Payload:    []byte(encoded),
	})
	if err != nil {
		log.Error(err)
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

// DeleteFriend 删除指定好友关系
// 1. 获取当前用户 ID
// 2. 获取目标好友 ID（通过 UniqueId 查询）
// 3. 调用 UseCase 删除好友关系（双向删除）
// 4. 返回处理结果
func (s *FriendService) DeleteFriend(ctx context.Context, req *pb.DeleteFriendRequest) (*pb.DeleteFriendResponse, error) {
	// 获取 user_id
	uidValue := ctx.Value("user_id")
	sid, ok := uidValue.(float64)
	if !ok {
		return nil, DeleteError(errors.New("invalid or missing user_id in context"))
	}

	// 调用用户服务获取目标 user_id（如果有）
	tid, err := s.UserClient.GetIdByUnique(ctx, &v1.GetIdByUniqueRequest{
		UniqueId: req.GetTargetUniqueId(),
	})
	if err != nil {
		return nil, DeleteError(err)
	}

	// 调用 UseCase 方法
	err = s.UC.DeleteFriend(ctx, uint(sid), uint(tid.GetUserId()))
	if err != nil {
		return nil, DeleteError(err)
	}

	// 原始数据
	data := map[string]interface{}{
		"type": 1,
	}

	// 转换为 JSON 字节
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// base64 编码
	encoded := base64.StdEncoding.EncodeToString(jsonBytes)
	//通知好友，如果失败不返回错误，打印日志
	_, err = s.PushClient.PushMsg(ctx, &v2.PushMsgRequest{
		ToUnique:   req.GetTargetUniqueId(),
		SelfUserId: uint64(sid),
		MsgType:    2, //2代表好友消息
		Payload:    []byte(encoded),
	})
	if err != nil {
		log.Error(err)
	}

	// 返回结果
	return &pb.DeleteFriendResponse{
		Msg: "delete success",
	}, nil
}

// GetFriendProfile 获取好友详细信息
// 1. 获取当前用户 ID
// 2. 获取目标好友 ID（通过 UniqueId 查询）
// 3. 调用 UseCase 获取双方好友关系对应的 Profile 信息
// 4. 返回好友 Profile
func (s *FriendService) GetFriendProfile(ctx context.Context, req *pb.GetFriendProfileRequest) (*pb.GetFriendProfileReply, error) {
	// 获取 user_id
	uidValue := ctx.Value("user_id")
	sid, ok := uidValue.(float64)
	if !ok {
		return nil, RequestSendError(errors.New("invalid or missing user_id in context"))
	}

	// 调用用户服务获取目标 user_id（如果有）
	tid, err := s.UserClient.GetIdByUnique(ctx, &v1.GetIdByUniqueRequest{
		UniqueId: req.GetUniqueId(),
	})
	if err != nil {
		return nil, RequestSendError(err)
	}
	//判断是否好友
	if uint64(sid) != tid.GetUserId() {
		//如果不相同，相同的话是查询自己，允许
		err := s.UC.IsFriend(ctx, uint(sid), uint(tid.GetUserId()))
		if err != nil {
			return nil, InternalError(err)
		}
	}
	getProfile, err := s.UserClient.GetProfile(ctx, &v1.GetProfileRequest{UniqueId: req.GetUniqueId()})
	if err != nil {
		return nil, InternalError(err)
	}

	publicKey, err := s.UserClient.GetPublicKey(ctx, &v1.GetPublicKeyRequest{
		UserId: tid.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	// 返回结果
	return &pb.GetFriendProfileReply{
		UniqueId: req.GetUniqueId(),
		UserProfile: &pb.UserProfile{
			Nickname: getProfile.Profile.Nickname,
			Bio:      getProfile.Profile.Bio,
			Gender:   getProfile.Profile.Gender,
			Birthday: getProfile.Profile.Birthday,
			Location: getProfile.Profile.Location,
			Other:    getProfile.Profile.Other,
		},
		PublicKey: publicKey.GetPublicKey(),
	}, nil
}

// FriendMark 给好友改备注
func (s *FriendService) FriendMark(ctx context.Context, req *pb.FriendMarkRequest) (*pb.FriendMarkReply, error) {
	// 获取 user_id
	uidValue := ctx.Value("user_id")
	sid, ok := uidValue.(float64)
	if !ok {
		return nil, RequestSendError(errors.New("invalid or missing user_id in context"))
	}

	// 调用用户服务获取目标 user_id（如果有）
	tid, err := s.UserClient.GetIdByUnique(ctx, &v1.GetIdByUniqueRequest{
		UniqueId: req.GetUniqueId(),
	})
	if err != nil {
		return nil, RequestSendError(err)
	}

	if uint64(sid) == tid.GetUserId() {
		return nil, InternalError(errors.New("不能给自己备注"))
	}

	//放到业务逻辑层处理
	////判断是否好友
	//err = s.UC.IsFriend(ctx, uint(sid), uint(tid.GetUserId()))
	//if err != nil {
	//	return nil, InternalError(err)
	//}

	//修改备注
	err = s.UC.UpdateMark(ctx, uint(sid), uint(tid.GetUserId()), req.GetMark())
	if err != nil {
		return nil, InternalError(err)
	}
	return &pb.FriendMarkReply{}, nil
}

// GetRequestPending RequestPending 获取需要审批的好友申请
func (s *FriendService) GetRequestPending(ctx context.Context, req *pb.GetRequestPendingRequest) (*pb.GetRequestPendingReply, error) {
	// 获取 user_id
	uidValue := ctx.Value("user_id")
	sid, ok := uidValue.(float64)
	if !ok {
		return nil, ListFetchError(errors.New("invalid or missing user_id in context"))
	}

	pend, err := s.UC.RequestPend(ctx, uint(sid))
	if err != nil {
		//随便返回一个错误
		return nil, RequestSendError(err)
	}
	var man = make([]*pb.FriendRequestInfo, 0)
	for p := range pend {
		many, err := s.UserClient.GetUniqueByIdMany(ctx, &v1.GetUniqueByIdManyRequest{
			UserId: uint64(pend[p].FromId),
		})
		if err != nil {
			return nil, err
		}
		man = append(man, &pb.FriendRequestInfo{
			FromId:      many.UniqueId,
			RequestTime: pend[p].Time.String(),
		})
	}

	return &pb.GetRequestPendingReply{
		Requests: man,
	}, nil
}
