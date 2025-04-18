package biz

import (
	"context"
	"errors"
	"github.com/WH-5/friend-service/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
)

type FriendRepo interface {
	IsFriend(ctx context.Context, self, target uint) (bool, error)
	HasRequest(ctx context.Context, self, target uint) (bool, error)
	MakeRequest(ctx context.Context, self, target uint) error
	AcceptRequest(ctx context.Context, self, target uint) error
	RejectRequest(ctx context.Context, self, target uint) error
	FriendList(ctx context.Context, self uint) ([]FriendInformation, int, error)
}

type FriendUsecase struct {
	repo FriendRepo
	log  *log.Helper
	CF   *conf.Bizfig
}

func NewFriendUsecase(cf *conf.Bizfig, repo FriendRepo, logger log.Logger) *FriendUsecase {
	return &FriendUsecase{CF: cf, repo: repo, log: log.NewHelper(logger)}
}

// SendFriend 发送好友请求
func (uc *FriendUsecase) SendFriend(ctx context.Context, Self, Target uint) error {
	if Self == Target {
		return errors.New("不能添加自己为好友")
	}
	fflag, err := uc.repo.IsFriend(ctx, Self, Target)
	if err != nil {
		return err
	}
	if fflag {
		return errors.New("已经是好友")
	}
	rflag, err := uc.repo.HasRequest(ctx, Self, Target)
	if err != nil {
		return err
	}
	if rflag {
		return errors.New("已经发送过请求")
	}
	r2flag, err := uc.repo.HasRequest(ctx, Target, Self)
	if err != nil {
		return err
	}
	if r2flag {
		return errors.New("请先处理对方的好友申请")
	}
	err = uc.repo.MakeRequest(ctx, Self, Target)
	if err != nil {
		return err
	}
	// TODO 通知好友
	return nil
}

func (uc *FriendUsecase) AcceptFriend(ctx context.Context, Self, Target uint) error {
	if Self == Target {
		return errors.New("不能添加自己为好友")
	}
	ex, err := uc.repo.HasRequest(ctx, Target, Self)
	if err != nil {
		return err
	}
	if !ex {
		return errors.New("好友请求不存在")
	}

	err = uc.repo.AcceptRequest(ctx, Self, Target)
	if err != nil {
		return err
	}
	//TODO通知好友
	return nil
}

func (uc *FriendUsecase) RejectFriend(ctx context.Context, Self, Target uint) error {
	if Self == Target {
		return errors.New("不能添加自己为好友")
	}
	//对方给自己发的请求，target在前，self在后
	ex, err := uc.repo.HasRequest(ctx, Target, Self)
	if err != nil {
		return err
	}
	if !ex {
		return errors.New("好友请求不存在")
	}

	err = uc.repo.RejectRequest(ctx, Self, Target)
	if err != nil {
		return err
	}
	//TODO通知好友
	return nil
}

func (uc *FriendUsecase) ListFriends(ctx context.Context, Self uint) ([]FriendInformation, int, error) {
	//返回好友信息切片、好友数量、err
	list, i, err := uc.repo.FriendList(ctx, Self)

	return list, i, err

}

type FriendInformation struct {
	FriendId uint
	Nickname string
}
