// 统一管理 API 和 WebSocket 地址
const isDev = process.env.NODE_ENV === 'development';

// 开发环境使用 Vite 代理，生产环境使用 Nginx 代理
export const API_BASE_URL = isDev ? '' : '/api';
export const WS_BASE_URL = isDev ? `ws://localhost:8080` : `ws://${window.location.host}/api/ws`;
