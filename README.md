# Gin-Ranking 投票系统

基于 Gin 框架开发的投票系统，支持用户注册、登录、投票等功能。

## 功能特点

- 用户注册和登录
- 选手展示
- 投票功能
- 排行榜显示
- 响应式设计

## 技术栈

- 后端：Gin + MySQL + Redis
- 前端：Vue.js + Axios
- 数据库：MySQL
- 缓存：Redis

## 安装和运行

1. 克隆项目
```bash
git clone https://github.com/你的用户名/gin-ranking.git
cd gin-ranking
```

2. 安装依赖
```bash
go mod tidy
```

3. 配置数据库
- 创建 MySQL 数据库
- 修改 config/database.go 中的数据库配置

4. 运行项目
```bash
go run api/main.go
```

5. 访问系统
打开浏览器访问：http://localhost:9999

## 项目结构

```
gin-ranking
├── api/            # API 接口
├── config/         # 配置文件
├── controllers/    # 控制器
├── middleware/     # 中间件
├── models/         # 数据模型
├── router/         # 路由配置
├── utils/          # 工具函数
└── web/           # 前端文件
    ├── css/       # 样式文件
    ├── js/        # JavaScript 文件
    └── images/    # 图片资源
```

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License 