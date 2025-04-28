# WebSocket 实时聊天室

基于 Go + Vue3 + WebSocket 实现的实时聊天应用，支持多人聊天、消息实时同步、消息编辑等功能。

## 功能特性

- 🔐 用户认证
  - 用户注册
  - 用户登录
  - 会话保持
- 💬 实时聊天
  - WebSocket 实时消息推送
  - 消息持久化存储
  - 支持编辑已发送的消息
  - 自动重连机制
- 🎨 现代化 UI
  - 响应式设计
  - 简洁美观的界面
  - 良好的用户体验

## 技术栈

### 后端
- Go
- Gin (Web 框架)
- Gorilla WebSocket
- GORM (ORM)
- SQLite (数据库)

### 前端
- Vue 3
- Vite
- WebSocket API
- 现代 CSS

### 部署
- Docker
- Docker Compose
- GoReleaser
- Nginx

## 快速开始

### 使用 Docker Compose（推荐）

1. 克隆仓库
```bash
git clone https://github.com/Sggoodman/websocket.chat.git
cd websocket.chat
```

2. 启动服务
```bash
docker-compose up --build
```

3. 访问应用
- 打开浏览器访问 http://localhost

### 本地开发

1. 启动后端
```bash
# 在项目根目录
go mod download
go run .
```

2. 启动前端
```bash
cd frontend-vue
npm install
npm run dev
```

## 项目结构

```
.
├── frontend-vue/          # Vue 3 前端项目
│   ├── src/
│   │   ├── api/          # API 接口
│   │   ├── components/   # Vue 组件
│   │   └── views/        # 页面
│   └── ...
├── *.go                  # Go 后端代码
├── docker-compose.yml    # Docker 编排配置
└── ...
```

## API 文档

### WebSocket 接口
- 连接地址：`ws://localhost:8080/message`
- 消息格式：
  ```json
  {
    "Content": "消息内容",
    "SenderID": 1
  }
  ```

### HTTP 接口
- POST `/register` - 用户注册
- POST `/login` - 用户登录
- PUT `/message` - 更新消息

## 开发计划

- [ ] 添加消息撤回功能
- [ ] 支持图片消息
- [ ] 添加在线用户列表
- [ ] 支持私聊功能
- [ ] 添加消息通知
- [ ] 支持表情消息

## 贡献指南

欢迎提交 Issue 和 Pull Request 来帮助改进项目。

## 许可证

MIT License
