package biz

import (
	"github.com/WH-5/friend-service/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
)

type FriendRepo interface {
}
type FriendUsecase struct {
	repo FriendRepo
	log  *log.Helper
	CF   *conf.Bizfig
}

func NewFriendUsecase(cf *conf.Bizfig, repo FriendRepo, logger log.Logger) *FriendUsecase {
	return &FriendUsecase{CF: cf, repo: repo, log: log.NewHelper(logger)}
}
