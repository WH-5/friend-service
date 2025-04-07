package data

import (
	"friend-service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type friendRepo struct {
	data *Data
	log  *log.Helper
}

func NewFriendRepo(data *Data, logger log.Logger) biz.FriendRepo {
	return &friendRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
