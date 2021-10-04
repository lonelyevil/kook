# Features

## Implemented

- [x] Websocket events
- [x] Webhook events
- [x] CardMessage builder
- [x] RolePermission
- [x] Injectable structural logger
  - [x] Integration
  - [x] `phuslu/log` adapter
  - [x] `zap` adapter
- [x] HTTP API

## WIP

## Planned
- [ ] TextMessage router
- [ ] Helper fucntion for messages could emit event

## HTTP API status

- [x] 服务器相关接口 `guild`
  - [x] 获取当前用户加入的服务器列表 `guild/list`
  - [x] 获取服务器详情 `guild/view`
  - [x] 获取服务器中的用户列表 `guild/user-list`
  - [x] 修改服务器中用户的昵称 `guild/nickname`
  - [x] 离开服务器 `guild/leave`
  - [x] 踢出服务器 `guild/kickout`
  - [x] 服务器静音闭麦列表 `guild-mute/list`
  - [x] 添加服务器静音或闭麦 `guild-mute/create`
  - [x] 删除服务器静音或闭麦 `guild-mute/delete`
- [x] 频道相关接口 `channel`
  - [ ] ~~发送频道聊天消息~~ `channel/message`
  - [x] 获取频道列表 `channel/list`
  - [x] 获取频道详情 `channel/view`
  - [x] 创建频道 `channel/create`
  - [x] 删除频道 `channel/delete`
  - [x] 语音频道之间移动用户 `channel/move-user`
  - [x] 获取频道角色权限详情 `channel-role/index`
  - [x] 创建频道角色权限 `channel-role/create`
  - [x] 更新频道角色权限 `channel-role/update`
  - [x] 删除频道角色权限 `channel-role/delete`
- [x] 频道消息相关接口 `message`
  - [x] 获取频道聊天消息列表 `message/list`
  - [x] 发送频道聊天消息 `message/create`
  - [x] 更新频道聊天消息 `message/update`
  - [x] 删除频道聊天消息 `message/delete`
  - [x] 获取频道消息某个回应的用户列表 `message/reaction-list`
  - [x] 给某个消息添加回应 `message/add-reaction`
  - [x] 删除消息的某个回应 `message/delete-reaction`
- [x] 私信聊天会话接口 `user-chat`
  - [x] 获取私信聊天会话列表 `user-chat/list`
  - [x] 获取私信聊天会话详情 `user-chat/view`
  - [x] 创建私信聊天会话 `user-chat/create`
  - [x] 删除私信聊天会话 `user-chat/delete`
- [x] 用户私聊消息接口 `direct-message`
  - [x] 获取私信聊天消息列表 `direct-message/list`
  - [x] 发送私信聊天消息 `direct-message/create`
  - [x] 更新私信聊天消息 `direct-message/update`
  - [x] 删除私信聊天消息 `direct-message/delete`
  - [x] 获取频道消息某个回应的用户列表 `direct-message/reaction-list`
  - [x] 给某个消息添加回应 `direct-message/add-reaction`
  - [x] 删除消息的某个回应 `direct-message/delete-reaction`
- [x] Gateway `gateway`
  - [x] 获取网关连接地址 `gateway/index`
- [x] 用户相关接口 `user`
  - [x] 获取当前用户信息 `user/me`
  - [x] 获取目标用户信息 `user/view`
- [x] 媒体模块 `asset`
  - [x] 上传文件/图片 `asset/create`
- [x] 服务器角色权限相关接口 `guild-role`
  - [x] 获取服务器角色列表 `guild-role/list`
  - [x] 创建服务器角色 `guild-role/create`
  - [x] 更新服务器角色 `guild-role/update`
  - [x] 删除服务器角色 `guild-role/delete`
  - [x] 赋予用户角色 `guild-role/grant`
  - [x] 删除用户角色 `guild-role/revoke`
- [x] 亲密度相关接口 `intimacy`
  - [x] 获取用户的亲密度 `intimacy/index`
  - [x] 更新用户的亲密度 `intimacy/update`
- [x] 服务器表情相关接口 `guild-emoji`
  - [x] 获取服务器表情列表 `guild-emoji/list`
  - [x] 创建服务器表情 `guild-emoji/create`
  - [x] 更新服务器表情 `guild-emoji/update`
  - [x] 删除服务器表情 `guild-emoji/delete`
- [x] 邀请相关接口 `invite`
  - [x] 获取邀请列表 `invite/list`
  - [x] 创建邀请链接 `invite/create`
  - [x] 删除邀请链接 `invite/delete`
