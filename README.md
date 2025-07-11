# Friend Service API

以下是好友服务的接口说明及相关校验逻辑。


好友列表
GET /friend/list
添加好友
POST /friend/request
请求通知页
你需要额外扩展接口，当前 proto 中未定义
接受好友请求
POST /friend/request/accept
拒绝好友请求
POST /friend/request/reject
好友详情页
GET /friend/profile/{unique_id}
修改备注
POST /friend/mark
删除好友
POST /friend/delete

---

## 1. 发送好友请求

**接口**：`POST /friends/request`  
**参数**：
- `target_id`：要添加的用户 unique ID

**校验**：
- 不能添加自己
- 不能已是好友
- 不能已经发送过请求（pending 状态）
- 不能存在对方发给你，但你还没处理的请求

**返回**：
- 请求结果（成功/失败）
- 通知对方

---

## 2. 接受好友请求

**接口**：`POST /friends/request/accept`  
**参数**：
- `from_id`：请求发起方的 unique ID

**校验**：
- 请求是否存在
- 状态是否为 pending

**返回**：
- 操作结果（成功/失败）
- 通知对方

---

## 3. 拒绝好友请求

**接口**：`POST /friends/request/reject`  
**参数**：
- `from_id`：请求发起方的 unique ID

**校验**：
- 请求是否存在
- 状态是否为 pending

**返回**：
- 操作结果（成功/失败）
- 通知对方

---

## 4. 查看好友列表

**接口**：`GET /friends`  
**返回**：
- 好友列表（含对方 ID、备注、昵称、头像等信息）

---

## 5. 删除好友

**接口**：`POST /friends/delete`  
**参数**：
- `target_id`：要删除的好友 unique ID
**校验**：
- 是否是好友关系

**返回**：
- 删除结果（成功/失败）
- 通知对方
---

## 6. 获取好友信息

**接口**：`GET /friends/profile`  
**参数**：
- `friend_id`：好友 unique ID

**校验**：
- 是否是好友关系

**返回**：
- 好友的基本信息
- 对好友的备注

---

## 7. 设置好友备注

**接口**：`POST /friends/remark`  
**参数**：
- `friend_id`：好友 unique ID
- `remark`：备注内容

**校验**：
- 是否是好友关系

**返回**：
- 修改后的备注内容

---

## 8. 获取需要审批的请求列表

**接口**：`GET /friends/requests/pending`  
**返回**：
- 需要审批的好友请求列表（含请求发起方 ID、请求时间、状态等信息）
- 返回格式示例：
  ```json
  [
    {
      "from_id": "user_unique_id_1",
      "request_time": "2025-04-19T00:00:00Z",
      "status": "pending"
    },
    {
      "from_id": "user_unique_id_2",
      "request_time": "2025-04-18T00:00:00Z",
      "status": "pending"
    }
  ]
  ```