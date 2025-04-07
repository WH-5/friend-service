package biz

import "github.com/go-kratos/kratos/v2/log"

type FriendRepo interface {
}
type FriendUsecase struct {
	repo FriendRepo
	log  *log.Helper
}

func NewFriendUsecase(repo FriendRepo, logger log.Logger) *FriendUsecase {
	return &FriendUsecase{repo: repo, log: log.NewHelper(logger)}
}
