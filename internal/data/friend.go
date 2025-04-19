package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/WH-5/friend-service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type friendRepo struct {
	data *Data
	log  *log.Helper
}

func (f *friendRepo) ModifyMark(ctx context.Context, self, target uint, mark string) error {
	err := f.data.DB.Model(&Friendship{}).Where("user_id = ? AND friend_id = ?", self, target).Update("nickname", mark).Error
	if err != nil {
		return err
	}
	return nil
}

func (f *friendRepo) DeleteFriend(ctx context.Context, self, target uint) error {
	//f.data.DB.Where("user_id = ? AND friend_id = ?",self,target).Delete(&Friendship{})

	err := f.data.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where("user_id = ? AND friend_id = ? OR user_id = ? AND friend_id = ?", self, target, target, self).Delete(&Friendship{}).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("删除好友错误: %s", err)
	}
	return nil
}

func (f *friendRepo) FriendList(ctx context.Context, self uint) ([]biz.FriendInformation, int, error) {
	//没有self这个用户返回空了
	var info []biz.FriendInformation
	err := f.data.DB.Model(&Friendship{}).Where("user_id=?", self).Select("friend_id", "nickname").Scan(&info).Error
	if err != nil {
		return nil, 0, err
	}
	return info, len(info), nil
}

func (f *friendRepo) AcceptRequest(ctx context.Context, self, target uint) error {
	req := &FriendRequest{}
	//接受请求
	err := f.data.DB.Model(req).Where("sender_id = ? AND receiver_id = ? AND status = ? ",
		target, self, "pending").Update("status", "accept").Error //需要修改pending状态的数据
	if err != nil {
		return err
	}
	ship := &Friendship{
		UserID:   self,
		FriendID: target,
		Nickname: "",
	}
	ship2 := &Friendship{
		UserID:   target,
		FriendID: self,
		Nickname: "",
	}
	// 创建好友关系（保持一致性）
	err = f.data.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(ship).Error; err != nil {
			return err
		}
		if err := tx.Create(ship2).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (f *friendRepo) RejectRequest(ctx context.Context, self, target uint) error {
	req := &FriendRequest{}
	//拒绝请求
	err := f.data.DB.Model(req).Where("sender_id = ? AND receiver_id = ? AND status = ? ",
		target, self, "pending").Update("status", "reject").Error //需要修改pending状态的数据
	if err != nil {
		return err
	}
	return nil
}

// IsFriend 判断两人是否有好友关系
func (f *friendRepo) IsFriend(ctx context.Context, self, target uint) (bool, error) {
	//创建好友关系的时候，需要保证数据是对称的
	ship := &Friendship{}
	err := f.data.DB.Model(ship).Where("user_id = ? AND friend_id = ?", self, target).First(ship).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil // 没找到
		}
		return false, err // 其他错误
	}
	return true, nil
}

// HasRequest 判断存在self给target创建的请求没且状态为pending
func (f *friendRepo) HasRequest(ctx context.Context, self, target uint) (bool, error) {
	req := &FriendRequest{}
	err := f.data.DB.Model(req).Where("sender_id = ? AND receiver_id = ? AND status = ? ", self, target, "pending").First(req).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil // 没找到
		}
		return false, err // 其他错误
	}
	return true, nil
}

// MakeRequest 创建请求
func (f *friendRepo) MakeRequest(ctx context.Context, self, target uint) error {
	//self发给target
	req := &FriendRequest{
		SenderID:   self,
		ReceiverID: target,
		Status:     "pending",
	}
	err := f.data.DB.Create(req).Error
	if err != nil {
		return err
	}
	return nil
}

func NewFriendRepo(data *Data, logger log.Logger) biz.FriendRepo {
	return &friendRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
