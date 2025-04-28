# Vue3 + Vite 聊天室前端

## 启动方式

```bash
cd frontend-vue
npm install
npm run dev
```

默认开发端口 http://localhost:5173

## 主要特性
- 登录/注册/聊天室分离
- API 地址统一管理
- WebSocket 实时通信
- 组件化开发，便于维护

## API 代理说明
已配置 Vite 代理，开发时所有 /login、/register、/message 请求自动转发到后端 8080 端口。

## 目录结构
- src/api/config.js 统一管理 API 地址
- src/components/ 业务组件
- src/views/ 页面
