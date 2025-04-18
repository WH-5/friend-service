package data

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type FriendRequest struct {
	gorm.Model
	SenderID   uint   `gorm:"not null"`                           // 发送请求的用户 ID
	ReceiverID uint   `gorm:"not null"`                           // 接收请求的用户 ID
	Status     string `gorm:"type:varchar(20);default:'pending'"` // 请求状态
	//双方只能存在一个好友请求
}
type Friendship struct {
	gorm.Model
	UserID   uint   `gorm:"not null"` // 用户A ID
	FriendID uint   `gorm:"not null"` // 用户B ID
	Nickname string `gorm:"size:255"` // A给B的备注
	//备注默认没有，如果为空，客户端会显示uniqueid
	//新增好友时做一个限制：好友上限为500
}
type UserBehaviorLog struct {
	gorm.Model
	UserID   uint            `gorm:"not null"`          // 关联用户 ID
	Action   string          `gorm:"size:255;not null"` // 用户执行的行为
	Metadata json.RawMessage `gorm:"type:jsonb"`        // 行为的相关数据 0注册、1登录、2其他
}

func MigrateDB(db *gorm.DB) error {
	// 执行自动迁移
	err := db.AutoMigrate(&FriendRequest{}, &Friendship{}, &UserBehaviorLog{})
	if err != nil {
		return fmt.Errorf("failed to migrate friend database: %w", err)
	}
	return nil
}
